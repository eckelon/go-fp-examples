package main

import (
	"errors"
	"fmt"
)

var NewError = func(errorMessage string) error {
	return errors.New(errorMessage)
}

var Sadd = func(a int) func(int) (int, error) {
	return func(b int) (int, error) {
		fmt.Println("sAdd Log", "a", a, "b", b)
		if b == 2 {
			return -1, NewError("error inesperado")
		}
		return a + b, nil
	}
}

var Add = Curry(func(a int, b int) (int, error) {
	return a + b, nil
	// return a + b, newError("error add")
})

var Multiply = Curry(func(a int, b int) (int, error) {
	return a * b, nil
})

var TraceInt = func(a int) (int, error) {
	fmt.Println("tracing a", a)
	return a, nil
}

var DoStuffToANumber = func() func(int) (int, error) {
	return Compose_(Multiply(2), Add(1))
}()
