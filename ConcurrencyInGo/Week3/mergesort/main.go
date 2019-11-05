package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"sync"
)

func merge(left, right []int) []int {
	res := make([]int, len(left)+len(right))
	i := 0
	for len(left) > 0 && len(right) > 0 {
		if left[0] < right[0] {
			res[i] = left[0]
			left = left[1:]
		} else {
			res[i] = right[0]
			right = right[1:]
		}
		i++
	}
	for j := 0; j < len(left); j++ {
		res[i] = left[j]
		i++
	}
	for j := 0; j < len(right); j++ {
		res[i] = right[j]
		i++
	}
	return res
}

func mySort(sli []int, wg *sync.WaitGroup) {
	sort.Ints(sli)
	wg.Done()
}

func main() {
	fmt.Println("Please input some numbers (more than 4) to be sorted:")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	inputs := strings.Fields(input)

	sli := make([]int, 0, 10)
	for _, i := range inputs {
		num, _ := strconv.Atoi(i)
		sli = append(sli, num)
	}

	// divide input into four parts and store in map
	m := map[int][]int{
		1: []int{},
		2: []int{},
		3: []int{},
		4: []int{},
	}
	i := 1
	for _, val := range sli {
		m[i] = append(m[i], val)
		i++
		if i > 4 {
			i = 1
		}
	}

	// go routine sort 4 parts concurrently
	var wg sync.WaitGroup
	wg.Add(4)
	for i := 1; i <= 4; i++ {
		go mySort(m[i], &wg)
	}
	wg.Wait()

	// merge results and output, also concurrently and using channel
	res1 := merge(m[1], m[2])
	res2 := merge(m[3], m[4])
	res := merge(res1, res2)
	fmt.Println(res)
}
