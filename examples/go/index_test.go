package main_test

import (
	"github.com/go/faker_example"
	"fmt"
	"testing"
	"github.com/bxcodec/faker"
)


func Test(t *testing.T) {
	a := main.SomeStruct{}
	faker.FakeData(&a)
	fmt.Println(a)
}