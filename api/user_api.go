package api

import (
	"fmt"
	"gogofly/dto"
	"gogofly/model"
	"gogofly/service"
	"gogofly/utils"
	"strconv"

	"github.com/kataras/iris/v12"
)

const (
	ERR_CODE_ADD_USER       = 10011
	ERR_CODE_GET_USER_BY_ID = 10012
	ERR_CODE_GET_USER_LIST  = 10013
	ERR_CODE_UPDATE_USER    = 10014
	ERR_CODE_DELETE_USER    = 10015
)

type UserApi struct {
	BaseApi
	Service *service.UserService
}

// =========================
// only controller
func NewUserApi() UserApi {
	return UserApi{
		BaseApi: NewBaseApi(),
		Service: service.NewUserService(),
	}
}

// @Tags User
// @Summary User Login
// @Description this api unfinished
// @Param name formData string true "User Name"
// @Param password formData string true "Password"
// @Success 200 {string} string "Login successful" {"data": S}
// @Failure 401 {string} string "Login failed"
// @Router /api/v1/public/user/login [post]
func (m UserApi) Login(ctx iris.Context) {

	var iUserLoginDto dto.UserLoginDto

	// 未封装写法
	// errs := ctx.ReadJSON(&iUserLoginDTO)
	// if errs != nil {
	// 	model.Fail(ctx, model.ReasponseJson{
	// 		Msg: errs.Error(),
	// 	})
	// }
	if err := m.BuildRequest(BuildRequestOption{Ctx: ctx, DTO: &iUserLoginDto}).GetError(); err != nil {
		return
	}

	iUser, err := m.Service.Login(iUserLoginDto)
	if err != nil {
		m.Fail(model.ReasponseJson{
			Msg: err.Error(),
		})
		return
	}

	token, _ := utils.GenerateToken(iUser.ID, iUser.Name)

	// 未封装写法
	// model.Ok(ctx, model.ReasponseJson{
	// 	Data: iUserLoginDTO,
	// })
	m.Ok(model.ReasponseJson{
		Data: map[string]any{
			"token": token,
		},
	})
}

// @Tags User
// @Summary User Add
// @Add a new user
// @Param name formData string true "User Name"
// @Param password formData string true "Password"
// @Success 200 {string} string "Add new user successful" {"data": S}
// @Failure 500 {string} string "User has been registered"
// @Router /api/v1/public/user/add [post]
func (m UserApi) AddUser(c iris.Context) {
	var iUserAddDto dto.UserAddDto
	if err := m.BuildRequest(BuildRequestOption{Ctx: c, DTO: &iUserAddDto}).GetError(); err != nil {
		return
	}

	// receive Avatar file
	// ! this parameter"uploadfile" should equal to formfield
	f, fh, err := c.FormFile("uploadfile")

	if err != nil {
		m.ServerFail(model.ReasponseJson{
			Code: ERR_CODE_ADD_USER,
			Msg:  err.Error(),
		})
		return
	}
	defer f.Close()

	stFilePath := fmt.Sprintf("./upload/%s", fh.Filename)
	fmt.Println("stFilePath is:", stFilePath)
	_, err = c.SaveFormFile(fh, stFilePath)
	if err != nil {
		m.Fail(model.ReasponseJson{
			Msg: err.Error(),
		})
		return
	}
	iUserAddDto.Avatar = stFilePath

	err = m.Service.AddUser(&iUserAddDto)

	if err != nil {
		m.ServerFail(model.ReasponseJson{
			Code: ERR_CODE_ADD_USER,
			Msg:  err.Error(),
		})
		return
	}

	m.Ok(model.ReasponseJson{
		Data: iUserAddDto,
	})

}

func (m UserApi) GetUserById(c iris.Context) {
	var iCommonIDDto dto.CommonIDDto
	if err := m.BuildRequest(BuildRequestOption{Ctx: c, DTO: &iCommonIDDto, BindUrl: true}).GetError(); err != nil {
		return
	}
	iUser, err := m.Service.GetUserById(&iCommonIDDto)
	if err != nil {
		m.ServerFail(model.ReasponseJson{
			Code: ERR_CODE_GET_USER_BY_ID,
			Msg:  err.Error(),
		})
		return
	}
	m.Ok(model.ReasponseJson{
		Data: iUser,
	})
}

func (m UserApi) GetUserList(c iris.Context) {
	var iUserListDto dto.UserListDto
	if err := m.BuildRequest(BuildRequestOption{Ctx: c, DTO: &iUserListDto}).GetError(); err != nil {
		fmt.Println("Page and limit is:", iUserListDto.Page)
		return
	}
	giUserList, nTotal, err := m.Service.GetUserList(&iUserListDto)
	fmt.Println("UserList is: ", giUserList)
	if err != nil {
		m.ServerFail(model.ReasponseJson{
			Code: ERR_CODE_GET_USER_LIST,
			Msg:  err.Error(),
		})
		return
	}
	m.Ok(model.ReasponseJson{
		Data:  giUserList,
		Total: nTotal,
	})
}

func (m UserApi) UpdateUser(c iris.Context) {
	var iUserUpdateDto dto.UserUpdateDto

	strID := c.Params().Get("id")
	id, _ := strconv.Atoi(strID)
	iUserUpdateDto.ID = uint(id)

	if err := m.BuildRequest(BuildRequestOption{Ctx: c, DTO: &iUserUpdateDto, BindAll: true}).GetError(); err != nil {
		return
	}

	err := m.Service.UpdateUser(&iUserUpdateDto)
	if err != nil {
		m.ServerFail(model.ReasponseJson{
			Code: ERR_CODE_UPDATE_USER,
			Msg:  err.Error(),
		})
		return
	}

	m.Ok(model.ReasponseJson{
		Data: iUserUpdateDto,
	})
}

func (m UserApi) DeleteUserById(c iris.Context) {
	var iCommonIDDto dto.CommonIDDto
	if err := m.BuildRequest(BuildRequestOption{Ctx: c, DTO: &iCommonIDDto, BindUrl: true}).GetError(); err != nil {
		return
	}

	err := m.Service.DeleteUserById(&iCommonIDDto)
	if err != nil {
		m.ServerFail(model.ReasponseJson{
			Code: ERR_CODE_DELETE_USER,
			Msg:  err.Error(),
		})
		return
	}

	m.Ok(model.ReasponseJson{})
}
