// Unified Concurrent DFS - Progressive Enhancement Architecture
// Combines proven semaphore pattern with advanced adaptive strategies
// 
// Modes:
//   Simple: Classic semaphore-based concurrent DFS (proven, reliable)
//   Advanced: Multi-strategy adaptive DFS with optimization
//   Auto: Automatically selects mode based on tree characteristics
//
// Author: Will Clingan
package main

import (
	"fmt"
	"math/rand"
	"os"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

// DFSMode determines the traversal complexity level
type DFSMode int

const (
	ModeSimple DFSMode = iota  // Classic semaphore pattern
	ModeAdvanced               // Multi-strategy with optimization
	ModeAuto                   // Automatic mode selection
)

// TraversalStrategy determines the order of node exploration (Advanced mode)
type TraversalStrategy int

const (
	StrategyDepthFirst TraversalStrategy = iota
	StrategyBreadthFirst
	StrategyRandom
	StrategyWorkStealing
	StrategyAdaptive
)

// Tree structure with unified design supporting both modes
type Tree struct {
	root *Node
	mode DFSMode
	
	// Advanced mode features
	nodeCount int64
	maxDepth  int32
	strategy  TraversalStrategy
	
	// Performance metrics
	executionTime time.Duration
	goroutinesUsed int64
}

// Node with progressive enhancement - supports both simple and advanced modes
type Node struct {
	key   int
	left  *Node
	right *Node
	
	// Advanced mode metadata (unused in simple mode)
	depth       int32
	subtreeSize int32
	visited     atomic.Bool
}

// WorkQueue for advanced work-stealing strategy
type WorkQueue struct {
	nodes    []*Node
	mu       sync.Mutex
	capacity int
}

// Global work queue for stealing (advanced mode only)
var globalQueue = &WorkQueue{
	nodes:    make([]*Node, 0, 1000),
	capacity: 1000,
}

// ============================================================================
// TREE CONSTRUCTION (Unified for both modes)
// ============================================================================

// NewTree creates a tree with specified DFS mode
func NewTree(mode DFSMode) *Tree {
	return &Tree{
		mode:     mode,
		strategy: StrategyAdaptive, // Default for advanced mode
	}
}

// Insert adds a node using the original working logic
func (tree *Tree) Insert(data int) {
	if tree.root == nil {
		tree.root = &Node{key: data, depth: 0}
	} else {
		tree.root.insert(data, 0)
	}
	atomic.AddInt64(&tree.nodeCount, 1)
}

// insert implements the original proven insertion logic with optional metadata
func (node *Node) insert(data int, depth int32) {
	if node.left == nil || node.right == nil {
		newNode := &Node{key: data, depth: depth + 1}
		if node.left == nil {
			node.left = newNode
		} else {
			node.right = newNode
		}
		
		// Update advanced mode metadata
		node.subtreeSize++
	} else {
		// Use original branching logic that works
		if (data % 2) == 0 {
			if ((node.left.key % 2) == 0) && node.left.left == nil {
				node.left.left = &Node{key: data, depth: depth + 1}
			} else if ((node.left.key % 2) == 0) && node.left.right == nil {
				node.left.right = &Node{key: data, depth: depth + 1}
			} else {
				node.right.insert(data, depth+1)
			}
		} else {
			if ((node.right.key % 2) != 0) && node.right.left == nil {
				node.right.left = &Node{key: data, depth: depth + 1}
			} else if ((node.right.key % 2) != 0) && node.right.right == nil {
				node.right.right = &Node{key: data, depth: depth + 1}
			} else {
				node.left.insert(data, depth+1)
			}
		}
		node.subtreeSize++
	}
}

// ============================================================================
// MODE SELECTION AND TRAVERSAL DISPATCH
// ============================================================================

// TraverseConcurrent is the unified entry point for all modes
func (tree *Tree) TraverseConcurrent() {
	mode := tree.selectMode()
	
	start := time.Now()
	var wg sync.WaitGroup
	
	// Create semaphore sized appropriately for mode
	semaphoreSize := tree.calculateSemaphoreSize(mode)
	semaphore := make(chan struct{}, semaphoreSize)
	
	// Reset visited flags for advanced mode
	if mode != ModeSimple {
		tree.resetVisited(tree.root)
	}
	
	wg.Add(1)
	
	switch mode {
	case ModeSimple:
		go tree.root.simpleTraversal(&wg, semaphore)
	case ModeAdvanced, ModeAuto:
		go tree.root.advancedTraversal(&wg, semaphore, tree.strategy)
	}
	
	wg.Wait()
	tree.executionTime = time.Since(start)
}

// selectMode chooses optimal mode based on tree characteristics
func (tree *Tree) selectMode() DFSMode {
	if tree.mode != ModeAuto {
		return tree.mode
	}
	
	// Auto-selection logic
	nodeCount := atomic.LoadInt64(&tree.nodeCount)
	
	if nodeCount < 100 {
		return ModeSimple // Small trees: simple is best
	} else if nodeCount < 10000 {
		return ModeAdvanced // Medium trees: advanced optimization
	} else {
		return ModeAdvanced // Large trees: definitely need optimization
	}
}

// calculateSemaphoreSize determines optimal goroutine limit with safety bounds
func (tree *Tree) calculateSemaphoreSize(mode DFSMode) int {
	maxCPUs := runtime.NumCPU()
	nodeCount := atomic.LoadInt64(&tree.nodeCount)
	
	var size int
	switch mode {
	case ModeSimple:
		// Conservative: proven value with safety limit
		size = min(20, maxCPUs*2)
	case ModeAdvanced:
		// Adaptive based on tree size, but bounded
		if nodeCount < 100 {
			size = min(10, maxCPUs)
		} else if nodeCount < 1000 {
			size = min(maxCPUs*2, 32)
		} else {
			size = min(maxCPUs*3, 64) // Cap at 64 even for huge trees
		}
	case ModeAuto:
		// Smart selection with conservative bounds
		if nodeCount < 100 {
			size = min(8, maxCPUs)
		} else if nodeCount < 1000 {
			size = min(16, maxCPUs*2)
		} else {
			size = min(32, maxCPUs*4)
		}
	}
	
	// Absolute safety: never exceed reasonable limits
	return max(1, min(size, 128))
}

// ============================================================================
// SIMPLE MODE - Original Proven Semaphore Pattern
// ============================================================================

// simpleTraversal implements the original working concurrent DFS with controlled acquisition
func (node *Node) simpleTraversal(wg *sync.WaitGroup, semaphore chan struct{}) {
	defer wg.Done()

	if node == nil {
		return
	}

	// Process current node
	fmt.Printf("Visiting Node: %d\n", node.key)

	// Launch recursive goroutines with controlled semaphore acquisition
	if node.left != nil {
		wg.Add(1)
		
		// Try to acquire semaphore with timeout control
		select {
		case semaphore <- struct{}{}:
			// Got semaphore: launch goroutine
			go func(left *Node) {
				defer func() { <-semaphore }() // Release semaphore slot
				left.simpleTraversal(wg, semaphore)
			}(node.left)
		default:
			// Semaphore full: execute synchronously (graceful degradation)
			node.left.simpleTraversal(wg, semaphore)
		}
	}

	if node.right != nil {
		wg.Add(1)
		
		// Try to acquire semaphore with timeout control
		select {
		case semaphore <- struct{}{}:
			// Got semaphore: launch goroutine
			go func(right *Node) {
				defer func() { <-semaphore }() // Release semaphore slot
				right.simpleTraversal(wg, semaphore)
			}(node.right)
		default:
			// Semaphore full: execute synchronously (graceful degradation)
			node.right.simpleTraversal(wg, semaphore)
		}
	}
}

// ============================================================================
// ADVANCED MODE - Multi-Strategy Adaptive System
// ============================================================================

// advancedTraversal implements sophisticated adaptive strategies
func (node *Node) advancedTraversal(wg *sync.WaitGroup, semaphore chan struct{}, strategy TraversalStrategy) {
	defer wg.Done()
	
	if node == nil || node.visited.Load() {
		return
	}
	
	// Mark as visited atomically (prevents duplicate work)
	if !node.visited.CompareAndSwap(false, true) {
		return
	}
	
	// Process current node
	fmt.Printf("Visiting Node: %d (Strategy: %s)\n", node.key, strategyName(strategy))
	
	// Select traversal approach based on strategy
	switch strategy {
	case StrategyDepthFirst:
		node.depthFirstTraversal(wg, semaphore)
	case StrategyBreadthFirst:
		node.breadthFirstTraversal(wg, semaphore)
	case StrategyRandom:
		node.randomTraversal(wg, semaphore)
	case StrategyWorkStealing:
		node.workStealingTraversal(wg, semaphore)
	case StrategyAdaptive:
		node.adaptiveTraversal(wg, semaphore)
	}
}

// depthFirstTraversal with path optimization
func (node *Node) depthFirstTraversal(wg *sync.WaitGroup, semaphore chan struct{}) {
	// Optimize path based on subtree characteristics
	leftFirst := node.shouldGoLeftFirst()
	
	first, second := node.left, node.right
	if !leftFirst {
		first, second = second, first
	}
	
	// Process first child
	if first != nil {
		wg.Add(1)
		select {
		case semaphore <- struct{}{}:
			// Non-blocking: launch goroutine
			go func() {
				defer func() { <-semaphore }()
				first.advancedTraversal(wg, semaphore, StrategyDepthFirst)
			}()
		default:
			// Semaphore full: process synchronously (graceful degradation)
			first.advancedTraversal(wg, semaphore, StrategyDepthFirst)
		}
	}
	
	// Process second child
	if second != nil {
		wg.Add(1)
		select {
		case semaphore <- struct{}{}:
			go func() {
				defer func() { <-semaphore }()
				second.advancedTraversal(wg, semaphore, StrategyDepthFirst)
			}()
		default:
			second.advancedTraversal(wg, semaphore, StrategyDepthFirst)
		}
	}
}

// breadthFirstTraversal prioritizes width over depth
func (node *Node) breadthFirstTraversal(wg *sync.WaitGroup, semaphore chan struct{}) {
	// Process children at current level first
	children := []*Node{}
	if node.left != nil {
		children = append(children, node.left)
	}
	if node.right != nil {
		children = append(children, node.right)
	}
	
	// Launch all children concurrently (breadth-first characteristic)
	for _, child := range children {
		wg.Add(1)
		select {
		case semaphore <- struct{}{}:
			go func(c *Node) {
				defer func() { <-semaphore }()
				c.advancedTraversal(wg, semaphore, StrategyBreadthFirst)
			}(child)
		default:
			child.advancedTraversal(wg, semaphore, StrategyBreadthFirst)
		}
	}
}

// randomTraversal introduces controlled chaos for load balancing
func (node *Node) randomTraversal(wg *sync.WaitGroup, semaphore chan struct{}) {
	children := []*Node{}
	if node.left != nil {
		children = append(children, node.left)
	}
	if node.right != nil {
		children = append(children, node.right)
	}
	
	// Shuffle children randomly
	for i := len(children) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		children[i], children[j] = children[j], children[i]
	}
	
	// Process in random order
	for _, child := range children {
		wg.Add(1)
		select {
		case semaphore <- struct{}{}:
			go func(c *Node) {
				defer func() { <-semaphore }()
				c.advancedTraversal(wg, semaphore, StrategyRandom)
			}(child)
		default:
			child.advancedTraversal(wg, semaphore, StrategyRandom)
		}
	}
}

