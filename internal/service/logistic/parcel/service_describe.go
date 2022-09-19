package parcel

import (
	"github.com/gempellm/bot/internal/model/logistic"
)

func (s *DummyParcelService) Describe(parcelID uint64) (*logistic.Parcel, error) {
	parcels := s.getParcels()

	for _, parcel := range parcels {
		if parcel.ParcelID == parcelID {
			return &parcel, nil
		}
	}

	return nil, logistic.ErrParcelNotFound(parcelID)
}
