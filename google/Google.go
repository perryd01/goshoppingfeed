package google

type DSAFeed struct {
	Rows []*DSAElement
}

type DSAElement struct {
	URL         string `csv:"Page URL" json:"Page URL"`
	CustomLabel string `csv:"Custom label,omitempty" json:"Custom label"`
}

type ShoppingFeed struct {
	Rows []*ShoppingElement
}

// ShoppingElement  https://support.google.com/merchants/answer/7052112?hl=en
type ShoppingElement struct {
	ID                              string            `csv:"id,omitempty" json:"id" xml:"id" validate:"required,min=1,max=50"`           // Your product’s unique identifier
	Title                           string            `csv:"title,omitempty" json:"title" xml:"title" validate:"required,min=1,max=150"` // Your product’s name
	Description                     string            `csv:"description,omitempty" validate:"required,min=1,max=5000"`                   // Your product’s description
	Link                            string            `csv:"link,omitempty" validate:"required,min=1,max=2000"`                          // Your product’s landing page
	ImageLink                       string            `csv:"image_link,omitempty" validate:"required,min=1,max=2000"`                    // The URL of your product’s main image
	AdditionalImageLink             *string           `csv:"additional_image_link,omitempty"`                                            // The URL of an additional image for your product
	MobileLink                      *string           `csv:"mobile_link,omitempty" validate:"min=1,max=2000"`
	Availability                    *AvailabilityType `csv:"availability,omitempty" validate:"required"`
	CostOfGoodsSold                 *NumericValue     `csv:"cost_of_goods_sold,omitempty"`
	AvailabilityDate                *GDate            `csv:"availability_date,omitempty"`
	ExpirationDate                  *GDate            `csv:"expiration_date,omitempty"`
	Price                           *NumericValue     `csv:"price,omitempty" validate:"required"`
	SalePrice                       *NumericValue     `csv:"sale_price,omitempty"`
	SalePriceEffectiveDate          *GDate
	UnitPricingMeasure              *NumericValue
	UnitPricingBaseMeasure          *NumericValue
	Installment                     *InstallmentType
	SubscriptionCost                *SubscriptionCostType
	LoyaltyPoints                   *LoyaltyPointsType
	GoogleProductCategory           *ProductCategoryType `csv:"google_product_category,omitempty"`
	ProductType                     *ProductCategoryType `csv:"product_type,omitempty"`
	Brand                           *string              `csv:"brand,omitempty"`
	GTIN                            *string              `csv:"gtin,omitempty"`
	MPN                             *string              `csv:"mpn,omitempty"`
	IdentifierExists                *ConvertibleBooleanType
	Condition                       *ConditionType
	Adult                           *ConvertibleBooleanType
	Multipack                       *uint64
	Bundle                          *ConvertibleBooleanType
	EnergyEfficiencyClass           *EnergyEfficiencyClassType
	MinimumEnergyEfficiencyClass    *EnergyEfficiencyClassType
	MaximumEnergyEfficiencyClass    *EnergyEfficiencyClassType
	AgeGroup                        *AgeGroupType
	Color                           *string
	Gender                          *GenderType
	Material                        *string
	Pattern                         *string
	Size                            *string
	SizeType                        *SizeType
	SizeSystem                      *string `validate:"max=3"`
	ItemGroupID                     *string `validate:"max=50"`
	ProductLength                   *NumericValue
	ProductWidth                    *NumericValue
	ProductHeight                   *NumericValue
	ProductWeight                   *NumericValue
	ProductDetail                   []*ProductDetailType
	ProductHighlight                *string `validate:"max=150"`
	AdsRedirect                     *string `validate:"max=2000"`
	CustomLabel0                    *string `validate:"max=100"`
	CustomLabel1                    *string `validate:"max=100"`
	CustomLabel2                    *string `validate:"max=100"`
	CustomLabel3                    *string `validate:"max=100"`
	CustomLabel4                    *string `validate:"max=100"`
	PromotionID                     *string `validate:"max=50"`
	ExcludedDestination             *AdsDestinationType
	IncludedDestination             *AdsDestinationType
	ExcludedCountriesForShoppingAds *string `validate:"max=2"`
	Shipping                        *ShippingType
	ShippingLabel                   *string `validate:"max=100"`
	ShippingWeight                  *NumericValue
	ShippingLength                  *NumericValue
	ShippingWidth                   *NumericValue
	ShippingHeight                  *NumericValue
	ShipsFromCountry                *string
	TransitTimeLabel                *string `validate:"max=100"`
	MinHandlingTime                 *uint64
	MaxHandlingTime                 *uint64
	Tax                             *TaxType
	TaxCategory                     *string `validate:"max=100"`
}
