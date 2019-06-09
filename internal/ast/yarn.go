package ast

import "strings"

type YARN string

func unescape(s string) YARN {
	var b strings.Builder

	escapeStarted := false
	for _, r := range s {
		switch {
		case escapeStarted:
			switch r {
			case ')':
				b.WriteRune('\n')
			case '>':
				b.WriteRune('\t')
			case 'o':
				b.WriteRune('\a')
			case '"':
				b.WriteRune('"')
			case ':':
				b.WriteRune(':')
			default:
				b.WriteRune(':')
				b.WriteRune(r)
			}
			escapeStarted = false
		case r == ':':
			escapeStarted = true
		default:
			b.WriteRune(r)
		}
	}
	return YARN(b.String())
}
