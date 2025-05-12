package repository

import (
	"context"
	"quizzy-classroom/entity"
)

func (r *repositoryImpl) GetUserProfilesByIDs(ctx context.Context, userIDs []string) (map[string]entity.UserProfileEntity, error) {
	var users []entity.UserProfileEntity
	if err := r.Provider.SupabaseClient.DB.Rpc("get_profiles_by_ids", map[string]any{
		"p_ids": userIDs,
	}).Execute(&users); err != nil {
		return nil, err
	}

	userMap := make(map[string]entity.UserProfileEntity)
	for _, user := range users {
		userMap[user.ID] = user
	}

	return userMap, nil
}
