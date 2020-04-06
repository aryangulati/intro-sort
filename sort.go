package main

import (
	"fmt"
	"math"
)

type IntroSort struct {
	data []int

	depthLimit int
}

func (s IntroSort) IntroSort() {
	begin := 0
	end := len(s.data) - 1

	s.depthLimit = 2 * int(math.Round(math.Log2(float64(end))))
	fmt.Println("data length is: ", len(s.data))
	fmt.Println("depth limit is: ", s.depthLimit)
	s.introSortUtil(begin, end)

}

func (s IntroSort) introSortUtil(begin, end int) {
	size := end - begin

	if size < 16 {
		fmt.Printf("use insert sort for %d - %d\n", begin, end)
		s.insertionSort(begin, end)
		return
	}

	if s.depthLimit == 0 {
		fmt.Printf("use heap sort for %d - %d\n", begin, end)
		s.heapSort()
		return
	}

	fmt.Printf("use quick sort for %d - %d\n", begin, end)

	pivot := s.medianOfThree(begin, begin+size/2, end)
	fmt.Printf("found pivot: %d -> %d\n", pivot, s.data[pivot])

	s.data[pivot], s.data[end] = s.data[end], s.data[pivot]

	partitionPoint := s.partition(begin, end)
	s.depthLimit -= 1

	fmt.Printf("partition %d -> %d for depth %d\n", partitionPoint, s.data[partitionPoint], s.depthLimit)

	s.introSortUtil(begin, partitionPoint-1)
	s.introSortUtil(partitionPoint+1, end)

}

func (s IntroSort) partition(low, high int) int {
	pivot := s.data[high]

	i := low - 1

	for j := low; j < high; j++ {
		if s.data[j] <= pivot {
			i = i + 1
			s.data[i], s.data[j] = s.data[j], s.data[i]
		}
	}
	s.data[i+1], s.data[high] = s.data[high], s.data[i+1]
	fmt.Printf("found pivot at %d -> %d, before: %+v, after: %+v\n", i+1, s.data[i+1], s.data[low:i+1], s.data[i+2:high+1])
	return i + 1

}

func (s IntroSort) medianOfThree(a, b, c int) int {
	da := s.data[a]
	db := s.data[b]
	dc := s.data[c]

	if da <= db && db <= dc {
		return b
	}

	if dc <= db && db <= da {
		return b
	}

	if db <= da && da <= dc {
		return a
	}
	if dc <= da && da <= db {
		return a
	}

	if da <= dc && dc <= db {
		return c
	}

	if db <= dc && dc <= da {
		return c
	}

	return a

}

// insertionSort use insertion sort to sort data
func (s IntroSort) insertionSort(begin, end int) {
	fmt.Println("insert sort start: ", s.data[begin:end+1])

	left := begin

	for i := left + 1; i <= end; i++ {
		key := s.data[i]
		j := i - 1

		for j >= left && s.data[j] > key {
			s.data[j+1] = s.data[j]
			j = j - 1
		}
		s.data[j+1] = key
	}

	fmt.Println("insert sort end: ", s.data[begin:end+1])

}

func (s IntroSort) heapSort() {
	heap := new(Heap)

	heap.HeapSort(s.data)

}

func main() {
	input := []int{2, 10, 24, 2, 10, 11, 27, 4, 2, 4, 28, 16, 9, 8, 28, 10, 13, 24, 22, 28, 0, 13, 27, 13, 3, 23, 18, 22, 8, 8, 88, 1200, 77, 94, 6, 29, 54, 3}
	fmt.Println("input is: ", input)
	is := IntroSort{data: input}
	is.IntroSort()
	fmt.Println("len result is: ", len(is.data))
	fmt.Println("Sort Result: ", is.data)
}
