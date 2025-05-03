package response

import "user_service/boundary/dto"

type CurrentUserRs struct {
	ID      string                `json:"userId"`
	Login   string                `json:"login"`
	Profile *ProfileCurrentUserRs `json:"profile"`
}

func CurrentUserRsFromDTO(dto *dto.UserDTO) *CurrentUserRs {
	var profile *ProfileCurrentUserRs
	if dto.Profile != nil {
		profile = ProfileCurrentUserRsFromDTO(dto.Profile)
	}
	return &CurrentUserRs{
		ID:      dto.ID,
		Login:   dto.Login,
		Profile: profile,
	}
}
