package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

type Split struct {
	category string
	memo     string
	amount   string
}

type Transaction struct {
	date     time.Time
	payee    string
	checkNum string
	cleared  string
	splits   []*Split
}

func NewTransaction() Transaction {
	return Transaction{splits: []*Split{}}
}

func main() {
	flag.Parse()
	filename := flag.Arg(0)
	account := flag.Arg(1)

	transactions := parse(filename)
	export(transactions, "Test Account")

	_ = account
}

func parse(filename string) []Transaction {
	transactions := []Transaction{}

	f, err := os.Open(filename)

	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(f)

	tx := NewTransaction()
	split := new(Split)
	tx.splits = append(tx.splits, split)
	for scanner.Scan() {
		line := scanner.Text()

		prefix, data := line[0], line[1:]

		switch prefix {
		case '!':
			break
		case 'D':
			t, err := time.Parse("01/02/2006", data)
			if err != nil {
				t, err = time.Parse("01/02/06", data)
			}
			if err != nil {
				log.Fatal(err)
			}

			tx.date = t
		case '^':
			// If there are splits, we also end up picking up the
			// overall total, which we don't want.
			if len(tx.splits) > 1 {
				tx.splits = tx.splits[1:]
			}
			transactions = append(transactions, tx)

			tx = NewTransaction()
			split = new(Split)
			tx.splits = append(tx.splits, split)
		case 'P':
			tx.payee = data
			if strings.HasPrefix(data, "Transfer :") {
				split.category = data[11:]
			}

		case 'E':
			fallthrough
		case 'M':
			split.memo = data

		case '$':
			fallthrough
		case 'T':
			split.amount = data

		case 'L':
			split.category = data
		case 'S':
			split = new(Split)
			tx.splits = append(tx.splits, split)

			split.category = data
		default:
			break

		}
	}

	return transactions
}

func export(txs []Transaction, acct string) {
	for _, tx := range txs {
		//fmt.Println(tx.splits[0])
		fmt.Printf("%s   %s\n", tx.date.Format("2006/01/02"), tx.payee)
		for _, s := range tx.splits {
			fmt.Printf("   %-40s  $%s\n", s.category, s.amount)
		}
		fmt.Printf("   %s\n", acct)
		fmt.Println()

	}
}
