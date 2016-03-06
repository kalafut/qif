package main

import (
	"flag"
	"fmt"
)
import "github.com/kalafut/qif"

func main() {
	flag.Parse()
	filename := flag.Arg(0)
	account := flag.Arg(1)

	transactions := qif.Parse(filename)
	export(transactions, "Test Account")

	_ = account
}

func export(txs []qif.Transaction, acct string) {
	for _, tx := range txs {
		//fmt.Println(tx.splits[0])
		fmt.Printf("%s   %s\n", tx.Date.Format("2006/01/02"), tx.Payee)
		for _, s := range tx.Splits {
			fmt.Printf("   %-40s  $%s\n", s.Category, s.Amount)
		}
		fmt.Printf("   %s\n", acct)
		fmt.Println()

	}
}
