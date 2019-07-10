package matrixlib

import (
	"errors"
	"fmt"
)

type matrix struct {
	size   int     "Size of matrix"
	values [][]int "Values of matrix"
}

type other matrix

func main() {
	var m1 matrix
	fmt.Println(m1)

	var m2 other = other(m1)
	fmt.Println(m2)

	m1 = matrix{3,
		[][]int{
			[]int{1, 2, 3},
			[]int{4, 5, 6},
			[]int{7, 8},
		},
	}
	fmt.Println(m1.values[2][1])
}

func NewMatrix(size int) (result *matrix, err error) {
	if size < 1 {
		err = errors.New("Invalid size")
		return
	}
	result = &matrix{}
	result.size = size
	result.values = make([][]int, size)

	for index, _ := range result.values {
		result.values[index] = make([]int, size)
	}
	// result = matrix{size: size}
	return
}

// First Class functions
// func (	,	int) int
// func ops( lavlue, r value int) value int () {}
// type matirxop func(int, int) int
// var m1 matrixop
type matrixop func(int, int) int

func Add(lhs matrix, rhs matrix) (result *matrix, err error) {
	return iterate(lhs, rhs, addop)
}
func addop(lval, rval int) int {
	return lval + rval
}
func subop(l, r int) int {
	return l - r
}
func (m *matrix) Set(row int, col int, value int) (err error) {
	if row < 0 || row > m.size && col < 0 || col > m.size {
		err = errors.New("Invalid size")
		return
	}
	m.values[row][col] = value
	return nil
}

func iterate(lhs, rhs matrix, operation matrixop) (result *matrix, err error) {
	if lhs.size != rhs.size {
		err = errors.New("Incompatible matix size")
		return
	}
	if lhs.values == nil || rhs.values == nil {
		err = errors.New("Nil values")
		return
	}
	result, _ = NewMatrix(lhs.size)

	for i, value := range lhs.values {
		for j, innervalue := range value {
			rhsvalue, _ := rhs.Get(i, j)
			if err != nil {
				err = errors.New("Get Error")
				return
			}
			err = result.Set(i, j, operation(innervalue, rhsvalue))
			if err != nil {
				err = errors.New("Set Error")
				return
			}
		}
	}
	return
}

func Substract(lhs, rhs matrix) (matrix, error) {
	return matrix{}, nil
}

func Set(m *matrix, row, col, values int) error {
	return nil
}

// func Get(m *matrix, row, col int) (int, error) {
// 	return 1, nil
// }
func (m *matrix) Get(row, col int) (element int, err error) {

	if row < 0 || row > m.size && col < 0 || col > m.size {
		err = errors.New("Invalid size")
		return
	}
	element = m.values[row][col]
	return
}

func (m *matrix) Size() int {
	return m.size
}
