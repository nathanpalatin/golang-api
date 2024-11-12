package model

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type UserResponse struct {
  Data  []User `json:"data"`
  Total int    `json:"total"`
}

func NewUserResponse(data []User, total int) *UserResponse {
  return &UserResponse{Data: data, Total: total}
}

