package chapter1

import "testing"

func TestEmbeddedStruct(t *testing.T) {
	s := struct3{}

	if s.method1() == "s.struct1.method1()" {
		t.Log("s.method1() correctly forwarded method call to s.struct1.method1()")
	} else {
		t.Errorf("s.method1() forwarded call incorrectly! %s", s.method1())
	}

	if s.method2() == "s.struct2.method2()" {
		t.Log("s.method2() correctly forwarded method call to s.struct2.method2()")
	} else {
		t.Errorf("s.method2() returned unexpected value: %s", s.method2())
	}

	if s.struct1.common() == "s.struct1.common()" {
		t.Log("s.struct1.common() correctly forwarded the method call to struct1.common()")
	} else {
		t.Errorf("Incorrect method call forwarding!: %s", s.struct1.common())
	}

	if s.common() == "s.struct2.common()" {
		t.Log("s.common() correctly forwarded the method call to struct2.common()")
	} else {
		t.Errorf("Incorrect method call forwarding!: %s", s.struct1.common())
	}
}
