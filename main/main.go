package main

import (
	"github.com/mike5535/uniq"
	"github.com/mike5535/uniq/uniq_read"
)

func main() {
	param := uniq_read.ParseFlags()
	input := uniq_read.UniqleRead()
	listPair := uniq.Uniqle(param, input)
	uniq.UniqWrite(param, listPair)
}
