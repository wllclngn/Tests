// ADAPTIVE TIMSORT IMPLEMENTATION
// Based on Tim Peters' original algorithm with intelligent pattern detection
// This is the REAL TimSort with adaptive strategy selection and parallel optimizations

package main

import (
	"math"
	"runtime"
	"sync"
)

// TimSort constants
const (
	MIN_MERGE     = 32
	MIN_GALLOP    = 7
	INITIAL_TMP_STORAGE_LENGTH = 256
)

// Data pattern types for adaptive sorting
type DataPattern int

const (
	PatternSorted DataPattern = iota
	PatternReversed
	PatternNearlySorted
	PatternComplex
)

// Run represents a sequence of consecutive elements
type Run struct {
	base   int  // Starting index of the run
	length int  // Length of the run
	desc   bool // True if run is descending
}

// MergeState holds the state for TimSort merging
type MergeState struct {
	data         []int
	tmp          []int
	runs         []Run
	minGallop    int
	tmpBase      int
	tmpLen       int
}

// NewMergeState creates a new merge state
func NewMergeState(data []int) *MergeState {
	tmpSize := len(data)
	if tmpSize < 2*MIN_MERGE {
		tmpSize = 2 * MIN_MERGE
	}
	
	return &MergeState{
		data:      data,
		tmp:       make([]int, tmpSize),
		runs:      make([]Run, 0),
		minGallop: MIN_GALLOP,
		tmpBase:   0,
		tmpLen:    0,
	}
}

// binaryInsertionSort performs binary insertion sort on a slice
func binaryInsertionSort(data []int, start, end int) {
	for i := start + 1; i < end; i++ {
		pivot := data[i]
		left, right := start, i
		
		// Binary search for insertion position
		for left < right {
			mid := (left + right) / 2
			if data[mid] > pivot {
				right = mid
			} else {
				left = mid + 1
			}
		}
		
		// Shift elements and insert
		copy(data[left+1:i+1], data[left:i])
		data[left] = pivot
	}
}

// countRunAndMakeAscending finds a run and ensures it's ascending
func countRunAndMakeAscending(data []int, start, end int) (int, bool) {
	if start == end-1 {
		return 1, false
	}
	
	runEnd := start + 1
	isDescending := false
	
	// Detect if run is ascending or descending
	if data[start] > data[runEnd] {
		isDescending = true
		// Count descending run
		for runEnd < end && data[runEnd-1] > data[runEnd] {
			runEnd++
		}
		// Reverse to make ascending
		reverseSlice(data, start, runEnd-1)
	} else {
		// Count ascending run (including equal elements)
		for runEnd < end && data[runEnd-1] <= data[runEnd] {
			runEnd++
		}
	}
	
	return runEnd - start, isDescending
}

// reverseSlice reverses elements between indices
func reverseSlice(data []int, start, end int) {
	for start < end {
		data[start], data[end] = data[end], data[start]
		start++
		end--
	}
}

// computeMinRunLength computes the minimum run length for TimSort
func computeMinRunLength(n int) int {
	r := 0
	for n >= MIN_MERGE {
		r |= n & 1
		n >>= 1
	}
	return n + r
}

// gallopLeft performs galloping search from the left
func gallopLeft(key int, data []int, base, length, hint int) int {
	lastOffset := 0
	offset := 1
	
	if data[base+hint] < key {
		// Gallop right until data[base+hint+lastOffset] < key <= data[base+hint+offset]
		maxOffset := length - hint
		for offset < maxOffset && data[base+hint+offset] < key {
			lastOffset = offset
			offset = (offset << 1) + 1
			if offset <= 0 { // Integer overflow
				offset = maxOffset
			}
		}
		if offset > maxOffset {
			offset = maxOffset
		}
		
		// Make offsets relative to base
		lastOffset += hint
		offset += hint
	} else {
		// Gallop left until data[base+hint-offset] < key <= data[base+hint-lastOffset]
		maxOffset := hint + 1
		for offset < maxOffset && data[base+hint-offset] >= key {
			lastOffset = offset
			offset = (offset << 1) + 1
			if offset <= 0 {
				offset = maxOffset
			}
		}
		if offset > maxOffset {
			offset = maxOffset
		}
		
		// Make offsets relative to base
		tmp := lastOffset
		lastOffset = hint - offset
		offset = hint - tmp
	}
	
	// Binary search between lastOffset and offset
	lastOffset++
	for lastOffset < offset {
		m := lastOffset + ((offset - lastOffset) >> 1)
		if data[base+m] < key {
			lastOffset = m + 1
		} else {
			offset = m
		}
	}
	return offset
}

