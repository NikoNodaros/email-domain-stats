package main

import (
	"fmt"
	"log"
)

func main() {
	stats, err := GetEmailDomainStats("./customer_data.csv")
	if err != nil {
		log.Fatal(err)
	}

	for _, domain := range stats {
		fmt.Printf("%s: %d\n", domain.Domain, domain.Count)
	}
}
