package main

type State interface {
	pull(wrapper *Fan)
}
