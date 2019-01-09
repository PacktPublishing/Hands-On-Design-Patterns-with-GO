package chapter1

import "testing"

func checkStructState(t *testing.T, actual, expected string) {
	t.Helper()

	if actual != expected {
		t.Errorf("Pointer in inconsistent state: %s != %s", actual, expected)
	} else {
		t.Log("Pointer in consistent state: " + actual)
	}
}

func TestPointerReceivers(t *testing.T) {
	ptr := pointer{"new-pointer", 100}                  // instantiate new struct
	checkStructState(t, ptr.read(), "new-pointer: 100") // check if state is as expected

	ptr.write("garbage", -100)                       // use value.write to update the state of struct
	checkStructState(t, ptr.read(), "garbage: -100") // check if original struct was updated with the write values

	ptr.reset()                            // reset the original struct
	checkStructState(t, ptr.read(), ": 0") // check if original struct has reset values
}

func TestValueReceivers(t *testing.T) {
	v := value{"new-value", 100}                    // instantiate new struct
	checkStructState(t, v.read(), "new-value: 100") // check if state is as expected

	copyOfV := v.write("ignored", 1234)             // use value.write to update the state of struct
	checkStructState(t, v.read(), "new-value: 100") // check if original struct was modified

	// check if the state of returned struct has the values we used to call value.write
	checkStructState(t, copyOfV.read(), "ignored: 1234")

	copyOfV = v.reset()                             // reset the original struct
	checkStructState(t, v.read(), "new-value: 100") // check if original struct was reset

	// check if the values on returned struct are reset
	checkStructState(t, copyOfV.read(), ": 0")
}
