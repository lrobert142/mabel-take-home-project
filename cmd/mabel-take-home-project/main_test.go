//go:build integration
// +build integration

package main

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"mabel-take-home-project/internal/model/collector"
	"testing"
)

var mockCollectorCollectError = errors.New("manually thrown Collect error")

type MockCollector struct {
	mustError bool
}

func (c *MockCollector) Collect(record []string) error {
	if c.mustError {
		return mockCollectorCollectError
	}
	return nil
}

func TestLoadFromCsvFile_Accounts(t *testing.T) {
	tests := map[string]struct {
		filePath      string
		collector     collector.Collector
		errorContains string
	}{
		"Collect the data without issue when the file can be read and no issues occur": {
			filePath:      "testdata/valid_data.csv",
			collector:     &MockCollector{},
			errorContains: "",
		},
		"Return an error when the file is not found": {
			filePath:      "testdata/404.csv",
			collector:     &MockCollector{},
			errorContains: "unable to read input file",
		},
		"Return an error when there is an issue parsing the file": {
			filePath:      "testdata/malformed_data.csv",
			collector:     &MockCollector{},
			errorContains: "unable to read CSV record",
		},
		"Return an error when the collector returns an error": {
			filePath: "testdata/valid_data.csv",
			collector: &MockCollector{
				mustError: true,
			},
			errorContains: mockCollectorCollectError.Error(),
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			err := loadFromCsvFile(test.filePath, test.collector)
			if test.errorContains == "" {
				assert.NoError(t, err)
			} else {
				assert.ErrorContains(t, err, test.errorContains)
			}
		})
	}
}
