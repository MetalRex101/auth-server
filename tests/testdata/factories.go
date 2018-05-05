package testdata

import (
	"github.com/bluele/factory-go/factory"
	"github.com/MetalRex101/auth-server/app/models"
	"time"
	"github.com/MetalRex101/auth-server/app/services"
)

var ClientFact = factory.NewFactory(
	&models.Client{},
).SeqInt("ID", func(n int) (interface{}, error) {
	return uint(n), nil
}).Attr("ClientID", func(args factory.Args) (interface{}, error) {
	clientID := "732173982718"
	return &clientID, nil
}).Attr("ClientSecret", func(args factory.Args) (interface{}, error) {
	clientSecret := "some_random_secret"
	return &clientSecret, nil
}).Attr("Name", func(args factory.Args) (interface{}, error) {
	name := "Some Client Name"
	return &name, nil
}).Attr("Status", func(args factory.Args) (interface{}, error) {
	return true, nil
}).Attr("IP", func(args factory.Args) (interface{}, error) {
	ip := "*"
	return &ip, nil
}).Attr("Url", func(args factory.Args) (interface{}, error) {
	url := "url"
	return &url, nil
}).Attr("Scope", func(args factory.Args) (interface{}, error) {
	scope := "oauth"
	return &scope, nil
})

var OauthSessFact = factory.NewFactory(
	&models.OauthSession{},
).SeqInt("ID", func(n int) (interface{}, error) {
	return uint(n), nil
}).Attr("AccessGrantedAt", func(args factory.Args) (interface{}, error) {
	time := time.Now()
	return &time, nil
}).Attr("AccessExpiresAt", func(args factory.Args) (interface{}, error) {
	time := time.Now().Add(time.Hour)
	return &time, nil
}).Attr("Offset", func(args factory.Args) (interface{}, error) {
	return 0, nil
}).Attr("RemoteAddr", func(args factory.Args) (interface{}, error) {
	addr := "*"
	return &addr, nil
})

var UserFact = factory.NewFactory(
	&models.User{},
).SeqInt("ID", func(n int) (interface{}, error) {
	return uint(n), nil
}).Attr("FirstName", func(args factory.Args) (interface{}, error) {
	firstName := "Some First Name"
	return &firstName, nil
}).Attr("LastName", func(args factory.Args) (interface{}, error) {
	firstName := "Some Last Name"
	return &firstName, nil
})

var PasswordFact = factory.NewFactory(
	&models.Password{},
).SeqInt("ID", func(n int) (interface{}, error) {
	return uint(n), nil
}).Attr("Password", func(args factory.Args) (interface{}, error) {
	pass := services.HashPassword("djslakdjka783279831")
	return &pass, nil
}).Attr("CreatedAt", func(args factory.Args) (interface{}, error) {
	time := time.Now()
	return &time, nil
})

var EmailFact = factory.NewFactory(
	&models.Email{},
).SeqInt("ID", func(n int) (interface{}, error) {
	return uint(n), nil
}).Attr("Email", func(args factory.Args) (interface{}, error) {
	email := "some_email@gmai.com"
	return &email, nil
}).Attr("IsDefault", func(args factory.Args) (interface{}, error) {
	return true, nil
}).Attr("Status", func(args factory.Args) (interface{}, error) {
	return true, nil
}).Attr("ConfirmDate", func(args factory.Args) (interface{}, error) {
	time := time.Now()
	return &time, nil
})