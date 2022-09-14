package parcel

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/gempellm/bot/internal/model/logistic"
)

type ParcelService interface {
	Describe(parcelID uint64) (*logistic.Parcel, error)
	List(cursor uint64, limit uint64) ([]logistic.Parcel, error)
	Create(logistic.Parcel) (uint64, error)
	Update(parcelID uint64, parcel logistic.Parcel) error
	Remove(parcelID uint64) (bool, error)
}

type DummyParcelService struct {
	Parcels map[uint64]logistic.Parcel
}

func NewDummyParcelService() *DummyParcelService {
	return &DummyParcelService{}
}

func (s *DummyParcelService) getParcels() []logistic.Parcel {
	parcelsData, _ := os.ReadFile("parcels.txt")
	var parcels []logistic.Parcel
	json.Unmarshal(parcelsData, parcels)

	return parcels
}

func (s *DummyParcelService) saveParcels(parcels []logistic.Parcel) {
	parcelsData, _ := json.Marshal(parcels)
	os.WriteFile("parcels.txt", parcelsData, 666)
}

func (s *DummyParcelService) Describe(parcelID uint64) (*logistic.Parcel, error) {
	parcels := s.getParcels()

	for _, parcel := range parcels {
		if parcel.ParcelID == parcelID {
			return &parcel, nil
		}
	}

	return nil, fmt.Errorf("parcel %d not found", parcelID)
}

func (s *DummyParcelService) List(cursor uint64, limit uint64) ([]logistic.Parcel, error) {
	parcels := s.getParcels()

	parcelsLimit := uint64(len(parcels) - 1)

	if cursor > parcelsLimit {
		return []logistic.Parcel{}, nil
	}

	requestedParcels := make([]logistic.Parcel, limit)
	j := 0

	for i := cursor; i < i+limit; i++ {
		if i > parcelsLimit {
			break
		}
		requestedParcels[j] = parcels[i]
		j++
	}

	return parcels, nil
}

func (s *DummyParcelService) Create(parcel logistic.Parcel) (uint64, error) {
	parcels := s.getParcels()
	parcels = append(parcels, parcel)
	s.saveParcels(parcels)

	return parcel.ParcelID, nil
}

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