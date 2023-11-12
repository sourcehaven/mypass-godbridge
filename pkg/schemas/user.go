package schemas

type UserReg struct {
	Username   string `json:"username" binding:"required" validate:"required" example:"mypass"`
	Email      string `json:"email" binding:"required" validate:"required,email" example:"mypass@mypass.com"`
	Passphrase string `json:"passphrase" binding:"required" validate:"required" example:"quick brown fox jumping"`
	Firstname  string `json:"firstname" example:"John"`
	Lastname   string `json:"lastname" example:"Doe"`
}

type UserLogin struct {
	Username string `json:"username" binding:"required" validate:"required" example:"mypass"`
	Password string `json:"password" binding:"required" validate:"required" example:"super-secret"`
}

type UserActivation struct {
	OldPassword string `json:"oldPassword" binding:"required" validate:"required"`
	NewPassword string `json:"newPassword" binding:"required" validate:"required"`
}

type UserRegOk struct {
	Token    string `json:"token" binding:"required"`
	Password string `json:"password" binding:"required"`
}
