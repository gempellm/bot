package logistic

import (
	"fmt"
)

var AllParcelsDemo = []Parcel{
	{ParcelID: 1, Title: "Clothing"},
	{ParcelID: 2, Title: "Computer"},
	{ParcelID: 3, Title: "Phone"},
	{ParcelID: 4, Title: "Jewerly"},
	{ParcelID: 5, Title: "Shoes"},
}

type Parcel struct {
	ParcelID uint64 `json:"parcelID"`
	Title    string `json:"title"`
}

func (p *Parcel) String() string {
	return fmt.Sprintf("Package{ParcelID: %d, Title: %s}", p.ParcelID, p.Title)
}
