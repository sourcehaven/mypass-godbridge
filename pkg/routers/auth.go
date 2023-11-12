package routers

import (
	"errors"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/sourcehaven/mypass-godbridge/pkg/app"
	"github.com/sourcehaven/mypass-godbridge/pkg/bytes"
	"github.com/sourcehaven/mypass-godbridge/pkg/controllers"
	"github.com/sourcehaven/mypass-godbridge/pkg/models"
	"github.com/sourcehaven/mypass-godbridge/pkg/schemas"
	"github.com/sourcehaven/mypass-godbridge/pkg/security/crypto"
	"github.com/sourcehaven/mypass-godbridge/pkg/security/jwt"
	"github.com/sourcehaven/mypass-godbridge/pkg/utils"
	"net/http"
	"strings"
)

var controller *controllers.UserController
var validating *validator.Validate

func init() {
	controller = &controllers.UserController{DB: app.DB}
	validating = app.Validator
}

// RegisterUser  godoc
// @Summary      Registration endpoint
// @Description  Responds with created status
// @Tags         auth
// @Param        user body schemas.UserReg true "Register user"
// @Produce      json
// @Success      201
// @Router       /auth/register [post]
func RegisterUser(ctx *fiber.Ctx) (err error) {
	userReg := &schemas.UserReg{}
	if err = ctx.BodyParser(&userReg); err != nil {
		return
	}

	if err = validating.Struct(userReg); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   utils.ValidatorErrors(err),
		})
	}
	randBytes, err := crypto.SecureRandomBytes(8)
	if err != nil {
		return
	}
	randPass := bytes.Std.EncodeToString(randBytes)
	randPass = strings.TrimSuffix(randPass, "=")
	if err = controller.Create(models.User{
		Username:  userReg.Username,
		Email:     userReg.Email,
		Password:  randPass,
		Firstname: userReg.Firstname,
		Lastname:  userReg.Lastname,
	}); err != nil {
		return
	}
	token, err := jwt.CreateAccessToken(userReg.Username, &jwt.TokenOptions{
		Fresh:            true,
		JwtAccessExpires: app.Cfg.JwtAccessExpires,
		JwtSigningMethod: app.Cfg.JwtSigningMethod,
		JwtAccessKey:     app.Cfg.JwtAccessKey,
	})
	if err != nil {
		return
	}
	if err = ctx.Status(http.StatusCreated).JSON(&schemas.UserRegOk{Token: token, Password: randPass}); err != nil {
		return
	}
	return
}

// ActivateUser  godoc
// @Summary      User activation endpoint
// @Description  Activates a freshly registered user based on activation link and initial password.
// @Tags         auth
// @Param        token path string true "ActivationToken token"
// @Param        activation body schemas.UserActivation true "User activation form containing old and new password"
// @Produce      json
// @Success      200
// @Router       /auth/activate/{token} [post]
func ActivateUser(ctx *fiber.Ctx) (err error) {
	tokenLink := &schemas.ActivationToken{}
	activation := &schemas.UserActivation{}
	if err = ctx.ParamsParser(tokenLink); err != nil {
		return
	}
	if err = ctx.BodyParser(activation); err != nil {
		return
	}
	if err = validating.Struct(tokenLink); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   utils.ValidatorErrors(err),
		})
	}
	if err = validating.Struct(activation); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   utils.ValidatorErrors(err),
		})
	}

	claims, err := jwt.ValidateAccessToken(tokenLink.Token, &jwt.TokenOptions{
		Fresh:            true,
		JwtAccessExpires: app.Cfg.JwtAccessExpires,
		JwtSigningMethod: app.Cfg.JwtSigningMethod,
		JwtAccessKey:     app.Cfg.JwtAccessKey,
	})
	if errors.Is(err, jwt.ErrTokenExpired) {
		// Token expired -> extract information if valid, then delete user if still exists
		claims, err = jwt.ValidateAccessToken(tokenLink.Token, &jwt.TokenOptions{
			Fresh:            true,
			JwtAccessExpires: app.Cfg.JwtAccessExpires,
			JwtSigningMethod: app.Cfg.JwtSigningMethod,
			JwtAccessKey:     app.Cfg.JwtAccessKey,
			AllowExpired:     true,
		})
		if err != nil {
			return
		}
		identity := claims.Subject

		if err = controller.DeleteByUsername(identity); err != nil {
			return
		}
		return
	}
	// Token has not been expired, move on
	identity := claims.Subject
	if err = controller.ActivateByUsername(identity, activation.OldPassword, activation.NewPassword); err != nil {
		return
	}
	if err = ctx.SendStatus(http.StatusOK); err != nil {
		return
	}
	return
}

// LoginUser     godoc
// @Summary      User login/authentication endpoint
// @Description  Authenticates user with given username and password. If correct, gives out access and refresh tokens.
// @Tags         auth
// @Param        user body schemas.UserLogin true "User login form containing username and plain password"
// @Produce      json
// @Success      201
// @Router       /auth/login [post]
func LoginUser(ctx *fiber.Ctx) (err error) {
	login := &schemas.UserLogin{}
	if err = ctx.BodyParser(login); err != nil {
		return
	}
	if err = validating.Struct(login); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   utils.ValidatorErrors(err),
		})
	}
	if err = controller.Authenticate(login.Username, login.Password); err != nil {
		return
	}
	if err = ctx.Status(http.StatusCreated).JSON(
		struct {
			accessToken  string
			refreshToken string
		}{"template-access", "template-refresh"}); err != nil {
		return
	}
	return
}
