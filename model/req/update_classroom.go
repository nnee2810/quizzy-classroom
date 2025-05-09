package req

type UpdateClassroomReq struct {
	ID          string `json:"id" validate:"required"`
	Name        string `json:"name"`
	Description string `json:"description"`
	AvatarUrl   string `json:"avatar_url"`
}
