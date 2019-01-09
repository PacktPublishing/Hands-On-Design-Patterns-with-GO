package chapter1

import (
	"fmt"
	"testing"
)

func TestFullName(t *testing.T) {
	expectedFullName := "Alice A."

	fcnFullName := GetFullName(Alice)
	mthdFullName := Alice.FullName()

	correctFcnFullName := fcnFullName == expectedFullName
	correctMthdFullName := mthdFullName == expectedFullName

	if correctFcnFullName && correctMthdFullName {
		msg := fmt.Sprintf("Correct Full Name returned by"+
			"GetFullName & Person.FullName: %s", expectedFullName)
		t.Log(msg)
	} else {
		msg := fmt.Sprintf(
			"Unexpected Full Name returned.\nExpected: %s. "+
				"vs.\nfcn: %s, mthd: %s",
			expectedFullName,
			fcnFullName,
			mthdFullName,
		)
		t.Error(msg)
	}
}

func TestChangeEmailID(t *testing.T) {
	newEmailID := "new_bob@email.com"

	oldEmailID := Bob.EmailID

	ChangeEmailID(&Bob, newEmailID)
	if Bob.EmailID == newEmailID {
		t.Log("ChangeEmailID returned correct Email ID: " + Bob.EmailID)
	} else {
		t.Error("ChangeEmailID: Email ID not changed as expected with" + Bob.EmailID)
	}

	Bob.ChangeEmailID(oldEmailID)
	if Bob.EmailID == oldEmailID {
		t.Log("Bob.ChangeEmailID returned correct Email ID: " + Bob.EmailID)
	} else {
		t.Error("Bob.ChangeEmailID: Email ID not changed as expected" + Bob.EmailID)
	}
}
