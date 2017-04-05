package main

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	if Sum(1, 3) != 4 {
		t.Error("ACHTUNG!", "expected", 4, "got", Sum(1, 3))
	}
}

func TestDevision(t *testing.T) {
	defer func() {
		recover()
	}()

	if Devision(2, 1) != 2 {
		t.Errorf("expected %d, got %d", 2, Devision(2, 1))
	}

	Devision(2, 0)
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
