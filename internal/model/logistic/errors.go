package logistic

import "fmt"

func ErrParcelNotFound(parcelID uint64) error {
	return fmt.Errorf("parcel %d not found", parcelID)
}

func StringParcelNotFound(parcelID uint64) string {
	return fmt.Sprintf("Parcel with ID %d not found.", parcelID)
}
