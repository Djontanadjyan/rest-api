package web

import (
	"api/internal/entity"
	"api/internal/infrastructure/repository"
	"api/internal/usecase"

	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/bxcodec/faker/v3"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestGet(t *testing.T) {
	gin.SetMode(gin.TestMode)

	t.Run("Success", func(t *testing.T) {
		uid := primitive.NewObjectID()

		var mockUser *entity.User
		var mockFriends []*entity.Friends
		err := faker.FakeData(&mockUser)
		assert.NoError(t, err)
		MockUserUsecase := new(usecase.UserUsecase)
		MockUserUsecase.On("GetByUUID", mock.Anything, uid.Hex()).Return(mockUser, nil)
		MockUserUsecase.On("GetAllFriends", mock.Anything, uid.Hex()).Return(mockFriends, nil)

		rr := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(rr)
		c.Params = []gin.Param{{Key: "id", Value: uid.Hex()}}

		c.Set(uid.String(), &entity.User{
			UUID: uid,
			Name: "Alex",
		},
		)
		router := gin.Default()

		NewRouter(router, MockUserUsecase)
		request, err := http.NewRequest(http.MethodGet, "/"+uid.Hex(), nil)
		assert.NoError(t, err)
		router.ServeHTTP(rr, request)
		handler := UserHandler{
			UserUseCase: MockUserUsecase,
		}
		handler.GetByUID(c)
		user, err := MockUserUsecase.GetByUUID(c, uid.Hex())
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, rr.Code)
		assert.Equal(t, user, mockUser)
		MockUserUsecase.AssertExpectations(t) // assert that UserService.Get was called
	})
}

func TestStart(t *testing.T) {
	gin.SetMode(gin.TestMode)

	t.Run("Success", func(t *testing.T) {

		mockUseCase := new(repository.MockUserRepo)
		rr := httptest.NewRecorder()

		router := gin.Default()
		NewRouter(router, mockUseCase)

		request, err := http.NewRequest(http.MethodGet, "/", nil)
		assert.NoError(t, err)

		router.ServeHTTP(rr, request)

		respBody, err := json.Marshal("Hello world")
		assert.NoError(t, err)

		assert.Equal(t, 200, rr.Code)
		assert.Equal(t, respBody, rr.Body.Bytes())
		mockUseCase.AssertExpectations(t) // assert that UserService.Get was called
	})
}
