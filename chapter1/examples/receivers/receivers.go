package chapter1

import "fmt"

type pointer struct {
	id    string
	value int
}

// read will return the current state of the struct sent to value receiver.
// It will also update the state of the original struct
func (p *pointer) read() string {
	return fmt.Sprintf("%s: %d", p.id, p.value)
}

// write will update the current state of the struct sent to value receiver
// It will also update the state of the original struct
func (p *pointer) write(id string, value int) {
	p.id = id
	p.value = value
}

// reset will reset the current state of the struct sent to value receiver
// It will also update the state of the original struct
func (p *pointer) reset() {
	p.id = ""
	p.value = 0
}

type value struct {
	id    string
	value int
}

// read will return the current state of the struct sent to value receiver.
func (v value) read() string {
	return fmt.Sprintf("%s: %d", v.id, v.value)
}

// write will update the current state of the struct sent to value receiver and also return it.
// The state of the original value receiver will not change because of this update.
func (v value) write(id string, value int) value {
	v.id = id
	v.value = value
	return v
}

// reset will reset the current state of the struct sent to value receiver and also return it.
// The state of the original value receiver will not change because of this update.
func (v value) reset() value {
	v.id = ""
	v.value = 0
	return v
}
