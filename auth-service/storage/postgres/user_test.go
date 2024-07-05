package postgres

import (
	"log"
	"testing"

	pb "root/genprotos"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestCreateUser(t *testing.T) {
	stg, err := NewPostgresStorage()
	if err != nil {
		log.Fatal("Error while connection on db: ", err.Error())
	}

	user := &pb.User{
		Id:       uuid.NewString(),
		UserName: uuid.NewString(),
		Password: uuid.NewString(),
		Email:    uuid.NewString(),
	}
	_, err = stg.User().CreateUser(user)

	assert.NoError(t, err)

}

func TestGetByIdUser(t *testing.T) {
	stg, err := NewPostgresStorage()
	if err != nil {
		log.Fatal("Error while connection on db: ", err.Error())
	}

	var Id pb.ById

	Id.Id = "b2a8c9d0-9d4b-4c2e-8d1f-9a8b1d0c9d1e"

	user, err := stg.User().GetByIdUser(&Id)

	assert.NoError(t, err)
	assert.NotNil(t, user)
}

func TestGetAllUser(t *testing.T) {
	stg, err := NewPostgresStorage()
	if err != nil {
		log.Fatal("Error while connection on db: ", err.Error())
	}
	users, err := stg.User().GetAllUser(&pb.User{})
	assert.NoError(t, err)
	assert.NotNil(t, users)
}

func TestUpdateUser(t *testing.T) {
	stg, err := NewPostgresStorage()
	if err != nil {
		log.Fatal("Error while connection on db: ", err.Error())
	}

	user := &pb.User{
		Id:       "b2a8c9d0-9d4b-4c2e-8d1f-9a8b1d0c9d1e",
		UserName: "updated_user",
		Email:    "updated_user@example.com",
	}
	_, err = stg.User().UpdateUser(user)
	assert.NoError(t, err)
	result, err := stg.User().GetByIdUser(&pb.ById{Id: user.Id})
	assert.NoError(t, err)
	assert.NotNil(t, result)
}

func TestLoginUser(t *testing.T) {
	stg, err := NewPostgresStorage()
	if err != nil {
		log.Fatal("Error while connection on db: ", err.Error())
	}
	user := &pb.User{
		Id:       uuid.NewString(),
		UserName: uuid.NewString(),
		Password: uuid.NewString(),
		Email:    uuid.NewString(),
	}
	_, err = stg.User().CreateUser(user)
	assert.NoError(t, err)
	userLogin, err := stg.User().LoginUser(&pb.User{Id: user.Id, UserName: user.UserName, Password: user.Password, Email: user.Email})

	assert.NoError(t, err)
	
	assert.Equal(t, userLogin.UserName, user.UserName)
	assert.Equal(t, userLogin.Email, user.Email)

}
