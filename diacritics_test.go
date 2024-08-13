package diacritics

import "testing"

func TestRemoveDiacritics(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "empty string",
			input:    "",
			expected: "",
		},
		{
			name:     "string without diacritics",
			input:    "hello",
			expected: "hello",
		},
		{
			name:     "string with diacritics",
			input:    "héllö",
			expected: "hello",
		},
		{
			name:     "string with diacritics",
			input:    "héllö, wörld",
			expected: "hello, world",
		},
		{
			name:     "string with diacritics",
			input:    "héllö, wörld!",
			expected: "hello, world!",
		},
		{
			name:     "string with diacritics",
			input:    "Ťĥůş şţŕîñĝ ĥâş ďîâçřîţîçş",
			expected: "Thus string has diacritics",
		},
		{
			name:     "string with chinese characters",
			input:    "你好，世界！",
			expected: "你好，世界！",
		},
		{
			name:     "string with chinese characters and diacritics",
			input:    "你好，世界！héllö, wörld!",
			expected: "你好，世界！hello, world!",
		},
		{
			name:     "string with german double s and diacritics",
			input:    "héllö ß",
			expected: "hello ß",
		},
		{
			name:     "string with korean characters and diacritics",
			input:    "안녕하세요, 세계! héllö, wörld!",
			expected: "안녕하세요, 세계! hello, world!",
		},
		{
			name:     "string with russian characters and diacritics",
			input:    "Привет, мир! héllö, wörld!",
			expected: "Привет, мир! hello, world!",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual, err := Remove(tt.input)
			if err != nil {
				t.Errorf("Remove(%q) error = %v", tt.input, err)
				return
			}
			if actual != tt.expected {
				t.Errorf("Remove(%q) = %q, want %q", tt.input, actual, tt.expected)
			}
		})
	}
}
