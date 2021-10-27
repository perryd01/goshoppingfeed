package google

import (
	"errors"
	"fmt"
	"strconv"
	"time"
)

type EnergyEfficiencyClassType string

const (
	Appp EnergyEfficiencyClassType = "A+++"
	App  EnergyEfficiencyClassType = "A++"
	Ap   EnergyEfficiencyClassType = "A+"
	A    EnergyEfficiencyClassType = "A"
	B    EnergyEfficiencyClassType = "B"
	C    EnergyEfficiencyClassType = "C"
	D    EnergyEfficiencyClassType = "D"
	E    EnergyEfficiencyClassType = "E"
	F    EnergyEfficiencyClassType = "F"
	G    EnergyEfficiencyClassType = "G"
)

func (eect *EnergyEfficiencyClassType) UnmarshalCSV(csv string) error {
	if csv == "" {
		return nil
	}
	newEect := EnergyEfficiencyClassType(csv)
	*eect = newEect
	return nil
}

func (eect *EnergyEfficiencyClassType) MarshalCSV() (string, error) {
	return string(*eect), nil
}

type SubscriptionCostPeriodType string

const (
	Month SubscriptionCostPeriodType = "month"
	Year  SubscriptionCostPeriodType = "year"
)

func (scpt *SubscriptionCostPeriodType) UnmarshalCSV(csv string) error {
	if csv == "" {
		return nil
	}
	newScpt := SubscriptionCostPeriodType(csv)
	*scpt = newScpt
	return nil
}

func (scpt *SubscriptionCostPeriodType) MarshalCSV() (string, error) {
	return string(*scpt), nil
}

// GDate YYYY-MM-DDThh:mmZ / ISO 8601 / RFC3339
type GDate struct {
	time.Time
}

func (gt *GDate) UnmarshalCSV(csv string) error {
	if csv == "" {
		return nil
	}
	t, err := time.Parse(time.RFC3339, csv)
	if err != nil {
		return err
	}
	gt.Time = t
	return nil
}
func (gt *GDate) MarshalCSV() (string, error) {
	return gt.Format(time.RFC3339), nil
}

type AdsDestinationType string

const (
	ShoppingAds         AdsDestinationType = "Shopping_ads"
	BuyOnGoogleListings AdsDestinationType = "Buy_on_Google_listings"
	DisplayAds          AdsDestinationType = "Display_ads"
	LocalInventoryAds   AdsDestinationType = "Local_inventory_ads"
	FreeListings        AdsDestinationType = "Free_listings"
	FreeLocalListings   AdsDestinationType = "Free_local_listings"
)

func (adt *AdsDestinationType) MarshalCSV() (string, error) {
	return string(*adt), nil
}

func (adt *AdsDestinationType) UnmarshalCSV(csv string) error {
	if csv == "" {
		return nil
	}
	newAdt := AdsDestinationType(csv)
	*adt = newAdt
	return nil
}

type SizeType string

const (
	Regular   SizeType = "regular"
	Petite    SizeType = "petite"
	Maternity SizeType = "maternity"
	Big       SizeType = "big"
	Tall      SizeType = "tall"
	Plus      SizeType = "plus"
)

func (st *SizeType) MarshalCSV() (string, error) {
	return string(*st), nil
}
func (st *SizeType) UnmarshalCSV(csv string) error {
	if csv == "" {
		return nil
	}
	newSt := SizeType(csv)
	*st = newSt
	return nil
}

type GenderType string

const (
	Male   GenderType = "male"
	Female GenderType = "female"
	Unisex GenderType = "unisex"
)

func (gt *GenderType) MarshalCSV() (string, error) {
	return string(*gt), nil
}

func (gt *GenderType) UnmarshalCSV(csv string) error {
	if csv == "" {
		return nil
	}
	newGt := GenderType(csv)
	*gt = newGt
	return nil
}

type AgeGroupType string

const (
	Newborn AgeGroupType = "newborn"
	Infant  AgeGroupType = "infant"
	Toddler AgeGroupType = "toddler"
	Kids    AgeGroupType = "kids"
	Adult   AgeGroupType = "adult"
)

func (agt *AgeGroupType) MarshalCSV() (string, error) {
	return string(*agt), nil
}
func (agt *AgeGroupType) UnmarshalCSV(csv string) error {
	if csv == "" {
		return nil
	}
	newAgt := AgeGroupType(csv)
	*agt = newAgt
	return nil
}

type ConvertibleBooleanType bool

func (b *ConvertibleBooleanType) MarshalCSV() (string, error) {
	return strconv.FormatBool(bool(*b)), nil
}

func (b *ConvertibleBooleanType) UnmarshalCSV(data []byte) error {
	asString := string(data)
	if asString == "1" || asString == "true" {
		*b = true
	} else if asString == "0" || asString == "false" {
		*b = false
	} else {
		return errors.New(fmt.Sprintf("Boolean unmarshal error: invalid input %s", asString))
	}
	return nil
}

type ConditionType string

const (
	New         ConditionType = "new"
	Refurbished ConditionType = "refurbished"
	Used        ConditionType = "used"
)

func (ct *ConditionType) MarshalCSV() (string, error) {
	return string(*ct), nil
}
func (ct *ConditionType) UnmarshalCSV(csv string) error {
	if csv == "" {
		return nil
	}

	newCt := ConditionType(csv)
	*ct = newCt
	return nil
}

type AvailabilityType string

const (
	InStock    AvailabilityType = "in_stock"
	OutOfStock AvailabilityType = "out_of_stock"
	Preorder   AvailabilityType = "preorder"
	Backorder  AvailabilityType = "backorder"
)

func (at *AvailabilityType) UnmarshalCSV(csv string) error {
	if csv == "" {
		return nil
	}
	var newAt = AvailabilityType(csv)
	*at = newAt
	return nil
}

func (at *AvailabilityType) MarshalCSV() (string, error) {
	return string(*at), nil
}
