package model

type CreateUserRequest struct {
	Username    string `json:"username" binding:"required"`
	Firstname   string `json:"firstname" binding:"required"`
	Lastname    string `json:"lastname" binding:"required"`
	Age         uint   `json:"age" binding:"required"`
	Email       string `json:"email" binding:"required,email"`
	PhoneNumber string `json:"phonenumber" binding:"e164"`
	Gender      string `json:"gender" binding:"required"`
}

type GetUserResponse struct {
	ID        uint   `json:"id"`
	Username  string `json:"username"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Age       uint   `json:"age"`
	Email     string `json:"email"`
	Gender    string `json:"gender"`
	Role      string `json:"role"`
}

type UpdateRoleRequest struct {
	UserID uint `json:"userId" binding:"required"`
	RoleID uint `json:"roleId" binding:"required"`
}
