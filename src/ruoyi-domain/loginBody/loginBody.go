package loginBody

type LoginBody struct {
	UserName string `json:"userName" gorm:"column:username"`
	Password string `json:"password" gorm:"column:password"`
	Code     string `json:"code"`
	Uuid     string `json:"uuid"`
}
