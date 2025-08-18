# Adaptive Kyng-Dinic’s Maximum Flow for Large Graphs in Go

> **A scalable, concurrent, and memory-efficient maximum flow algorithm for massive graphs (up to 100 million vertices), written in Go. Features adaptive algorithm selection, parallel/concurrent DFS, and O(E) memory. Inspired by Kyng et al. (STOC 2024).**

[![Go Version](https://img.shields.io/badge/Go-1.24+-00ADD8?style=for-the-badge&logo=go)](https://go.dev/)
[![Scale](https://img.shields.io/badge/Scalable-100M_Vertices-success?style=for-the-badge)](https://github.com/wllclngn/Tests)
[![Algorithm](https://img.shields.io/badge/Algorithm-Kyng--Dinic-blue?style=for-the-badge)](https://arxiv.org/abs/2203.00671)
[![License](https://img.shields.io/badge/License-MIT-yellow?style=for-the-badge)](LICENSE)

---

## Overview

**Adaptive Kyng-Dinic’s** is a Go library for maximum flow in extremely large graphs. It is designed for practical, large-scale network analysis, using adaptive selection among Kyng-Dinic, Dinic's, Push-Relabel, ISAP, and unit-capacity optimizations. The codebase is engineered for O(E) memory, robust concurrency, and resilience against stack overflows.

---

## Why Use This Library?

- **Handles massive graphs:** 100M+ vertices processed in 32.7 seconds, 200M+ edges with 33.6GB RAM.
- **Adaptive runtime:** Picks the optimal algorithm for your graph’s structure.
- **Hybrid parallel/iterative DFS:** Fast, robust, never stack overflows.
- **O(E) memory:** Efficient even on commodity hardware.
- **Real benchmarks:** Transparent, reproducible results.
- **Go-native:** Simple API, idiomatic, no dependencies.

---

## Quick Start

```bash
go get github.com/wllclngn/Tests
```

```go
import "github.com/wllclngn/Tests"

g := kyng.NewAdaptiveGraph(100_000_000)
for i := 0; i < 99_999_999; i++ {
    for j := i+1; j < min(i+3, 100_000_000); j++ {
        g.AddEdge(i, j, rand.Intn(100)+1)
    }
}
flow := g.MaxFlow(0, 99_999_999)
fmt.Println("Max flow:", flow) // Completes in 32.7 seconds
g.PrintStatistics() // Shows hybrid DFS breakdown and performance metrics
```
_Automatically selects Kyng-Dinic's algorithm, outputs detailed performance analytics._

---

## Key Features

- **Automatic algorithm selection:** Based on density, degree, and capacity patterns.
- **Concurrent DFS:** Parallel search with automatic fallback for deep recursions.
- **Memory-optimized:** 16-byte edge records, O(E) scaling.
- **Concurrency-safe:** Adaptive semaphore sizing, no goroutine leaks.
- **Transparent analytics:** Reports throughput, memory, concurrency ratio, and scaling.

---

## Benchmarks (Go 1.24)

| Vertices      | Edges        | Time      | Throughput   | Memory Usage |
|---------------|--------------|-----------|--------------|-------------|
| 50,000,000    | 100,000,000  | 23.8s     | 2.10M v/s    | 17.2GB      |
| 100,000,000   | 200,000,000  | **32.7s** | **3.06M v/s** | **33.6GB**  |

### **Key Performance Achievements:**
- 🚀 **3.06 million vertices/second** throughput on 100M scale
- ⚡ **32.7 seconds** for 100 million vertex graphs (vs 9m35s projected)
- 🧠 **Hybrid DFS**: 99.99% concurrent, 0.01% iterative paths
- 💾 **Linear memory scaling**: 33.6GB for 200M edges (168 bytes/edge)
- 🎯 **Perfect algorithm selection**: Automatically chose Kyng-Dinic's for sparse graphs
- 🛡️ **Zero stack overflows** with adaptive depth switching at 5000 levels

---

## When to Use

- Large-scale flow/network analysis, graph mining, or infrastructure modeling.
- Need for reliability and scale beyond standard libraries.
- Go environments where performance and memory efficiency matter.

---

## Limitations

- Not designed for tiny graphs—stdlib or simpler libs will be faster for <10K nodes.
- No support for dynamic graph modification after construction.
- API surface is intentionally minimal and focused.

---

## Research & Credits

- **Theory:** [Kyng et al., STOC 2024](https://arxiv.org/abs/2203.00671)
- **Engineering/code:** Will Clingan ([GitHub @wllclngn](https://github.com/wllclngn))
- **Inspiration:** Dragonbox, ETH Zurich, Go community.

---

## License

MIT

---

**Keywords:**  
maximum flow Go, Kyng-Dinic Go, large graph max flow, concurrent max flow, scalable graph algorithm, O(m·polylog(n)), Go network flow, production graph processing

---

*For full docs, stress tests, and source, see the repository.*