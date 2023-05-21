package md

import (
	"strings"
)

type HeadingLevel int

const (
	L1 HeadingLevel = iota + 1
	L2
	L3
	L4
	L5
	L6
)

func Heading(level HeadingLevel, s string) string {
	if s == "" {
		return ""
	}
	return strings.Repeat("#", int(level)) + " " + s
}

func Bold(s string) string {
	if s == "" {
		return ""
	}
	return "**" + s + "**"
}

func Italic(s string) string {
	if s == "" {
		return ""
	}
	return "*" + s + "*"
}

func UnorderedListItem(s string) string {
	if s == "" {
		return ""
	}
	return "* " + s
}

func Code(s string) string {
	if s == "" {
		return ""
	}
	return "`" + s + "`"
}

func CodeBlock(s string, lang ...string) string {
	if s == "" {
		return ""
	}
	if len(lang) > 0 {
		return "```" + lang[0] + "\n" + s + "\n```"
	}
	return "```\n" + s + "\n```"
}

func Link(url string, text ...string) string {
	if url == "" {
		return ""
	}

	txt := ""
	if len(text) > 0 {
		txt = text[0]
	}
	if txt == "" {
		txt = url
	}

	return "[" + txt + "](" + url + ")"
}

func Anchor(s string) string {
	if s == "" {
		return ""
	}
	return "#" + strings.ReplaceAll(s, ":", "")
}
