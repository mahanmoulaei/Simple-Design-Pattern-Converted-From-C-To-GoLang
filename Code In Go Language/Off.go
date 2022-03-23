package main

import (
	"fmt"
)

type Off struct {
}

func (object *Off) pull(wrapper *Fan) {
	wrapper.setState(new(Low))
	//fmt.Println("object: ", reflect.TypeOf(object), " wrapper:", reflect.TypeOf(wrapper)) //for testing
	fmt.Println("Turning Low Speed")
}
