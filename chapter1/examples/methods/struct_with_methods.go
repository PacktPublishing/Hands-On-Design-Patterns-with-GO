package chapter1

import "fmt"

// SafeVoucher defines all interface methods
// we want to use to get and set values on the Voucher
type SafeVoucher interface {
	ID() string
	ValidForDays() int
	HasExpired() bool
	DecrementValidDays() error
}

// Voucher has two fields id & validForDays
// but the fields are not accessible outside of the package chapter1.
type Voucher struct {
	id           string
	validForDays int
}

// ID returns the id of Voucher through a query method.
func (v *Voucher) ID() string {
	return v.id
}

// ValidForDays returns the remaining days of validity for the Voucher.
func (v *Voucher) ValidForDays() int {
	return v.validForDays
}

// HasExpired returns true if Voucher.validForDays is 0.
func (v *Voucher) HasExpired() bool {
	return v.ValidForDays() <= 0
}

// DecrementValidDays decreases the Voucher validity by one day.
// If the voucher has expired, then the method returns an error.
func (v *Voucher) DecrementValidDays() (err error) {
	if !v.HasExpired() {
		v.validForDays--
	} else {
		err = fmt.Errorf("The voucher has expired! Unable to decrement further")
	}
	return
}

// NewVoucher allows us to instantiate a Voucher instance
// whose fields are not visible outside of package for instantiation purpose.
func NewVoucher(id string, days int) *Voucher {
	return &Voucher{id: id, validForDays: days}
}

// DigitalVoucher also implements SafeVoucher interface.
type DigitalVoucher struct {
	id           string
	validForDays int
	refreshIndex int
}

// ID returns the id of Voucher through a query method.
func (dv *DigitalVoucher) ID() string {
	return dv.id
}

// ValidForDays returns the remaining days of validity for the Voucher.
func (dv *DigitalVoucher) ValidForDays() int {
	return dv.validForDays
}

// HasExpired returns true if Voucher.validForDays is 0.
func (dv *DigitalVoucher) HasExpired() bool {
	return dv.ValidForDays() <= 0
}

// DecrementValidDays decreases the Voucher validity by one day.
// NOTE: This method will never return an error
// but we need to comply with the `SafeVoucher` interface.
func (dv *DigitalVoucher) DecrementValidDays() error {
	if !dv.HasExpired() {
		dv.validForDays--
	} else {
		dv.RefreshVoucher()
	}
	return nil
}

// RefreshVoucher updates the DigitalVoucher with new id and validForDays
func (dv *DigitalVoucher) RefreshVoucher() {
	dv.id = fmt.Sprintf("DigitalVoucher-%2d", dv.refreshIndex)
	dv.refreshIndex++
	dv.validForDays = 2
}

// NewDigitalVoucher creates an instance of DigitalVoucher with
func NewDigitalVoucher(id string, days int) *DigitalVoucher {
	return &DigitalVoucher{id: id, validForDays: days, refreshIndex: 1}
}
