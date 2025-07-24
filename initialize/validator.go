package initialize

import (
	"fmt"
	myvalidator "my-ops-admin/pkg/validator"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func InitValidator() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		err := v.RegisterValidation("dateformat", myvalidator.ValidateDateFormat)
		if err != nil {
			fmt.Println("validator 注册失败！")
		}
	}
}
