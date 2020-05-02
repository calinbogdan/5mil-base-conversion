package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

type NumberData struct {
	Number int
	Base   int
	Target int
}

func main() {
	start := time.Now()

	numbers := make([]NumberData, 0)
	for i := 0; i < 5000000; i++ {
		randomNumber := rand.Intn(50000)
		randomCurrentBase := rand.Intn(8) + 2
		randomTargetBase := rand.Intn(8) + 2

		numbers = append(numbers, NumberData{randomNumber, randomCurrentBase, randomTargetBase})
	}

	convertRandomNumbers(numbers)
	fmt.Printf("It took %f seconds to convert 5M numbers.", time.Since(start).Seconds())
}

func convertRandomNumbers(numbers []NumberData) {
	results := make([]int, 0)
	resultChannel := make(chan []int, 4)

	step := 5000
	for i := 0; i < len(numbers); i += step {
		go func(numbersBatch []NumberData) {
			resultsBatch := make([]int, 0)
			for _, value := range numbersBatch {
				resultsBatch = append(resultsBatch, convertFromBaseTo(value.Number, value.Base, value.Target))
			}
			resultChannel <- resultsBatch
		}(numbers[i : i+step])
	}

	for i := 0; i < len(numbers); i += step {
		resultBatch := <-resultChannel
		results = append(results, resultBatch...)
	}
}

func convertFromBaseTo(number, base, nextBase int) int {
	return convertToBase(convertToBase10(number, base), nextBase)
}

func convertToBase10(number, currentBase int) int {
	numberInBase10, _ := strconv.ParseInt(fmt.Sprintf("%d", int64(number)), currentBase, 32)
	return int(numberInBase10)
}

func convertToBase(number, targetBase int) int {
	numberInTargetBase := strconv.FormatInt(int64(number), targetBase)
	actualNumber, _ := strconv.ParseInt(numberInTargetBase, 10, 32)
	return int(actualNumber)
}

func toDigits(number, base int) string {
	digits := ""
	for number > 0 {
		digits = fmt.Sprintf("%d", number%base) + digits
		number = number / base
	}
	return digits
}

func fromDigits(digits string, base int) int {
	number := 0
	for _, char := range digits {
		number = number*base + int(char-'0')
	}
	return number
}

func convertBaseWithDigits(number string, base, target int) string {
	return toDigits(fromDigits(number, base), target)
}
