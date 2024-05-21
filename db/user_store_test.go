package db

import (
	"testing"

	"github.com/Admiral-Simo/shortly_backend/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type UserStoreTestSuite struct {
	suite.Suite
	DB        *gorm.DB
	UserStore UserStorer
}

func (suite *UserStoreTestSuite) SetupTest() {
	database, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	suite.NoError(err)

	err = database.AutoMigrate(&models.User{})
	suite.NoError(err)

	suite.DB = database
	suite.UserStore = NewUserStore(database)
}

func (suite *UserStoreTestSuite) TestCreateUser() {
	user, err := suite.UserStore.CreateUser("testuser", "password123")
	assert.NoError(suite.T(), err)
	assert.NotNil(suite.T(), user)
	assert.Equal(suite.T(), "testuser", user.Username)
	assert.Empty(suite.T(), user.Password)

	// Try creating a user with the same username
	user, err = suite.UserStore.CreateUser("testuser", "password123")
	assert.Error(suite.T(), err)
	assert.Nil(suite.T(), user)
	assert.Equal(suite.T(), "username already taken", err.Error())
}

func (suite *UserStoreTestSuite) TestCheckUser() {
	// Create a user to test CheckUser
	_, err := suite.UserStore.CreateUser("testuser", "password123")
	assert.NoError(suite.T(), err)

	user, err := suite.UserStore.CheckUser("testuser", "password123")
	assert.NoError(suite.T(), err)
	assert.NotNil(suite.T(), user)
	assert.Equal(suite.T(), "testuser", user.Username)
	assert.Empty(suite.T(), user.Password)

	// Check with wrong password
	user, err = suite.UserStore.CheckUser("testuser", "wrongpassword")
	assert.Error(suite.T(), err)
	assert.Nil(suite.T(), user)
	assert.Equal(suite.T(), "invalid password", err.Error())

	// Check with non-existing user
	user, err = suite.UserStore.CheckUser("nonexisting", "password123")
	assert.Error(suite.T(), err)
	assert.Nil(suite.T(), user)
	assert.Equal(suite.T(), "user not found", err.Error())
}

func (suite *UserStoreTestSuite) TestGetUserById() {
	// Create a user to test GetUserById
	createdUser, err := suite.UserStore.CreateUser("testuser", "password123")
	assert.NoError(suite.T(), err)

	user, err := suite.UserStore.GetUserById(createdUser.ID)
	assert.NoError(suite.T(), err)
	assert.NotNil(suite.T(), user)
	assert.Equal(suite.T(), "testuser", user.Username)
	assert.Empty(suite.T(), user.Password)

	// Get non-existing user
	user, err = suite.UserStore.GetUserById(999)
	assert.Error(suite.T(), err)
	assert.Nil(suite.T(), user)
	assert.Equal(suite.T(), "user not found", err.Error())
}

func TestUserStoreTestSuite(t *testing.T) {
	suite.Run(t, new(UserStoreTestSuite))
}
