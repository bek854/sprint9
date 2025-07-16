package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	SIZE    = 100_000_000
	CHUNKS  = 8
	MAX_NUM = 1_000_000_000
)

// generateRandomElements generates random elements.
func generateRandomElements(size int) []int {
	// ваш код здесь
	if size <= 0 {
		return nil
	}
	data := make([]int, size)
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	for i := 0; i < size; i++ {
		data[i] = r.Intn(MAX_NUM)
	}
	return data
}

// maximum returns the maximum number of elements.
func maximum(data []int) int {
	// ваш код здесь
	if len(data) == 0 {
		return 0
	}
	max := data[0]
	for _, num := range data[1:] {
		if num > max {
			max = num
		}
	}
	return max
}

// maxChunks returns the maximum number of elements in a chunks.
func maxChunks(data []int) int {
	// ваш код здесь
	if len(data) == 0 {
		return 0
	}

	chunkSize := len(data) / CHUNKS
	results := make([]int, CHUNKS)
	var wg sync.WaitGroup

	for i := 0; i < CHUNKS; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			start := i * chunkSize
			end := start + chunkSize
			if i == CHUNKS-1 {
				end = len(data)
			}

			chunkMax := data[start]
			for _, num := range data[start+1 : end] {
				if num > chunkMax {
					chunkMax = num
				}
			}
			results[i] = chunkMax
		}(i)
	}

	wg.Wait()
	return maximum(results)
}

func main() {
	fmt.Printf("Генерируем %d целых чисел", SIZE)
	// ваш код здесь
	data := generateRandomElements(SIZE)

	fmt.Println("Ищем максимальное значение в один поток")
	// ваш код здесь
	start := time.Now()
	max := maximum(data)
	elapsed := time.Since(start).Milliseconds()

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed)

	fmt.Printf("Ищем максимальное значение в %d потоков", CHUNKS)
	// ваш код здесь
	start = time.Now()
	max = maxChunks(data)
	elapsed = time.Since(start).Milliseconds()

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed)
}
