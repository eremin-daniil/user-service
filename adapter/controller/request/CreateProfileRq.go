package request

import (
	"time"
	"user_service/boundary/dto"
)

type CreateProfileRq struct {
	LastName  string `json:"lastName"`
	FirstName string `json:"firstName"`
}

func NewEmptyCreateProfileRq() *CreateProfileRq {
	return &CreateProfileRq{
		LastName:  "",
		FirstName: "",
	}
}

func (rq *CreateProfileRq) ToDTO() *dto.ProfileDTO {
	return &dto.ProfileDTO{
		ID:              "",
		LastName:        rq.LastName,
		FirstName:       rq.FirstName,
		CreatedDateTime: time.Time{},
		UpdatedDateTime: time.Time{},
	}
}
