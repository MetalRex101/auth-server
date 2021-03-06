package resources

import (
	"github.com/MetalRex101/auth-server/app/models"
	"github.com/MetalRex101/auth-server/app/services"
	"github.com/labstack/echo"
	"time"
	"strconv"
)

func NewRegisteredResource(user *models.User, um *services.UserManager, c echo.Context) *RegisteredUserResource {
	var gender string
	var phone string

	phoneModel, _ := um.GetDefaultPhone(user.ID, true, c)
	if phoneModel != nil {
		phone = strconv.Itoa(int(*phoneModel.Phone))
	} else {
		phone = ""
	}

	email, _ := um.GetDefaultEmail(user.ID, true, c)

	if user.Gender != nil {
		if *user.Gender == "m" {
			gender = "male"
		} else if *user.Gender == "f" {
			gender = "female"
		}
	} else {
		gender = ""
	}

	lastVisit := ""
	if user.LastVisit != nil {
		lastVisit = user.LastVisit.Format(time.RFC1123Z)
	}

	birthDate := ""
	if user.BirthDate != nil {
		birthDate = user.BirthDate.Format(time.RFC1123Z)
	}

	merged := ""
	if user.Merged != nil {
		merged = *user.Merged
	}

	nickname := ""
	if user.Nickname != nil {
		nickname = *user.Nickname
	}

	givenName := ""
	if user.FirstName != nil {
		givenName = *user.FirstName
	}

	lastName := ""
	if user.LastName != nil {
		lastName = *user.LastName
	}

	parentName := ""
	if user.FatherName != nil {
		parentName = *user.FatherName
	}

	return &RegisteredUserResource{
		Id: strconv.Itoa(int(user.ID)),
		Created: user.CreatedAt.Format(time.RFC1123Z),
		Updated: user.UpdatedAt.Format(time.RFC1123Z),
		Entered: lastVisit,
		Name: nickname,
		GivenName: givenName,
		FamilyName: lastName,
		ParentName: parentName,
		BirthDate: birthDate,
		Gender: gender,
		Merged: merged,
		Phone: phone,
		Email: *email.Email,
	}
}

type RegisteredUserResource struct {
	Id string
	Created string
	Updated string
	Entered string
	Name string
	GivenName string
	ParentName string
	FamilyName string
	BirthDate string
	Gender string
	Merged string
	Phone string
	Email string
}