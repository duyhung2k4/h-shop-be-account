package controller

import (
	"app/dto/request"
	"app/model"
	"app/service"
	"app/utils"
	"encoding/json"
	"net/http"
	"time"

	"github.com/go-chi/render"
	"github.com/google/uuid"
)

type accessController struct {
	loginGoogleService service.LoginGoogleService
	jwtUtils           utils.JwtUtils
}

type AccessController interface {
	LoginGoogle(w http.ResponseWriter, r *http.Request)
}

func (a *accessController) LoginGoogle(w http.ResponseWriter, r *http.Request) {
	var userRequest request.LoginGoogleRequest
	if err := json.NewDecoder(r.Body).Decode(&userRequest); err != nil {
		badRequest(w, r, err)
		return
	}

	isExist, user, err := a.loginGoogleService.CheckExistUser(userRequest)

	if err != nil {
		internalServerError(w, r, err)
		return
	}

	var mapData map[string]interface{}
	var profileResponse model.Profile

	if isExist {
		mapData = map[string]interface{}{
			"profile_id": user.Profile.ID,
			"email":      user.Profile.Email,
			"role":       user.Profile.User.Role.Code,
			"sub":        user.Profile.Sub,
		}

		profileResponse = *user.Profile
	} else {
		profile, err := a.loginGoogleService.CreateProfile(userRequest, model.USER)
		if err != nil {
			internalServerError(w, r, err)
			return
		}

		mapData = map[string]interface{}{
			"profile_id": profile.ID,
			"email":      profile.Email,
			"role":       profile.User.Role,
			"sub":        profile.Sub,
		}

		profileResponse = *profile
	}

	accessData := mapData
	accessData["uuid"] = uuid.New()
	accessData["exp"] = time.Now().Add(24 * time.Hour).Unix()
	accessToken, errAccessToken := a.jwtUtils.JwtEncode(accessData)
	if errAccessToken != nil {
		internalServerError(w, r, errAccessToken)
		return
	}

	refreshData := mapData
	refreshData["uuid"] = uuid.New()
	refreshData["exp"] = time.Now().Add(24 * time.Hour).Unix()
	refreshToken, errRefreshToken := a.jwtUtils.JwtEncode(refreshData)
	if errRefreshToken != nil {
		internalServerError(w, r, errRefreshToken)
		return
	}

	res := Response{
		Data: map[string]interface{}{
			"profile":      profileResponse,
			"accessToken":  accessToken,
			"refreshToken": refreshToken,
		},
	}

	render.JSON(w, r, res)
}

func NewAccess() AccessController {
	return &accessController{
		loginGoogleService: service.NewGoginGoogleService(),
		jwtUtils:           utils.NewJwtUtils(),
	}
}
