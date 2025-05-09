package req

type CreateClassroomRequest struct {
	Name        *string `json:"name" validate:"required"`
	Description *string `json:"description" validate:"required"`
	AvatarUrl   *string `json:"avatar_url" validate:"required,url"`
}
