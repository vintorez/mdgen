package md

import (
	"reflect"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewBuilder(t *testing.T) {
	tests := []struct {
		name     string
		expected *Builder
	}{
		{
			name: "empty builder",
			expected: &Builder{
				indents: make([]int, 0, 1),
				buff:    new(strings.Builder),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := NewBuilder()
			assert.True(t, reflect.DeepEqual(b, tt.expected))
		})
	}
}

func TestBuilder_Paragraph(t *testing.T) {
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
			expected: "qwerty\n",
		},
		{
			name:     "a few words",
			input:    "qwerty qwerty qwerty",
			expected: "qwerty qwerty qwerty\n",
		},
		{
			name:     "a sentence",
			input:    "Qwerty qwerty. Qwerty qwerty.",
			expected: "Qwerty qwerty. Qwerty qwerty.\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := NewBuilder()
			b = b.Paragraph(tt.input)
			assert.Equal(t, tt.expected, b.String())
		})
	}
}

func TestBuilder_H1(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "empty input",
			input:    "",
			expected: "\n",
		},
		{
			name:     "one word",
			input:    "qwerty",
			expected: "# qwerty\n\n",
		},
		{
			name:     "a few words",
			input:    "qwerty qwerty qwerty",
			expected: "# qwerty qwerty qwerty\n\n",
		},
		{
			name:     "a sentence",
			input:    "Qwerty qwerty. Qwerty qwerty.",
			expected: "# Qwerty qwerty. Qwerty qwerty.\n\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := NewBuilder()
			b = b.H1(tt.input)
			assert.Equal(t, tt.expected, b.String())
		})
	}
}

func TestBuilder_H2(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "empty input",
			input:    "",
			expected: "\n",
		},
		{
			name:     "one word",
			input:    "qwerty",
			expected: "## qwerty\n\n",
		},
		{
			name:     "a few words",
			input:    "qwerty qwerty qwerty",
			expected: "## qwerty qwerty qwerty\n\n",
		},
		{
			name:     "a sentence",
			input:    "Qwerty qwerty. Qwerty qwerty.",
			expected: "## Qwerty qwerty. Qwerty qwerty.\n\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := NewBuilder()
			b = b.H2(tt.input)
			assert.Equal(t, tt.expected, b.String())
		})
	}
}

func TestBuilder_H3(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "empty input",
			input:    "",
			expected: "\n",
		},
		{
			name:     "one word",
			input:    "qwerty",
			expected: "### qwerty\n\n",
		},
		{
			name:     "a few words",
			input:    "qwerty qwerty qwerty",
			expected: "### qwerty qwerty qwerty\n\n",
		},
		{
			name:     "a sentence",
			input:    "Qwerty qwerty. Qwerty qwerty.",
			expected: "### Qwerty qwerty. Qwerty qwerty.\n\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := NewBuilder()
			b = b.H3(tt.input)
			assert.Equal(t, tt.expected, b.String())
		})
	}
}

func TestBuilder_H4(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "empty input",
			input:    "",
			expected: "\n",
		},
		{
			name:     "one word",
			input:    "qwerty",
			expected: "#### qwerty\n\n",
		},
		{
			name:     "a few words",
			input:    "qwerty qwerty qwerty",
			expected: "#### qwerty qwerty qwerty\n\n",
		},
		{
			name:     "a sentence",
			input:    "Qwerty qwerty. Qwerty qwerty.",
			expected: "#### Qwerty qwerty. Qwerty qwerty.\n\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := NewBuilder()
			b = b.H4(tt.input)
			assert.Equal(t, tt.expected, b.String())
		})
	}
}

func TestBuilder_H5(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "empty input",
			input:    "",
			expected: "\n",
		},
		{
			name:     "one word",
			input:    "qwerty",
			expected: "##### qwerty\n\n",
		},
		{
			name:     "a few words",
			input:    "qwerty qwerty qwerty",
			expected: "##### qwerty qwerty qwerty\n\n",
		},
		{
			name:     "a sentence",
			input:    "Qwerty qwerty. Qwerty qwerty.",
			expected: "##### Qwerty qwerty. Qwerty qwerty.\n\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := NewBuilder()
			b = b.H5(tt.input)
			assert.Equal(t, tt.expected, b.String())
		})
	}
}

