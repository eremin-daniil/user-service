package gateway

import (
	"context"
	"user_service/boundary/dto"
)

type ProfileEventDrivenGatewayInterface interface {
	SendCreateProfileEvent(ctx context.Context, profile *dto.ProfileDTO) error
	SendUpdateProfileEvent(ctx context.Context, profile *dto.ProfileDTO) error
}
