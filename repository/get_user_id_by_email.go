package repository

import (
	"context"
	qerror "quizzy-classroom/error"
)

func (r *repositoryImpl) GetUserIDByEmail(ctx context.Context, email string) (string, error) {
	var users []struct {
		ID string `json:"id"`
	}

	if err := r.Provider.SupabaseClient.DB.Rpc("get_user_id_by_email", map[string]interface{}{
		"email": email,
	}).Execute(&users); err != nil {
		return "", err
	}

	if len(users) == 0 {
		return "", qerror.ErrEmailNotFound
	}

	return users[0].ID, nil
}
