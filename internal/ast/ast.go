package ast

//go:generate ragel-go -G2 -o ast_parser.go ast_parser.rl

type Program struct{}
