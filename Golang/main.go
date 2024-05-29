package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"

	"gonum.org/v1/gonum/mat"
)

func parseEquation(equation string) ([]float64, error) {
	coeffs := make([]float64, 4) // a, b, c, and d
	// Regex to match terms with coefficients
	termRegex := regexp.MustCompile(`([+-]?\d*\.?\d*)([xyz])`)
	// Regex to match the constant term
	constantRegex := regexp.MustCompile(`=\s*([+-]?\d*\.?\d*)`)

	lhs := strings.Split(equation, "=")[0]
	rhs := constantRegex.FindStringSubmatch(equation)
	if len(rhs) != 2 {
		return nil, fmt.Errorf("equation must contain '=' sign")
	}
	constant, err := strconv.ParseFloat(rhs[1], 64)
	if err != nil {
		return nil, err
	}
	coeffs[3] = constant

	terms := termRegex.FindAllStringSubmatch(lhs, -1)
	for _, term := range terms {
		coeffStr, variable := term[1], term[2]
		if coeffStr == "" || coeffStr == "+" {
			coeffStr = "1"
		} else if coeffStr == "-" {
			coeffStr = "-1"
		}
		coeff, err := strconv.ParseFloat(coeffStr, 64)
		if err != nil {
			return nil, err
		}
		switch variable {
		case "x":
			coeffs[0] = coeff
		case "y":
			coeffs[1] = coeff
		case "z":
			coeffs[2] = coeff
		}
	}
	return coeffs, nil
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter 3 equations in the form 'ax + by + cz = d':")

	var AData []float64
	var BData []float64

	for i := 0; i < 3; i++ {
		fmt.Printf("Enter equation %d: ", i+1)
		scanner.Scan()
		equation := scanner.Text()
		coeffs, err := parseEquation(equation)
		if err != nil {
			fmt.Println("Error parsing equation:", err)
			return
		}
		AData = append(AData, coeffs[0], coeffs[1], coeffs[2])
		BData = append(BData, coeffs[3])
	}

	A := mat.NewDense(3, 3, AData)
	B := mat.NewVecDense(3, BData)

	fmt.Println("\nMatrix A:")
	matPrint(A)

	fmt.Println("\nVector B:")
	matPrint(B)

	var X mat.VecDense
	err := X.SolveVec(A, B)
	if err != nil {
		fmt.Println("The system of equations does not have a unique solution:", err)
		return
	}

	fmt.Println("\nSolution for X:")
	fmt.Printf("x = %.2f\n", X.AtVec(0))
	fmt.Printf("y = %.2f\n", X.AtVec(1))
	fmt.Printf("z = %.2f\n", X.AtVec(2))
}

func matPrint(X mat.Matrix) {
	fa := mat.Formatted(X, mat.Prefix(""), mat.Excerpt(0))
	fmt.Printf("%v\n", fa)
}
