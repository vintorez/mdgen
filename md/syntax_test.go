package md

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBold(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "empty input",
			input:    "",
			expected: "",
		},
		{
			name:     "one word",
			input:    "qwerty",
			expected: "**qwerty**",
		},
		{
			name:     "a few words",
			input:    "qwerty qwerty qwerty",
			expected: "**qwerty qwerty qwerty**",
		},
		{
			name:     "a sentence",
			input:    "Qwerty qwerty. Qwerty qwerty.",
			expected: "**Qwerty qwerty. Qwerty qwerty.**",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Bold(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestItalic(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "empty input",
			input:    "",
			expected: "",
		},
		{
			name:     "one word",
			input:    "qwerty",
			expected: "*qwerty*",
		},
		{
			name:     "a few words",
			input:    "qwerty qwerty qwerty",
			expected: "*qwerty qwerty qwerty*",
		},
		{
			name:     "a sentence",
			input:    "Qwerty qwerty. Qwerty qwerty.",
			expected: "*Qwerty qwerty. Qwerty qwerty.*",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Italic(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestAnchor(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "empty input",
			input:    "",
			expected: "",
		},
		{
			name:     "one word",
			input:    "qwerty",
			expected: "#qwerty",
		},
		{
			name:     "column-separated word",
			input:    ":qwe:rty:",
			expected: "#qwerty",
		},
		{
			name:     "space-separated word",
			input:    "qwe rty",
			expected: "#qwe rty",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Anchor(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestCode(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "empty input",
			input:    "",
			expected: "",
		},
		{
			name:     "one word",
			input:    "qwerty",
			expected: "`qwerty`",
		},
		{
			name:     "a few words",
			input:    "qwerty qwerty qwerty",
			expected: "`qwerty qwerty qwerty`",
		},
		{
			name:     "a sentence",
			input:    "Qwerty qwerty. Qwerty qwerty.",
			expected: "`Qwerty qwerty. Qwerty qwerty.`",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Code(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestCodeBlock(t *testing.T) {
	type args struct {
		input string
		lang  []string
	}
	tests := []struct {
		name     string
		args     args
		expected string
	}{
		{
			name: "empty input + no language",
			args: args{
				input: "",
			},
			expected: "",
		},
		{
			name: "one word + no language",
			args: args{
				input: "qwerty",
			},
			expected: "```\nqwerty\n```",
		},
		{
			name: "a few words + no language",
			args: args{
				input: "qwerty qwerty qwerty",
			},
			expected: "```\nqwerty qwerty qwerty\n```",
		},
		{
			name: "a sentence + no language",
			args: args{
				input: "Qwerty qwerty. Qwerty qwerty.",
			},
			expected: "```\nQwerty qwerty. Qwerty qwerty.\n```",
		},
		{
			name: "empty input + empty language",
			args: args{
				input: "",
				lang:  []string{""},
			},
			expected: "",
		},
		{
			name: "one word + empty language",
			args: args{
				input: "qwerty",
				lang:  []string{""},
			},
			expected: "```\nqwerty\n```",
		},
		{
			name: "a few words + empty language",
			args: args{
				input: "qwerty qwerty qwerty",
				lang:  []string{""},
			},
			expected: "```\nqwerty qwerty qwerty\n```",
		},
		{
			name: "a sentence + empty language",
			args: args{
				input: "Qwerty qwerty. Qwerty qwerty.",
				lang:  []string{""},
			},
			expected: "```\nQwerty qwerty. Qwerty qwerty.\n```",
		},
		{
			name: "empty input + language",
			args: args{
				input: "",
				lang:  []string{"go"},
			},
			expected: "",
		},
		{
			name: "one word + language",
			args: args{
				input: "qwerty",
				lang:  []string{"go"},
			},
			expected: "```go\nqwerty\n```",
		},
		{
			name: "a few words + language",
			args: args{
				input: "qwerty qwerty qwerty",
				lang:  []string{"go"},
			},
			expected: "```go\nqwerty qwerty qwerty\n```",
		},
		{
			name: "a sentence + language",
			args: args{
				input: "Qwerty qwerty. Qwerty qwerty.",
				lang:  []string{"go"},
			},
			expected: "```go\nQwerty qwerty. Qwerty qwerty.\n```",
		},

		{
			name: "empty input + multi language",
			args: args{
				input: "",
				lang:  []string{"go", "sql"},
			},
			expected: "",
		},
		{
			name: "one word + multi language",
			args: args{
				input: "qwerty",
				lang:  []string{"go", "sql"},
			},
			expected: "```go\nqwerty\n```",
		},
		{
			name: "a few words + multi language",
			args: args{
				input: "qwerty qwerty qwerty",
				lang:  []string{"go", "sql"},
			},
			expected: "```go\nqwerty qwerty qwerty\n```",
		},
		{
			name: "a sentence + multi language",
			args: args{
				input: "Qwerty qwerty. Qwerty qwerty.",
				lang:  []string{"go", "sql"},
			},
			expected: "```go\nQwerty qwerty. Qwerty qwerty.\n```",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := CodeBlock(tt.args.input, tt.args.lang...)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestHeadingLevel1(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "empty input",
			input:    "",
			expected: "",
		},
		{
			name:     "one word",
			input:    "qwerty",
			expected: "# qwerty",
		},
		{
			name:     "a few words",
			input:    "qwerty qwerty qwerty",
			expected: "# qwerty qwerty qwerty",
		},
		{
			name:     "a sentence",
			input:    "Qwerty qwerty. Qwerty qwerty.",
			expected: "# Qwerty qwerty. Qwerty qwerty.",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Heading(L1, tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestHeadingLevel2(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "empty input",
			input:    "",
			expected: "",
		},
		{
			name:     "one word",
			input:    "qwerty",
			expected: "## qwerty",
		},
		{
			name:     "a few words",
			input:    "qwerty qwerty qwerty",
			expected: "## qwerty qwerty qwerty",
		},
		{
			name:     "a sentence",
			input:    "Qwerty qwerty. Qwerty qwerty.",
			expected: "## Qwerty qwerty. Qwerty qwerty.",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Heading(L2, tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestHeadingLevel3(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "empty input",
			input:    "",
			expected: "",
		},
		{
			name:     "one word",
			input:    "qwerty",
			expected: "### qwerty",
		},
		{
			name:     "a few words",
			input:    "qwerty qwerty qwerty",
			expected: "### qwerty qwerty qwerty",
		},
		{
			name:     "a sentence",
			input:    "Qwerty qwerty. Qwerty qwerty.",
			expected: "### Qwerty qwerty. Qwerty qwerty.",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Heading(L3, tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestHeadingLevel4(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "empty input",
			input:    "",
			expected: "",
		},
		{
			name:     "one word",
			input:    "qwerty",
			expected: "#### qwerty",
		},
		{
			name:     "a few words",
			input:    "qwerty qwerty qwerty",
			expected: "#### qwerty qwerty qwerty",
		},
		{
			name:     "a sentence",
			input:    "Qwerty qwerty. Qwerty qwerty.",
			expected: "#### Qwerty qwerty. Qwerty qwerty.",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Heading(L4, tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestHeadingLevel5(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "empty input",
			input:    "",
			expected: "",
		},
		{
			name:     "one word",
			input:    "qwerty",
			expected: "##### qwerty",
		},
		{
			name:     "a few words",
			input:    "qwerty qwerty qwerty",
			expected: "##### qwerty qwerty qwerty",
		},
		{
			name:     "a sentence",
			input:    "Qwerty qwerty. Qwerty qwerty.",
			expected: "##### Qwerty qwerty. Qwerty qwerty.",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Heading(L5, tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestHeadingLevel6(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "empty input",
			input:    "",
			expected: "",
		},
		{
			name:     "one word",
			input:    "qwerty",
			expected: "###### qwerty",
		},
		{
			name:     "a few words",
			input:    "qwerty qwerty qwerty",
			expected: "###### qwerty qwerty qwerty",
		},
		{
			name:     "a sentence",
			input:    "Qwerty qwerty. Qwerty qwerty.",
			expected: "###### Qwerty qwerty. Qwerty qwerty.",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Heading(L6, tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestUnorderedListItem(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "empty input",
			input:    "",
			expected: "",
		},
		{
			name:     "one word",
			input:    "qwerty",
			expected: "- qwerty",
		},
		{
			name:     "a few words",
			input:    "qwerty qwerty qwerty",
			expected: "- qwerty qwerty qwerty",
		},
		{
			name:     "a sentence",
			input:    "Qwerty qwerty. Qwerty qwerty.",
			expected: "- Qwerty qwerty. Qwerty qwerty.",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ListItem(Unordered, tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestOrderedListItem(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "empty input",
			input:    "",
			expected: "",
		},
		{
			name:     "one word",
			input:    "qwerty",
			expected: "1. qwerty",
		},
		{
			name:     "a few words",
			input:    "qwerty qwerty qwerty",
			expected: "1. qwerty qwerty qwerty",
		},
		{
			name:     "a sentence",
			input:    "Qwerty qwerty. Qwerty qwerty.",
			expected: "1. Qwerty qwerty. Qwerty qwerty.",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ListItem(Ordered, tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestInvalidListItem(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "empty input",
			input:    "",
			expected: "",
		},
		{
			name:     "one word",
			input:    "qwerty",
			expected: "",
		},
		{
			name:     "a few words",
			input:    "qwerty qwerty qwerty",
			expected: "",
		},
		{
			name:     "a sentence",
			input:    "Qwerty qwerty. Qwerty qwerty.",
			expected: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ListItem(ListType(2), tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestLink(t *testing.T) {
	type args struct {
		url  string
		text string
	}
	tests := []struct {
		name     string
		args     args
		expected string
	}{
		{
			name: "empty url + empty text",
			args: args{
				url:  "",
				text: "",
			},
			expected: "",
		},
		{
			name: "url + empty text",
			args: args{
				url:  "https://example.com",
				text: "",
			},
			expected: "[https://example.com](https://example.com)",
		},
		{
			name: "empty url + text",
			args: args{
				url:  "",
				text: "qwerty",
			},
			expected: "",
		},
		{
			name: "url + text",
			args: args{
				url:  "https://example.com",
				text: "qwerty",
			},
			expected: "[qwerty](https://example.com)",
		},
		{
			name: "url + long text",
			args: args{
				url:  "https://example.com",
				text: "qwerty qwerty qwerty",
			},
			expected: "[qwerty qwerty qwerty](https://example.com)",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Link(tt.args.url, tt.args.text)
			assert.Equal(t, tt.expected, result)
		})
	}
}