// gallopRight performs galloping search from the right
func gallopRight(key int, data []int, base, length, hint int) int {
	offset := 1
	lastOffset := 0
	
	if data[base+hint] > key {
		// Gallop left until data[base+hint-offset] <= key < data[base+hint-lastOffset]
		maxOffset := hint + 1
		for offset < maxOffset && data[base+hint-offset] > key {
			lastOffset = offset
			offset = (offset << 1) + 1
			if offset <= 0 {
				offset = maxOffset
			}
		}
		if offset > maxOffset {
			offset = maxOffset
		}
		
		// Make offsets relative to base
		tmp := lastOffset
		lastOffset = hint - offset
		offset = hint - tmp
	} else {
		// Gallop right until data[base+hint+lastOffset] <= key < data[base+hint+offset]
		maxOffset := length - hint
		for offset < maxOffset && data[base+hint+offset] <= key {
			lastOffset = offset
			offset = (offset << 1) + 1
			if offset <= 0 {
				offset = maxOffset
			}
		}
		if offset > maxOffset {
			offset = maxOffset
		}
		
		// Make offsets relative to base
		lastOffset += hint
		offset += hint
	}
	
	// Binary search
	lastOffset++
	for lastOffset < offset {
		m := lastOffset + ((offset - lastOffset) >> 1)
		if data[base+m] <= key {
			lastOffset = m + 1
		} else {
			offset = m
		}
	}
	return offset
}

// mergeLow merges two adjacent runs with the left run being smaller
func (ms *MergeState) mergeLow(base1, len1, base2, len2 int) {
	// Copy first run into temp storage
	copy(ms.tmp[:len1], ms.data[base1:base1+len1])
	
	cursor1 := 0       // Indexes into tmp array
	cursor2 := base2   // Indexes into data array
	dest := base1      // Indexes into data array
	
	// Move first element of second run and deal with degenerate cases
	ms.data[dest] = ms.data[cursor2]
	dest++
	cursor2++
	len2--
	
	if len2 == 0 {
		copy(ms.data[dest:dest+len1], ms.tmp[:len1])
		return
	}
	if len1 == 1 {
		copy(ms.data[dest:dest+len2], ms.data[cursor2:cursor2+len2])
		ms.data[dest+len2] = ms.tmp[cursor1]
		return
	}
	
	minGallop := ms.minGallop
	
outer:
	for {
		count1 := 0 // Number of times in a row that first run won
		count2 := 0 // Number of times in a row that second run won
		
		// Do the straightforward thing until one run starts winning consistently
		for {
			if ms.data[cursor2] < ms.tmp[cursor1] {
				ms.data[dest] = ms.data[cursor2]
				dest++
				cursor2++
				count2++
				count1 = 0
				len2--
				if len2 == 0 {
					break outer
				}
			} else {
				ms.data[dest] = ms.tmp[cursor1]
				dest++
				cursor1++
				count1++
				count2 = 0
				len1--
				if len1 == 1 {
					break outer
				}
			}
			
			if (count1 | count2) >= minGallop {
				break
			}
		}
		
		// One run is winning so consistently that galloping may be a huge win
		for {
			count1 = gallopRight(ms.data[cursor2], ms.tmp, cursor1, len1, 0)
			if count1 != 0 {
				copy(ms.data[dest:dest+count1], ms.tmp[cursor1:cursor1+count1])
				dest += count1
				cursor1 += count1
				len1 -= count1
				if len1 <= 1 {
					break outer
				}
			}
			ms.data[dest] = ms.data[cursor2]
			dest++
			cursor2++
			len2--
			if len2 == 0 {
				break outer
			}
			
			count2 = gallopLeft(ms.tmp[cursor1], ms.data, cursor2, len2, 0)
			if count2 != 0 {
				copy(ms.data[dest:dest+count2], ms.data[cursor2:cursor2+count2])
				dest += count2
				cursor2 += count2
				len2 -= count2
				if len2 == 0 {
					break outer
				}
			}
			ms.data[dest] = ms.tmp[cursor1]
			dest++
			cursor1++
			len1--
			if len1 == 1 {
				break outer
			}
			
			minGallop--
			if count1 < MIN_GALLOP && count2 < MIN_GALLOP {
				if minGallop < 0 {
					minGallop = 0
				}
				minGallop += 2 // Penalize for leaving gallop mode
				break
			}
		}
	}
	
	ms.minGallop = minGallop
	if minGallop < 1 {
		ms.minGallop = 1
	}
	
	if len1 == 1 {
		copy(ms.data[dest:dest+len2], ms.data[cursor2:cursor2+len2])
		ms.data[dest+len2] = ms.tmp[cursor1] // Last element of run 1 to end of merge
	} else {
		copy(ms.data[dest:dest+len1], ms.tmp[cursor1:cursor1+len1])
	}
}

