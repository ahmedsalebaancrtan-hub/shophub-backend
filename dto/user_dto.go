
package dto

type RegisterRequest struct {
    FullName string `json:"full_name" binding:"required"`
    Email    string `json:"email" binding:"required,email"`
    Password string `json:"password" binding:"required,min=8"`
}

type RegisterResponse struct {
    ID        uint   `json:"id"`
    FullName  string `json:"full_name"`
    Email     string `json:"email"`
    Role      string `json:"role"`
    CreatedAt string `json:"created_at"`
}