// workStealingTraversal implements work-stealing for load balancing
func (node *Node) workStealingTraversal(wg *sync.WaitGroup, semaphore chan struct{}) {
	// Add children to global work queue
	if node.left != nil {
		globalQueue.Push(node.left)
	}
	if node.right != nil {
		globalQueue.Push(node.right)
	}
	
	// Try to steal and process work
	for i := 0; i < 3; i++ { // Limited attempts to prevent infinite loops
		if stolen := globalQueue.Steal(); stolen != nil {
			wg.Add(1)
			select {
			case semaphore <- struct{}{}:
				go func(n *Node) {
					defer func() { <-semaphore }()
					n.advancedTraversal(wg, semaphore, StrategyWorkStealing)
				}(stolen)
			default:
				stolen.advancedTraversal(wg, semaphore, StrategyWorkStealing)
			}
		} else {
			break
		}
	}
}

// adaptiveTraversal chooses optimal strategy per subtree
func (node *Node) adaptiveTraversal(wg *sync.WaitGroup, semaphore chan struct{}) {
	strategy := node.selectOptimalStrategy()
	
	switch strategy {
	case StrategyDepthFirst:
		node.depthFirstTraversal(wg, semaphore)
	case StrategyBreadthFirst:
		node.breadthFirstTraversal(wg, semaphore)
	case StrategyWorkStealing:
		node.workStealingTraversal(wg, semaphore)
	default:
		node.randomTraversal(wg, semaphore)
	}
}

