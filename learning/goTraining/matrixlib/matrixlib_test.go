package matrixlib_test

import (
	"fmt"
	"testing"
)

import "github.com/parinay/golang/goTraining/matrixlib"

func TestNewMatrix(t *testing.T) {
	m1, err := matrixlib.NewMatrix(3)
	fmt.Printf("%T\n", m1)
	if err != nil {
		t.Log("Should not be error")
		t.Fail()
	}

	if m1.Size() != 3 {
		t.Log("Size not properly set.")
		t.Fail()
	}
	if value, err := m1.Get(2, 2); err != nil || value != 0 {
		t.Log("Could not correctly access last value in matirx")
		t.Fail()
	}
}

func TestNewMatrixNegSize(t *testing.T) {
	m2, err := matrixlib.NewMatrix(-1)
	fmt.Printf("%T\n", m2)
	if err == nil {
		t.Log("Invalid size should be an error")
		t.Fail()
	}
}

func TestAdd(t *testing.T) {

}
