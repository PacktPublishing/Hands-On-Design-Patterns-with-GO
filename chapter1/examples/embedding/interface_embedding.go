package chapter1

type interfaceA interface {
	name() string
}

type interfaceB interface {
	hasPhone() bool
}

// structs need to implement name() & hasPhone()
// to be classified as type interfaceC
type interfaceC interface {
	interfaceA
	interfaceB
}

type structA struct{}

// structA satisfies the interface interfaceA
func (structA) name() string {
	return "A"
}

type structB struct{}

// structB satisfies the interface interfaceB
func (structB) hasPhone() bool {
	return false
}

type structC struct{}

// structC satisfies the interface interfaceA
func (structC) name() string {
	return "C"
}

// structC satisfies the interface interfaceB
func (structC) hasPhone() bool {
	return true
}

// NOTE: structC satisfies interface interfaceC
// because it has methods hasPhone() & name() defined on it.