func TestBuilder_H6(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "empty input",
			input:    "",
			expected: "\n",
		},
		{
			name:     "one word",
			input:    "qwerty",
			expected: "###### qwerty\n\n",
		},
		{
			name:     "a few words",
			input:    "qwerty qwerty qwerty",
			expected: "###### qwerty qwerty qwerty\n\n",
		},
		{
			name:     "a sentence",
			input:    "Qwerty qwerty. Qwerty qwerty.",
			expected: "###### Qwerty qwerty. Qwerty qwerty.\n\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := NewBuilder()
			b = b.H6(tt.input)
			assert.Equal(t, tt.expected, b.String())
		})
	}
}

func TestBuilder_CodeBlock(t *testing.T) {
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
			expected: "```\nqwerty\n```\n",
		},
		{
			name:     "a few words",
			input:    "qwerty qwerty qwerty",
			expected: "```\nqwerty qwerty qwerty\n```\n",
		},
		{
			name:     "a sentence",
			input:    "Qwerty qwerty. Qwerty qwerty.",
			expected: "```\nQwerty qwerty. Qwerty qwerty.\n```\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := NewBuilder()
			b = b.CodeBlock(tt.input)
			assert.Equal(t, tt.expected, b.String())
		})
	}
}

func TestBuilder_Indent(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			name:     "no input",
			input:    nil,
			expected: []int{defaultIndent},
		},
		{
			name:     "empty input",
			input:    []int{},
			expected: []int{defaultIndent},
		},
		{
			name:     "zero input",
			input:    []int{0},
			expected: []int{},
		},
		{
			name:     "specific input",
			input:    []int{8},
			expected: []int{8},
		},
		{
			name:     "multi input",
			input:    []int{2, 4, 6},
			expected: []int{2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := NewBuilder()
			assert.Equal(t, 0, len(b.indents))
			b = b.Indent(tt.input...)
			assert.Equal(t, tt.expected, b.indents)
		})
	}
}

func TestBuilder_NoIndent(t *testing.T) {
	tests := []struct {
		name     string
		indents  []int
		expected []int
	}{
		{
			name:     "no indents",
			indents:  []int{},
			expected: []int{},
		},
		{
			name:     "default indent",
			indents:  []int{defaultIndent},
			expected: []int{},
		},
		{
			name:     "zero indent",
			indents:  []int{0},
			expected: []int{},
		},
		{
			name:     "one indent",
			indents:  []int{8},
			expected: []int{},
		},
		{
			name:     "two indents",
			indents:  []int{2, 4},
			expected: []int{2},
		},
		{
			name:     "three indents",
			indents:  []int{2, 4, 6},
			expected: []int{2, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := NewBuilder()
			b.indents = tt.indents
			b = b.NoIndent()
			assert.Equal(t, tt.expected, b.indents)
		})
	}
}

func TestBuilder_Indent_NoIndent(t *testing.T) {
	b := NewBuilder()
	assert.Equal(t, []int{}, b.indents)

	b = b.Indent()
	assert.Equal(t, []int{defaultIndent}, b.indents)

	b = b.Indent(8)
	assert.Equal(t, []int{defaultIndent, defaultIndent + 8}, b.indents)

	b = b.Indent(2)
	assert.Equal(t, []int{defaultIndent, defaultIndent + 8, defaultIndent + 8 + 2}, b.indents)

	b = b.NoIndent()
	assert.Equal(t, []int{defaultIndent, defaultIndent + 8}, b.indents)

	b = b.NoIndent()
	assert.Equal(t, []int{defaultIndent}, b.indents)

	b = b.NoIndent()
	assert.Equal(t, []int{}, b.indents)
}

func TestBuilder_UnorderedListItem(t *testing.T) {
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
			expected: "- qwerty\n",
		},
		{
			name:     "a few words",
			input:    "qwerty qwerty qwerty",
			expected: "- qwerty qwerty qwerty\n",
		},
		{
			name:     "a sentence",
			input:    "Qwerty qwerty. Qwerty qwerty.",
			expected: "- Qwerty qwerty. Qwerty qwerty.\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := NewBuilder()
			b = b.UnorderedListItem(tt.input)
			assert.Equal(t, tt.expected, b.String())
		})
	}
}

