package router

// =================================
// set router with apis

import (
	"gogofly/api"

	irisRouter "github.com/kataras/iris/v12/core/router"
)

type PingResponse struct {
	Message string `json:"message"`
}

func InitUserRoutes() {
	userApi := api.NewUserApi()

	RegisterRoute(func(rgPublic irisRouter.Party, rgAuth irisRouter.Party) {
		rgPublicUser := rgPublic.Party("/user")
		{ // use blacket to organize router register code
			rgPublicUser.Post("/login", userApi.Login)
			rgPublicUser.Post("/add", userApi.AddUser)
			rgPublicUser.Get("/{id:uint}", userApi.GetUserById)
		}

		rgAuthUser := rgAuth.Party("/user")
		{
			rgAuthUser.Post("/list", userApi.GetUserList)
			rgAuthUser.Put("/{id:uint}", userApi.UpdateUser)
			rgAuthUser.Delete("/{id:uint}", userApi.DeleteUserById)

		}
	})
}
