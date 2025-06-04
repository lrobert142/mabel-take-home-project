package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"mabel-take-home-project/internal/model/collector"
	"os"
)

func main() {
	//Note: os.Args[0] is the name of the program, so start from 1 instead.

	fmt.Println("Loading accounts...")
	accountsCollector := collector.NewAccountsCollector()
	if err := loadFromCsvFile(os.Args[1], accountsCollector); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Loading transactions...")
	transactionsCollector := collector.NewTransactionsCollector(accountsCollector.GetAccounts())
	if err := loadFromCsvFile(os.Args[2], transactionsCollector); err != nil {
		log.Fatal(err)
	}

	for i, transaction := range transactionsCollector.GetTransactions() {
		fmt.Printf("Running transaction %v...\n", i)
		if err := transaction.Transact(); err != nil {
			log.Fatal(err)
		}
	}

	fmt.Println("Successfully completed all transactions!")
}

// loadFromCsvFile loads data from a CSV file, passing it to a collector.Collector to transform it into something more
// useful
func loadFromCsvFile(filePath string, collector collector.Collector) error {
	f, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("unable to read input file %q: %w", filePath, err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	for {
		record, err := csvReader.Read()
		if err == io.EOF {
			// Reached the end of the file? Time to stop processing
			break
		} else if err != nil {
			return fmt.Errorf("unable to read CSV record: %w", err)
		}

		if err := collector.Collect(record); err != nil {
			return err
		}
	}

	return nil
}
