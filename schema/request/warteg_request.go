package request

type Warteg struct {
	WartegName        string `validate:"required" json:"warteg_name"`
	WartegDesc        string `json:"warteg_desc"`
	WartegAddr        string `validate:"required" json:"warteg_addr"`
	WartegContactName string `validate:"required" json:"warteg_contact_name"`
	WartegPhone       string `validate:"required" json:"warteg_phone"`
}

type WartegUpdate struct {
	WartegName        string `validate:"required" json:"warteg_name"`
	WartegDesc        string `json:"warteg_desc"`
	WartegAddr        string `validate:"required" json:"warteg_addr"`
	WartegContactName string `validate:"required" json:"warteg_contact_name"`
	WartegPhone       string `validate:"required" json:"warteg_phone"`
}
