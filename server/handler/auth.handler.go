package handler

import (
	"encoding/json"
	"fmt"
	"forum/config"
	"forum/models"
	"forum/server/service"
	"forum/utils"
	"io"
	"net/http"
	"strings"
	"text/template"
	"time"
)

var authService = service.AuthSrvice

func SignUpHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		tmpl, err := template.ParseFiles("./templates/signup.html")
		if err != nil {
			RenderErrorPage(http.StatusInternalServerError, w)
			return
		}
		err = tmpl.Execute(w, nil)
		if err != nil {
			fmt.Println(err)
		}
	case http.MethodPost:
		body := r.Body
		content, err := io.ReadAll(body)
		if err != nil {
			if err == io.EOF {
				w.WriteHeader(http.StatusBadRequest)
				return
			} else {
				RenderErrorPage(http.StatusInternalServerError, w)
				return
			}
		}
		var user models.User
		err = json.Unmarshal(content, &user)
		if err != nil {
			RenderErrorPage(http.StatusInternalServerError, w)
			return
		} else {
			if err := utils.VerifyUsername(user.Username); err != nil {
				w.WriteHeader(http.StatusUnprocessableEntity)
				json.NewEncoder(w).Encode(map[string]string{"msg": err.Error()})
				return
			} else if err := utils.IsValidEmail(user.Email); err != nil {
				w.WriteHeader(http.StatusUnprocessableEntity)
				json.NewEncoder(w).Encode(map[string]string{"msg": err.Error()})
				return
			} else if err := utils.VerifyName(user.Firstname); err != nil {
				w.WriteHeader(http.StatusUnprocessableEntity)
				json.NewEncoder(w).Encode(map[string]string{"msg": err.Error()})
				return
			} else if err := utils.VerifyName(user.Lastname); err != nil {
				w.WriteHeader(http.StatusUnprocessableEntity)
				json.NewEncoder(w).Encode(map[string]string{"msg": err.Error()})
				return
			}
		}
		user.CreatedAt = time.Now().Format(string(config.Get("TIME_FORMAT")))
		user.UpdatedAt = user.CreatedAt

		err = authService.CreateNewUser(&user)
		if err != nil {
			w.WriteHeader(http.StatusUnprocessableEntity)
			json.NewEncoder(w).Encode(map[string]string{"msg": err.Error()})
		} else {
			newSess, cookie := authService.GenCookieSession(w, user, r)
			newSess.UserId = user.UserId
			http.SetCookie(w, &cookie)
			err := authService.SessRepo.SaveSession(newSess)
			if err != nil {
				json.NewEncoder(w).Encode(map[string]string{"msg": err.Error()})
				return
			}
			http.Redirect(w, r, "/", http.StatusPermanentRedirect)

		}
	}
}

func SignInHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		tmpl,err := template.ParseFiles("./templates/login.html")
		if err != nil {
			RenderErrorPage(http.StatusInternalServerError, w)
			return
		}
		tmpl.Execute(w, nil)
	case http.MethodPost:
		var credentials = make(map[string]string, 2)
		if content, err := io.ReadAll(r.Body); err != nil {
			w.WriteHeader(http.StatusUnprocessableEntity)
			json.NewEncoder(w).Encode(map[string]string{"msg": err.Error()})
			return
		} else {
			err := json.Unmarshal(content, &credentials)
			if err != nil {
				w.WriteHeader(http.StatusUnprocessableEntity)
				json.NewEncoder(w).Encode(map[string]string{"msg": err.Error()})
				return
			}
			identifiant := strings.TrimSpace(credentials["identifiant"])
			credentials["identifiant"]=identifiant
			if err := utils.VerifyUsername(identifiant); err != nil && !strings.Contains(identifiant, "@") {
				w.WriteHeader(http.StatusUnprocessableEntity)
				json.NewEncoder(w).Encode(map[string]string{"msg": err.Error()})
				return
			} else if err := utils.IsValidEmail(identifiant); err != nil && strings.Contains(identifiant, "@") {
				w.WriteHeader(http.StatusUnprocessableEntity)
				json.NewEncoder(w).Encode(map[string]string{"msg": err.Error()})
				return
			}
		}

		user, err := authService.CheckCredentials(credentials)
		if err != nil {
			w.WriteHeader(401)
			fmt.Fprintln(w, err)
			return
		}
		newSess, cookie := authService.GenCookieSession(w, user, r)
		http.SetCookie(w, &cookie)
		authService.SessRepo.SaveSession(newSess)

		authService.RemExistingUsrSession(user.UserId.String())
		http.Redirect(w, r, "/", http.StatusMovedPermanently)

	default:
		RenderErrorPage(http.StatusMethodNotAllowed, w)
		return
	}
}

func SignOutHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodDelete:

		cookie := &http.Cookie{
			Name:     config.Get("COOKIE_NAME").ToString(),
			Value:    "",
			Path:     "/",
			MaxAge:   -1,
			HttpOnly: true,
		}
		http.SetCookie(w, cookie)

		tokenvalue, err := authService.VerifyToken(r)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		authService.SessRepo.DeleteSession(tokenvalue.SessId)
		w.Header().Add("HX-Redirect", "/")

	default:
		RenderErrorPage(http.StatusMethodNotAllowed, w)
	}
}

func VerifyEmailHandler(w http.ResponseWriter, r *http.Request) {
	email := r.URL.Query().Get("email")
	_, err := authService.UserRepo.GetUserByEmail(email)
	if err == nil {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{"msg": "email already in use"})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"msg": "valid"})

}

func VerifyUsernameHandler(w http.ResponseWriter, r *http.Request) {
	username := r.URL.Query().Get("username")
	_, err := authService.UserRepo.GetUserByUsername(username)
	if err == nil {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{"msg": "username already in use"})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"msg": "valid"})

}
