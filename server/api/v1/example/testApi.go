package example

import (
	"encoding/json"
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/gin-gonic/gin"
)

type TestApi struct{}

func (tt *TestApi)TestApi(c *gin.Context)  {
	body, _ := c.GetRawData()

	jsonFormData := make(map[string]interface{})
	_ = json.Unmarshal(body, &jsonFormData)
	fmt.Println(jsonFormData)

	response.Ok(c)
}