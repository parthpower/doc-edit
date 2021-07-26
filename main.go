package main

import (
	"fmt"
	"time"

	"github.com/nguyenthenguyen/docx"
)

var in string
var out string
var amount0 string
var amount1 string

func init() {
	return
	// flag.StringVar(&in, "f", "", "docx file to read")
	// flag.StringVar(&out, "o", "", "docx file to read")
	// flag.IntVar(&amount0, "amount0", 0, "value for #__AMOUNT_0__#")
	// flag.IntVar(&amount0, "amount1", 0, "value for #__AMOUNT_1__#")
}

func main() {
	// flag.Parse()
	getInput()
	r, err := docx.ReadDocxFile(in)
	if err != nil {
		panic(err)
	}
	docx1 := r.Editable()

	docx1.Replace("#__AMOUNT_0__#", amount0, -1)
	fmt.Printf("Writing #__AMOUNT_0__# to %v\n", amount0)
	docx1.Replace("#__AMOUNT_1__#", amount1, -1)
	fmt.Printf("Writing #__AMOUNT_1__# to %v\n", amount1)
	date := time.Now().Local().String()
	docx1.Replace("#__DATE__#", date, -1)
	fmt.Printf("Writing #__DATE__# to %v\n", date)
	docx1.WriteToFile(out)
	fmt.Printf("Done writing file: %v\n", out)
	fmt.Printf("Press any key to exit...\n")
	fmt.Scanln()
}

func getInput() {
	fmt.Println("File To Edit:")
	fmt.Scanln(&in)
	fmt.Println("Output File:")
	fmt.Scanln(&out)
	fmt.Println("#__AMOUNT_0__# value:")
	fmt.Scanln(&amount0)
	fmt.Println("#__AMOUNT_1__# value:")
	fmt.Scanln(&amount1)
}
