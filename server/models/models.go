package models

type UpsertUserPhotoResonse struct {
	Message  string `json:"message" example:"1"`
	PhotoUrl string `json:"photo_url" example:"http://localhost:8080/static/1.jpg"`
}

type GetUserPhotoResponse struct {
	PhotoUrl string `json:"photo_url" example:"http://localhost:8080/static/1.jpg"`
}

type ErrorResponse struct {
	Message string `json:"message" example:"Error message"`
}
