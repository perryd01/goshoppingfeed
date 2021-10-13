package goshoppingfeed

import (
	"encoding/xml"
	"errors"
	"fmt"
)

type ArukeresoDataFeed struct {
	XMLName xml.Name `xml:"Products"`
	Rows    []*ArukeresoDataElement
}

// ArukeresoDataElement https://www.arukereso.hu/admin/Login#feed-format
type ArukeresoDataElement struct {
	XMLName                 xml.Name           `xml:"Product,omitempty" csv:"-"`
	Identifier              string             `xml:"Identifier,omitempty" csv:"identifier"`
	Manufacturer            string             `xml:"Manufacturer,omitempty" csv:"manufacturer"`
	Name                    string             `xml:"Name,omitempty" csv:"name"`
	ProductUrl              string             `xml:"ProductUrl,omitempty" csv:"producturl"`
	Price                   string             `xml:"Price,omitempty" csv:"price"`
	NetPrice                string             `xml:"NetPrice,omitempty" csv:"netprice"`
	ProductNumber           string             `xml:"ProductNumber,omitempty" csv:"productnumber"`
	ImageUrl                string             `xml:"ImageUrl,omitempty" csv:"imageurl"`
	ImageUrl2               string             `xml:"ImageUrl2,omitempty"`
	Category                string             `xml:"Category,omitempty" csv:"category"`
	Description             string             `xml:"Description,omitempty" csv:"description"`
	DeliveryTime            string             `xml:"DeliveryTime,omitempty" csv:"deliverytime"`
	DeliveryCost            string             `xml:"DeliveryCost,omitempty" csv:"deliverycost"`
	EanCode                 string             `xml:"EanCode,omitempty" csv:"eancode"`
	EnergyEfficiencyAG      string             `xml:"EnergyEfficiencyA-G,omitempty"`
	EnergyLabelAG           string             `xml:"EnergyLabelAG,omitempty"`
	DetailedSpecificationAG string             `xml:"DetailedSpecificationA-G,omitempty"`
	Color                   string             `xml:"Color,omitempty"`
	Size                    string             `xml:"Size,omitempty"`
	GroupId                 string             `xml:"GroupId,omitempty"`
	BasketDisabled          ConvertibleBoolean `xml:"BasketDisabled,omitempty"`
	FreeDelivery            ConvertibleBoolean `xml:"FreeDelivery,omitempty"`
	Oversize                ConvertibleBoolean `xml:"Oversize,omitempty"`
	PPPMaxQuantity          uint32             `xml:"PPPMaxQuantity,omitempty"`
	MPLPPMaxQuantity        uint32             `xml:"MPLPPMaxQuantity,omitempty"`
	MPLCSAMaxQuantity       uint32             `xml:"MPLCSAMaxQuantity,omitempty"`
	FOXPOSTMaxQuantity      uint32             `xml:"FOXPOSTMaxQuantity,omitempty"`
	GLSMaxQuantity          uint32             `xml:"GLSMaxQuantity,omitempty"`
	DeliveryCostByProduct   int                `xml:"DeliveryCostByProduct,omitempty" csv:"deliverycostbyproduct,omitempty"`
	Warranty                uint8              `xml:"Warranty,omitempty" csv:"warranty,omitempty"`
	//Weight ArukeresoWeight `xml:"Weight,omitempty"`

	Attributes struct { // TODO custom marshal?
		Attribute []struct {
			AttributeName  string `xml:"Attribute_name"`
			AttributeValue string `xml:"Attribute_value"`
		} `xml:"Attribute"`
	} `xml:"Attributes"`
	MaxCPCMultiplier string `xml:"MaxCPCMultiplier"`
}

type ConvertibleBoolean bool

func (b *ConvertibleBoolean) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var asString string
	err := d.DecodeElement(&asString, &start)
	if err != nil {
		return err
	}
	if asString == "1" || asString == "true" {
		*b = true
	} else if asString == "0" || asString == "false" {
		*b = false
	} else {
		return errors.New(fmt.Sprintf("Boolean unmarshal error: invalid input %s", asString))
	}
	return nil
}

func (b ConvertibleBoolean) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if b {
		return e.EncodeElement("1", start)
	} else {
		return e.EncodeElement("0", start)
	}
}
