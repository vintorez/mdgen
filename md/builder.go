package md

import (
	"strings"
)

const (
	defaultIndent = 4
)

type Builder struct {
	indents []int
	buff    *strings.Builder
}

func NewBuilder() *Builder {
	return &Builder{
		indents: make([]int, 0, 1),
		buff:    new(strings.Builder),
	}
}

func (b *Builder) String() string {
	return b.buff.String()
}

func (b *Builder) H1(s string) *Builder {
	return b.heading(L1, s)
}

func (b *Builder) H2(s string) *Builder {
	return b.heading(L2, s)
}

func (b *Builder) H3(s string) *Builder {
	return b.heading(L3, s)
}

func (b *Builder) H4(s string) *Builder {
	return b.heading(L4, s)
}

func (b *Builder) H5(s string) *Builder {
	return b.heading(L5, s)
}

func (b *Builder) H6(s string) *Builder {
	return b.heading(L6, s)
}

func (b *Builder) Paragraph(s string) *Builder {
	return b.write(s)
}

func (b *Builder) Ln(count ...int) *Builder {
	return b.ln(count...)
}

func (b *Builder) CodeBlock(s string) *Builder {
	return b.write(CodeBlock(s))
}

type ListItemFn func(s string) string

func (b *Builder) UnorderedList(list []string, fns ...ListItemFn) *Builder {
	return b.list(Unordered, list, fns...)
}

func (b *Builder) UnorderedListItem(s string) *Builder {
	return b.listItem(Unordered, s)
}

func (b *Builder) OrderedList(list []string, fns ...ListItemFn) *Builder {
	return b.list(Ordered, list, fns...)
}

func (b *Builder) OrderedListItem(s string) *Builder {
	return b.listItem(Ordered, s)
}

func (b *Builder) WithIndent(fn func(b *Builder)) *Builder {
	b.Indent()
	fn(b)
	b.NoIndent()
	return b
}

func (b *Builder) Indent(counts ...int) *Builder {
	count := 0
	if len(counts) > 0 {
		count = counts[0]
	} else {
		count = defaultIndent
	}
	if count <= 0 {
		return b
	}
	if len(b.indents) > 0 {
		count += b.indents[len(b.indents)-1]
	}
	b.indents = append(b.indents, count)
	return b
}

func (b *Builder) NoIndent() *Builder {
	if len(b.indents) > 0 {
		b.indents = b.indents[:len(b.indents)-1]
	}
	return b
}

func (b *Builder) list(t ListType, list []string, fns ...ListItemFn) *Builder {
	var fn ListItemFn
	if len(fns) > 0 {
		fn = fns[0]
	}

	if len(list) == 0 {
		return b
	}

	b.buff.Grow(len(list))
	for _, item := range list {
		if fn != nil {
			item = fn(item)
		}
		b.listItem(t, item)
	}
	return b.ln()
}

func (b *Builder) listItem(t ListType, s string) *Builder {
	return b.write(ListItem(t, s))
}

func (b *Builder) heading(level HeadingLevel, s string) *Builder {
	return b.write(Heading(level, s)).ln()
}

func (b *Builder) write(s string) *Builder {
	if s == "" {
		return b
	}
	var indent string
	if len(b.indents) > 0 {
		indent = strings.Repeat(" ", b.indents[len(b.indents)-1])
	}
	b.buff.WriteString(indent + s)
	return b.ln()
}

func (b *Builder) ln(count ...int) *Builder {
	n := "\n"
	if len(count) > 0 {
		n = strings.Repeat(n, count[0])
	}
	b.buff.WriteString(n)
	return b
}
