package account

type EventOpen struct {
	ID string `json:"id"`
}

type AccountDetails struct {
	Email string       `json:"email"`
	State AccountState `json:"state"`
}

type EventUpdate struct {
	NewDetails AccountDetails
}
