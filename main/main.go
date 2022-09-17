package main

import (
	"os"
	"fmt"
	"github.com/mike5535/calc"
)

func main() {
	simpExpr := calc.ParseSimpleExpr(calc.ParseExpr(os.Args[1]))
	fmt.Println(calc.CalcSimpleExpr(simpExpr))
}
