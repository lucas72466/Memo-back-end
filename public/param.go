package public

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func DefaultParamsBindAndValidate(c *gin.Context, params interface{}, autoSuccessLog bool) error {
	if err := c.ShouldBind(params); err != nil {
		LogWithContext(c, ErrorLevel, err, nil)
		if invalid, ok := err.(*validator.InvalidValidationError); ok {
			return errors.New("input param is invalid:" + invalid.Error())
		}

		return ParseValidatorErr(params, err)
	}

	if autoSuccessLog {
		LogWithContext(c, InfoLevel, JsonString(params), nil)
	}

	return nil
}
