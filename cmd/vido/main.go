package main

import (
	"github.com/tcm5343/vido/pkg/parser"
	"github.com/tcm5343/vido/pkg/status"
)

func main() {
	parser.Parser()
	status.Status()
}
