package main

func Contains(numbers []int, target int) bool {
	for _, number := range numbers {
		if number == target {
			return true
		}
	}
	return false
}
