package parcel

func (s *DummyParcelService) Ids() ([]uint64, error) {
	parcels := s.getParcels()
	ids := make([]uint64, len(parcels))

	for i, parcel := range parcels {
		ids[i] = parcel.ParcelID
	}

	return ids, nil
}
