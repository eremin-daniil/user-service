package response

import "user_service/boundary/dto"

type ProfileOtherUserRs struct {
	ID        string `json:"profileID"`
	LastName  string `json:"lastName"`
	FirstName string `json:"firstName"`
}

func ProfileOtherUserRsFromDTO(dto *dto.ProfileDTO) *ProfileOtherUserRs {
	return &ProfileOtherUserRs{
		ID:        dto.ID,
		LastName:  dto.LastName,
		FirstName: dto.FirstName,
	}
}
