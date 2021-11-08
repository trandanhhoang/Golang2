package main

import (
	"fmt"
	"test/lalalla"

	"github.com/leekchan/accounting"
)

func main() {
	ac := accounting.Accounting{Symbol: "$", Precision: 2}
	fmt.Println(ac.FormatMoney(12345.1234))
	lalalla.Hehe()
	lalalla.Haha()
}