func TestBuilder_UnorderedList(t *testing.T) {
	tests := []struct {
		name     string
		input    []string
		fns      []ListItemFn
		expected string
	}{
		{
			name:     "empty input + no functions",
			input:    nil,
			fns:      nil,
			expected: "",
		},
		{
			name:     "one line + no functions",
			input:    []string{"qwerty qwerty"},
			fns:      nil,
			expected: "- qwerty qwerty\n\n",
		},
		{
			name:     "multi lines + no functions",
			input:    []string{"qwerty qwerty", "qwerty qwerty", "qwerty qwerty"},
			fns:      nil,
			expected: "- qwerty qwerty\n- qwerty qwerty\n- qwerty qwerty\n\n",
		},
		{
			name:     "empty input + bold function",
			input:    nil,
			fns:      []ListItemFn{Bold},
			expected: "",
		},
		{
			name:     "one line + bold function",
			input:    []string{"qwerty qwerty"},
			fns:      []ListItemFn{Bold},
			expected: "- **qwerty qwerty**\n\n",
		},
		{
			name:     "multi lines + bold function",
			input:    []string{"qwerty qwerty", "qwerty qwerty", "qwerty qwerty"},
			fns:      []ListItemFn{Bold},
			expected: "- **qwerty qwerty**\n- **qwerty qwerty**\n- **qwerty qwerty**\n\n",
		},
		{
			name:     "multi lines + multi functions",
			input:    []string{"qwerty qwerty", "qwerty qwerty", "qwerty qwerty"},
			fns:      []ListItemFn{Bold, Italic},
			expected: "- **qwerty qwerty**\n- **qwerty qwerty**\n- **qwerty qwerty**\n\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := NewBuilder()
			b = b.UnorderedList(tt.input, tt.fns...)
			assert.Equal(t, tt.expected, b.String())
		})
	}
}

func TestBuilder_OrderedListItem(t *testing.T) {
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
			expected: "1. qwerty\n",
		},
		{
			name:     "a few words",
			input:    "qwerty qwerty qwerty",
			expected: "1. qwerty qwerty qwerty\n",
		},
		{
			name:     "a sentence",
			input:    "Qwerty qwerty. Qwerty qwerty.",
			expected: "1. Qwerty qwerty. Qwerty qwerty.\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := NewBuilder()
			b = b.OrderedListItem(tt.input)
			assert.Equal(t, tt.expected, b.String())
		})
	}
}

func TestBuilder_OrderedList(t *testing.T) {
	tests := []struct {
		name     string
		input    []string
		fns      []ListItemFn
		expected string
	}{
		{
			name:     "empty input + no functions",
			input:    nil,
			fns:      nil,
			expected: "",
		},
		{
			name:     "one line + no functions",
			input:    []string{"qwerty qwerty"},
			fns:      nil,
			expected: "1. qwerty qwerty\n\n",
		},
		{
			name:     "multi lines + no functions",
			input:    []string{"qwerty qwerty", "qwerty qwerty", "qwerty qwerty"},
			fns:      nil,
			expected: "1. qwerty qwerty\n1. qwerty qwerty\n1. qwerty qwerty\n\n",
		},
		{
			name:     "empty input + bold function",
			input:    nil,
			fns:      []ListItemFn{Bold},
			expected: "",
		},
		{
			name:     "one line + bold function",
			input:    []string{"qwerty qwerty"},
			fns:      []ListItemFn{Bold},
			expected: "1. **qwerty qwerty**\n\n",
		},
		{
			name:     "multi lines + bold function",
			input:    []string{"qwerty qwerty", "qwerty qwerty", "qwerty qwerty"},
			fns:      []ListItemFn{Bold},
			expected: "1. **qwerty qwerty**\n1. **qwerty qwerty**\n1. **qwerty qwerty**\n\n",
		},
		{
			name:     "multi lines + multi functions",
			input:    []string{"qwerty qwerty", "qwerty qwerty", "qwerty qwerty"},
			fns:      []ListItemFn{Bold, Italic},
			expected: "1. **qwerty qwerty**\n1. **qwerty qwerty**\n1. **qwerty qwerty**\n\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := NewBuilder()
			b = b.OrderedList(tt.input, tt.fns...)
			assert.Equal(t, tt.expected, b.String())
		})
	}
}

