package main

type Fan struct {
	currentState State
}

func (object *Fan) Initialize() {
	object.currentState = new(Off)
}

func (object *Fan) setState(s State) {
	object.currentState = s
}

func (object *Fan) pull() {
	object.currentState.pull(object)
}
