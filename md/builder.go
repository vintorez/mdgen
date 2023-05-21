package md

import (
	"strings"
)

type Builder struct {
	buff *strings.Builder
}

func NewBuilder() *Builder {
	return &Builder{
		buff: new(strings.Builder),
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

func (b *Builder) ListItem(s string) *Builder {
	return b.write(UnorderedListItem(s))
}

func (b *Builder) heading(level HeadingLevel, s string) *Builder {
	return b.write(Heading(level, s)).ln()
}

func (b *Builder) write(s string) *Builder {
	if s == "" {
		return b
	}
	b.buff.WriteString(s)
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
