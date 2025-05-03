package gateway

import (
	"context"
	"user_service/boundary/dto"
)

type UserEventDrivenGateway struct {
}

func (u *UserEventDrivenGateway) SendCreateUserEvent(ctx context.Context, user *dto.UserDTO) error {
	return nil
}
