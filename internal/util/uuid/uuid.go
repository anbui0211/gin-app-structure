package puuid

import (
	"github.com/google/uuid"
)

func ValidID(id string) bool {
	_, err := uuid.Parse(id)
	return err != nil
}
