package validate

import (
	"reflect"
	"unicode/utf8"

	"regexp"

	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zhTrans "github.com/go-playground/validator/v10/translations/zh"
)

var (
	Validate *validator.Validate
	Trans    ut.Translator
)

func init() {
	zh2 := zh.New()
	uni := ut.New(zh2, zh2)
	Trans, _ = uni.GetTranslator("zh")
	Validate = validator.New()
	Validate.RegisterTagNameFunc(func(field reflect.StructField) string {
		str := field.Tag.Get("json")
		if len(str) == 0 {
			str = field.Name
		}
		return str
	})

	registerTagMsg("userNameIgnoreEmpty", userNameIgnoreEmpty, "{0}必须由5-12位(数字、字母、_、-)组成，且必须以字母开头")
	registerTagMsg("emailIgnoreEmpty", emailIgnoreEmpty, "邮箱格式不符合规范")

	_ = zhTrans.RegisterDefaultTranslations(Validate, Trans)
}

func registerTagMsg(tag string, function validator.Func, msg string) {
	_ = Validate.RegisterValidation(tag, function)
	_ = Validate.RegisterTranslation(tag, Trans, func(ut ut.Translator) error {
		return ut.Add(tag, msg, true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T(tag, fe.Field())
		return t
	})
}

func Valid(object interface{}) []string {
	var msgs []string
	err := Validate.Struct(object)
	if err != nil {
		errs := err.(validator.ValidationErrors).Translate(Trans)
		for _, v := range errs {
			msgs = append(msgs, v)
		}
	}
	return msgs
}

func ValidFirst(object interface{}) string {
	errs := Valid(object)
	if len(errs) > 0 {
		return errs[0]
	} else {
		return ""
	}
}

// custom validation

func userNameIgnoreEmpty(f validator.FieldLevel) bool {
	str := f.Field().String()
	count := utf8.RuneCountInString(str)
	if count == 0 {
		return true
	}
	pattern := `^[a-zA-Z][0-9a-zA-Z_-]{4,11}$`
	matched, _ := regexp.MatchString(pattern, str)
	return matched
}

func emailIgnoreEmpty(f validator.FieldLevel) bool {
	str := f.Field().String()
	count := utf8.RuneCountInString(str)
	if count == 0 {
		return true
	}
	pattern := `^([A-Za-z0-9_\-\.])+\@([A-Za-z0-9_\-\.])+\.([A-Za-z]{2,4})$`
	matched, _ := regexp.MatchString(pattern, str)
	return matched
}
