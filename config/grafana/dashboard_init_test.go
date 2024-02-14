package grafana

import "testing"

func TestReplaceInterpolation(t *testing.T) {
	cases := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "no interpolation",
			input:    "no interpolation",
			expected: "no interpolation",
		},
		{
			name:     "interpolation",
			input:    "interpolation ${}",
			expected: "interpolation $${}",
		},
		{
			name:     "escaped interpolation",
			input:    "escaped interpolation $${}",
			expected: "escaped interpolation $${}",
		},
		{
			name:     "multiple interpolations",
			input:    "multiple ${} interpolations ${}",
			expected: "multiple $${} interpolations $${}",
		},
		{
			name:     "multiple escaped interpolations",
			input:    "interpolation $${} with escaped $${}",
			expected: "interpolation $${} with escaped $${}",
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			actual := replaceInterpolation(tc.input)
			if actual != tc.expected {
				t.Errorf("expected %s, got %s", tc.expected, actual)
			}
		})
	}
}
