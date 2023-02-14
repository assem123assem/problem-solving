package problem6

type Matrix = [][]int

func Transpose(original Matrix) Matrix {

	var transposed Matrix = make(Matrix, len(original))

	for i := 0; i < len(original); i++ {
		for j := 0; j < len(original[0]); j++ {
			transposed[j] = append(transposed[j], original[i][j])

		}
	}

	return transposed
}
