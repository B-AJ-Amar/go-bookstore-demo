package utils

import (
	"log"
	"regexp"

	"github.com/go-playground/validator/v10"
)

var Validate *validator.Validate

func InitValidators() {
	Validate = validator.New()

	Validate.RegisterValidation("alpha_space", func(fl validator.FieldLevel) bool {
		value := fl.Field().String()
		re := regexp.MustCompile(`^[a-zA-Z ]{2,30}$`)
		
		return re.MatchString(value)
	})

	log.Println("Validators initialized successfully")
}
