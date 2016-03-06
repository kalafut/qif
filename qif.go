package qif

import (
	"bufio"
	"log"
	"os"
	"strings"
	"time"
)

type Split struct {
	Category string
	Memo     string
	Amount   string
}

type Transaction struct {
	Date     time.Time
	Payee    string
	CheckNum string
	Cleared  string
	Splits   []*Split
}

func NewTransaction() Transaction {
	return Transaction{Splits: []*Split{}}
}

func Parse(filename string) []Transaction {
	transactions := []Transaction{}

	f, err := os.Open(filename)

	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(f)

	tx := NewTransaction()
	split := new(Split)
	tx.Splits = append(tx.Splits, split)
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

			tx.Date = t
		case '^':
			// If there are splits, we also end up picking up the
			// overall total, which we don't want.
			if len(tx.Splits) > 1 {
				tx.Splits = tx.Splits[1:]
			}
			transactions = append(transactions, tx)

			tx = NewTransaction()
			split = new(Split)
			tx.Splits = append(tx.Splits, split)
		case 'P':
			tx.Payee = data
			if strings.HasPrefix(data, "Transfer :") {
				split.Category = data[11:]
			}

		case 'E':
			fallthrough
		case 'M':
			split.Memo = data

		case '$':
			fallthrough
		case 'T':
			split.Amount = data

		case 'L':
			split.Category = data
		case 'S':
			split = new(Split)
			tx.Splits = append(tx.Splits, split)

			split.Category = data
		default:
			break

		}
	}

	return transactions
}
