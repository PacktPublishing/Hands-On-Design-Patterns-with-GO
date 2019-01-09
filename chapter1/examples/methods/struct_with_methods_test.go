package chapter1

import (
	"testing"
)

func TestSafeVoucherWithValidVoucher(t *testing.T) {
	days := 1
	voucherID := "voucher1"

	// We declare the variable `voucher` to be of type `SafeVoucher`,
	// which is an interface instead of the concrete type `Voucher`
	var voucher SafeVoucher = NewVoucher(voucherID, days)

	// Working against the interface method, SafeVoucher.ID()
	if voucher.ID() != voucherID {
		t.Error("Voucher.ID() returned unexpected Id: " + voucher.ID())
	}

	// Working against the interface method, SafeVoucher.HasExpired()
	if voucher.HasExpired() {
		t.Error("Voucher has already expired!")
	}

	// Working against the interface method, SafeVoucher.ValidForDays()
	if voucher.ValidForDays() != days {
		t.Errorf(
			"Voucher has unexpected number of days: %d",
			voucher.ValidForDays(),
		)
	}

	// Working against the interface method, SafeVoucher.DecrementValidDays()
	if err := voucher.DecrementValidDays(); err != nil {
		t.Errorf("Unexpected error while decrementing days(%d): %s",
			voucher.ValidForDays(),
			err,
		)
	}

	if err := voucher.DecrementValidDays(); err != nil {
		if err.Error() != "The voucher has expired! Unable to decrement further" {
			t.Errorf("Unexpected error while decrementing days: %s", err)
		} else {
			t.Log("Voucher.DecrementValidDays() returned error as expected")
		}
	} else {
		t.Errorf("Expected error while decrementing!")
	}
}

func TestSafeVoucherForExpiredVoucher(t *testing.T) {
	voucherID := "voucher0"
	days := 0

	var voucher = NewVoucher(voucherID, days)

	if voucher.ID() != voucherID {
		t.Error("Voucher.ID() returned unexpected Id: " + voucher.ID())
	}

	if !voucher.HasExpired() {
		t.Error("Voucher should have expired!")
	}

	if voucher.ValidForDays() != days {
		t.Errorf("Voucher has unexpected number of days: %d", voucher.ValidForDays())
	}

	if err := voucher.DecrementValidDays(); err != nil {
		if err.Error() != "The voucher has expired! Unable to decrement further" {
			t.Errorf("Unexpected error while decrementing days: %s", err)
		} else {
			t.Log("Voucher.DecrementValidDays() returned error as expected for expired Voucher")
		}
	} else {
		t.Errorf("Expected error while decrementing!")
	}
}

func TestSafeVoucherWithDigitalVoucher(t *testing.T) {
	days := 1
	voucherID := "voucher1"

	// We again declare voucher as type `SafeVoucher` but use `DigitalVoucher`
	var voucher SafeVoucher = NewDigitalVoucher(voucherID, days)

	// Working against the interface method, SafeVoucher.ID()
	if voucher.ID() != voucherID {
		t.Error("Voucher.ID() returned unexpected Id: " + voucher.ID())
	}

	// Working against the interface method, SafeVoucher.HasExpired()
	if voucher.HasExpired() {
		t.Error("Voucher has already expired!")
	}

	// Working against the interface method, SafeVoucher.ValidForDays()
	if voucher.ValidForDays() != days {
		t.Errorf(
			"Voucher has unexpected number of days: %d",
			voucher.ValidForDays(),
		)
	}

	// Working against the interface method, SafeVoucher.DecrementValidDays()
	if err := voucher.DecrementValidDays(); err != nil {
		t.Errorf("Unexpected error while decrementing days!"+
			" Error should not occur!\nError: %s", err)
	}

	if err := voucher.DecrementValidDays(); err != nil {
		t.Errorf("Unexpected error while decrementing days!"+
			" Error should not occur!\nError: %s", err)
	} else if voucher.ID() != "DigitalVoucher-01" &&
		voucher.ValidForDays() != 2 {
		t.Errorf(
			"Unexpected values returned for either/or"+
				" - voucher.ID(): %s, voucher.ValidForDays(): %d",
			voucher.ID(),
			voucher.ValidForDays(),
		)
	}
}
