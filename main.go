package main

import (
	"errors"
	"fmt"
	"os"
)

func AddMatrices(a, b [][]float64) ([][]float64, error) {
	if len(a) != len(b) || len(a[0]) != len(b[0]) {
		return nil, errors.New("матрицы должны быть одинакового размера")
	}

	size := len(a)
	result := make([][]float64, size)
	for i := range result {
		result[i] = make([]float64, size)
		for j := range result[i] {
			result[i][j] = a[i][j] + b[i][j]
		}
	}
	return result, nil
}

func MultiplyMatrixByScalar(matrix [][]float64, scalar float64) [][]float64 {
	size := len(matrix)
	result := make([][]float64, size)
	for i := range result {
		result[i] = make([]float64, size)
		for j := range result[i] {
			result[i][j] = matrix[i][j] * scalar
		}
	}
	return result
}

func MultiplyMatrices(a, b [][]float64) ([][]float64, error) {
	size := len(a)
	if len(a) != len(b) || len(a[0]) != len(b[0]) {
		return nil, errors.New("матрицы должны быть одинакового размера")
	}

	result := make([][]float64, size)
	for i := range result {
		result[i] = make([]float64, size)
		for j := range result[i] {
			for k := 0; k < size; k++ {
				result[i][j] += a[i][k] * b[k][j]
			}
		}
	}
	return result, nil
}

func printMatrix(matrix [][]float64) {
	size := len(matrix)
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			fmt.Printf("%8.2f ", matrix[i][j])
		}
		fmt.Println()
	}
}

func inputMatrix(size int, name string) [][]float64 {
	fmt.Printf("\n%s (%dx%d):\n", name, size, size)
	matrix := make([][]float64, size)
	for i := range matrix {
		matrix[i] = make([]float64, size)
		for j := range matrix[i] {
			fmt.Printf("Элемент [%d][%d]: ", i+1, j+1)
			var val float64
			fmt.Scanln(&val)
			matrix[i][j] = val
		}
	}
	return matrix
}

func getMatrixSize() int {
	var size int
	for {
		fmt.Print("Введите размер матрицы (2 или 3): ")
		fmt.Scanln(&size)
		if size == 2 || size == 3 {
			break
		}
		fmt.Println("Неверный размер. Введите 2 или 3.")
	}
	return size
}

func handleMatrixAddition() {
	size := getMatrixSize()
	a := inputMatrix(size, "Первая матрица")
	b := inputMatrix(size, "Вторая матрица")

	result, err := AddMatrices(a, b)
	if err != nil {
		fmt.Printf("Ошибка: %s\n", err)
		return
	}

	fmt.Println("\nРезультат сложения:")
	printMatrix(result)
}

func handleScalarMultiplication() {
	size := getMatrixSize()
	matrix := inputMatrix(size, "Матрица")
	var scalar float64
	fmt.Print("Введите число для умножения: ")
	fmt.Scanln(&scalar)

	result := MultiplyMatrixByScalar(matrix, scalar)
	fmt.Println("\nРезультат умножения на число:")
	printMatrix(result)
}

func handleMatrixMultiplication() {
	fmt.Println("Умножение матриц: A * B")
	fmt.Print("Размер первой матрицы (2 или 3): ")
	var sizeA int
	fmt.Scanln(&sizeA)

	fmt.Print("Размер второй матрицы (2 или 3): ")
	var sizeB int
	fmt.Scanln(&sizeB)

	if sizeA != sizeB {
		fmt.Println("Ошибка: размеры матриц должны совпадать для умножения 2x2 или 3x3.")
		return
	}

	a := inputMatrix(sizeA, "Первая матрица")
	b := inputMatrix(sizeB, "Вторая матрица")

	result, err := MultiplyMatrices(a, b)
	if err != nil {
		fmt.Printf("Ошибка: %s\n", err)
		return
	}

	fmt.Println("\nРезультат умножения:")
	printMatrix(result)
}

func main() {
	fmt.Println("=== Калькулятор матриц Тощев ИС-323 💾 ===")
	fmt.Println()

	for {
		fmt.Println("Выберите операцию:")
		fmt.Println("1. Сложение матриц")
		fmt.Println("2. Умножение матрицы на число")
		fmt.Println("3. Умножение двух матриц")
		fmt.Println("4. Выход")

		var choice string
		fmt.Print("Ваш выбор: ")
		fmt.Scanln(&choice)

		switch choice {
		case "1":
			handleMatrixAddition()
		case "2":
			handleScalarMultiplication()
		case "3":
			handleMatrixMultiplication()
		case "4":
			fmt.Println("До свидания!")
			os.Exit(0)
		default:
			fmt.Println("Неверный выбор. Попробуйте снова.")
		}
		fmt.Println()
	}
}
