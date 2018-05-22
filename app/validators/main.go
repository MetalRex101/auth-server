package validators

var Client *ClientValidator
var Request *RequestValidator
var Url *UrlValidator
var Email *EmailValidator

func init() {
	base := GetValidator()

	Client = &ClientValidator{Base: base}
	Request = &RequestValidator{Base: base}
	Url = &UrlValidator{Base: base}
	Email = &EmailValidator{Base: base}
}
