package handlers

import (
	"github.com/MetalRex101/auth-server/app/validators"
	"github.com/MetalRex101/auth-server/app/services"
	"github.com/MetalRex101/auth-server/app/models"
	"github.com/labstack/echo"
	"net/http"
	"github.com/jinzhu/gorm"
	"github.com/elgs/gostrgen"
	"github.com/MetalRex101/auth-server/app/resources"
)

type RegisterHandler struct{
	osm  services.IOauthSessionManager
	ocm  services.IOauthClientManager
	um   services.IUserManager
	em   services.IEmailManager
	uMer services.IUserMerger
	db *gorm.DB
}

func NewRegisterHandler(
	osm services.IOauthSessionManager,
	ocm services.IOauthClientManager,
	um services.IUserManager,
	em services.IEmailManager,
	uMer services.IUserMerger,
	db *gorm.DB,
) *RegisterHandler {
	return &RegisterHandler{osm, ocm, um, em, uMer, db}
}

func (reg *RegisterHandler) Handle (c echo.Context) error {
	var client *models.Client
	var authUser *models.User

	if err := validators.Request.OauthTID(c); err != nil {
		return err
	}

	accessToken, err := validators.Request.GetAccessToken(false, c)
	if err != nil {
		return err
	}

	if accessToken != "" {
		oauthSess, err := reg.osm.FindByToken(accessToken)
		if err != nil {
			return err
		}

		client, err = reg.ocm.GetClientFromSession(oauthSess)
		if err != nil {
			return err
		}

		authUser, err = reg.um.GetUserFromSession(oauthSess)
		if err != nil {
			return err
		}
	} else {
		clientID, err := validators.Request.GetClientId(true, c)
		if err != nil {
			return err
		}

		clientSecret, err := validators.Request.GetClientSecret(true, c)
		if err != nil {
			return err
		}

		client, err = reg.ocm.GetForApi(clientID, clientSecret, c.Request().RemoteAddr)
		if err != nil {
			return err
		}
	}

	if err = validators.Client.HasScope(client, []string{"oauth"}); err != nil {
		return err
	}

	password, err := validators.Request.GetPassword(false, c, false)
	if err != nil {
		return err
	}

	if err = validators.Password.ValidatePassword(password); err != nil {
		return err
	}

	gender,err := validators.Request.GetGender(false, c)

	emailAddr, _ := validators.Request.GetEmail(false, c, false)

	if err := validators.Email.ValidateEmail(emailAddr); err != nil {
		return err
	}

	if err = reg.em.EmailNotUsed(emailAddr); err != nil {
		return err
	}

	nick := c.QueryParam("name")
	firstName := c.QueryParam("given_name")
	lastName := c.QueryParam("family_name")
	fatherName := c.QueryParam("parent_name")
	birthDate, err := validators.Request.GetBirthDate(false, c)
	if err != nil {
		return err
	}
	genderLetter := string([]rune(gender)[0])

	user := &models.User{
		Nickname: &nick,
		FirstName: &firstName,
		LastName: &lastName,
		FatherName: &fatherName,
		Status: true,
		BirthDate: birthDate,
		Gender: &genderLetter,
	}
	if err = reg.db.Create(user).Error; err != nil {
		return err
	}

	emailCode, _ := gostrgen.RandGen(64, gostrgen.All, "", "")

	userEmail := &models.Email{Email: &emailAddr, Status: true, Code: &emailCode}

	reg.db.Model(user).Association("Emails").Append(userEmail)

	reg.um.GetDefaultEmail(user.ID, true, c)

	hashedPass := services.HashPassword(password)
	userPass := &models.Password{Password: &hashedPass}

	reg.db.Model(user).Association("Passwords").Append(userPass)

	if authUser != nil {
		reg.uMer.MergerUsers(user, authUser, c)
	}

	return c.JSON(http.StatusOK, resources.NewRegisteredResource(user, reg.um.(*services.UserManager), c))
}
