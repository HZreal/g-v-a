package example

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/gin-gonic/gin"
)

type TestRouter struct{}

func (e *TestRouter) InitTestRouter(Router *gin.RouterGroup) {
	testRouter := Router.Group("testGroup")
	testApi := v1.ApiGroupApp.ExampleApiGroup.TestApi
	{
		testRouter.POST("test111", testApi.TestApi)
	}
}

