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

type SubscriptionCostPeriodType string

const (
	Month SubscriptionCostPeriodType = "month"
	Year  SubscriptionCostPeriodType = "year"
)

// GDate YYYY-MM-DDThh:mmZ
type GDate time.Time

type AdsDestinationType string

const (
	ShoppingAds         AdsDestinationType = "Shopping_ads"
	BuyOnGoogleListings AdsDestinationType = "Buy_on_Google_listings"
	DisplayAds          AdsDestinationType = "Display_ads"
	LocalInventoryAds   AdsDestinationType = "Local_inventory_ads"
	FreeListings        AdsDestinationType = "Free_listings"
	FreeLocalListings   AdsDestinationType = "Free_local_listings"
)

type SizeType string

const (
	Regular   SizeType = "regular"
	Petite    SizeType = "petite"
	Maternity SizeType = "maternity"
	Big       SizeType = "big"
	Tall      SizeType = "tall"
	Plus      SizeType = "plus"
)

type GenderType string

const (
	Male   GenderType = "male"
	Female GenderType = "female"
	Unisex GenderType = "unisex"
)

type AgeGroupType string

const (
	Newborn AgeGroupType = "newborn"
	Infant  AgeGroupType = "infant"
	Toddler AgeGroupType = "toddler"
	Kids    AgeGroupType = "kids"
	Adult   AgeGroupType = "adult"
)

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