func TestBuilder_Ln(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected string
	}{
		{
			name:     "default number of new lines",
			input:    []int{},
			expected: "\n",
		},
		{
			name:     "zero new line",
			input:    []int{0},
			expected: "",
		},
		{
			name:     "one new line",
			input:    []int{1},
			expected: "\n",
		},
		{
			name:     "five new lines",
			input:    []int{5},
			expected: "\n\n\n\n\n",
		},
		{
			name:     "a couple of numbers set",
			input:    []int{2, 3, 4},
			expected: "\n\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := NewBuilder()
			b = b.Ln(tt.input...)
			assert.Equal(t, tt.expected, b.String())
		})
	}
}

func TestBuilder_String(t *testing.T) {
	type fields struct {
		indents []int
		buff    *strings.Builder
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &Builder{
				indents: tt.fields.indents,
				buff:    tt.fields.buff,
			}
			if got := b.String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBuilder_heading(t *testing.T) {
	tests := []struct {
		name     string
		level    HeadingLevel
		input    string
		expected string
	}{
		{
			name:     "empty input",
			level:    L1,
			input:    "",
			expected: "\n",
		},
		{
			name:     "one word",
			level:    L1,
			input:    "qwerty",
			expected: "# qwerty\n\n",
		},
		{
			name:     "a few words",
			level:    L1,
			input:    "qwerty qwerty qwerty",
			expected: "# qwerty qwerty qwerty\n\n",
		},
		{
			name:     "a sentence",
			level:    L1,
			input:    "Qwerty qwerty. Qwerty qwerty.",
			expected: "# Qwerty qwerty. Qwerty qwerty.\n\n",
		},
		{
			name:     "empty input",
			level:    L2,
			input:    "",
			expected: "\n",
		},
		{
			name:     "one word",
			level:    L2,
			input:    "qwerty",
			expected: "## qwerty\n\n",
		},
		{
			name:     "a few words",
			level:    L2,
			input:    "qwerty qwerty qwerty",
			expected: "## qwerty qwerty qwerty\n\n",
		},
		{
			name:     "a sentence",
			level:    L2,
			input:    "Qwerty qwerty. Qwerty qwerty.",
			expected: "## Qwerty qwerty. Qwerty qwerty.\n\n",
		},
		{
			name:     "empty input",
			level:    L3,
			input:    "",
			expected: "\n",
		},
		{
			name:     "one word",
			level:    L3,
			input:    "qwerty",
			expected: "### qwerty\n\n",
		},
		{
			name:     "a few words",
			level:    L3,
			input:    "qwerty qwerty qwerty",
			expected: "### qwerty qwerty qwerty\n\n",
		},
		{
			name:     "a sentence",
			level:    L3,
			input:    "Qwerty qwerty. Qwerty qwerty.",
			expected: "### Qwerty qwerty. Qwerty qwerty.\n\n",
		},
		{
			name:     "empty input",
			level:    L4,
			input:    "",
			expected: "\n",
		},
		{
			name:     "one word",
			level:    L4,
			input:    "qwerty",
			expected: "#### qwerty\n\n",
		},
		{
			name:     "a few words",
			level:    L4,
			input:    "qwerty qwerty qwerty",
			expected: "#### qwerty qwerty qwerty\n\n",
		},
		{
			name:     "a sentence",
			level:    L4,
			input:    "Qwerty qwerty. Qwerty qwerty.",
			expected: "#### Qwerty qwerty. Qwerty qwerty.\n\n",
		},
		{
			name:     "empty input",
			level:    L5,
			input:    "",
			expected: "\n",
		},
		{
			name:     "one word",
			level:    L5,
			input:    "qwerty",
			expected: "##### qwerty\n\n",
		},
		{
			name:     "a few words",
			level:    L5,
			input:    "qwerty qwerty qwerty",
			expected: "##### qwerty qwerty qwerty\n\n",
		},
		{
			name:     "a sentence",
			level:    L5,
			input:    "Qwerty qwerty. Qwerty qwerty.",
			expected: "##### Qwerty qwerty. Qwerty qwerty.\n\n",
		},
		{
			name:     "empty input",
			level:    L6,
			input:    "",
			expected: "\n",
		},
		{
			name:     "one word",
			level:    L6,
			input:    "qwerty",
			expected: "###### qwerty\n\n",
		},
		{
			name:     "a few words",
			level:    L6,
			input:    "qwerty qwerty qwerty",
			expected: "###### qwerty qwerty qwerty\n\n",
		},
		{
			name:     "a sentence",
			level:    L6,
			input:    "Qwerty qwerty. Qwerty qwerty.",
			expected: "###### Qwerty qwerty. Qwerty qwerty.\n\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := NewBuilder()
			b = b.heading(tt.level, tt.input)
			assert.Equal(t, tt.expected, b.String())
		})
	}
}

func TestBuilder_listItem(t *testing.T) {
	tests := []struct {
		name     string
		listType ListType
		input    string
		expected string
	}{
		{
			name:     "empty input",
			listType: Unordered,
			input:    "",
			expected: "",
		},
		{
			name:     "one word",
			listType: Unordered,
			input:    "qwerty",
			expected: "- qwerty\n",
		},
		{
			name:     "a few words",
			listType: Unordered,
			input:    "qwerty qwerty qwerty",
			expected: "- qwerty qwerty qwerty\n",
		},
		{
			name:     "a sentence",
			listType: Unordered,
			input:    "Qwerty qwerty. Qwerty qwerty.",
			expected: "- Qwerty qwerty. Qwerty qwerty.\n",
		},
		{
			name:     "empty input",
			listType: Ordered,
			input:    "",
			expected: "",
		},
		{
			name:     "one word",
			listType: Ordered,
			input:    "qwerty",
			expected: "1. qwerty\n",
		},
		{
			name:     "a few words",
			listType: Ordered,
			input:    "qwerty qwerty qwerty",
			expected: "1. qwerty qwerty qwerty\n",
		},
		{
			name:     "a sentence",
			listType: Ordered,
			input:    "Qwerty qwerty. Qwerty qwerty.",
			expected: "1. Qwerty qwerty. Qwerty qwerty.\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := NewBuilder()
			b = b.listItem(tt.listType, tt.input)
			assert.Equal(t, tt.expected, b.String())
		})
	}
}

