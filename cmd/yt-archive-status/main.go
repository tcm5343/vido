package main

import (
	"github.com/tcm5343/yt-archive-status/pkg/parser"
	"github.com/tcm5343/yt-archive-status/pkg/status"
)

func main() {
	parser.Parser()
	status.Status()
}
