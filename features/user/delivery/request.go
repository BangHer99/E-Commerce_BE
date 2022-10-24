package delivery

import "project/e-commerce/features/user"

type UserRequest struct {
	Name string `json:"name" form:"name"`
	// Username string `json:"username" form:"username"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

func toCore(data UserRequest) user.Core {
	return user.Core{
		Name:     data.Name,
		Email:    data.Email,
		Password: data.Password,
	}
}
