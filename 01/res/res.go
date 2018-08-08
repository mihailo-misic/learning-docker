package res

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type Error struct {
	Title         string `json:"title"`
	Detail        string `json:"detail"`
	OriginalError error  `json:"original-error"`
}

func Data(data interface{}) gin.H {
	// Handle db data
	if dbData, ok := data.(*gorm.DB); ok {
		data = dbData.Value
	}

	return gin.H{
		"data": data,
	}
}

func Err(errs ...Error) gin.H {
	var allErrs []Error
	for _, err := range errs {
		allErrs = append(allErrs, err)
	}

	return gin.H{
		"errors": allErrs,
	}
}
