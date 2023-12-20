package validation

import (
	"fmt"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
)

type Validation struct{}

var (
	validate *validator.Validate
)

func (v *Validation) setValidation(value interface{}) []error {
	validate = validator.New()

	english := en.New()
	uni := ut.New(english, english)

	trans, _ := uni.GetTranslator("en")
	_ = enTranslations.RegisterDefaultTranslations(validate, trans)

	err := validate.Struct(value)
	errs := v.translateError(err, trans)

	return errs
}

func (v *Validation) translateError(err error, trans ut.Translator) (errs []error) {
	if err == nil {
		return nil
	}

	validatorErrs := err.(validator.ValidationErrors)
	for _, e := range validatorErrs {
		translatedErr := fmt.Errorf(e.Translate(trans))
		errs = append(errs, translatedErr)
	}
	return errs
}
