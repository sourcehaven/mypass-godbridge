package schemas

type ActivationToken struct {
	Token string `json:"token" binding:"required" validate:"required"`
}
