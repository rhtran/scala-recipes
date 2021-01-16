package visit

import (
	"time"

	"github.com/rhtran/go-rest-template/api/rest/pet"

	"github.com/rhtran/go-rest-template/model"
)

type Visit struct {
	model.Base
	VisitDate   time.Time
	Description string
	PetID       int
	Pet         pet.Pet
}

func (v *Visit) ToVisitResponse(visit *Visit) *VisitResponse {
	return &VisitResponse{
		ID:          visit.ID,
		VisitDate:   visit.VisitDate,
		Description: visit.Description,
		PetResponse: *visit.Pet.ToPetResponse(&visit.Pet),
	}
}

func (v Visit) ToVisitResponses(visits []Visit) []VisitResponse {
	visitResponses := make([]VisitResponse, len(visits))
	for i, v := range visits {
		visitResponses[i] = *v.ToVisitResponse(&v)
	}
	return visitResponses
}