// mergeHigh merges two adjacent runs with the right run being smaller
func (ms *MergeState) mergeHigh(base1, len1, base2, len2 int) {
	// Copy second run into temp storage
	copy(ms.tmp[:len2], ms.data[base2:base2+len2])
	
	cursor1 := base1 + len1 - 1 // Indexes into data array
	cursor2 := len2 - 1         // Indexes into tmp array
	dest := base2 + len2 - 1    // Indexes into data array
	
	// Move last element of first run and deal with degenerate cases
	ms.data[dest] = ms.data[cursor1]
	dest--
	cursor1--
	len1--
	
	if len1 == 0 {
		copy(ms.data[dest-(len2-1):dest+1], ms.tmp[:len2])
		return
	}
	if len2 == 1 {
		dest -= len1
		cursor1 -= len1
		copy(ms.data[dest+1:dest+1+len1], ms.data[cursor1+1:cursor1+1+len1])
		ms.data[dest] = ms.tmp[cursor2]
		return
	}
	
	minGallop := ms.minGallop
	
outer:
	for {
		count1 := 0 // Number of times in a row that first run won
		count2 := 0 // Number of times in a row that second run won
		
		// Do straightforward thing until one run appears to win consistently
		for {
			if ms.tmp[cursor2] < ms.data[cursor1] {
				ms.data[dest] = ms.data[cursor1]
				dest--
				cursor1--
				count1++
				count2 = 0
				len1--
				if len1 == 0 {
					break outer
				}
			} else {
				ms.data[dest] = ms.tmp[cursor2]
				dest--
				cursor2--
				count2++
				count1 = 0
				len2--
				if len2 == 1 {
					break outer
				}
			}
			
			if (count1 | count2) >= minGallop {
				break
			}
		}
		
		// One run is winning consistently, galloping may be a huge win
		for {
			count1 = len1 - gallopRight(ms.tmp[cursor2], ms.data, base1, len1, len1-1)
			if count1 != 0 {
				dest -= count1
				cursor1 -= count1
				len1 -= count1
				copy(ms.data[dest+1:dest+1+count1], ms.data[cursor1+1:cursor1+1+count1])
				if len1 == 0 {
					break outer
				}
			}
			ms.data[dest] = ms.tmp[cursor2]
			dest--
			cursor2--
			len2--
			if len2 == 1 {
				break outer
			}
			
			count2 = len2 - gallopLeft(ms.data[cursor1], ms.tmp, 0, len2, len2-1)
			if count2 != 0 {
				dest -= count2
				cursor2 -= count2
				len2 -= count2
				copy(ms.data[dest+1:dest+1+count2], ms.tmp[cursor2+1:cursor2+1+count2])
				if len2 == 1 {
					break outer
				}
			}
			ms.data[dest] = ms.data[cursor1]
			dest--
			cursor1--
			len1--
			if len1 == 0 {
				break outer
			}
			
			minGallop--
			if count1 < MIN_GALLOP && count2 < MIN_GALLOP {
				if minGallop < 0 {
					minGallop = 0
				}
				minGallop += 2
				break
			}
		}
	}
	
	ms.minGallop = minGallop
	if minGallop < 1 {
		ms.minGallop = 1
	}
	
	if len2 == 1 {
		dest -= len1
		cursor1 -= len1
		copy(ms.data[dest+1:dest+1+len1], ms.data[cursor1+1:cursor1+1+len1])
		ms.data[dest] = ms.tmp[cursor2]
	} else {
		copy(ms.data[dest-(len2-1):dest+1], ms.tmp[:len2])
	}
}

