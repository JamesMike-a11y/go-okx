package rest

import (
	"github.com/JamesMike-a11y/go-okx/examples"
	rc "github.com/JamesMike-a11y/go-okx/rest"
)

// 敏感信息申请的模拟盘KEY，不确定何时会删除
var TestClient = rc.New("", examples.Auth, nil)
