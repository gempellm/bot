package parcel

import "github.com/gempellm/bot/internal/model/logistic"

func (s *DummyParcelService) Create(parcel logistic.Parcel) (uint64, error) {
	parcels := s.getParcels()
	parcels = append(parcels, parcel)
	s.saveParcels(parcels)

	return parcel.ParcelID, nil
}
