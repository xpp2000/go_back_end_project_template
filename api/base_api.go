package api

import (
	"gogofly/global"
	"gogofly/model"
	"gogofly/utils"

	"github.com/go-playground/validator/v10"
	"github.com/kataras/iris/v12"
	"go.uber.org/zap"
)

var validate *validator.Validate = validator.New()

type BaseApi struct {
	Ctx    iris.Context
	Errors error
	Logger *zap.SugaredLogger
}

type BuildRequestOption struct {
	Ctx      iris.Context
	DTO      any
	BindUrl  bool
	BindAll  bool
	BindForm bool
}

func NewBaseApi() BaseApi {
	return BaseApi{
		Logger: global.Logger,
	}
}

func (m *BaseApi) AddError(errNew error) {
	m.Errors = utils.AppendError(m.Errors, errNew)
}

func (m BaseApi) GetError() error {
	return m.Errors
}

func (m *BaseApi) BuildRequest(option BuildRequestOption) *BaseApi {
	var errResult error
	// bind context
	m.Ctx = option.Ctx
	// bind data
	// ! the BuildRequestOption.DTO must be ptr
	// ! m.Ctx.ReadJSON(&option.DTO)  and  m.Ctx.ReadJSON(option.DTO) both can work
	// ! However, m.Ctx.ReadParams(&option.DTO) can't worked and m.Ctx.ReadParams(option.DTO)
	if option.DTO != nil {
		if option.BindAll || option.BindUrl {
			errResult = utils.AppendError(errResult, m.Ctx.ReadParams(option.DTO))
		}
		// if option.BindAll || !option.BindUrl {
		// 	errResult = utils.AppendError(errResult, m.Ctx.ReadJSON(option.DTO))
		// }
		if option.BindAll || !option.BindForm {
			errResult = utils.AppendError(errResult, m.Ctx.ReadForm(option.DTO))
		}
		if errResult != nil {
			m.AddError(errResult)
			m.Fail(model.ReasponseJson{
				Msg: m.GetError().Error(),
			})
		}

		// = validate data
		if err := validate.Struct(option.DTO); err != nil {
			// 提取校验错误信息
			m.AddError(err)
			m.Fail(model.ReasponseJson{
				Msg: m.GetError().Error(),
			})
		}

	}
	return m
}

func (m *BaseApi) Fail(resp model.ReasponseJson) {
	model.Fail(m.Ctx, resp)
}

func (m *BaseApi) Ok(resp model.ReasponseJson) {
	model.Ok(m.Ctx, resp)
}

func (m *BaseApi) ServerFail(resp model.ReasponseJson) {
	model.ServerFail(m.Ctx, resp)
}
