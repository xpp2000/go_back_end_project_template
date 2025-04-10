package model

import (
	"reflect"

	"github.com/kataras/iris/v12"
)

type ReasponseJson struct {
	Status int    `json:"-"`
	Code   int    `json:"code,omitempty"`
	Msg    string `json:"msg,omitempty"`
	Data   any    `json:"data,omitempty"`
	Total  int64  `json:"total,omitempty"`
}

func buildStatus(resp ReasponseJson, nDefaultStatus int) int {
	if resp.Status == 0 {
		return nDefaultStatus
	}
	return resp.Status
}

func (m ReasponseJson) IsEmpty() bool {
	return reflect.DeepEqual(m, ReasponseJson{})
}

func Ok(ctx iris.Context, resp ReasponseJson) {
	HttpResponse(ctx, buildStatus(resp, 200), resp)
	// = alternative
	// ctx.StopWithJSON(buildStatus(resp, 200), resp)
}

func Fail(ctx iris.Context, resp ReasponseJson) {
	HttpResponse(ctx, buildStatus(resp, 400), resp)
	// = alternative
	// ctx.StopWithJSON(buildStatus(resp, 400), resp)
}

func ServerFail(ctx iris.Context, resp ReasponseJson) {
	HttpResponse(ctx, buildStatus(resp, 500), resp)
}

func HttpResponse(ctx iris.Context, status int, resp ReasponseJson) {
	if resp.IsEmpty() {
		ctx.StopWithJSON(status, nil)
		return
	} else {
		ctx.StopWithJSON(status, resp)
	}
}
