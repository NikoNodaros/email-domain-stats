package main

import (
	"fmt"
	"log"

	"emaildomainstats"
)

func main() {
	stats, err := emaildomainstats.GetEmailDomainStats("customer_data.csv")
	if err != nil {
		log.Fatal(err)
	}

	for _, domain := range stats {
		fmt.Printf("%s: %d\n", domain.Domain, domain.Count)
	}
}
