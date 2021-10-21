package google

import (
	"fmt"
	"github.com/gocarina/gocsv"
	"net/url"
	"strconv"
	"strings"
)

type MultipleGUrl []*GUrl

type GUrl struct {
	*url.URL
}

func (gurl *GUrl) UnmarshalCSV(csv string) error {
	if csv == "" {
		return nil
	}
	url, err := url.Parse(csv)
	if err == nil {
		gurl.URL = url
	}
	return err
}
func (gurl *GUrl) MarshalCSV() (string, error) {
	return gurl.String(), nil
}

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
	Country     *string                `xml:"g:country,omitempty"`
	Area        *string                `xml:"g:region,omitempty"`
	Rate        *float64               `xml:"g:rate,omitempty"`
	ShippingTax ConvertibleBooleanType `xml:"g:tax_ship,omitempty"`
}

func (tt *TaxType) UnmarshalCSV(csv string) error {
	if csv == "" {
		return nil
	}
	splitted := strings.Split(csv, ":")
	var boolean ConvertibleBooleanType
	err := gocsv.UnmarshalString(splitted[3], &boolean)
	if err != nil {
		return err
	}
	rate, err := strconv.ParseFloat(splitted[2], 32)
	if err != nil {
		return err
	}
	*tt.Country = splitted[0]
	*tt.Area = splitted[1]
	*tt.Rate = rate
	tt.ShippingTax = boolean
	return nil
}

func (tt *TaxType) MarshalCSV() (string, error) {
	var country string = *tt.Country
	var area string = *tt.Area
	var rate string = fmt.Sprintf("%.2f", *tt.Rate)
	shippingtax, _ := tt.ShippingTax.MarshalCSV()
	return strings.Join([]string{country, area, rate, shippingtax}, ":"), nil
}

// TODO unmarshal and marshal
type ShippingType struct {
	Country           *string       `xml:"g:country,omitempty"`
	Region            *string       `xml:"g:region,omitempty"`
	PostalCode        *string       `xml:"g:postal_code,omitempty"`
	LocationID        *string       `xml:"g:location_id,omitempty"`
	LocationGroupName *string       `xml:"g:location_group_name,omitempty"`
	Service           *string       `xml:"g:service,omitempty"`
	Price             *NumericValue `xml:"g:price,omitempty"`
	MinHandlingTime   *string       `xml:"g:min_handling_time,omitempty"`
	MaxHandlingTime   *string       `xml:"g:max_handling_time,omitempty"`
	MinTransitTime    *string       `xml:"g:min_transit_time,omitempty"`
	MaxTransitTime    *string       `xml:"g:max_transit_time,omitempty"`
}

type MultipleProductDetailType []*ProductDetailType

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
