package response

type SwaggerWartegAdd struct {
	Base
	Data DataWarteg `json:"data"`
}

type DataWarteg struct {
	WartegName        string `json:"warteg_name"`
	WartegDesc        string `json:"warteg_desc"`
	WartegAddr        string `json:"warteg_addr"`
	WartegContactName string `json:"warteg_contact_name"`
	WartegPhone       string `json:"warteg_phone"`
}

type SwaggerWartegDetail struct {
	Base
	Data DataWartegDetail `json:"data"`
}

type DataWartegDetail struct {
	WartegId          string `json:"warteg_id"`
	WartegName        string `json:"warteg_name"`
	WartegDesc        string `json:"warteg_desc"`
	WartegAddr        string `json:"warteg_addr"`
	WartegContactName string `json:"warteg_contact_name"`
	WartegPhone       string `json:"warteg_phone"`
}

type SwaggerWartegList struct {
	Base
	Data []DataWartegList `json:"data"`
}

type DataWartegList struct {
	WartegId   string `json:"warteg_id"`
	WartegName string `json:"warteg_name"`
	WartegAddr string `json:"warteg_addr"`
}
