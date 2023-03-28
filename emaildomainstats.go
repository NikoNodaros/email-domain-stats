package main

import (
	"encoding/csv"
	"io"
	"os"
	"sort"
	"strings"
)

type DomainStats struct {
	Domain string
	Count  int
}

func GetEmailDomainStats(filename string) ([]DomainStats, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	stats := make([]DomainStats, 0)

	// Read each record from CSV file and update stats slice
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}

		// Extract domain from email address
		email := record[2]
		parts := strings.Split(email, "@")
		if len(parts) != 2 {
			continue
		}
		domain := parts[1]

		// Update stats slice with domain count
		found := false
		for i := range stats {
			if stats[i].Domain == domain {
				stats[i].Count++
				found = true
				break
			}
		}
		if !found {
			stats = append(stats, DomainStats{Domain: domain, Count: 1})
		}
	}

	// Sort stats slice by domain name
	sort.Slice(stats, func(i, j int) bool {
		return stats[i].Domain < stats[j].Domain
	})

	return stats, nil
}
