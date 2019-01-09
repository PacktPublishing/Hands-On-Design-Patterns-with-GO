package chapter1

import (
	"reflect"
	"testing"
)

func TestVariadic(t *testing.T) {
	single := variadic(1)
	if reflect.DeepEqual(single, stats{length: 1, sum: 1, average: 1}) {
		t.Logf("variadic returned expected stats for variadic(1): %#v", single)
	} else {
		t.Errorf("variadic returned unexpected stats for variadic(1): %#v", single)
	}

	double := variadic(1, 2)
	if reflect.DeepEqual(double, stats{length: 2, sum: 3, average: 1.5}) {
		t.Logf("variadic returned expected stats for variadic(1, 2): %#v", double)
	} else {
		t.Errorf("variadic returned unexpected stats for variadic(1, 2): %#v", double)
	}

	none := variadic()
	if reflect.DeepEqual(none, stats{length: 0, sum: 0, average: 0}) {
		t.Logf("variadic returned expected stats for variadic(): %#v", none)
	} else {
		t.Errorf("variadic returned unexpected stats for variadic(): %#v", none)
	}

	list := variadic([]int{1, 2, 3, 4, 5}...)
	if reflect.DeepEqual(list, stats{length: 5, sum: 15, average: 3}) {
		t.Logf("variadic returned expected stats for variadic(): %#v", list)
	} else {
		t.Errorf("variadic returned unexpected stats for variadic(): %#v", list)
	}
}
