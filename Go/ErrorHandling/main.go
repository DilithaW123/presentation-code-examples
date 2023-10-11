package main

import (
	"errors"
	"log"
)

func doSomethingCool(errorPls bool) (int, error) {
	if errorPls {
		return 0, errors.New("Something went wrong")
	}
	return 1, nil
}

func main() {
	val, err := doSomethingCool(false)
	if err != nil {
		panic(err)
	}
	println(val)
	_, err = doSomethingCool(true) // _ is used to ignore the return value
	if err != nil {
		log.Fatal(err)
	}
}
