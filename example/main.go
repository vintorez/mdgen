package main

import (
	"fmt"

	"github.com/vintorez/mdgen/md"
)

var numbers = []string{"one", "two", "three"}

func main() {
	b := md.NewBuilder()

	b.H1("Primary title")
	b.Paragraph("Some " + md.Bold("bold") + " text.").Ln()

	b.H2("Secondary title")
	// use UnorderedList instead of UnorderedListItem where it is possible
	b.UnorderedList(numbers, func(s string) string {
		return md.Link(md.Anchor(s), md.Code(s))
	})

	// the same as above
	b.Indent()
	for _, number := range numbers {
		b.UnorderedListItem(md.Link(md.Anchor(number), md.Code(number)))
	}
	b.NoIndent()
	b.Ln()

	b.H3("Title")
	b.OrderedList(numbers)
	b.Indent()
	for _, number := range numbers {
		b.OrderedListItem(number)
	}
	b.NoIndent()
	b.Ln()

	for _, number := range numbers {
		b.H4(number)
		b.Paragraph("Description.")
		b.Ln()
	}

	fmt.Println(b.String())
}
