package google

import (
	"fmt"
	"github.com/gocarina/gocsv"
	"github.com/shopspring/decimal"
	"net/url"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

type MultipleGUrl struct {
	Urls []*GUrl
}

func (mg *MultipleGUrl) UnmarshalCSV(csv string) error {
	if csv == "" {
		return nil
	}
	split := strings.Split(csv, ",")
	for _, s := range split {
		var u *GUrl
		err := u.UnmarshalCSV(s)
		if err != nil {
			return err
		}
		mg.Urls = append(mg.Urls, u)
	}
	return nil

}
func (mg *MultipleGUrl) MarshalCSV() (string, error) {
	if len(mg.Urls) == 0 {
		return "", nil
	}

	var s string
	for i, gUrl := range mg.Urls {
		csv, err := gUrl.MarshalCSV()
		if err != nil {
			return "", err
		}
		s += csv

		if i < len(mg.Urls)-1 {
			s += ","
		}
	}
	return s, nil
}

type GUrl struct {
	*url.URL
}

func (gurl *GUrl) UnmarshalCSV(csv string) error {
	if csv == "" {
		return nil
	}
	newUrl, err := url.Parse(csv)
	if err == nil {
		gurl.URL = newUrl
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

func (cat *ProductCategoryType) UnmarshalCSV(csv string) error {
	if csv == "" {
		return nil
	}
	r, _ := regexp.Compile("^[0-9]+$")

	isNumeric := r.MatchString(csv)
	if isNumeric {
		parseUint, err := strconv.ParseUint(csv, 10, 64)
		if err != nil {
			return err
		}
		cat.ID = &parseUint
	} else {
		split := strings.Split(csv, ">")
		for _, c := range split {
			trim := strings.TrimSpace(c)
			cat.Categories = append(cat.Categories, &trim)
		}
	}
	return nil
}
func (cat *ProductCategoryType) MarshalCSV() (string, error) {
	var s string
	if cat.ID != nil {
		s = strconv.FormatUint(*cat.ID, 10)
		return s, nil
	}
	if len(cat.Categories) != 0 {
		for i, v := range cat.Categories {
			s += *v
			if i < len(cat.Categories)-1 {
				s += " > "
			}
		}
		return s, nil
	}
	return "", nil
}

type LoyaltyPointsType struct {
	Name        *string
	PointsValue *uint64  // TODO decimal
	Ratio       *float32 // TODO decimal
}

func (lpt *LoyaltyPointsType) UnmarshalCSV(csv string) error {
	if csv == "" {
		return nil
	}
	split := strings.Split(csv, ":")
	*lpt.Name = split[0]
	points, err := strconv.ParseUint(split[1], 10, 64)
	if err != nil {
		return err
	}
	*lpt.PointsValue = points
	ratio, err := strconv.ParseFloat(split[2], 32)
	if err != nil {
		return err
	}
	*lpt.Ratio = float32(ratio)
	return nil
}
func (lpt *LoyaltyPointsType) MarshalCSV() (string, error) {
	points := strconv.FormatUint(*lpt.PointsValue, 10)
	ratio := fmt.Sprintf("%0.2f", *lpt.Ratio)
	return fmt.Sprintf("%s:%s:%s", *lpt.Name, points, ratio), nil
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

func (it *InstallmentType) UnmarshalCSV(csv string) error {
	if csv == "" {
		return nil
	}
	split := strings.Split(csv, ",")
	var nv NumericValue
	err := nv.UnmarshalCSV(split[1])
	if err != nil {
		return err
	}
	months, err := strconv.ParseUint(split[0], 10, 64)
	if err != nil {
		return err
	}
	*it.Months = months
	*it.Amount = nv
	return nil
}
func (it *InstallmentType) MarshalCSV() (string, error) {
	months := strconv.FormatUint(*it.Months, 10)
	amount, _ := it.Amount.MarshalCSV()
	return fmt.Sprintf("%s,%s", months, amount), nil
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

// TODO
type MultipleProductDetailType []*ProductDetailType

type ProductDetailType struct {
	Section        *string
	AttributeName  *string
	AttributeValue *string
}

func (pdt *ProductDetailType) UnmarshalCSV(csv string) error {
	if csv == "" {
		return nil
	}
	split := strings.Split(csv, ":")
	*pdt.Section = split[0]
	*pdt.AttributeName = split[1]
	*pdt.AttributeValue = split[2]
	return nil
}
func (pdt *ProductDetailType) MarshalCSV() (string, error) {
	return fmt.Sprintf("%s:%s:%s", *pdt.Section, *pdt.AttributeName, *pdt.AttributeValue), nil
}

type NumericValue struct {
	Value     decimal.NullDecimal
	Dimension string
	HasSpace  bool
}

func (nv *NumericValue) MarshalCSV() (string, error) {
	text, err := nv.Value.MarshalText()
	if err != nil {
		return "", err
	}
	if nv.HasSpace {
		return fmt.Sprintf("%s %s", text, nv.Dimension), nil
	} else {
		return fmt.Sprintf("%s%s", text, nv.Dimension), nil
	}
}

func (nv *NumericValue) UnmarshalCSV(csv string) error {
	if csv == "" {
		return nil
	}
	nv.HasSpace = strings.Contains(csv, " ")
	var fSlice []string
	var numPart string
	var dimension string
	if nv.HasSpace {
		fSlice = strings.Split(csv, " ")
		numPart = fSlice[0]
		dimension = fSlice[1]
	} else {
		whereToSplit := 0
		for i := 0; i < len(csv); i++ {
			if unicode.IsDigit(rune(csv[i])) {
				whereToSplit = i
			}
		}
		numPart = csv[:whereToSplit]
		dimension = csv[whereToSplit : len(csv)-1]
	}

	value, _ := decimal.NewFromString(numPart)
	nv.Value = decimal.NewNullDecimal(value)
	nv.Dimension = dimension

	return nil
}
