package main

import (
	"fmt"
)

type Medium struct {
}

func (object *Medium) pull(wrapper *Fan) {
	wrapper.setState(new(High))
	//fmt.Println("object: ", reflect.TypeOf(object), " wrapper:", reflect.TypeOf(wrapper)) //for testing
	fmt.Println("Turning High Speed")
}
