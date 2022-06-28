package models

type Client struct {
	ClientID    uint   `json:"clientid" gorm:"primary_key;auto_increment;not_null"`
	ClientName  string `json:"clientName"`
	ClientRegin string `json:"clientRegion"`
	GST         uint   `json:"GST"`
}
