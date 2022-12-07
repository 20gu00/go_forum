package param

type RegisterInput struct {
	Username   string `json:"username" binding:"required"` //必填
	Password   string `json:"password" binding:"required"`
	RePassword string `json:"re_password" binding:"required,eqfield=Password"` //eq
}
