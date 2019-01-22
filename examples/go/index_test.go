package main_test

import (
	"github.com/bxcodec/faker"
	"github.com/go/faker_example"
	"testing"
)


func Test(t *testing.T) {
	a := main.SomeStruct{}
	faker.FakeData(&a)
	main.SomeAwesomeFunction(a)
}