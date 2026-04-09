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
		upper := strings.ToUpper(p)
		switch {
		case acronymSet[upper]:
			parts[i] = upper
		case strings.HasSuffix(p, "s") && acronymSet[strings.ToUpper(strings.TrimSuffix(p, "s"))]:
			parts[i] = strings.ToUpper(strings.TrimSuffix(p, "s")) + "s"
		default:
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
		trimmed := strings.TrimSuffix(p, "s")
		isPluralAcronym := strings.HasSuffix(p, "s") && acronymSet[strings.ToUpper(trimmed)]
		switch {
		case i == 0 && acronymSet[upper]:
			parts[i] = strings.ToLower(p)
		case i == 0 && isPluralAcronym:
			parts[i] = strings.ToLower(trimmed) + "s"
		case i == 0:
			parts[i] = p
		case acronymSet[upper]:
			parts[i] = upper
		case isPluralAcronym:
			parts[i] = strings.ToUpper(trimmed) + "s"
		default:
			parts[i] = strings.ToUpper(p[:1]) + p[1:]
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
