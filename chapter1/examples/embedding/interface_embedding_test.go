package chapter1

import (
	"fmt"
	"testing"
)

func TestInterfaceA(t *testing.T) {
	interfaceTester := func(a interfaceA) {
		t.Log("returned name: " + a.name())
	}

	for _, tc := range []struct {
		tcName string
		a      interfaceA
	}{
		{"structA", structA{}},
		{"structC", structC{}},
	} {
		tcName := fmt.Sprintf("What is the name of %s?", tc.tcName)
		t.Run(tcName, func(t *testing.T) {
			interfaceTester(tc.a)
		})
	}
}
func TestInterfaceB(t *testing.T) {
	interfaceTester := func(b interfaceB) {
		if b.hasPhone() {
			t.Log("has a phone!")
		} else {
			t.Log("does not have a phone!")
		}
	}

	for _, tc := range []struct {
		tcName string
		b      interfaceB
	}{
		{"structA", structB{}},
		{"structC", structC{}},
	} {
		tcName := fmt.Sprintf("Does %s have a phone?", tc.tcName)
		t.Run(tcName, func(t *testing.T) {
			interfaceTester(tc.b)
		})
	}
}

func TestInterfaceC(t *testing.T) {
	interfaceTester := func(c interfaceC) {
		t.Log("interfaceC returned name: " + c.name())
		if c.hasPhone() {
			t.Log("interfaceC has a phone!")
		} else {
			t.Error("interfaceC does not have a phone!")
		}
	}

	interfaceTester(structC{})
}
