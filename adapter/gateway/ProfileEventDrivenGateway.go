package gateway

import (
	"context"
	"user_service/boundary/dto"
)

type ProfileEventDrivenGateway struct {
}

func (p *ProfileEventDrivenGateway) SendCreateProfileEvent(ctx context.Context, profile *dto.ProfileDTO) error {
	return nil
}

func (p *ProfileEventDrivenGateway) SendUpdateProfileEvent(ctx context.Context, profile *dto.ProfileDTO) error {
	return nil
}
