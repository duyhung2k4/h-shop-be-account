package controller

import (
	"app/config"
	"app/dto/request"
	"app/model"
	"app/service"
	"app/utils"
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/go-chi/render"
	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
)

type accessController struct {
	loginGoogleService service.LoginGoogleService
	jwtUtils           utils.JwtUtils
	rdb                *redis.Client
	utils              utils.JwtUtils
}

type AccessController interface {
	LoginGoogle(w http.ResponseWriter, r *http.Request)
	RefreshToken(w http.ResponseWriter, r *http.Request)
}

func (a *accessController) LoginGoogle(w http.ResponseWriter, r *http.Request) {
	var userRequest request.LoginGoogleRequest
	if err := json.NewDecoder(r.Body).Decode(&userRequest); err != nil {
		badRequest(w, r, err)
		return
	}

	isExist, user, err := a.loginGoogleService.CheckExistUser(userRequest)

	if err != nil {
		log.Println("Error check exit user")
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
		profile, err := a.loginGoogleService.CreateProfile(userRequest)
		if err != nil {
			log.Println("Error create profile")
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
	accessData["exp"] = time.Now().Add(3 * time.Hour).Unix()
	accessToken, errAccessToken := a.jwtUtils.JwtEncode(accessData)
	if errAccessToken != nil {
		log.Println("Error create accessToken")
		internalServerError(w, r, errAccessToken)
		return
	}

	refreshData := mapData
	refreshData["uuid"] = uuid.New()
	refreshData["exp"] = time.Now().Add(3 * 3 * time.Hour).Unix()
	refreshToken, errRefreshToken := a.jwtUtils.JwtEncode(refreshData)
	if errRefreshToken != nil {
		log.Println("Error create refreshToken")
		internalServerError(w, r, errRefreshToken)
		return
	}

	errSetKeyAccessToken := a.rdb.Set(context.Background(), "access_token:"+strconv.Itoa(int(profileResponse.ID)), accessToken, 24*time.Hour).Err()
	if errSetKeyAccessToken != nil {
		log.Println("Error save accessToken")
		internalServerError(w, r, errSetKeyAccessToken)
		return
	}
	errSetKeyRefreshToken := a.rdb.Set(context.Background(), "refresh_token:"+strconv.Itoa(int(profileResponse.ID)), refreshToken, 3*24*time.Hour).Err()
	if errSetKeyRefreshToken != nil {
		log.Println("Error save refreshToken")
		internalServerError(w, r, errSetKeyRefreshToken)
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

func (a *accessController) RefreshToken(w http.ResponseWriter, r *http.Request) {
	tokenString := strings.Split(r.Header.Get("Authorization"), " ")[1]
	mapDataRequest, errMapData := a.utils.JwtDecode(tokenString)

	if errMapData != nil {
		internalServerError(w, r, errMapData)
		return
	}

	profileId := uint(mapDataRequest["profile_id"].(float64))
	profileResponse, errProfile := a.loginGoogleService.GetProfile(uint(profileId))
	if errProfile != nil {
		internalServerError(w, r, errProfile)
		return
	}

	mapData := map[string]interface{}{
		"profile_id": profileResponse.ID,
		"email":      profileResponse.Email,
		"role":       profileResponse.User.Role.Code,
		"sub":        profileResponse.Sub,
	}

	accessData := mapData
	accessData["uuid"] = uuid.New()
	accessData["exp"] = time.Now().Add(3 * time.Hour).Unix()
	accessToken, errAccessToken := a.jwtUtils.JwtEncode(accessData)
	if errAccessToken != nil {
		internalServerError(w, r, errAccessToken)
		return
	}

	refreshData := mapData
	refreshData["uuid"] = uuid.New()
	refreshData["exp"] = time.Now().Add(3 * 3 * time.Hour).Unix()
	refreshToken, errRefreshToken := a.jwtUtils.JwtEncode(refreshData)
	if errRefreshToken != nil {
		internalServerError(w, r, errRefreshToken)
		return
	}

	errSetKeyAccessToken := a.rdb.Set(context.Background(), "access_token:"+strconv.Itoa(int(profileResponse.ID)), accessToken, 24*time.Hour).Err()
	if errSetKeyAccessToken != nil {
		internalServerError(w, r, errSetKeyAccessToken)
		return
	}
	errSetKeyRefreshToken := a.rdb.Set(context.Background(), "refresh_token:"+strconv.Itoa(int(profileResponse.ID)), refreshToken, 3*24*time.Hour).Err()
	if errSetKeyRefreshToken != nil {
		internalServerError(w, r, errSetKeyRefreshToken)
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
		rdb:                config.GetRDB(),
		utils:              utils.NewJwtUtils(),
	}
}
