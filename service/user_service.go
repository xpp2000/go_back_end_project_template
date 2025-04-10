package service

import (
	"errors"
	"gogofly/dao"
	"gogofly/dto"
	"gogofly/model"
)

var userService *UserService

type UserService struct {
	BaseService
	Dao *dao.UserDao
}

func NewUserService() *UserService {
	if userService == nil {
		userService = &UserService{
			Dao: dao.NewUserDao(),
		}
	}

	return userService
}

func (m *UserService) Login(iUserDto dto.UserLoginDto) (model.User, error) {
	var errResult error
	iUser := m.Dao.GetUserByNameAndPassword(iUserDto.Name, iUserDto.Password)

	if iUser.ID == 0 {
		errResult = errors.New("invalid UserName or Passwoed")
	}
	return iUser, errResult
}

func (m *UserService) AddUser(iUserAddDto *dto.UserAddDto) error {
	if m.Dao.CheckUserNameExist(iUserAddDto.Name) {
		return errors.New("user Name Exist")
	}
	return m.Dao.AddUser(iUserAddDto)
}

func (m *UserService) GetUserById(iCommonIDDto *dto.CommonIDDto) (model.User, error) {
	return m.Dao.GetUserById(iCommonIDDto.ID)
}

func (m *UserService) GetUserList(iUserListDto *dto.UserListDto) ([]model.User, int64, error) {
	return m.Dao.GetUserList(iUserListDto)
}

func (m *UserService) UpdateUser(iUserUpdateDto *dto.UserUpdateDto) error {
	if iUserUpdateDto.ID == 0 {
		return errors.New("invalid User ID")
	}
	// Maybe: Name collision check
	return m.Dao.UpdateUser(iUserUpdateDto)
}

func (m *UserService) DeleteUserById(iCommonIDDto *dto.CommonIDDto) error {
	return m.Dao.DeleteUserById(iCommonIDDto.ID)
}
