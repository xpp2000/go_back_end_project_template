package dto

import "gogofly/model"

// 结构体标签的写法和框架有关
// dto在ReadJson()时并不会检查是否有多余的字段传入
type UserLoginDto struct {
	Name     string `form:"name" json:"name" validate:"required"`
	Password string `form:"password" json:"password" validate:"required"`
}

//===========================================
// =
type UserAddDto struct {
	ID       uint
	Name     string `json:"name" form:"name" validate:"required"`
	RealName string `json:"real_name" form:"real_name"`
	Avatar   string `json:"avatar" form:"avatar"`
	Mobile   string `json:"mobile" form:"mobile"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password,omitempty" form:"password" validate:"required"`
}

func (m *UserAddDto) ConvertToModel(iUser *model.User) {
	iUser.Name = m.Name
	iUser.RealName = m.RealName
	iUser.Mobile = m.Mobile
	iUser.Email = m.Email
	iUser.Password = m.Password
	iUser.Avatar = m.Avatar
}

//===========================================
// =
type UserListDto struct {
	PagingDto
}

//==========================================
// =
type UserUpdateDto struct {
	ID       uint   `json:"id" param:"id" form:"name"`
	Name     string `json:"name" form:"name"`
	RealName string `json:"real_name" form:"real_name"`
	Mobile   string `json:"mobile" form:"mobile"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password,omitempty" form:"password"`
}

func (m *UserUpdateDto) ConverttoModel(iUser *model.User) {
	iUser.ID = m.ID
	iUser.Name = m.Name
	iUser.RealName = m.RealName
	iUser.Mobile = m.Mobile
	iUser.Email = m.Email
}
