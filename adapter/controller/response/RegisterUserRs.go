package response

type RegisterUserRs struct {
	ID string `json:"userID"`
}

func RegisterUserRsFromID(id string) *RegisterUserRs {
	return &RegisterUserRs{
		ID: id,
	}
}