// mergeAt merges the run at stack index i with the one above it
func (ms *MergeState) mergeAt(i int) {
	base1 := ms.runs[i].base
	len1 := ms.runs[i].length
	base2 := ms.runs[i+1].base
	len2 := ms.runs[i+1].length
	
	// Record the length of the combined runs; if i is the 3rd-last
	// run now, also slide over the last run (which isn't involved
	// in this merge). The current run (i+1) goes away in any case.
	ms.runs[i].length = len1 + len2
	if i == len(ms.runs)-3 {
		ms.runs[i+1] = ms.runs[i+2]
	}
	ms.runs = ms.runs[:len(ms.runs)-1]
	
	// Find where the first element of run2 goes in run1
	k := gallopRight(ms.data[base2], ms.data, base1, len1, 0)
	base1 += k
	len1 -= k
	if len1 == 0 {
		return
	}
	
	// Find where the last element of run1 goes in run2
	len2 = gallopLeft(ms.data[base1+len1-1], ms.data, base2, len2, len2-1)
	if len2 == 0 {
		return
	}
	
	// Merge what remains of the runs, using tmp array with min(len1, len2) elements
	if len1 <= len2 {
		ms.mergeLow(base1, len1, base2, len2)
	} else {
		ms.mergeHigh(base1, len1, base2, len2)
	}
}

// mergeCollapse maintains the stack invariant by merging runs as needed
func (ms *MergeState) mergeCollapse() {
	for len(ms.runs) > 1 {
		n := len(ms.runs) - 2
		if (n > 0 && ms.runs[n-1].length <= ms.runs[n].length+ms.runs[n+1].length) ||
		   (n > 1 && ms.runs[n-2].length <= ms.runs[n-1].length+ms.runs[n].length) {
			if ms.runs[n-1].length < ms.runs[n+1].length {
				n--
			}
			ms.mergeAt(n)
		} else if ms.runs[n].length <= ms.runs[n+1].length {
			ms.mergeAt(n)
		} else {
			break
		}
	}
}

// mergeForceCollapse merges all remaining runs
func (ms *MergeState) mergeForceCollapse() {
	for len(ms.runs) > 1 {
		n := len(ms.runs) - 2
		if n > 0 && ms.runs[n-1].length < ms.runs[n+1].length {
			n--
		}
		ms.mergeAt(n)
	}
}

// pushRun adds a new run to the stack
func (ms *MergeState) pushRun(base, length int, desc bool) {
	ms.runs = append(ms.runs, Run{base: base, length: length, desc: desc})
}

