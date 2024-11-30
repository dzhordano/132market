package mapper

import (
	"github.com/dzhordano/132market/services/users/internal/application/model"
	"github.com/dzhordano/132market/services/users/pkg/pb/user_v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func ToUserResponse(userResult *model.UserResult) *user_v1.User {
	return &user_v1.User{
		Id:         userResult.ID.String(),
		Name:       userResult.Name,
		Email:      userResult.Email,
		Roles:      userResult.Roles,
		Status:     userResult.Status,
		State:      userResult.State,
		CreatedAt:  timestamppb.New(userResult.CreatedAt),
		LastSeenAt: timestamppb.New(userResult.LastSeenAt),
	}
}

func ToUserListResponse(userResults []*model.UserResult) []*user_v1.User {
	var userResponses []*user_v1.User

	for _, userResult := range userResults {
		userResponses = append(userResponses, &user_v1.User{
			Id:         userResult.ID.String(),
			Name:       userResult.Name,
			Email:      userResult.Email,
			Roles:      userResult.Roles,
			Status:     userResult.Status,
			State:      userResult.State,
			CreatedAt:  timestamppb.New(userResult.CreatedAt),
			LastSeenAt: timestamppb.New(userResult.LastSeenAt),
		})
	}

	return userResponses
}
