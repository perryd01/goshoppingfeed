package goshoppingfeed

type GoogleDSAFeed struct {
	Rows []*GoogleDSAElement
}

type GoogleDSAElement struct {
	URL         string `csv:"Page URL" json:"Page URL"`
	CustomLabel string `csv:"Custom label,omitempty" json:"Custom label"`
}
