package service

import (
	"github.com/Vinicius-Madeira/go-web-app/src/configuration/rest_err"
	"github.com/Vinicius-Madeira/go-web-app/src/model"
	"github.com/Vinicius-Madeira/go-web-app/src/tests/mocks"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/mock/gomock"
	"math/rand"
	"strconv"
	"testing"
)

func TestUserDomainService_FindUserByIDServices(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repository := mocks.NewMockUserRepository(ctrl)
	service := NewUserDomainService(repository)

	t.Run("when_exists_user_returns_success", func(t *testing.T) {
		id := primitive.NewObjectID().Hex()

		userDomain := model.NewUserDomain("test@test.com", "testing123$", "test", 30)
		userDomain.SetID(id)

		repository.EXPECT().FindUserByID(id).Return(userDomain, nil)

		userDomainReturn, err := service.FindUserByIDServices(id)

		assert.Nil(t, err)
		assert.EqualValues(t, userDomainReturn.GetID(), userDomain.GetID())
		assert.EqualValues(t, userDomainReturn.GetEmail(), userDomain.GetEmail())
		assert.EqualValues(t, userDomainReturn.GetPassword(), userDomain.GetPassword())
		assert.EqualValues(t, userDomainReturn.GetName(), userDomain.GetName())
		assert.EqualValues(t, userDomainReturn.GetAge(), userDomain.GetAge())
	})

	t.Run("when_does_not_exists_an_user_returns_error", func(t *testing.T) {
		id := primitive.NewObjectID().Hex()

		repository.EXPECT().FindUserByID(id).Return(nil, rest_err.NewNotFoundError("user not found"))
		userDomainReturn, err := service.FindUserByIDServices(id)

		assert.NotNil(t, err)
		assert.EqualValues(t, err.Message, "user not found")
		assert.Nil(t, userDomainReturn)
	})
}

func TestUserDomainService_FindUserByEmailServices(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repository := mocks.NewMockUserRepository(ctrl)
	service := NewUserDomainService(repository)

	t.Run("when_exists_user_returns_success", func(t *testing.T) {
		id := primitive.NewObjectID().Hex()
		email := "test@success.com"

		userDomain := model.NewUserDomain(email, "testing123$", "test", 30)
		userDomain.SetID(id)

		repository.EXPECT().FindUserByEmail(email).Return(userDomain, nil)

		userDomainReturn, err := service.FindUserByEmailServices(email)

		assert.Nil(t, err)
		assert.EqualValues(t, userDomainReturn.GetID(), userDomain.GetID())
		assert.EqualValues(t, userDomainReturn.GetEmail(), userDomain.GetEmail())
		assert.EqualValues(t, userDomainReturn.GetPassword(), userDomain.GetPassword())
		assert.EqualValues(t, userDomainReturn.GetName(), userDomain.GetName())
		assert.EqualValues(t, userDomainReturn.GetAge(), userDomain.GetAge())
	})

	t.Run("when_does_not_exists_an_user_returns_error", func(t *testing.T) {
		email := "test@fail.com"

		repository.EXPECT().FindUserByEmail(email).Return(nil, rest_err.NewNotFoundError("user not found"))
		userDomainReturn, err := service.FindUserByEmailServices(email)

		assert.NotNil(t, err)
		assert.EqualValues(t, err.Message, "user not found")
		assert.Nil(t, userDomainReturn)
	})
}
func TestUserDomainService_FindUserByEmailAndPasswordServices(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repository := mocks.NewMockUserRepository(ctrl)
	service := &userDomainService{repository}

	t.Run("when_exists_user_returns_success", func(t *testing.T) {
		id := primitive.NewObjectID().Hex()
		email := "test@success.com"
		password := strconv.FormatInt(rand.Int63n(8), 10)

		userDomain := model.NewUserDomain(email, password, "test", 30)
		userDomain.SetID(id)

		repository.EXPECT().FindUserByEmailAndPassword(email, password).Return(userDomain, nil)

		userDomainReturn, err := service.findUserByEmailAndPasswordServices(email, password)

		assert.Nil(t, err)
		assert.EqualValues(t, userDomainReturn.GetID(), userDomain.GetID())
		assert.EqualValues(t, userDomainReturn.GetEmail(), userDomain.GetEmail())
		assert.EqualValues(t, userDomainReturn.GetPassword(), userDomain.GetPassword())
		assert.EqualValues(t, userDomainReturn.GetName(), userDomain.GetName())
		assert.EqualValues(t, userDomainReturn.GetAge(), userDomain.GetAge())
	})

	t.Run("when_does_not_exists_an_user_returns_error", func(t *testing.T) {
		email := "test@fail.com"
		password := strconv.FormatInt(rand.Int63n(8), 10)

		repository.EXPECT().FindUserByEmailAndPassword(email, password).Return(nil, rest_err.NewNotFoundError("user not found"))
		userDomainReturn, err := service.findUserByEmailAndPasswordServices(email, password)

		assert.NotNil(t, err)
		assert.EqualValues(t, err.Message, "user not found")
		assert.Nil(t, userDomainReturn)
	})
}
