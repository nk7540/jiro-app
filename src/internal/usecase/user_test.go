package usecase

import (
	"reflect"
	"testing"

	"artics-api/src/internal/domain"
	"artics-api/src/internal/domain/content"
	"artics-api/src/internal/domain/user"
	"artics-api/src/internal/usecase/request"
	"artics-api/src/internal/usecase/response"
	mock_content "artics-api/src/mock/domain/content"
	mock_user "artics-api/src/mock/domain/user"
	mock_pkg "artics-api/src/mock/pkg"
	mock_validation "artics-api/src/mock/usecase/validation"
	"artics-api/src/pkg"

	"github.com/golang/mock/gomock"
)

func TestUserUsecase_Create(t *testing.T) {
	tests := map[string]struct {
		Request *request.CreateUser
	}{
		"ok": {
			Request: &request.CreateUser{
				Email:                "test@blail.co.jp",
				Password:             "12345678",
				PasswordConfirmation: "12345678",
			},
		},
	}

	for result, test := range tests {
		// Defined variables
		ves := make([]*domain.ValidationError, 0)
		u := &user.User{
			Status:               "provisional",
			Email:                test.Request.Email,
			Password:             test.Request.Password,
			PasswordConfirmation: test.Request.PasswordConfirmation,
		}

		// Defined mocks
		ctrl := gomock.NewController(t)
		m := NewUserMocks(ctrl)

		m.RequestValidator.EXPECT().Run(m.Context, u).Return(ves)

		m.UserService.EXPECT().Create(m.Context, u).Return(nil)

		// Run test
		t.Run(result, func(t *testing.T) {
			target := NewUserUsecase(m.RequestValidator, m.UserService, m.ContentService)

			if err := target.Create(m.Context, test.Request); err != nil {
				t.Fatalf("failed to create a user: %v", err)
				return
			}
		})
	}
}

func TestUserUsecase_Show(t *testing.T) {
	tests := map[string]struct {
		ID       int
		Expected *response.ShowUser
	}{
		"ok": {
			ID: 1,
			Expected: &response.ShowUser{
				ID:             1,
				Nickname:       "test",
				Email:          "test@blail.co.jp",
				Followingcount: 999,
				Followercount:  99,
				FavoriteContents: []*response.Content{
					{
						ID:    1,
						Title: "test content 1",
					},
					{
						ID:    2,
						Title: "test content 2",
					},
				},
			},
		},
	}

	for result, test := range tests {
		// Defined variables
		u := &user.User{
			ID:             test.Expected.ID,
			Nickname:       test.Expected.Nickname,
			Email:          test.Expected.Email,
			FollowingCount: test.Expected.Followingcount,
			FollowerCount:  test.Expected.Followercount,
		}
		favoriteContents := make([]*content.Content, len(test.Expected.FavoriteContents))
		for i, c := range test.Expected.FavoriteContents {
			favoriteContents[i] = &content.Content{
				ID:    c.ID,
				Title: c.Title,
			}
		}

		// Defined mocks
		ctrl := gomock.NewController(t)
		m := NewUserMocks(ctrl)

		m.UserService.EXPECT().Show(m.Context, test.ID).Return(u, nil)
		m.ContentService.EXPECT().GetFavoriteContents(m.Context, test.ID, 3).Return(favoriteContents, nil)

		// Run test
		t.Run(result, func(t *testing.T) {
			target := NewUserUsecase(m.RequestValidator, m.UserService, m.ContentService)

			got, err := target.Show(m.Context, test.ID)
			if err != nil {
				t.Fatalf("failed to get a user: %v", err)
				return
			}

			if !reflect.DeepEqual(got, test.Expected) {
				t.Fatalf("expected %#v, got %#v", test.Expected, got)
				return
			}
		})
	}
}

func TestUserUsecase_Followings(t *testing.T) {
	tests := map[string]struct {
		ID       int
		Expected *response.Users
	}{
		"ok": {
			ID: 1,
			Expected: &response.Users{Users: []*response.User{
				{
					ID:       2,
					Nickname: "test user 2",
				},
				{
					ID:       3,
					Nickname: "test user 3",
				},
			}},
		},
	}

	for result, test := range tests {
		// Defined variables
		us := make([]*user.User, len(test.Expected.Users))

		// Defined mocks
		ctrl := gomock.NewController(t)
		m := NewUserMocks(ctrl)

		m.UserService.EXPECT().Followings(m.Context, test.ID).Return(us, nil)

		// Run test
		t.Run(result, func(t *testing.T) {
			target := NewUserUsecase(m.RequestValidator, m.UserService, m.ContentService)

			got, err := target.Followings(m.Context, test.ID)
			if err != nil {
				t.Fatalf("failed to get followings: %v", err)
				return
			}

			if !reflect.DeepEqual(got, test.Expected) {
				t.Fatalf("expected %#v, got %#v", test.Expected, got)
				return
			}
		})
	}
}

