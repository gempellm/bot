package parcel

import (
	"encoding/json"
	"os"

	"github.com/gempellm/bot/internal/model/logistic"
)

type ParcelService interface {
	Describe(parcelID uint64) (*logistic.Parcel, error)
	List(cursor uint64, limit uint64) ([]logistic.Parcel, error)
	Create(logistic.Parcel) (uint64, error)
	Update(parcelID uint64, parcel logistic.Parcel) error
	Remove(parcelID uint64) (bool, error)
	Ids() ([]uint64, error)
}

type DummyParcelService struct{}

func NewDummyParcelService() *DummyParcelService {
	return &DummyParcelService{}
}

func (s *DummyParcelService) getParcels() []logistic.Parcel {
	parcelsData, _ := os.ReadFile("parcels.txt")
	var parcels []logistic.Parcel
	json.Unmarshal(parcelsData, &parcels)

	return parcels
}

func (s *DummyParcelService) saveParcels(parcels []logistic.Parcel) {
	parcelsData, _ := json.Marshal(parcels)
	os.WriteFile("parcels.txt", parcelsData, 0666)
}
