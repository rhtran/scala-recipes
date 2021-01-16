package pet

import "time"

type PetResponse struct {
	ID        int        `json:"id"`
	Name      string     `json:"name"`
	BirthDate *time.Time `json:"birthDate"`
	Type      string     `json:"type"`
}
