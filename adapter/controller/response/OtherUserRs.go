package response

import "user_service/boundary/dto"

type OtherUserRs struct {
	ID      string              `json:"userId"`
	Login   string              `json:"login"`
	Profile *ProfileOtherUserRs `json:"profile"`
}

func OtherUserRsFromDTO(dto *dto.UserDTO) *OtherUserRs {
	var profile *ProfileOtherUserRs
	if dto.Profile != nil {
		profile = ProfileOtherUserRsFromDTO(dto.Profile)
	}
	return &OtherUserRs{
		ID:      dto.ID,
		Login:   dto.Login,
		Profile: profile,
	}
}
