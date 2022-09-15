package logistic

import (
	"fmt"
	"time"
)

var AllParcelsDemo = []Parcel{
	{ParcelID: 1, Title: "Clothing"},
	{ParcelID: 2, Title: "Computer"},
	{ParcelID: 3, Title: "Phone"},
	{ParcelID: 4, Title: "Jewerly"},
	{ParcelID: 5, Title: "Shoes"},
}

type Parcel struct {
	ParcelID  uint64 `json:"parcelID"`
	Title     string `json:"title"`
	Timestamp int64  `json:"timestamp"`
}

func (p *Parcel) String() string {
	return fmt.Sprintf("Parcel{ParcelID: %d, Title: %s, Timestamp: %s}", p.ParcelID, p.Title, time.Unix(p.Timestamp, 0))
}
