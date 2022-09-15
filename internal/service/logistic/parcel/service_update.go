package parcel

import (
	"fmt"

	"github.com/gempellm/bot/internal/model/logistic"
)

func (s *DummyParcelService) Update(parcelID uint64, parcel logistic.Parcel) error {
	parcels := s.getParcels()

	for i := range parcels {
		if parcels[i].ParcelID == parcelID {
			parcels[i] = parcel
			s.saveParcels(parcels)
			return nil
		}
	}

	return fmt.Errorf("parcel %d not found", parcelID)
}
