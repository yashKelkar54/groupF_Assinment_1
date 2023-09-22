package main

import "fmt"

main() {
	(calculateAverage())
	(calcAv())
	(cube())
	(fibseq())
	(even())
	(max())
	(sum())
	(sum2())
}

// yash function // 500223746
func fibseq() {
    for i := 0; i < 10; i++ {
        fmt.Println(fibonaciseq(i), " ")
    }
}

func fibonaciseq(n int) int {
    if n <= 1 {
        return n
    }
    return fibonaciseq(n-1) + fibonaciseq(n-2)
}

// aarti function 500224115
func even() {

    var num int = 8

    if num%2 == 0 {
        fmt.Println(num, "is even.")
    } else {
        fmt.Println(num, "is not even.")
    }
}
 
// hardik function
func max() (num1, num2 int) int {
	/* local variable declaration */
	var result int

	if num1 > num2 {
		result = num1
	} else {
		result = num2
	}
	return result
}

func calcAv() {
	var a int = 100
	var b int = 200
	var ret int

	/* calling a function to get max value */
	ret = max(a, b)

	fmt.Printf("Max value is : %d\n", ret)

	// Oweipadei Joshua Bayefa: Function
	numbers := []float64{1.0, 2.0, 3.0, 4.0, 5.0}
	avg := calculateAverage(numbers)
	fmt.Printf("Average: %.2f\n", avg)

	var num float32 = 7.6
	cubee := cube(num)
	fmt.Printf("Cube: %.2f\n", cubee)
	result := add(3, 5)
	fmt.Printf("3 + 5 = %d\n", result)
}

func calculateAverage(numbers []float64) float64 {
	// Avoid division by zero for empty slices
	if len(numbers) == 0 {
		return 0.0
	}

	var sum float64
	for _, num := range numbers {
		sum += num
	}

	return sum / float64(len(numbers))
}
 // karan code
func cube(x float64) float64 {
	return x * x * x
}

// sumit code
func add(a, b int) int {
	return a + b
}

// Sukhneet function 500224802
func sum(num1, num2 int) int {
	return num1 + num2
}

func sum2() {
	result := sum(9, 1)

	fmt.Println(result)
}
