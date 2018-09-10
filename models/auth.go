package models

type (
	// FormLoginStruct represents Login form data
	FormLoginStruct struct {
		Email    string `form:"email" json:"email" binding:"required"`
		Password string `form:"password" json:"password" binding:"required"`
	}

	// FormRegisterStruct represents Register form data
	FormRegisterStruct struct {
		Email     string `form:"email" json:"email" binding:"required"`
		Password  string `form:"password" json:"password" binding:"required"`
		FirstName string `form:"first-name" json:"first-name"`
		LastName  string `form:"last-name" json:"last-name"`
		Image     string `form:"image" json:"image"`
	}
)
