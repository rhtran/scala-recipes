package visit

import (
	"time"

	"github.com/rhtran/go-rest-template/api/rest/pet"
)

type VisitResponse struct {
	ID          int             `json:"id"`
	VisitDate   time.Time       `json:"visitDate"`
	Description string          `json:"description"`
	PetResponse pet.PetResponse `json:"pet"`
}
