package main

import (
	"encoding/json"
	"fmt"
	"os"

	"go.dedis.ch/kyber/v4/util/test"
)

var (
	outputFile = "../docs/benchmark-app/src/data/data.json"
	signatures = []string{"anon", "bls"}
)

// BenchmarkGroup runs benchmarks for the given group and writes the results to a JSON file.
func benchmarkGroup(name string, description string, gb *test.GroupBench) map[string]any {
	_ = "STUB: not implemented"
	return nil
}

// Scalar operations

// Point operations

// BenchmarkSign runs benchmarks for the some signature schemes.
func benchmarkSign(sigType string) map[string]any { _ = "STUB: not implemented"; return nil }

// Generate keys

// Signing

// Verification

// Key generation

// Signing

// Verification

func main() {
	// Write results to JSON file
	results := make(map[string]map[string]map[string]any)

	file, err := os.Create(outputFile)
	if err != nil {
		fmt.Println("Error creating output file:", err)
		return
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")

	// Run benchmarks for each group
	results["groups"] = make(map[string]map[string]any)
	for _, suite := range suites {
		groupBench := test.NewGroupBench(suite)
		result := benchmarkGroup(suite.String(), "Description", groupBench)
		results["groups"][suite.String()] = result
	}

	// Run benchmarks for signatures
	results["sign"] = make(map[string]map[string]any)
	for _, sigType := range signatures {
		result := benchmarkSign(sigType)
		results["sign"][sigType] = result
	}

	if err := encoder.Encode(results); err != nil {
		fmt.Println("Error encoding JSON:", err)
		return
	}
	fmt.Printf("Benchmark results written to %s\n", outputFile)
}
