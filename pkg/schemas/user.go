package schemas

type UserReg struct {
	Username   string `json:"username" binding:"required"`
	Email      string `json:"email" binding:"required"`
	Password   string `json:"password" binding:"required"`
	Passphrase string `json:"passphrasw" binding:"required"`
	Firstname  string `json:"firstname"`
	Lastname   string `json:"lastname"`
}