func TestBuilder_list(t *testing.T) {
	tests := []struct {
		name     string
		listType ListType
		input    []string
		fns      []ListItemFn
		expected string
	}{
		{
			name:     "empty input + no functions",
			listType: Unordered,
			input:    nil,
			fns:      nil,
			expected: "",
		},
		{
			name:     "one line + no functions",
			listType: Unordered,
			input:    []string{"qwerty qwerty"},
			fns:      nil,
			expected: "- qwerty qwerty\n\n",
		},
		{
			name:     "multi lines + no functions",
			listType: Unordered,
			input:    []string{"qwerty qwerty", "qwerty qwerty", "qwerty qwerty"},
			fns:      nil,
			expected: "- qwerty qwerty\n- qwerty qwerty\n- qwerty qwerty\n\n",
		},
		{
			name:     "empty input + bold function",
			listType: Unordered,
			input:    nil,
			fns:      []ListItemFn{Bold},
			expected: "",
		},
		{
			name:     "one line + bold function",
			listType: Unordered,
			input:    []string{"qwerty qwerty"},
			fns:      []ListItemFn{Bold},
			expected: "- **qwerty qwerty**\n\n",
		},
		{
			name:     "multi lines + bold function",
			listType: Unordered,
			input:    []string{"qwerty qwerty", "qwerty qwerty", "qwerty qwerty"},
			fns:      []ListItemFn{Bold},
			expected: "- **qwerty qwerty**\n- **qwerty qwerty**\n- **qwerty qwerty**\n\n",
		},
		{
			name:     "multi lines + multi functions",
			listType: Unordered,
			input:    []string{"qwerty qwerty", "qwerty qwerty", "qwerty qwerty"},
			fns:      []ListItemFn{Bold, Italic},
			expected: "- **qwerty qwerty**\n- **qwerty qwerty**\n- **qwerty qwerty**\n\n",
		},
		{
			name:     "empty input + no functions",
			listType: Ordered,
			input:    nil,
			fns:      nil,
			expected: "",
		},
		{
			name:     "one line + no functions",
			listType: Ordered,
			input:    []string{"qwerty qwerty"},
			fns:      nil,
			expected: "1. qwerty qwerty\n\n",
		},
		{
			name:     "multi lines + no functions",
			listType: Ordered,
			input:    []string{"qwerty qwerty", "qwerty qwerty", "qwerty qwerty"},
			fns:      nil,
			expected: "1. qwerty qwerty\n1. qwerty qwerty\n1. qwerty qwerty\n\n",
		},
		{
			name:     "empty input + bold function",
			listType: Ordered,
			input:    nil,
			fns:      []ListItemFn{Bold},
			expected: "",
		},
		{
			name:     "one line + bold function",
			listType: Ordered,
			input:    []string{"qwerty qwerty"},
			fns:      []ListItemFn{Bold},
			expected: "1. **qwerty qwerty**\n\n",
		},
		{
			name:     "multi lines + bold function",
			listType: Ordered,
			input:    []string{"qwerty qwerty", "qwerty qwerty", "qwerty qwerty"},
			fns:      []ListItemFn{Bold},
			expected: "1. **qwerty qwerty**\n1. **qwerty qwerty**\n1. **qwerty qwerty**\n\n",
		},
		{
			name:     "multi lines + multi functions",
			listType: Ordered,
			input:    []string{"qwerty qwerty", "qwerty qwerty", "qwerty qwerty"},
			fns:      []ListItemFn{Bold, Italic},
			expected: "1. **qwerty qwerty**\n1. **qwerty qwerty**\n1. **qwerty qwerty**\n\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := NewBuilder()
			b = b.list(tt.listType, tt.input, tt.fns...)
			assert.Equal(t, tt.expected, b.String())
		})
	}
}

