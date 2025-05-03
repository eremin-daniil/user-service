package request

import (
	"time"
	"user_service/boundary/dto"
)

type UpdateProfileRq struct {
	LastName  string `json:"lastName"`
	FirstName string `json:"firstName"`
}

func NewEmptyUpdateProfileRq() *UpdateProfileRq {
	return &UpdateProfileRq{
		LastName:  "",
		FirstName: "",
	}
}

func (rq *UpdateProfileRq) ToDTO() *dto.ProfileDTO {
	return &dto.ProfileDTO{
		ID:              "",
		LastName:        rq.LastName,
		FirstName:       rq.FirstName,
		CreatedDateTime: time.Time{},
		UpdatedDateTime: time.Time{},
	}
}
