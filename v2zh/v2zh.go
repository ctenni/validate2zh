package v2zh

import (
	"errors"
	"strings"

	zhLocales "github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zhTranslations "github.com/go-playground/validator/v10/translations/zh"
)

var (
	validate *validator.Validate
	trans    ut.Translator
)

func init() {
	validate = validator.New()
	zh := zhLocales.New()
	uni := ut.New(zh, zh)
	trans, _ = uni.GetTranslator("zh")
	zhTranslations.RegisterDefaultTranslations(validate, trans)

	// 注册 oneofci（大小写不敏感的 oneof）
	validate.RegisterValidation("oneofci", func(fl validator.FieldLevel) bool {
		val := strings.ToLower(fl.Field().String())
		params := strings.Split(fl.Param(), " ")
		for _, p := range params {
			if val == strings.ToLower(strings.TrimSpace(p)) {
				return true
			}
		}
		return false
	})

	// 注册 oneofci 的中文翻译
	validate.RegisterTranslation("oneofci", trans, func(ut ut.Translator) error {
		return ut.Add("oneofci", "{0}必须是[{1}]中的一个（忽略大小写）", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("oneofci", fe.Field(), fe.Param())
		return t
	})
}

func ValidateStruct(s any) error {
	err := validate.Struct(s)
	if err != nil {
		if validationErrors, ok := err.(validator.ValidationErrors); ok && len(validationErrors) > 0 {
			return errors.New(validationErrors[0].Translate(trans))
		}
		return err
	}
	return nil
}

// ValidateStructAll 校验结构体，返回所有中文错误信息，以 "; " 拼接。
// 如果校验通过则返回 nil。
func ValidateStructAll(s any) error {
	err := validate.Struct(s)
	if err == nil {
		return nil
	}
	if ves, ok := err.(validator.ValidationErrors); ok {
		msgs := make([]string, len(ves))
		for i, ve := range ves {
			msgs[i] = ve.Translate(trans)
		}
		return errors.New(strings.Join(msgs, "; "))
	}
	return err
}
