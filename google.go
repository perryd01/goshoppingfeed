package goshoppingfeed

type GoogleDSAFeed struct {
	Rows []*GoogleDSAElement
}

type GoogleDSAElement struct {
	URL         string `csv:"Page URL" json:"Page URL"`
	CustomLabel string `csv:"Custom label,omitempty" json:"Custom label"`
}

type GoogleShoppingFeed struct {
	Rows []*GoogleShoppingElement
}

// GoogleShoppingElement https://support.google.com/merchants/answer/7052112?hl=en
type GoogleShoppingElement struct {
	ID                  string `csv:"id" json:"id" xml:"id"`          // Your product’s unique identifier
	Title               string `csv:"title" json:"title" xml:"title"` // Your product’s name
	Description         string `csv:"description"`                    // Your product’s description
	Link                string `csv:"link"`                           // Your product’s landing page
	ImageLink           string `csv:"image_link"`                     // The URL of your product’s main image
	AdditionalImageLink string `csv:"additional_image_link"`          // The URL of an additional image for your product
	MobileLink          string `csv:"mobile_link"`
}
