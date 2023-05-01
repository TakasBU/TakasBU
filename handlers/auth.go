package handlers

import (
	"net/http"
	"strings"
	"time"

	"github.com/TakasBU/TakasBU/databases"
	"github.com/TakasBU/TakasBU/initializers"
	"github.com/TakasBU/TakasBU/models"
	"github.com/TakasBU/TakasBU/utils"
	"github.com/thanhpk/randstr"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type AuthController struct {
	DB *gorm.DB
}

func NewAuthController() *AuthController {
	db := databases.InitDb()
	db.AutoMigrate(&models.User{})
	return &AuthController{DB: db}
}
func (ac *AuthController) SignUpUser(ctx echo.Context) error {
	var payload *models.SignUpInput

	if err := ctx.Bind(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, ErrorResponse{Code: http.StatusBadRequest, Message: "Passwords do not match"})
		return err
	}

	if payload.Password != payload.PasswordConfirm {
		ctx.JSON(http.StatusBadRequest, ErrorResponse{Code: http.StatusBadRequest, Message: "ben alttaki"})
		return nil
	}

	hashedPassword, err := utils.HashPassword(payload.Password)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, ErrorResponse{Code: http.StatusBadGateway, Message: err.Error()})
		return err

	}

	now := time.Now()
	newUser := models.User{
		Name:      payload.Name,
		Email:     strings.ToLower(payload.Email),
		Password:  hashedPassword,
		Role:      "user",
		Verified:  false,
		Photo:     payload.Photo,
		Provider:  "local",
		CreatedAt: now,
		UpdatedAt: now,
	}

	result := ac.DB.Create(&newUser)

	if result.Error != nil && strings.Contains(result.Error.Error(), "duplicate key value violates unique") {
		ctx.JSON(http.StatusConflict, ErrorResponse{Code: http.StatusConflict, Message: "this email already exits"})
		return result.Error

	} else if result.Error != nil {
		ctx.JSON(http.StatusBadGateway, ErrorResponse{Code: http.StatusBadGateway, Message: "kotu seyler oldu"})
		return result.Error
	}

	config, _ := initializers.LoadConfig(".")

	// Generate Verification Code
	code := randstr.String(20)

	verification_code := utils.Encode(code)

	// Update User in Database
	newUser.VerificationCode = verification_code
	ac.DB.Save(newUser)

	var firstName = newUser.Name

	if strings.Contains(firstName, " ") {
		firstName = strings.Split(firstName, " ")[1]
	}

	// 👇 Send Email
	emailData := utils.EmailData{
		URL:       config.ClientOrigin + "/verifyemail/" + code,
		FirstName: firstName,
		Subject:   "Your account verification code",
	}

	utils.SendEmail(&newUser, &emailData)

	message := "We sent an email with a verification code to " + newUser.Email
	//TODO Email tamamlanıcak
	ctx.JSON(http.StatusCreated, ErrorResponse{Code: http.StatusCreated, Message: message})
	return ctx.JSON(http.StatusOK, payload)

}

func (ac *AuthController) SignInUser(ctx echo.Context) error {
	var payload *models.SignInInput

	if err := ctx.Bind(&payload); err != nil {
		ctx.JSON(http.StatusInternalServerError, ErrorResponse{Code: http.StatusBadRequest, Message: "ben malım"})
		return err
	}

	var user models.User
	result := ac.DB.First(&user, "email = ?", strings.ToLower(payload.Email))
	if result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, ErrorResponse{Code: http.StatusBadRequest, Message: "ben malım2"})
		return result.Error
	}

	// if !user.Verified {
	// 	ctx.JSON(http.StatusInternalServerError, ErrorResponse{Code: http.StatusBadRequest, Message: "ben malım 3"})
	// 	/*ctx.JSON(http.StatusForbidden, gin.H{"status": "fail", "message": "Please verify your email"}) mesajları nasıl yazıcam bilimiyorum*/
	// 	return result.Error
	// }

	if err := utils.VerifyPassword(user.Password, payload.Password); err != nil {
		ctx.JSON(http.StatusInternalServerError, ErrorResponse{Code: http.StatusBadRequest, Message: "ben malım 4"})
		return result.Error
	}

	config, _ := initializers.LoadConfig(".")

	// Generate Token
	token, err := utils.GenerateToken(config.TokenExpiresIn, user.ID, config.TokenSecret)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, ErrorResponse{Code: http.StatusBadRequest, Message: err.Error()})
		return result.Error
	}

	cookie := http.Cookie{}
	cookie.Name = "token"
	cookie.Value = token
	cookie.Expires = time.Now().Add(time.Duration(config.TokenMaxAge * 60))
	cookie.Domain = "localhost"
	cookie.Secure = false
	cookie.HttpOnly = true
	cookie.Path = "/"

	ctx.SetCookie(&cookie)

	ctx.JSON(http.StatusOK, ErrorResponse{Code: http.StatusOK, Message: token})
	return ctx.JSON(http.StatusOK, payload)
}
