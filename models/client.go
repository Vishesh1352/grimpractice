package models

type Client struct {
	ClientID    uint    `json:"clientId" gorm:"primary_key;auto_increment;not_null"`
	ClientName  string  `json:"clientName"`
	ClientRegin string  `json:"clientRegion"`
	GST         float32 `json:"gst"`
}
