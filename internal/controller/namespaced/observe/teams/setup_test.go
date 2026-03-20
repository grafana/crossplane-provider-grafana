package teams

import (
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"strings"
	"testing"
	"time"

	"github.com/grafana/crossplane-provider-grafana/v2/internal/clients"
	"github.com/grafana/grafana-openapi-client-go/models"
)

// dockerMode is true when TestMain successfully started a Grafana container.
// Integration tests can use this to make deterministic assertions.
var dockerMode bool

// seededTeams holds the names of the teams seeded into the Docker container.
var seededTeams = struct {
	All       []string // all seeded team names
	OpsQuery  []string // teams matching query="ops"
	ExactName string   // name used for exact-name filter test
}{
	All:       []string{"alpha-team", "beta-team", "ops-east", "ops-west"},
	OpsQuery:  []string{"ops-east", "ops-west"},
	ExactName: "alpha-team",
}

const (
	grafanaDockerImage   = "grafana/grafana:11.0.0"
	grafanaContainerName = "grafana-crossplane-test"
	grafanaLocalPort     = "13000"
	grafanaLocalURL      = "http://localhost:" + grafanaLocalPort
	grafanaAdminAuth     = "admin:admin"
	dockerStartTimeout   = 60 * time.Second
)

func TestMain(m *testing.M) {
	// Only activate when GRAFANA_DOCKER=1 and GRAFANA_AUTH is not already set.
	if os.Getenv("GRAFANA_DOCKER") != "1" || os.Getenv("GRAFANA_AUTH") != "" {
		os.Exit(m.Run())
	}

	started, err := startGrafanaContainer()
	if err != nil {
		fmt.Fprintf(os.Stderr, "WARNING: could not start Grafana Docker container: %v\n", err)
		fmt.Fprintf(os.Stderr, "         Integration tests will be skipped.\n")
		os.Exit(m.Run())
	}
	if started {
		defer stopGrafanaContainer()
	}

	if err := waitForGrafana(grafanaLocalURL, dockerStartTimeout); err != nil {
		fmt.Fprintf(os.Stderr, "WARNING: Grafana did not become healthy: %v\n", err)
		fmt.Fprintf(os.Stderr, "         Integration tests will be skipped.\n")
		os.Exit(m.Run())
	}

	if err := seedTestTeams(grafanaLocalURL, grafanaAdminAuth); err != nil {
		fmt.Fprintf(os.Stderr, "WARNING: could not seed test teams: %v\n", err)
		fmt.Fprintf(os.Stderr, "         Integration tests will be skipped.\n")
		os.Exit(m.Run())
	}

	// Wire the env vars so grafanaClientFromEnv picks them up.
	os.Setenv("GRAFANA_URL", grafanaLocalURL)   //nolint:errcheck
	os.Setenv("GRAFANA_AUTH", grafanaAdminAuth) //nolint:errcheck
	dockerMode = true

	os.Exit(m.Run())
}

// startGrafanaContainer starts a Grafana Docker container.
// Returns (true, nil) if the container was started, (false, nil) if it was
// already running, or (false, err) if start failed.
func startGrafanaContainer() (bool, error) {
	cmd := exec.Command(
		"docker", "run",
		"--detach",
		"--rm",
		"--name", grafanaContainerName,
		"-p", grafanaLocalPort+":3000",
		"-e", "GF_AUTH_ANONYMOUS_ENABLED=false",
		grafanaDockerImage,
	)
	out, err := cmd.CombinedOutput()
	if err != nil {
		// If the container name is already in use, that's a hard error (leftover
		// from a previous crashed run). The caller will skip integration tests.
		return false, fmt.Errorf("docker run: %w\noutput: %s", err, strings.TrimSpace(string(out)))
	}
	fmt.Fprintf(os.Stderr, "Started Grafana container: %s\n", strings.TrimSpace(string(out)))
	return true, nil
}

// stopGrafanaContainer stops (and auto-removes via --rm) the test container.
func stopGrafanaContainer() {
	cmd := exec.Command("docker", "stop", grafanaContainerName)
	if out, err := cmd.CombinedOutput(); err != nil {
		fmt.Fprintf(os.Stderr, "WARNING: docker stop failed: %v\noutput: %s\n", err, string(out))
	} else {
		fmt.Fprintf(os.Stderr, "Stopped Grafana container %s\n", grafanaContainerName)
	}
}

// waitForGrafana polls GET <url>/api/health until it returns HTTP 200 or the
// timeout elapses.
func waitForGrafana(url string, timeout time.Duration) error {
	deadline := time.Now().Add(timeout)
	healthURL := url + "/api/health"
	for time.Now().Before(deadline) {
		resp, err := http.Get(healthURL) //nolint:noctx
		if err == nil {
			resp.Body.Close()
			if resp.StatusCode == http.StatusOK {
				fmt.Fprintf(os.Stderr, "Grafana is healthy at %s\n", url)
				return nil
			}
		}
		time.Sleep(time.Second)
	}
	return fmt.Errorf("timed out waiting for Grafana at %s after %s", url, timeout)
}

// seedTestTeams creates the known test teams in the running Grafana instance.
func seedTestTeams(grafanaURL, auth string) error {
	c, err := clients.NewOAPIClient(grafanaURL, auth)
	if err != nil {
		return fmt.Errorf("NewOAPIClient: %w", err)
	}

	seed := []models.CreateTeamCommand{
		{Name: "alpha-team", Email: "alpha@example.com"},
		{Name: "beta-team", Email: "beta@example.com"},
		{Name: "ops-east", Email: "ops-east@example.com"},
		{Name: "ops-west", Email: "ops-west@example.com"},
	}

	for i := range seed {
		if _, err := c.Teams.CreateTeam(&seed[i]); err != nil {
			return fmt.Errorf("CreateTeam %q: %w", seed[i].Name, err)
		}
		fmt.Fprintf(os.Stderr, "Seeded team: %s\n", seed[i].Name)
	}
	return nil
}
