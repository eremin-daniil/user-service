package request

import "user_service/boundary/dto"

type RegisterUserRq struct {
	Login string `json:"login"`
}

func NewEmptyRegisterUserRq() *RegisterUserRq {
	return &RegisterUserRq{
		Login: "",
	}
}

func (rq *RegisterUserRq) ToDTO() *dto.UserDTO {
	return &dto.UserDTO{
		ID:      "",
		Login:   rq.Login,
		Profile: nil,
	}
}
