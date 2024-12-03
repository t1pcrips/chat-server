package tests

import "math/rand"

const (
	successDelete = "success delete test"
	failedDelete  = "failed delete test"
	successCreate = "success create test"
	failedCreate  = "failed create test"
	successSend   = "success send test"
	failedSend    = "failed send test"
)

func names(numNames int) []string {
	// Массив имен
	names := []string{"Alice", "Bob", "Charlie", "David", "Eve", "Frank", "Grace", "Heidi", "Tom", "Keyle"}

	// Генерация массива случайных имен
	randomNames := make([]string, numNames)
	for i := 0; i < numNames; i++ {
		randomNames[i] = names[rand.Intn(len(names))]
	}
	return randomNames
}
