package main

import (
	"fmt"

	"github.com/vintorez/mdgen/md"
)

var numbers = []string{"one", "two", "three"}

func main() {
	b := md.NewBuilder()

	b.H1("First title")
	b.Paragraph("Some " + md.Bold("bold") + " and " + md.Italic("italic") + " text.").Ln()

	b.H3("An unordered list with no indent")
	b.UnorderedList(numbers, func(s string) string {
		return md.Link(md.Anchor(s), md.Code(s))
	})

	// add default indent
	b.WithIndent(func(b *md.Builder) {
		b.H3("An unordered list with a single indent")
		b.UnorderedList(numbers, func(s string) string {
			return md.Link(md.Anchor(s), md.Code(s))
		})

		b.WithIndent(func(b *md.Builder) {
			b.H3("An unordered list with a double indent")
			b.UnorderedList(numbers, func(s string) string {
				return md.Link(md.Anchor(s), md.Code(s))
			})
		})
	})

	// the same as WithIndent but the length of indent can be set manually
	b.Indent(2)
	// use UnorderedList instead of UnorderedListItem where it is possible
	for _, number := range numbers {
		b.UnorderedListItem(md.Link(md.Anchor(number), md.Code(number)))
	}
	b.NoIndent()
	b.Ln()

	b.H3("Second title")
	b.OrderedList(numbers)

	for _, number := range numbers {
		b.H4(number)
		b.Paragraph("Description.")
		b.Ln()
	}

	fmt.Println(b.String())
}
