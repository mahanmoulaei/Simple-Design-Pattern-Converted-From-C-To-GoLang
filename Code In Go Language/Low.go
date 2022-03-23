package main

import (
	"fmt"
)

type Low struct {
}

func (object *Low) pull(wrapper *Fan) {
	wrapper.setState(new(Medium))
	//fmt.Println("object: ", reflect.TypeOf(object), " wrapper:", reflect.TypeOf(wrapper)) //for testing
	fmt.Println("Turning Medium Speed")
}
