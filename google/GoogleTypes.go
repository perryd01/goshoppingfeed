package google

import (
	"strconv"
	"strings"
)

type ProductCategoryType struct {
	ID         *uint64
	Categories []*string
}

type LoyaltyPointsType struct {
	Name        *string
	PointsValue *uint64
	Ratio       *float32
}

type SubscriptionCostType struct {
	Period       *SubscriptionCostPeriodType
	PeriodLength *uint64
	Amount       *NumericValue
}

type InstallmentType struct {
	Months *uint64
	Amount *NumericValue
}

type TaxType struct {
	Country     *string
	Area        *string
	Rate        *float32
	ShippingTax ConvertibleBooleanType
}

// TODO unmarshal and marshal
type ShippingType struct {
	Country           *string
	Region            *string
	PostalCode        *string
	LocationID        *string
	LocationGroupName *string
	Service           *string
	Price             *NumericValue
	MinHandlingTime   *string
	MaxHandlingTime   *string
	MinTransitTime    *string
	MaxTransitTime    *string
}

// TODO marshal and unmarshal
type ProductDetailType struct {
	Section        *string
	AttributeName  *string
	AttributeValue *string
}

type NumericValue struct {
	Int       *uint64
	Fraction  *uint64
	Dimension *string
}

func (nv *NumericValue) UnmarshalCSV(csv string) error {
	if csv == "" {
		return nil
	}
	firstSlice := strings.Split(csv, " ")
	hasFraction := strings.Contains(csv, ".")
	numPart := firstSlice[0]
	nv.Dimension = &firstSlice[1]

	if hasFraction {
		secondSlice := strings.Split(numPart, ".")
		integer, err := strconv.ParseUint(secondSlice[0], 10, 64)
		if err != nil {
			return err
		}
		*nv.Int = integer
		fraction, err := strconv.ParseUint(secondSlice[1], 10, 64)
		if err != nil {
			return err
		}
		*nv.Fraction = fraction
	} else {
		integer, err := strconv.ParseUint(firstSlice[0], 10, 64)
		if err != nil {
			return err
		}
		*nv.Int = integer
	}
	return nil
}
