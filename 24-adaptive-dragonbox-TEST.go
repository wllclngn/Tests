package main

import (
	"fmt"
	"math"
	"testing"
	"time"
)

// Test helper - uses the unified Dragonbox
var testDB = NewUnifiedDragonbox()

func ConvertSingle(f float64) string {
	return testDB.Convert(f)
}

// ============================================================================
// CORRECTNESS TESTS
// ============================================================================

func TestSpecialValues(t *testing.T) {
	tests := []struct {
		name     string
		input    float64
		expected string
	}{
		{"Zero", 0.0, "0"},
		{"Negative Zero", math.Copysign(0, -1), "-0"},
		{"Positive Infinity", math.Inf(1), "+Inf"},
		{"Negative Infinity", math.Inf(-1), "-Inf"},
		{"NaN", math.NaN(), "NaN"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ConvertSingle(tt.input)
			if result != tt.expected {
				t.Errorf("ConvertSingle(%v) = %s, want %s", tt.input, result, tt.expected)
			}
		})
	}
}

func TestCommonValues(t *testing.T) {
	tests := []struct {
		input    float64
		expected string
	}{
		{0.5, "0.5"},
		{0.25, "0.25"},
		{0.2, "0.2"},
		{0.1, "0.1"},
		{0.01, "0.01"},
		{0.001, "0.001"},
		{-0.5, "-0.5"},
		{-0.25, "-0.25"},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("Common_%v", tt.input), func(t *testing.T) {
			result := ConvertSingle(tt.input)
			if result != tt.expected {
				t.Errorf("ConvertSingle(%v) = %s, want %s", tt.input, result, tt.expected)
			}
		})
	}
}

func TestIntegers(t *testing.T) {
	tests := []struct {
		input    float64
		expected string
	}{
		{1.0, "1"},
		{10.0, "10"},
		{100.0, "100"},
		{1000.0, "1000"},
		{-1.0, "-1"},
		{-10.0, "-10"},
		{0.0, "0"},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("Integer_%v", tt.input), func(t *testing.T) {
			result := ConvertSingle(tt.input)
			if result != tt.expected {
				t.Errorf("ConvertSingle(%v) = %s, want %s", tt.input, result, tt.expected)
			}
		})
	}
}

func TestBatchConversion(t *testing.T) {
	db := NewUnifiedDragonbox()
	testValues := []float64{1.0, 2.0, 3.0, 0.5, 0.25}
	
	results := db.BatchConvert(testValues)
	
	if len(results) != len(testValues) {
		t.Errorf("BatchConvert returned %d results, expected %d", len(results), len(testValues))
	}
	
	// Check individual results
	expected := []string{"1", "2", "3", "0.5", "0.25"}
	for i, result := range results {
		if result != expected[i] {
			t.Errorf("BatchConvert[%d] = %s, want %s", i, result, expected[i])
		}
	}
}

func TestPatternDetection(t *testing.T) {
	db := NewUnifiedDragonbox()
	
	// Test pattern detection by running conversions and checking stats
	testData := []float64{
		0.0, math.Inf(1), math.NaN(),  // Special values
		1.0, 10.0, 100.0,              // Integers
		0.5, 0.25, 0.1,                // Simple decimals
		1e-10, 1e20,                   // Scientific
		math.Pi, math.E,               // Complex
	}
	
	for _, f := range testData {
		_ = db.Convert(f)
	}
	
	// Check that patterns were detected
	totalPatterns := uint64(0)
	for _, count := range db.patternStats {
		totalPatterns += count
	}
	
	if totalPatterns == 0 {
		t.Error("No patterns were detected")
	}
	
	// Should have detected at least some special values, integers, and simple decimals
	if db.patternStats[PatternSpecialValue] == 0 {
		t.Error("Should have detected special values")
	}
	if db.patternStats[PatternInteger] == 0 {
		t.Error("Should have detected integers")
	}
	if db.patternStats[PatternSimpleDecimal] == 0 {
		t.Error("Should have detected simple decimals")
	}
}

