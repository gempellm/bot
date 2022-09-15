package parcel

import (
	"github.com/gempellm/bot/internal/model/logistic"
)

func (s *DummyParcelService) List(offset uint64, limit uint64) ([]logistic.Parcel, error) {
	parcels := s.getParcels()

	parcelsCount := len(parcels)
	if parcelsCount < 1 {
		return []logistic.Parcel{}, nil
	}

	parcelsMaxIndex := uint64(parcelsCount - 1)

	if offset > parcelsMaxIndex {
		return []logistic.Parcel{}, nil
	}

	if limit == 1 {
		return []logistic.Parcel{parcels[offset]}, nil
	}

	requestedParcels := make([]logistic.Parcel, 0)
	parcelsLimit := offset + limit

	for i := offset; i < parcelsLimit; i++ {
		if i > parcelsMaxIndex {
			break
		}
		requestedParcels = append(requestedParcels, parcels[i])
	}

	return requestedParcels, nil
}