func TestBuilder_ln(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected string
	}{
		{
			name:     "default number of new lines",
			input:    []int{},
			expected: "\n",
		},
		{
			name:     "zero new line",
			input:    []int{0},
			expected: "",
		},
		{
			name:     "one new line",
			input:    []int{1},
			expected: "\n",
		},
		{
			name:     "five new lines",
			input:    []int{5},
			expected: "\n\n\n\n\n",
		},
		{
			name:     "a couple of numbers set",
			input:    []int{2, 3, 4},
			expected: "\n\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := NewBuilder()
			b = b.ln(tt.input...)
			assert.Equal(t, tt.expected, b.String())
		})
	}
}

func TestBuilder_write(t *testing.T) {
	tests := []struct {
		name     string
		indents  []int
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
			expected: "qwerty\n",
		},
		{
			name:     "a few words",
			input:    "qwerty qwerty qwerty",
			expected: "qwerty qwerty qwerty\n",
		},
		{
			name:     "a sentence",
			input:    "Qwerty qwerty. Qwerty qwerty.",
			expected: "Qwerty qwerty. Qwerty qwerty.\n",
		},
		{
			name:     "empty input + indent",
			indents:  []int{4},
			input:    "",
			expected: "",
		},
		{
			name:     "one word + indent",
			indents:  []int{4},
			input:    "qwerty",
			expected: "    qwerty\n",
		},
		{
			name:     "a few words + indent",
			indents:  []int{4},
			input:    "qwerty qwerty qwerty",
			expected: "    qwerty qwerty qwerty\n",
		},
		{
			name:     "a sentence + indent",
			indents:  []int{4},
			input:    "Qwerty qwerty. Qwerty qwerty.",
			expected: "    Qwerty qwerty. Qwerty qwerty.\n",
		},
		{
			name:     "empty input + multi indent",
			indents:  []int{4, 6},
			input:    "",
			expected: "",
		},
		{
			name:     "one word + multi indent",
			indents:  []int{4, 6},
			input:    "qwerty",
			expected: "      qwerty\n",
		},
		{
			name:     "a few words + multi indent",
			indents:  []int{4, 6},
			input:    "qwerty qwerty qwerty",
			expected: "      qwerty qwerty qwerty\n",
		},
		{
			name:     "a sentence + multi indent",
			indents:  []int{4, 6},
			input:    "Qwerty qwerty. Qwerty qwerty.",
			expected: "      Qwerty qwerty. Qwerty qwerty.\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := NewBuilder()
			b.indents = tt.indents
			b = b.write(tt.input)
			assert.Equal(t, tt.expected, b.String())
		})
	}
}
