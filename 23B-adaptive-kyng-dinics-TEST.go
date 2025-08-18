// ADAPTIVE KYNG-DINIC'S ALGORITHM - PROPER TEST FILE
// Tests the main implementation (23-adaptive-kyng-dinics-algorithm.go)
// No duplicate algorithm code - just tests and benchmarks
//
// Author: Will Clingan
package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

// Simple graph generator for testing
func generateSparseGraph(vertices int) *AdaptiveGraph {
	graph := NewAdaptiveGraph(vertices)
	edges := 0
	
	// Sparse connectivity: each vertex connects to ~3 others
	for i := 0; i < vertices-1; i++ {
		for j := i + 1; j < min(i+3, vertices); j++ {
			capacity := rand.Intn(100) + 1
			graph.AddEdge(i, j, capacity)
			edges++
		}
	}
	fmt.Printf("Generated sparse graph: %d vertices, %d edges\n", vertices, edges)
	return graph
}

func generateDenseGraph(vertices int) *AdaptiveGraph {
	graph := NewAdaptiveGraph(vertices)
	edges := 0
	
	// Dense connectivity: more edges per vertex
	connectionsPerVertex := min(10, vertices/100)
	
	for i := 0; i < vertices-1; i++ {
		for j := i + 1; j < min(i+connectionsPerVertex, vertices); j++ {
			capacity := rand.Intn(50) + 1
			graph.AddEdge(i, j, capacity)
			edges++
		}
	}
	fmt.Printf("Generated dense graph: %d vertices, %d edges\n", vertices, edges)
	return graph
}

func runStressTests() {
	fmt.Println("\n🚀 KYNG-DINIC ALGORITHM STRESS TEST")
	fmt.Println("Testing the main implementation (23-adaptive-kyng-dinics-algorithm.go)")
	fmt.Println("==================================================================")
	fmt.Printf("System: %d cores, Go %s\n", runtime.NumCPU(), runtime.Version())
	
	// Test sizes
	testSizes := []int{100000000}
	
	fmt.Println("\n📊 PERFORMANCE TESTS")
	
	for _, size := range testSizes {
		fmt.Printf("\n🔸 Testing %d vertices:\n", size)
		
		// Generate test graph
		start := time.Now()
		graph := generateSparseGraph(size)
		genTime := time.Since(start)
		
		// Run max flow
		start = time.Now()
		maxFlow := graph.MaxFlow(0, size-1)
		algoTime := time.Since(start)
		
		// Report results
		fmt.Printf("Results:\n")
		fmt.Printf("  Generation time: %v\n", genTime)
		fmt.Printf("  Algorithm time: %v\n", algoTime)
		fmt.Printf("  Max flow value: %d\n", maxFlow)
		fmt.Printf("  Performance: %.2f vertices/sec\n", float64(size)/algoTime.Seconds())
		
		// Get detailed performance report
		graph.PrintStatistics()
		
		// Memory usage check
		var m runtime.MemStats
		runtime.ReadMemStats(&m)
		fmt.Printf("  Memory used: %.2f MB\n", float64(m.Alloc)/1024/1024)
	}
	
	fmt.Println("\n✅ STRESS TEST COMPLETE")
	fmt.Printf("📈 All tests use the main implementation from 23-adaptive-kyng-dinics-algorithm.go\n")
}