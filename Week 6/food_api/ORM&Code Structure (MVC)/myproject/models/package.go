// models/package.go
package models

import "gorm.io/gorm"

type Package struct {
	gorm.Model
	Name             string  `json:"name"`
	Sender           string  `json:"sender"`
	Receiver         string  `json:"receiver"`
	SenderLocation   string  `json:"sender_location"`
	ReceiverLocation string  `json:"receiver_location"`
	Fee              float64 `json:"fee"`
	Weight           float64 `json:"weight"`
}
