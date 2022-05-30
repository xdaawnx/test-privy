package constant

// DataPaging is used to create a page data struct
type DataPaging struct {
	Total      int         `json:"total"`
	PageActive int         `json:"active_page"`
	Showing    string      `json:"showing"`
	PageList   []int       `json:"page_list"`
	Rows       interface{} `json:"rows"`
	Count      interface{} `json:"count,omitempty"`
}
