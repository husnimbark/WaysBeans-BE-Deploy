package usersdto

type UserResponse struct {
	ID       int    `json:"id"`
	Name     string `json:"name" form:"name" validate:"required"`
	Email    string `json:"email" form:"email" validate:"required"`
	Password string `json:"password" form:"password" validate:"required"`
	Image    string `json:"image" form:"image" gorm:"type: varchar(255)"`
	Address  string `json:"address" gorm:"type: text"`
	Status   string `json:"status" form:"password"`
}
