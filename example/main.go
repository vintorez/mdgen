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
	for _, number := range numbers {
		b.ListItem(md.Link(md.Anchor(number), md.Code(number)))
	}
	b.Ln()

	for _, number := range numbers {
		b.H3(number)
		b.Paragraph("Description.")
		b.Ln()
	}

	fmt.Println(b.String())
}