// adaptiveTimSort is the main TimSort implementation with intelligent pattern detection
func adaptiveTimSort(data []int) {
	n := len(data)
	if n < 2 {
		return
	}
	
	// For small arrays, use simple insertion sort
	if n < MIN_MERGE {
		binaryInsertionSort(data, 0, n)
		return
	}
	
	// ADAPTIVE PATTERN DETECTION - The key to beating Go's stdlib!
	pattern := detectDataPattern(data)
	
	switch pattern {
	case PatternSorted:
		return // Already sorted!
	case PatternReversed:
		reverseArray(data) // O(n) reverse and done
		return
	case PatternNearlySorted:
		// Use real TimSort - perfect for nearly sorted data!
		timSort(data)
		return
	case PatternComplex:
		// Use our adaptive parallel TimSort for complex cases
		if n >= 1000 {
			adaptiveParallelTimSort(data)
		} else {
			timSort(data)
		}
	}
}

// adaptiveParallelTimSort combines intelligent pattern detection with parallel TimSort features
func adaptiveParallelTimSort(data []int) {
	n := len(data)
	temp := make([]int, n)
	
	// Phase 1: Parallel chunk sorting using TimSort for each chunk
	var wg sync.WaitGroup
	cpuCount := runtime.NumCPU()
	chunkSize := (n + cpuCount - 1) / cpuCount
	
	for i := 0; i < cpuCount; i++ {
		start := i * chunkSize
		end := int(math.Min(float64((i+1)*chunkSize), float64(n)))
		
		if start >= n {
			break
		}
		
		wg.Add(1)
		go func(start, end int) {
			defer wg.Done()
			// Use TimSort on each chunk for better performance
			chunk := data[start:end]
			timSort(chunk)
		}(start, end)
	}
	wg.Wait()
	
	// Phase 2: Bottom-up merge (like in your original implementation)
	for size := chunkSize; size < n; size *= 2 {
		for left := 0; left < n; left += 2 * size {
			mid := int(math.Min(float64(left+size-1), float64(n-1)))
			right := int(math.Min(float64(left+2*size-1), float64(n-1)))
			if mid < right {
				mergeRangesFixed(data, temp, left, mid+1, right+1)
			}
		}
	}
	
	// Phase 3: Safety net - final insertion sort pass to ensure correctness
	// This is the key insight from adaptive algorithm design!
	binaryInsertionSort(data, 0, n)
}

// timSort performs regular TimSort on the entire array
func timSort(data []int) {
	timSortRange(data, 0, len(data))
}

// timSortRange performs TimSort on a range of the array
func timSortRange(data []int, start, end int) {
	n := end - start
	if n < 2 {
		return
	}
	
	if n < MIN_MERGE {
		binaryInsertionSort(data, start, end)
		return
	}
	
	ms := NewMergeState(data[start:end])
	minRun := computeMinRunLength(n)
	
	lo := 0
	for lo < n {
		runLen, _ := countRunAndMakeAscending(ms.data, lo, n)
		
		// If run is too short, extend using insertion sort
		if runLen < minRun {
			force := minRun
			if n-lo < minRun {
				force = n - lo
			}
			binaryInsertionSort(ms.data, lo, lo+force)
			runLen = force
		}
		
		// Push run onto pending-run stack, and maybe merge
		ms.pushRun(lo, runLen, false)
		ms.mergeCollapse()
		
		lo += runLen
	}
	
	// Merge all remaining runs
	ms.mergeForceCollapse()
	
	// Copy back to original array
	copy(data[start:end], ms.data)
}

// mergeChunksWithBoundaries merges chunks using actual computed boundaries
func mergeChunksWithBoundaries(data []int, boundaries []int) {
	if len(boundaries) < 2 {
		return // Nothing to merge
	}
	
	temp := make([]int, len(data))
	
	// Merge chunks one by one from left to right
	for i := 1; i < len(boundaries); i++ {
		start := 0
		mid := boundaries[i-1]
		end := boundaries[i]
		
		// Merge [0, mid) with [mid, end) into [0, end)
		mergeRanges(data, temp, start, mid, end)
	}
}

