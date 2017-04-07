package main

import (
	"reflect"
	"testing"
	"github.com/pkg/errors"
	"time"
)

func TestSum(t *testing.T) {
	if Sum(1, 3) != 4 {
		t.Error("ACHTUNG!", "expected", 4, "got", Sum(1, 3))
	}
}

func TestDevision(t *testing.T) {
	t.Parallel()
	defer func() {
		recover()
	}()

	if Division(2, 1) != 2 {
		t.Errorf("expected %d, got %d", 2, Division(2, 1))
	}

	Division(2, 0)
	t.Error("Panic expected")
}

func TestAddUser(t *testing.T) {
	users := []User{}

	AddUser(&users, "Vasya", 32)

	if len(users) == 0 {
		t.Fatal("Empty slice")
	}

	expected := []User{
		{
			Name: "Vasya",
			Age:  32,
		},
	}

	if !reflect.DeepEqual(users, expected) {
		t.Errorf("expected %+v, got %+v", expected, users)
	}

}

// Table testing
// Удобен, когда есть много входных и выходных значений
// Чтобы не копипастить код сравнения
func TestTable(t *testing.T) {
	data := []struct {
		A      int
		B      int
		Result int
		Err    error
	}{
		{
			A:      1,
			B:      1,
			Result: 1,
			Err:    nil,
		},
		{
			A:      0,
			B:      1,
			Result: 0,
			Err:    nil,
		},
		{
			A:      1,
			B:      0,
			Result: 0,
			Err:    errors.New("Division by zero"),
		},
	}

	for _, testCase := range data {
		res, err := TrueDivision(testCase.A, testCase.B)
		if res != testCase.Result {
			t.Errorf("Expected %d, got %d", testCase.Result, res)
		}

		if err != nil && err.Error() != testCase.Err.Error() {
			t.Errorf("Expected %+v, got %+v", testCase.Err.Error(), err.Error())
		}
	}
}
