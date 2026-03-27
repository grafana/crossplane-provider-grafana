/*
Copyright 2026 Grafana Labs
*/

package generateobserved

import "strings"

func snakeToCamel(acronymSet map[string]bool, s string) string {
	parts := strings.Split(s, "_")
	for i, p := range parts {
		if p == "" {
			continue
		}
		if acronymSet[strings.ToUpper(p)] {
			parts[i] = strings.ToUpper(p)
		} else {
			parts[i] = strings.ToUpper(p[:1]) + p[1:]
		}
	}
	return strings.Join(parts, "")
}

func snakeToCamelJSON(acronymSet map[string]bool, s string) string {
	parts := strings.Split(s, "_")
	for i, p := range parts {
		if p == "" {
			continue
		}
		upper := strings.ToUpper(p)
		if i == 0 {
			if acronymSet[upper] {
				parts[i] = strings.ToLower(p)
			} else {
				parts[i] = p
			}
		} else {
			if acronymSet[upper] {
				parts[i] = upper
			} else {
				parts[i] = strings.ToUpper(p[:1]) + p[1:]
			}
		}
	}
	return strings.Join(parts, "")
}

func buildAcronymSet(cfg Config) map[string]bool {
	m := make(map[string]bool, len(cfg.Acronyms))
	for _, a := range cfg.Acronyms {
		m[a] = true
	}
	return m
}
