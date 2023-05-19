package main

import "fmt"

func mergeArrays(array1 []string, array2 []string) ([]string, []string) {
	mergedArray := make([]string, len(array1))
	combinedArray := make([]string, 0)

	copy(mergedArray, array1)
	for i := 0; i < len(array2); i++ {
		exists := false
		for j := 0; j < len(mergedArray); j++ {
			if array2[i] == mergedArray[j] {
				exists = true
				break
			}
		}
		if !exists {
			mergedArray = append(mergedArray, array2[i])
		}
	}

	for i := 0; i < len(array1); i++ {
		if contains(array1[i], array2) && !contains(array1[i], combinedArray) {
			combinedArray = append(combinedArray, array1[i])
		}
	}
	for i := 0; i < len(array2); i++ {
		if contains(array2[i], array1) && !contains(array2[i], combinedArray) {
			combinedArray = append(combinedArray, array2[i])
		}
	}

	return mergedArray, combinedArray
}

func contains(elem string, array []string) bool {
	for i := 0; i < len(array); i++ {
		if elem == array[i] {
			return true
		}
	}
	return false
}

func main() {
	array1 := []string{"a", "b", "c", "d"}
	array2 := []string{"a", "d", "e", "f"}

	merged, combined := mergeArrays(array1, array2)

	fmt.Println("Merged Array:", merged)
	fmt.Println("Combined Array:", combined)
}