// ============================================================================
// OPTIMIZATION HEURISTICS
// ============================================================================

// selectOptimalStrategy uses heuristics to choose best approach
func (node *Node) selectOptimalStrategy() TraversalStrategy {
	if node.subtreeSize < 10 {
		return StrategyDepthFirst // Small subtrees: simple DFS
	}
	
	if node.depth > 15 {
		return StrategyBreadthFirst // Deep trees: BFS for parallelism
	}
	
	// Calculate balance factor
	leftSize, rightSize := int32(0), int32(0)
	if node.left != nil {
		leftSize = node.left.subtreeSize
	}
	if node.right != nil {
		rightSize = node.right.subtreeSize
	}
	
	imbalance := abs(leftSize - rightSize)
	if imbalance > node.subtreeSize/3 {
		return StrategyWorkStealing // Imbalanced: work-stealing
	}
	
	return StrategyDepthFirst // Default to proven strategy
}

// shouldGoLeftFirst determines path optimization
func (node *Node) shouldGoLeftFirst() bool {
	if node.left == nil {
		return false
	}
	if node.right == nil {
		return true
	}
	
	// Prefer larger subtree first for better parallelism
	return node.left.subtreeSize >= node.right.subtreeSize
}

// ============================================================================
// UTILITY FUNCTIONS
// ============================================================================

