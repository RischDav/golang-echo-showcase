package response

type ApiResponse struct {
    Success bool        `json:"success"`
    Data    interface{} `json:"data,omitempty"`
    Error   string      `json:"error,omitempty"`
    Message string      `json:"message,omitempty"`
}

type UserRequest struct {
    Firstname string `json:"firstname" validate:"required,min=2"`
    Lastname  string `json:"lastname" validate:"required,min=2"`
}

type UserResponse struct {
    ID        int32  `json:"id"`
    Firstname string `json:"firstname"`
    Lastname  string `json:"lastname"`
    CreatedAt string `json:"created_at,omitempty"`
    UpdatedAt string `json:"updated_at,omitempty"`
}

type ErrorResponse struct {
    Success bool   `json:"success"`
    Error   string `json:"error"`
    Code    string `json:"code,omitempty"`
}