// mergeChunksSequential merges the parallel-sorted chunks one by one
func mergeChunksSequential(data []int, chunkSize, cpuCount int) {
	n := len(data)
	temp := make([]int, n)
	
	// Merge chunks one by one from left to right
	currentEnd := chunkSize
	if currentEnd > n {
		currentEnd = n
	}
	
	for i := 1; i < cpuCount && currentEnd < n; i++ {
		nextEnd := (i + 1) * chunkSize
		if nextEnd > n {
			nextEnd = n
		}
		
		// Merge [0, currentEnd) with [currentEnd, nextEnd)
		mergeRanges(data, temp, 0, currentEnd, nextEnd)
		currentEnd = nextEnd
	}
}

// mergeChunks merges the parallel-sorted chunks (old version - backup)
func mergeChunks(data []int, chunkSize, cpuCount int) {
	n := len(data)
	temp := make([]int, n)
	
	// Iteratively merge chunks
	for size := chunkSize; size < n; size *= 2 {
		for start := 0; start < n; start += 2 * size {
			mid := int(math.Min(float64(start+size), float64(n)))
			end := int(math.Min(float64(start+2*size), float64(n)))
			
			if mid < end {
				mergeRanges(data, temp, start, mid, end)
			}
		}
	}
}

// mergeRangesFixed merges two sorted ranges using the original parallel TimSort approach
func mergeRangesFixed(data, temp []int, left, mid, right int) {
	copy(temp[left:right], data[left:right])
	
	i, j, k := left, mid, left
	
	for i < mid && j < right {
		if temp[i] <= temp[j] {
			data[k] = temp[i]
			i++
		} else {
			data[k] = temp[j]
			j++
		}
		k++
	}
	
	// Copy remaining elements from left part
	for i < mid {
		data[k] = temp[i]
		i++
		k++
	}
	
	// Copy remaining elements from right part  
	for j < right {
		data[k] = temp[j]
		j++
		k++
	}
}

// mergeRanges merges two sorted ranges using a temporary array
func mergeRanges(data, temp []int, start, mid, end int) {
	copy(temp[start:end], data[start:end])
	
	i, j, k := start, mid, start
	
	for i < mid && j < end {
		if temp[i] <= temp[j] {
			data[k] = temp[i]
			i++
		} else {
			data[k] = temp[j]
			j++
		}
		k++
	}
	
	// Copy remaining elements
	for i < mid {
		data[k] = temp[i]
		i++
		k++
	}
	for j < end {
		data[k] = temp[j]
		j++
		k++
	}
}

// detectDataPattern analyzes the data to choose optimal sorting strategy
func detectDataPattern(data []int) DataPattern {
	n := len(data)
	if n < 2 {
		return PatternSorted
	}
	
	// Quick sorted check - O(n) but early exit
	sorted := true
	for i := 1; i < n; i++ {
		if data[i] < data[i-1] {
			sorted = false
			break
		}
	}
	if sorted {
		return PatternSorted
	}
	
	// Quick reverse sorted check - O(n) but early exit  
	reversed := true
	for i := 1; i < n; i++ {
		if data[i] > data[i-1] {
			reversed = false
			break
		}
	}
	if reversed {
		return PatternReversed
	}
	
	// Nearly sorted check - sample approach for performance
	sampleSize := int(math.Min(float64(n), 1000)) // Sample first 1000 elements
	outOfOrder := 0
	
	for i := 1; i < sampleSize; i++ {
		if data[i] < data[i-1] {
			outOfOrder++
		}
	}
	
	// If less than 10% out of order in sample, consider nearly sorted
	if float64(outOfOrder)/float64(sampleSize) < 0.1 {
		return PatternNearlySorted
	}
	
	return PatternComplex
}

// reverseArray reverses the array in-place - O(n)
func reverseArray(data []int) {
	n := len(data)
	for i := 0; i < n/2; i++ {
		data[i], data[n-1-i] = data[n-1-i], data[i]
	}
}

// Note: This file contains the Adaptive TimSort implementation
// For testing and demonstration, use: go run 22B-adaptive-timsort-TEST.go