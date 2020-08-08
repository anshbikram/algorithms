package main

import (
	"algorithms/sort"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter comma separated integers to sort:")
	values, _ := reader.ReadString('\n')
	values = strings.TrimSuffix(values, "\n")
	var numsStr []string = strings.Split(values, ",")
	var numsInt []int
	for _, num := range numsStr {
		n, _ := strconv.Atoi(num)
		numsInt = append(numsInt, n)
	}

	fmt.Println("Enter sort type (1. Merge):")
	var t int
	fmt.Scan(&t)

	if t == 1 {
		fmt.Println("Sorted nums:", sort.MergeSort(numsInt))
	}
}