func TestUserUsecase_Followers(t *testing.T) {
	tests := map[string]struct {
		ID       int
		Expected *response.Users
	}{
		"ok": {
			ID: 1,
			Expected: &response.Users{Users: []*response.User{
				{
					ID:       2,
					Nickname: "test user 2",
				},
				{
					ID:       3,
					Nickname: "test user 3",
				},
			}},
		},
	}

	for result, test := range tests {
		// Defined variables
		us := make([]*user.User, len(test.Expected.Users))

		// Defined mocks
		ctrl := gomock.NewController(t)
		m := NewUserMocks(ctrl)

		m.UserService.EXPECT().Followers(m.Context, test.ID).Return(us, nil)

		// Run test
		t.Run(result, func(t *testing.T) {
			target := NewUserUsecase(m.RequestValidator, m.UserService, m.ContentService)

			got, err := target.Followers(m.Context, test.ID)
			if err != nil {
				t.Fatalf("failed to get followers: %v", err)
				return
			}

			if !reflect.DeepEqual(got, test.Expected) {
				t.Fatalf("expected %#v, got %#v", test.Expected, got)
				return
			}
		})
	}
}

func TestUserUsecase_Update(t *testing.T) {
	tests := map[string]struct {
		Request     *request.UpdateUser
		CurrentUser *user.User
		Expected    *response.UpdateUser
	}{}

	for result, test := range tests {
		// Defined variables
		ves := make([]*domain.ValidationError, 0)
		u := test.CurrentUser
		u.Nickname = test.Request.Nickname
		u.Email = test.Request.Email
		thumbnail, _ := test.Request.Thumbnail.Open()

		// Defined mocks
		ctrl := gomock.NewController(t)
		m := NewUserMocks(ctrl)

		m.Context.Locals("user", test.CurrentUser)
		m.RequestValidator.EXPECT().Run(m.Context, u).Return(ves)
		m.UserService.EXPECT().UpdateThumbnail(m.Context, thumbnail).Return(test.Expected.ThumbnailURL)
		m.UserService.EXPECT().Update(m.Context, u).Return(nil)

		// Run test
		t.Run(result, func(t *testing.T) {
			target := NewUserUsecase(m.RequestValidator, m.UserService, m.ContentService)

			got, err := target.Update(m.Context, test.Request)
			if err != nil {
				t.Fatalf("failed to update user: %v", err)
				return
			}

			if !reflect.DeepEqual(got, test.Expected) {
				t.Fatalf("expected %#v, got %#v", test.Expected, got)
				return
			}
		})
	}
}

func TestUserUsecase_Suspend(t *testing.T) {
	tests := map[string]struct {
		CurrentUser *user.User
	}{
		"ok": {
			CurrentUser: &user.User{
				ID:       1,
				Email:    "test@blail.co.jp",
				Nickname: "test",
			},
		},
	}

	for result, test := range tests {
		// Defined variables
		ves := make([]*domain.ValidationError, 0)
		u := test.CurrentUser
		u.Status = "suspended"

		// Defined mocks
		ctrl := gomock.NewController(t)
		m := NewUserMocks(ctrl)

		m.Context.Locals("user", u)
		m.RequestValidator.EXPECT().Run(m.Context, u).Return(ves)
		m.UserService.EXPECT().Suspend(m.Context, u).Return(nil)

		// Run test
		t.Run(result, func(t *testing.T) {
			target := NewUserUsecase(m.RequestValidator, m.UserService, m.ContentService)

			if err := target.Suspend(m.Context); err != nil {
				t.Fatalf("failed to suspend user: %v", err)
				return
			}
		})
	}
}

type UserMocks struct {
	Context          pkg.Context
	RequestValidator *mock_validation.MockRequestValidator
	UserService      *mock_user.MockUserService
	ContentService   *mock_content.MockContentService
}

func NewUserMocks(ctrl *gomock.Controller) *UserMocks {
	return &UserMocks{
		Context:          mock_pkg.NewMockContext(nil),
		RequestValidator: mock_validation.NewMockRequestValidator(ctrl),
		UserService:      mock_user.NewMockUserService(ctrl),
		ContentService:   mock_content.NewMockContentService(ctrl),
	}
}
