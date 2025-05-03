package gateway

import (
	"context"
	"user_service/boundary/dto"
)

type UserEventDrivenGatewayInterface interface {
	SendCreateUserEvent(ctx context.Context, user *dto.UserDTO) error
}
