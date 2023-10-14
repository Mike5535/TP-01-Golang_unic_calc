package uniq_read

import (
	"fmt"
	"bufio"
	"flag"
	"io"
	"os"
	"github.com/mike5535/uniq/uniq_types"
)

func ParseFlags() (*uniq_types.Options) {
	opt := new(uniq_types.Options)
	opt.FieldOptions = new(uniq_types.OptFields)

	flag.BoolVar(&opt.Counting, "c", false, "strings with number of repeat")
	flag.BoolVar(&opt.Repeat, "d", false, "repeating strings")
	flag.BoolVar(&opt.Uniq, "u", false, "uniq strings")
	flag.BoolVar(&opt.FieldOptions.AnyRegister, "i", false, "ignore register")
	flag.IntVar(&opt.FieldOptions.NumFields, "f", 0, "number of fields to ignore")
	flag.IntVar(&opt.FieldOptions.NumChars, "s", 0, "number of chars to ignore")
	flag.Parse()

	return opt
}

func UniqleRead() ([]string) {
	var in io.Reader
	if filename := flag.Arg(0); filename != "" {
		f, err := os.Open(filename)
		if err != nil {
			fmt.Println("error opening file: err:", err)
			os.Exit(1)
		}
		defer f.Close()

		in = f
	} else {
		in = os.Stdin
	}

	buf := bufio.NewScanner(in)

	input := []string{}

	for buf.Scan() {
		input = append(input, buf.Text())
	}

	if err := buf.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "error reading: err:", err)
	}
	return input
}
