package parcel

import "github.com/gempellm/bot/internal/model/logistic"

func (s *DummyParcelService) Remove(parcelID uint64) (bool, error) {
	parcels := s.getParcels()

	for i := range parcels {
		if parcels[i].ParcelID == parcelID {
			parcels = append(parcels[:i], parcels[i+1:]...)
			s.saveParcels(parcels)
			return true, nil
		}
	}

	return false, logistic.ErrParcelNotFound(parcelID)
}
