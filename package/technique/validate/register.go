package validate

import (
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
)

const defaultErrMsg = " field did not pass validation"

/*var Trans ut.Translator*/

func RegisterTranslations(validate *validator.Validate, trans ut.Translator, builtinMessages map[string]string) {
	for s, message := range builtinMessages {
		validate.RegisterTranslation(s, trans, func(ut ut.Translator) error {
			return ut.Add(s, message, true)
		}, func(ut ut.Translator, fe validator.FieldError) string {
			t, _ := ut.T(s, fe.StructNamespace())
			return t
		})
	}

}
func RegisterDefaultTranslations(validate *validator.Validate, trans ut.Translator) {
	en_translations.RegisterDefaultTranslations(validate, trans)
}

var builtinMessages = map[string]string{
	"_": "{field}" + defaultErrMsg, // default message
	// builtin
	"_validate": "{field} did not pass validation", // default validate message
	"_filter":   "{field} data is invalid",         // data filter error
	// int value
	"min": "{field} min value is %v",
	"max": "{field} max value is %v",
	// type check: int
	"isInt":  "{field} value must be an integer",
	"isInt1": "{field} value must be an integer and mix value is %d",      // has min check
	"isInt2": "{field} value must be an integer and in the range %d - %d", // has min, max check
	"isInts": "{field} value must be an int slice",
	"isUint": "{field} value must be an unsigned integer(>= 0)",
	// type check: string
	"isString":  "{field} value must be a string",
	"isString1": "{field} value must be a string and min length is %d", // has min len check
	// length
	"minLength": "{field} min length is %d",
	"maxLength": "{field} max length is %d",
	// string length. calc rune
	"stringLength":  "{field} length must be in the range %d - %d",
	"stringLength1": "{field} min length is %d",
	"stringLength2": "{field} length must be in the range %d - %d",

	"isURL":     "{field} must be a valid URL address",
	"isFullURL": "{field} must be a valid full URL address",
	"regexp":    "{field} must match pattern %s",

	"isFile":  "{field} must be an uploaded file",
	"isImage": "{field} must be an uploaded image file",

	"enum":  "{field} value must be in the enum %v",
	"range": "{field} value must be in the range %d - %d",
	// int compare
	"lt": "{field} value should be less than %v",
	"gt": "{field} value should be greater than %v",
	// required
	"required":           "{field} is required to not be empty",
	"requiredIf":         "{field} is required when {args0} is {args1end}",
	"requiredUnless":     "{field} field is required unless {args0} is in {args1end}",
	"requiredWith":       "{field} field is required when {values} is present",
	"requiredWithAll":    "{field} field is required when {values} is present",
	"requiredWithout":    "{field} field is required when {values} is not present",
	"requiredWithoutAll": "{field} field is required when none of {values} are present",
	// field compare
	"eqField":  "{field} value must be equal the field %s",
	"neField":  "{field} value cannot be equal to the field %s",
	"ltField":  "{field} value should be less than the field %s",
	"lteField": "{field} value should be less than or equal to the field %s",
	"gtField":  "{field} value must be greater than the field %s",
	"gteField": "{field} value should be greater or equal to the field %s",
	// data type
	"bool":    "{field} value must be a bool",
	"float":   "{field} value must be a float",
	"slice":   "{field} value must be a slice",
	"map":     "{field} value must be a map",
	"array":   "{field} value must be an array",
	"strings": "{field} value must be a []string",
	"notIn":   "{field} value must not be in the given enum list %d",
	//
	"contains":    "{field} value does not contain %s",
	"notContains": "{field} value contains %s",
	"startsWith":  "{field} value does not start with %s",
	"endsWith":    "{field} value does not end with %s",
	"email":       "{field} value is an invalid email address",
	"regex":       "{field} value does not pass the regex check",
	"file":        "{field} value must be a file",
	"image":       "{field} value must be an image",
	// date
	"date":    "{field} value should be a date string",
	"gtDate":  "{field} value should be after %s",
	"ltDate":  "{field} value should be before %s",
	"gteDate": "{field} value should be after or equal to %s",
	"lteDate": "{field} value should be before or equal to %s",
	// check char
	"hasWhitespace":  "{field} value should contains spaces",
	"ascii":          "{field} value should be an ASCII string",
	"alpha":          "{field} value contains only alpha char",
	"alphaNum":       "{field} value contains only alpha char and num",
	"alphaDash":      "{field} value contains only letters, num, dashes (-) and underscores (_)",
	"multiByte":      "{field} value should be a multiByte string",
	"base64":         "{field} value should be a base64 string",
	"dnsName":        "{field} value should be a DNS string",
	"dataURI":        "{field} value should be a DataURL string",
	"empty":          "{field} value should be empty",
	"hexColor":       "{field} value should be a color string in hexadecimal",
	"hexadecimal":    "{field} value should be a hexadecimal string",
	"json":           "{field} value should be a json string",
	"lat":            "{field} value should be a latitude coordinate",
	"lon":            "{field} value should be a longitude coordinate",
	"num":            "{field} value should be a num (>=0) string",
	"mac":            "{field} value should be a MAC address",
	"cnMobile":       "{field} value should be string of Chinese 11-digit mobile phone numbers",
	"printableASCII": "{field} value should be a printable ASCII string",
	"rgbColor":       "{field} value should be a RGB color string",
	"fullURL":        "{field} value should be a complete URL string",
	"full":           "{field} value should be a URL string",
	"ip":             "{field} value should be an IP (v4 or v6) string",
	"ipv4":           "{field} value should be an IPv4 string",
	"ipv6":           "{field} value should be an IPv6 string",
	"CIDR":           "{field} value should be a CIDR string",
	"CIDRv4":         "{field} value should be a CIDRv4 string",
	"CIDRv6":         "{field} value should be a CIDRv6 string",
	"uuid":           "{field} value should be a UUID string",
	"uuid3":          "{field} value should be a UUID3 string",
	"uuid4":          "{field} value should be a UUID4 string",
	"uuid5":          "{field} value should be a UUID5 string",
	"filePath":       "{field} value should be an existing file path",
	"unixPath":       "{field} value should be a unix path string",
	"winPath":        "{field} value should be a windows path string",
	"isbn10":         "{field} value should be a isbn10 string",
	"isbn13":         "{field} value should be a isbn13 string",
}
