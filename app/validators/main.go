package validators

var Client *ClientValidator
var Request *RequestValidator
var Url *UrlValidator

func init() {
	base := GetValidator()

	Client = &ClientValidator{Base: base}
	Request = &RequestValidator{Base: base}
	Url = &UrlValidator{Base: base}
}
