package main

import (
	"assignment1/data"
	"fmt"
	"os"
	"strconv"
)

func main() {
	arg := os.Args[1]

	noAbsen, err := strconv.Atoi(arg)
	if err != nil {
		fmt.Println("Number argument Required")
		return
	}

	totalClassMate := data.GetTotalClassMate()
	if noAbsen < 0 || noAbsen > totalClassMate {
		fmt.Println("No absen tidak ditemukan")
		return
	}

	classMate := data.GetClassMate(noAbsen)
	fmt.Printf(
		"No Absen \t: %d\nNama \t\t: %s\nAlamat \t\t: %s\nPekerjaan \t: %s\nAlasan \t\t: %s\n",
		classMate.GetNo(), classMate.GetNama(), classMate.GetAlamat(), classMate.GetPekerjaan(), classMate.GetAlasan())
}
