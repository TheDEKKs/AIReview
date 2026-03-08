package main

import (
	"fmt"
	"flag"
)

func main () {
	fmt.Println("Hello, World!")

	CurrentBranch := flag.String("cb", "", "Current Branch")
	MainBranch := flag.String("mb", "main", "Main Branch")
	CustomPromt := flag.Bool("cp", false, "Custom Promt")
	OutFile := flag.String("o", "a.txt", "Out File")

	flag.Parse()

	fmt.Println(*CurrentBranch, *MainBranch, *CustomPromt, *OutFile)
}
