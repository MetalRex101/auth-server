package testdata

import (
	"github.com/bluele/factory-go/factory"
	"github.com/MetalRex101/auth-server/app/models"
	"time"
	"github.com/MetalRex101/auth-server/app/services"
	"github.com/Pallinder/go-randomdata"
	"strconv"
)

var ClientFact = factory.NewFactory(
	&models.Client{},
).SeqInt("ID", func(n int) (interface{}, error) {
	return uint(n), nil
}).Attr("ClientID", func(args factory.Args) (interface{}, error) {
	clientID := strconv.Itoa(randomdata.Number(1000000000, 9999999999))
	return &clientID, nil
}).Attr("ClientSecret", func(args factory.Args) (interface{}, error) {
	clientSecret := randomdata.RandStringRunes(16)
	return &clientSecret, nil
}).Attr("Name", func(args factory.Args) (interface{}, error) {
	name := randomdata.SillyName()
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
	firstName := randomdata.FirstName(randomdata.Male)
	return &firstName, nil
}).Attr("LastName", func(args factory.Args) (interface{}, error) {
	firstName := randomdata.LastName()
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
	email := randomdata.Email()
	return &email, nil
}).Attr("IsDefault", func(args factory.Args) (interface{}, error) {
	return true, nil
}).Attr("Status", func(args factory.Args) (interface{}, error) {
	return true, nil
}).Attr("ConfirmDate", func(args factory.Args) (interface{}, error) {
	time := time.Now()
	return &time, nil
})

var PhoneFact = factory.NewFactory(
	&models.Phone{},
).SeqInt("ID", func(n int) (interface{}, error) {
	return uint(n), nil
}).Attr("Phone", func(args factory.Args) (interface{}, error) {
	phone := randomdata.Number(1000000000, 9999999999)
	return &phone, nil
}).Attr("IsDefault", func(args factory.Args) (interface{}, error) {
	return true, nil
}).Attr("Status", func(args factory.Args) (interface{}, error) {
	return true, nil
}).Attr("ConfirmDate", func(args factory.Args) (interface{}, error) {
	time := time.Now()
	return &time, nil
})