// WorkQueue methods for work-stealing
func (wq *WorkQueue) Push(node *Node) bool {
	wq.mu.Lock()
	defer wq.mu.Unlock()
	
	if len(wq.nodes) < wq.capacity {
		wq.nodes = append(wq.nodes, node)
		return true
	}
	return false
}

func (wq *WorkQueue) Steal() *Node {
	wq.mu.Lock()
	defer wq.mu.Unlock()
	
	if len(wq.nodes) > 0 {
		node := wq.nodes[len(wq.nodes)-1]
		wq.nodes = wq.nodes[:len(wq.nodes)-1]
		return node
	}
	return nil
}

// resetVisited clears visited flags for advanced mode
func (tree *Tree) resetVisited(node *Node) {
	if node == nil {
		return
	}
	node.visited.Store(false)
	tree.resetVisited(node.left)
	tree.resetVisited(node.right)
}

// Helper functions
func abs(x int32) int32 {
	if x < 0 {
		return -x
	}
	return x
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func strategyName(s TraversalStrategy) string {
	names := []string{"DepthFirst", "BreadthFirst", "Random", "WorkStealing", "Adaptive"}
	if int(s) < len(names) {
		return names[s]
	}
	return "Unknown"
}

// ============================================================================
// DEMONSTRATION AND BENCHMARKING
// ============================================================================

func runDemoModes() {
	fmt.Println("UNIFIED CONCURRENT DFS - Progressive Enhancement")
	fmt.Println("=" + string(make([]byte, 50)))
	
	// Demonstrate all modes
	modes := []struct {
		mode DFSMode
		name string
		size int
	}{
		{ModeSimple, "Simple Mode (Proven Semaphore)", 50},
		{ModeAdvanced, "Advanced Mode (Multi-Strategy)", 100},
		{ModeAuto, "Auto Mode (Intelligent Selection)", 200},
	}
	
	for _, test := range modes {
		fmt.Printf("\n=== %s ===\n", test.name)
		
		tree := NewTree(test.mode)
		tree.strategy = StrategyAdaptive
		
		// Build test tree
		start := time.Now()
		for i := 0; i < test.size; i++ {
			tree.Insert(i)
		}
		buildTime := time.Since(start)
		
		// Execute traversal
		tree.TraverseConcurrent()
		
		// Report results
		fmt.Printf("Tree Size: %d nodes\n", tree.nodeCount)
		fmt.Printf("Build Time: %v\n", buildTime)
		fmt.Printf("Traversal Time: %v\n", tree.executionTime)
		fmt.Printf("Mode Selected: %s\n", modeName(tree.selectMode()))
	}
	
	fmt.Println("\n🎯 KEY ACHIEVEMENTS:")
	fmt.Println("✅ Preserved proven semaphore pattern")
	fmt.Println("✅ Added sophisticated adaptive strategies")
	fmt.Println("✅ Unified architecture supporting both approaches")
	fmt.Println("✅ Automatic mode selection based on tree characteristics")
	fmt.Println("✅ Graceful degradation under high load")
	fmt.Println("✅ Educational progression from simple → advanced")
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	
	// Check if we should run stress tests
	if len(os.Args) > 1 && os.Args[1] == "stress" {
		runAllTests()
	} else {
		runDemoModes()
	}
}

func modeName(m DFSMode) string {
	names := []string{"Simple", "Advanced", "Auto"}
	if int(m) < len(names) {
		return names[m]
	}
	return "Unknown"
}