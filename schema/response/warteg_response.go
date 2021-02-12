package response

type WartegAdd struct {
	WartegName        string `json:"warteg_name"`
	WartegDesc        string `json:"warteg_desc"`
	WartegAddr        string `json:"warteg_addr"`
	WartegContactName string `json:"warteg_contact_name"`
	WartegPhone       string `json:"warteg_phone"`
}

type WartegDelete struct {
	WartegId string `json:"warteg_id"`
}

type WartegUpdate struct {
	WartegId          string `json:"warteg_id"`
	WartegName        string `json:"warteg_name"`
	WartegDesc        string `json:"warteg_desc"`
	WartegAddr        string `json:"warteg_addr"`
	WartegContactName string `json:"warteg_contact_name"`
	WartegPhone       string `json:"warteg_phone"`
}

type WartegList struct {
	WartegId   string `json:"warteg_id"`
	WartegName string `json:"warteg_name"`
	WartegAddr string `json:"warteg_addr"`
}

type WartegDetail struct {
	WartegId          string `json:"warteg_id"`
	WartegName        string `json:"warteg_name"`
	WartegDesc        string `json:"warteg_desc"`
	WartegAddr        string `json:"warteg_addr"`
	WartegContactName string `json:"warteg_contact_name"`
	WartegPhone       string `json:"warteg_phone"`
}
