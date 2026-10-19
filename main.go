package main

import (
	"fmt"
	"math"
)

const delta = 1e-12

// Sqrt computes the square root of x using Newton's method.
// It stops once the value changes by less than delta.
func Sqrt(x float64) float64 {
	res, _ := SqrtWithGuess(x, 1.0, false)
	return res
}

// SqrtWithGuess computes sqrt(x) starting with an initial guess z.
// It returns the calculated root and the number of iterations taken.
func SqrtWithGuess(x, z float64, verbose bool) (float64, int) {
	if x < 0 {
		return math.NaN(), 0
	}
	if x == 0 {
		return 0, 0
	}
	if z <= 0 {
		z = 1.0
	}

	iterations := 0
	for {
		iterations++
		next := z - (z*z-x)/(2*z)
		diff := math.Abs(next - z)
		if verbose {
			fmt.Printf("    iteration %2d: %-18.15f (change: %e)\n", iterations, next, diff)
		}
		if diff <= delta {
			return next, iterations
		}
		z = next
	}
}

func main() {
	testValues := []float64{1, 2, 3, 4, 9, 25, 100, 10000}

	fmt.Println("=== Comparing Initial Guesses (Stopping condition: |change| <= 1e-12) ===")
	for _, x := range testValues {
		fmt.Printf("\n--- x = %v (math.Sqrt = %v) ---\n", x, math.Sqrt(x))

		guesses := []struct {
			name string
			val  float64
		}{
			{"z0 = 1", 1.0},
			{"z0 = x", x},
			{"z0 = x/2", x / 2},
		}

		for _, g := range guesses {
			res, iters := SqrtWithGuess(x, g.val, false)
			diff := math.Abs(res - math.Sqrt(x))
			fmt.Printf("  %-10s -> result: %.15f, iterations: %2d, diff: %e\n",
				g.name, res, iters, diff)
		}
	}

	fmt.Println("\n=== Detailed step-by-step for x = 2 (z0 = 1) ===")
	SqrtWithGuess(2, 1.0, true)

	fmt.Println("\n=== Detailed step-by-step for x = 100 (z0 = 1 vs z0 = x/2) ===")
	fmt.Println("With z0 = 1:")
	SqrtWithGuess(100, 1.0, true)
	fmt.Println("With z0 = x/2 (50):")
	SqrtWithGuess(100, 50.0, true)
}
