package dao

import (
	"gogofly/dto"
	"gogofly/model"
)

var userDao *UserDao

type UserDao struct {
	BaseDao
}

func NewUserDao() *UserDao {
	if userDao == nil {
		userDao = &UserDao{
			BaseDao: NewBaseDao(),
		}
	}
	return userDao
}

func (m *UserDao) GetUserByNameAndPassword(stUserName, stPassword string) model.User {
	var iUser model.User
	m.Orm.Model(&iUser).Where("name=? and password=?", stUserName, stPassword).Find(&iUser)
	return iUser
}

func (m *UserDao) CheckUserNameExist(stUserName string) bool {
	var nTotal int64
	m.Orm.Model(&model.User{}).Where("name = ?", stUserName).Count(&nTotal)

	return nTotal > 0
}

func (m *UserDao) AddUser(iUserAddDto *dto.UserAddDto) error {
	var iUser model.User
	iUserAddDto.ConvertToModel(&iUser)

	err := m.Orm.Save(&iUser).Error
	if err == nil {
		iUserAddDto.ID = iUser.ID
		iUserAddDto.Password = ""
	}
	return err
}

func (m *UserDao) GetUserById(id uint) (model.User, error) {
	var iUser model.User
	err := m.Orm.First(&iUser, id).Error
	return iUser, err
}

func (m *UserDao) GetUserList(iUserListDto *dto.UserListDto) ([]model.User, int64, error) {
	var giUserList []model.User
	var nTotal int64
	err := m.Orm.Model(&model.User{}).
		Scopes(Paginate(iUserListDto.PagingDto)).
		Find(&giUserList).
		Offset(-1).Limit(-1).
		Count(&nTotal).Error

	return giUserList, nTotal, err
}

func (m *UserDao) UpdateUser(iUserUpdateDto *dto.UserUpdateDto) error {
	var iUser model.User
	m.Orm.First(&iUser, iUserUpdateDto.ID)

	iUserUpdateDto.ConverttoModel(&iUser)

	return m.Orm.Save(&iUser).Error

}

func (m *UserDao) DeleteUserById(id uint) error {
	return m.Orm.Delete(&model.User{}, id).Error
}