func TestRangeSelection(t *testing.T) {
	db := NewUnifiedDragonbox()
	
	// Test range selection by using values that should hit different ranges
	testData := []float64{
		1.0,      // Should use compact range
		1e50,     // Should use medium range
		1e200,    // Should use full range
	}
	
	for _, f := range testData {
		_ = db.Convert(f)
	}
	
	// Check that different ranges were used
	totalRanges := uint64(0)
	for _, count := range db.rangeStats {
		totalRanges += count
	}
	
	if totalRanges == 0 {
		t.Error("No range selection occurred")
	}
}

func TestCaching(t *testing.T) {
	db := NewUnifiedDragonbox()
	
	// Convert same value multiple times
	testValue := 3.14159
	
	// First conversion - should miss cache
	result1 := db.Convert(testValue)
	
	// Second conversion - should hit cache
	result2 := db.Convert(testValue)
	
	if result1 != result2 {
		t.Errorf("Cached results should match: %s != %s", result1, result2)
	}
	
	if db.cacheHits == 0 {
		t.Error("Should have at least one cache hit")
	}
}

func BenchmarkDragonboxConversion(b *testing.B) {
	db := NewUnifiedDragonbox()
	testValues := []float64{
		1.0, 3.14159, 0.5, 1e10, 1e-10, math.Pi, math.E,
	}
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, v := range testValues {
			_ = db.Convert(v)
		}
	}
}

func BenchmarkBatchConversion(b *testing.B) {
	db := NewUnifiedDragonbox()
	testValues := make([]float64, 1000)
	for i := range testValues {
		testValues[i] = float64(i) * 3.14159
	}
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = db.BatchConvert(testValues)
	}
}

// Run comprehensive test showing all adaptive features
func TestComprehensiveAdaptiveFeatures(t *testing.T) {
	db := NewUnifiedDragonbox()
	
	fmt.Println("\n=== COMPREHENSIVE ADAPTIVE DRAGONBOX TEST ===")
	
	// Test different pattern types
	testCases := map[string][]float64{
		"Special Values": {0.0, math.Inf(1), math.Inf(-1), math.NaN()},
		"Integers":      {1.0, 10.0, 100.0, 1000.0, -42.0},
		"Simple Decimals": {0.5, 0.25, 0.1, 0.01, -0.5},
		"Scientific":    {1e-10, 1e20, 1e-100, 1e50},
		"Complex":       {math.Pi, math.E, math.Sqrt(2), math.Log(10)},
	}
	
	start := time.Now()
	totalConversions := 0
	
	for category, values := range testCases {
		fmt.Printf("\nTesting %s:\n", category)
		for _, v := range values {
			result := db.Convert(v)
			fmt.Printf("  %v -> %s\n", v, result)
			totalConversions++
		}
	}
	
	duration := time.Since(start)
	
	fmt.Printf("\n=== PERFORMANCE SUMMARY ===\n")
	fmt.Printf("Total conversions: %d\n", totalConversions)
	fmt.Printf("Total time: %v\n", duration)
	fmt.Printf("Average per conversion: %v\n", duration/time.Duration(totalConversions))
	
	fmt.Printf("\n%s\n", db.GetPerformanceReport())
	
	// Verify adaptive features worked
	patternsUsed := 0
	for _, count := range db.patternStats {
		if count > 0 {
			patternsUsed++
		}
	}
	
	rangesUsed := 0
	for _, count := range db.rangeStats {
		if count > 0 {
			rangesUsed++
		}
	}
	
	if patternsUsed < 3 {
		t.Errorf("Should use at least 3 different patterns, used %d", patternsUsed)
	}
	
	if rangesUsed < 2 {
		t.Errorf("Should use at least 2 different ranges, used %d", rangesUsed)
	}
	
	fmt.Printf("✅ ADAPTIVE SUCCESS: Used %d patterns and %d ranges\n", patternsUsed, rangesUsed)
	fmt.Printf("✅ CACHE PERFORMANCE: %.1f%% hit rate\n", 
		float64(db.cacheHits)*100/float64(db.cacheHits+db.cacheMisses+1))
}