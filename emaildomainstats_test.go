package emaildomainstats_test

import (
	"bytes"
	"emaildomainstats"
	"os"
	"reflect"
	"testing"
)

func TestGetDomainStatsFromFile(t *testing.T) {
	// Define the path to the sample CSV file
	filePath := "./customer_data.csv"

	// Open the CSV file for reading
	file, err := os.Open(filePath)
	if err != nil {
		t.Errorf("Failed to open CSV file: %v", err)
	}
	defer file.Close()

	// Call the GetDomainStatsFromFile function with the CSV file
	stats, err := emaildomainstats.GetDomainStatsFromFile(file)
	if err != nil {
		t.Errorf("GetDomainStatsFromFile failed with error: %v", err)
	}

	// Define the expected output for the sample CSV file
	expectedStats := []emaildomainstats.DomainStats{
		{Domain: "example.com", Count: 3},
		{Domain: "example.net", Count: 2},
		{Domain: "example.org", Count: 2},
	}

	// Compare the expected output with the actual output
	if !reflect.DeepEqual(stats, expectedStats) {
		t.Errorf("GetDomainStatsFromFile failed: expected %v, but got %v", expectedStats, stats)
	}
}

func TestGetDomainStatsFromFile(t *testing.T) {
	// Define test cases for different input scenarios and edge cases
	testCases := []struct {
		name          string
		input         string
		expectedStats []emaildomainstats.DomainStats
		expectedError error
	}{
		{
			name:          "Empty file",
			input:         "",
			expectedStats: []emaildomainstats.DomainStats{},
			expectedError: nil,
		},
		{
			name: "Invalid CSV format",
			input: "first_name,last_name,email,gender,ip_address\n" +
				"John,Doe,john.doe@example.com,Male,\n" +
				"Jane,Doe,jane.doe@example.com,Female,\n" +
				"James,Smith,james.smith@example.org,Male,\n" +
				"Jessica,Baker,jessica.baker@example.org,Female,\n" +
				"Jacob,Jones,jacob.jones@example.net,Male,\n" +
				"Julie,White,julie.white@example.net,Female,\n" +
				"Jack,Anderson,jack.anderson@example.com,Male,\n" +
				"Kate,Williams,kate.williams@example.co.uk,Female,\n",
			expectedStats: nil,
			expectedError: emaildomainstats.ErrInvalidCSVFormat,
		},
		{
			name: "Valid CSV file",
			input: "first_name,last_name,email,gender,ip_address\n" +
				"John,Doe,john.doe@example.com,Male,127.0.0.1\n" +
				"Jane,Doe,jane.doe@example.com,Female,127.0.0.2\n" +
				"James,Smith,james.smith@example.org,Male,127.0.0.3\n" +
				"Jessica,Baker,jessica.baker@example.org,Female,127.0.0.4\n" +
				"Jacob,Jones,jacob.jones@example.net,Male,127.0.0.5\n" +
				"Julie,White,julie.white@example.net,Female,127.0.0.6\n" +
				"Jack,Anderson,jack.anderson@example.com,Male,127.0.0.7\n",
			expectedStats: []emaildomainstats.DomainStats{
				{Domain: "example.com", Count: 2},
				{Domain: "example.net", Count: 2},
				{Domain: "example.org", Count: 2},
			},
			expectedError: nil,
		},
		{
			name: "CSV file with duplicate emails",
			input: "first_name,last_name,email,gender,ip_address\n" +
				"John,Doe,john.doe@example.com,Male,127.0.0.1\n" +
				"Jane,Doe,jane.doe@example.com,Female,127.0.0.2\n" +
				"James,Smith,james.smith@example.org,Male,127.0.0.3\n" +
				"Jessica,Baker,jessica.baker@example.org,Female,127.0.0.4\n" +
				"Jacob,Jones,jacob.jones@example.net,Male,127.0.0.5\n" +
				"Julie,White,julie.white@example.net,Female,127.0.0.6\n" +
				"Jack,Anderson,jack.anderson@example.com,Male,127.0.0.7\n" +
				"Kate,Williams,kate.williams@example.net,Female,127.0.0.8\n" +
				"John,Smith,john.smith@example.com,Male,127.0.0.9\n" +
				"Jane,Williams,jane.williams@example.com,Female,127.0.0.10\n" +
				"James,Johnson,james.johnson@example.org,Male,127.0.0.11\n" +
				"Jessica,Brown,jessica.brown@example.org,Female,127.0.0.12\n",
			expectedStats: []emaildomainstats.DomainStats{
				{Domain: "example.com", Count: 2},
				{Domain: "example.net", Count: 3},
				{Domain: "example.org", Count: 2},
			},
			expectedError: nil,
		},
		{
			name:          "CSV file with only header",
			input:         "first_name,last_name,email,gender,ip_address\n",
			expectedStats: []emaildomainstats.DomainStats{},
			expectedError: nil,
		},
		{
			name: "CSV file with only one row",
			input: "first_name,last_name,email,gender,ip_address\n" +
				"John,Doe,john.doe@example.com,Male,127.0.0.1\n",
			expectedStats: []emaildomainstats.DomainStats{
				{Domain: "example.com", Count: 1},
			},
			expectedError: nil,
		},
	}

	// Iterate over the test cases and run the tests
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			// Convert the input string to a byte buffer
			buffer := bytes.NewBufferString(testCase.input)

			// Call the GetDomainStatsFromFile function with the byte buffer
			stats, err := emaildomainstats.GetDomainStatsFromFile(buffer)

			// Compare the expected output with the actual output
			if !reflect.DeepEqual(stats, testCase.expectedStats) {
				t.Errorf("GetDomainStatsFromFile failed: expected %v, but got %v", testCase.expectedStats, stats)
			}

			// Compare the expected error with the actual error
			if err != testCase.expectedError {
				t.Errorf("GetDomainStatsFromFile failed: expected error %v, but got error %v", testCase.expectedError, err)
			}
		})
	}
}
