package main

import (
	"fmt"
)

type High struct {
}

func (object *High) pull(wrapper *Fan) {
	wrapper.setState(new(Off))
	//fmt.Println("object: ", reflect.TypeOf(object), " wrapper:", reflect.TypeOf(wrapper)) //for testing
	fmt.Println("Turning Off")
}
