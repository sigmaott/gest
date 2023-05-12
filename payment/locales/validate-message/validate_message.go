package validate_message

var ValidateMessage = map[string]map[string]string{}

func init() {
	ValidateMessage["en"] = enMessages
	ValidateMessage["vi"] = viMessages
}
