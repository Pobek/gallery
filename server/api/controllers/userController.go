package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/pobek/gallery/server/api/models"
	"github.com/pobek/gallery/server/api/responses"
	"github.com/pobek/gallery/server/utils"
)

// UserSignUp - the signup controller for creating new users
func (app *App) UserSignUp(respWriter http.ResponseWriter, req *http.Request) {
	var resp = map[string]interface{}{
		"status":  "success",
		"message": "Registed Successfully",
	}

	user := &models.User{}
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		responses.ERROR(respWriter, http.StatusBadRequest, err)
		return
	}

	err = json.Unmarshal(body, &user)
	if err != nil {
		responses.ERROR(respWriter, http.StatusBadRequest, err)
	}

	usr, _ := user.GetUser(app.DB)
	if usr != nil {
		resp["status"] = "failed"
		resp["message"] = "User already registered, please login"
		responses.JSON(respWriter, http.StatusBadRequest, resp)
		return
	}

	user.Prepare()

	err = user.Validate("")
	if err != nil {
		responses.ERROR(respWriter, http.StatusBadRequest, err)
		return
	}

	userCreated, err := user.SaveUser(app.DB)
	if err != nil {
		responses.ERROR(respWriter, http.StatusBadRequest, err)
		return
	}

	resp["user"] = userCreated
	responses.JSON(respWriter, http.StatusCreated, resp)
	return
}

// Login - the login controller
func (app *App) Login(respWriter http.ResponseWriter, req *http.Request) {
	var resp = map[string]interface{}{
		"status":  "Success",
		"message": "Login successfully",
	}
	// if req.Method == "OPTIONS" {
	// 	resp["status"] = "OPTIONS"
	// 	resp["message"] = "POST, OPTIONS"
	// 	responses.JSON(respWriter, http.StatusOK, resp)
	// 	return
	// }

	user := &models.User{}
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		responses.ERROR(respWriter, http.StatusBadRequest, err)
		return
	}

	err = json.Unmarshal(body, &user)
	if err != nil {
		responses.ERROR(respWriter, http.StatusBadRequest, err)
		return
	}

	user.Prepare()

	err = user.Validate("login")
	if err != nil {
		responses.ERROR(respWriter, http.StatusBadRequest, err)
		return
	}

	usr, err := user.GetUser(app.DB)
	if err != nil {
		responses.ERROR(respWriter, http.StatusInternalServerError, err)
		return
	}

	if usr == nil { // user is not registered
		resp["status"] = "failed"
		resp["message"] = "Login failed, please signup"
		responses.JSON(respWriter, http.StatusBadRequest, resp)
		return
	}

	err = models.CheckPasswordHash(user.Password, usr.Password)
	if err != nil {
		resp["status"] = "failed"
		resp["message"] = "Login failed, please try again"
		responses.JSON(respWriter, http.StatusForbidden, resp)
		return
	}

	token, err := utils.EncodeAuthToken(usr.ID)
	if err != nil {
		responses.ERROR(respWriter, http.StatusBadRequest, err)
		return
	}

	resp["token"] = token
	responses.JSON(respWriter, http.StatusOK, resp)
	return
}
