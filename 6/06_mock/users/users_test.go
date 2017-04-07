package users

import (
	"testing"
	"github.com/golang/mock/gomock"
	"msu-go-11/6/06_mock/mocks"
)

func TestDoSomethingWithUsers(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	users := mock_users.NewMockUsers(ctrl)

	users.EXPECT().Put(500, "name", 7)

	DoSomethingWithUsers(&users)
}
