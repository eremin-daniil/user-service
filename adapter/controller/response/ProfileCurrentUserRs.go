package response

import "user_service/boundary/dto"

type ProfileCurrentUserRs struct {
	ID        string `json:"profileID"`
	LastName  string `json:"lastName"`
	FirstName string `json:"firstName"`
}

func ProfileCurrentUserRsFromDTO(dto *dto.ProfileDTO) *ProfileCurrentUserRs {
	return &ProfileCurrentUserRs{
		ID:        dto.ID,
		LastName:  dto.LastName,
		FirstName: dto.FirstName,
	}
}
