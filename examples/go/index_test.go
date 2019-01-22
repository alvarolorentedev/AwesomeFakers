package main_test

import (
	"github.com/bxcodec/faker"
	"github.com/go/faker_example"
	"testing"
)


func Test_1(t *testing.T) {
	a := main.SomeStruct{}
	faker.FakeData(&a)
	main.SomeAwesomeFunction(a)
}

func Test_2(t *testing.T) {
	a := main.SomeStruct{}
	a.String = faker.GetLorem().Sentence()
	main.SomeAwesomeFunction(a)
}