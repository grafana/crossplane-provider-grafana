## ⚠️ Manual intervention required

The automated generation steps did not complete successfully. This PR
contains the version bump commit and any partial generated output.
Please check out the branch, resolve the issues, and push additional
commits before merging.

**Failed steps:**
${FAILED_STEPS}

Common cause: a new resource category was added upstream, which makes
`make generate` panic. Add an entry to `categoryConfig` in
[config/groups.go](https://github.com/grafana/crossplane-provider-grafana/blob/main/config/groups.go).
See [docs/contributing.md#update-resources](https://github.com/grafana/crossplane-provider-grafana/blob/main/docs/contributing.md#update-resources) for guidance.

**Workflow run**: ${RUN_URL}

---

