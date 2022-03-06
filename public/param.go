package public

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"log"
)

func DefaultParamsBindAndValidate(c *gin.Context, params interface{}) error {
	if err := c.ShouldBind(params); err != nil {
		return err
	}

	content, err := json.Marshal(params)
	if err != nil {
		log.Printf("marshal param err, err:%v", err)
	}
	log.Printf("param:%v", content)

	return nil
}
