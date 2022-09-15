package parcel

import "fmt"

func (s *DummyParcelService) Remove(parcelID uint64) (bool, error) {
	parcels := s.getParcels()

	for i := range parcels {
		if parcels[i].ParcelID == parcelID {
			parcels = append(parcels[:i], parcels[i+1:]...)
			s.saveParcels(parcels)
			return true, nil
		}
	}

	return false, fmt.Errorf("parcel %d not found", parcelID)
}
