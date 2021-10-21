package google

type DSAFeed struct {
	Rows []*DSAElement
}

type DSAElement struct {
	URL         *GUrl  `csv:"Page URL" json:"Page URL"`
	CustomLabel string `csv:"Custom label,omitempty" json:"Custom label"`
}

type ShoppingFeed struct {
	Rows []*ShoppingElement
}

// ShoppingElement  https://support.google.com/merchants/answer/7052112?hl=en
type ShoppingElement struct {
	ID                              *string                    `csv:"id,omitempty" json:"id" xml:"id" validate:"required,min=1,max=50"`           // Your product’s unique identifier
	Title                           *string                    `csv:"title,omitempty" json:"title" xml:"title" validate:"required,min=1,max=150"` // Your product’s name
	Description                     *string                    `csv:"description,omitempty" validate:"required,min=1,max=5000"`                   // Your product’s description
	Link                            *GUrl                      `csv:"link,omitempty" validate:"required,min=1,max=2000"`                          // Your product’s landing page
	ImageLink                       *GUrl                      `csv:"image_link,omitempty" validate:"required,min=1,max=2000"`                    // The URL of your product’s main image
	AdditionalImageLink             *MultipleGUrl              `csv:"additional_image_link,omitempty"`                                            // The URL of an additional image for your product
	MobileLink                      *GUrl                      `csv:"mobile_link,omitempty" validate:"min=1,max=2000"`
	Availability                    *AvailabilityType          `csv:"availability,omitempty" validate:"required"`
	CostOfGoodsSold                 *NumericValue              `csv:"cost_of_goods_sold,omitempty"`
	AvailabilityDate                *GDate                     `csv:"availability_date,omitempty"`
	ExpirationDate                  *GDate                     `csv:"expiration_date,omitempty"`
	Price                           *NumericValue              `csv:"price,omitempty" validate:"required"`
	SalePrice                       *NumericValue              `csv:"sale_price,omitempty"`
	SalePriceEffectiveDate          *GDate                     `csv:"sale_price_effecitve_date,omitempty"`
	UnitPricingMeasure              *NumericValue              `csv:"unit_pricing_measure,omitempty"`
	UnitPricingBaseMeasure          *NumericValue              `csv:"unit_pricing_base_measure,omitempty"`
	Installment                     *InstallmentType           `csv:"installment,omitempty"`
	SubscriptionCost                *SubscriptionCostType      `csv:"subscription_cost,omitempty"`
	LoyaltyPoints                   *LoyaltyPointsType         `csv:"loyalty_points,omitempty"`
	GoogleProductCategory           *ProductCategoryType       `csv:"google_product_category,omitempty"`
	ProductType                     *ProductCategoryType       `csv:"product_type,omitempty"`
	Brand                           *string                    `csv:"brand,omitempty"`
	GTIN                            *string                    `csv:"gtin,omitempty"`
	MPN                             *string                    `csv:"mpn,omitempty"`
	IdentifierExists                *ConvertibleBooleanType    `csv:"identifier_exists,omitempty"`
	Condition                       *ConditionType             `csv:"condition,omitempty"`
	Adult                           *ConvertibleBooleanType    `csv:"adult,omitempty"`
	Multipack                       *uint64                    `csv:"multipack,omitempty"`
	Bundle                          *ConvertibleBooleanType    `csv:"is_bundle,omitempty"`
	EnergyEfficiencyClass           *EnergyEfficiencyClassType `csv:"energy_efficiency_class,omitempty"`
	MinimumEnergyEfficiencyClass    *EnergyEfficiencyClassType `csv:"min_energy_efficiency_class,omitempty"`
	MaximumEnergyEfficiencyClass    *EnergyEfficiencyClassType `csv:"max_energy_efficiency_class,omitempty"`
	AgeGroup                        *AgeGroupType              `csv:"age_group,omitempty"`
	Color                           *string                    `csv:"color,omitempty"`
	Gender                          *GenderType                `csv:"gender,omitempty"`
	Material                        *string                    `csv:"material,omitempty"`
	Pattern                         *string                    `csv:"pattern,omitempty"`
	Size                            *string                    `csv:"size,omitempty"`
	SizeType                        *SizeType                  `csv:"size_type,omitempty"`
	SizeSystem                      *string                    `csv:"size_system,omitempty" validate:"max=3"`
	ItemGroupID                     *string                    `csv:"item_group_id,omitempty" validate:"max=50"`
	ProductLength                   *NumericValue              `csv:"product_length,omitempty"`
	ProductWidth                    *NumericValue              `csv:"product_width,omitempty"`
	ProductHeight                   *NumericValue              `csv:"product_height,omitempty"`
	ProductWeight                   *NumericValue              `csv:"product_weight,omitempty"`
	ProductDetail                   *MultipleProductDetailType `csv:"product_detail,omitempty"`
	ProductHighlight                *string                    `csv:"product_highlight,omitempty" validate:"max=150"`
	AdsRedirect                     *GUrl                      `csv:"ads_redirect,omitempty"`
	CustomLabel0                    *string                    `csv:"custom_label_0,omitempty" validate:"max=100"`
	CustomLabel1                    *string                    `csv:"custom_label_1,omitempty" validate:"max=100"`
	CustomLabel2                    *string                    `csv:"custom_label_2,omitempty" validate:"max=100"`
	CustomLabel3                    *string                    `csv:"custom_label_3,omitempty" validate:"max=100"`
	CustomLabel4                    *string                    `csv:"custom_label_4,omitempty" validate:"max=100"`
	PromotionID                     *string                    `csv:"promotion_id,omitempty" validate:"max=50"`
	ExcludedDestination             *AdsDestinationType        `csv:"excluded_destination,omitempty"`
	IncludedDestination             *AdsDestinationType        `csv:"included_destination,omitempty"`
	ExcludedCountriesForShoppingAds *string                    `csv:"shopping_ads_excluded_country,omitempty" validate:"max=2"`
	Shipping                        *ShippingType              `csv:"shipping,omitempty"`
	ShippingLabel                   *string                    `csv:"shipping_label,omitempty" validate:"max=100"`
	ShippingWeight                  *NumericValue              `csv:"shipping_weight,omitempty"`
	ShippingLength                  *NumericValue              `csv:"shipping_length,omitempty"`
	ShippingWidth                   *NumericValue              `csv:"shipping_width,omitempty"`
	ShippingHeight                  *NumericValue              `csv:"shipping_height,omitempty"`
	ShipsFromCountry                *string                    `csv:"ships_from_country,omitempty"`
	TransitTimeLabel                *string                    `csv:"transit_time_label,omitempty" validate:"max=100"`
	MinHandlingTime                 *uint64                    `csv:"min_handling_time,omitempty"`
	MaxHandlingTime                 *uint64                    `csv:"max_handling_time,omitempty"`
	Tax                             *TaxType                   `csv:"tax,omitempty" json:"tax,omitempty" xml:"g:tax,omitempty"`
	TaxCategory                     *string                    `csv:"tax_category,omitempty" json:"tax_category,omitempty" xml:"g:tax_category" validate:"max=100"`
}
