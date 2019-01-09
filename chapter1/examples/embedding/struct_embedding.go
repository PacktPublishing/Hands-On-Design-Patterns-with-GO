package chapter1

type struct1 struct{}

func (struct1) method1() string {
	return "s.struct1.method1()"
}

func (struct1) common() string {
	return "s.struct1.common()"
}

type struct2 struct{}

func (struct2) method2() string {
	return "s.struct2.method2()"
}
func (struct2) common() string {
	return "s.struct2.common()"
}

type struct3 struct {
	struct1
	struct2
}

// struct1 & struct2 both provide the method common()
// Go compiler will fail to build the file if we try to
// call method struct.common() In order to avoid build fail,
//  we can explicitly forward the call to one of the embedded structs
// or provide a method on struct3 with same method name.
func (s struct3) common() string {
	return s.struct2.common()
}
