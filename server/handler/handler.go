package handler

import (
	"encoding/json"
	"fmt"
	"forum/dto"
	"forum/models"
	"forum/server/repositories"
	"net/http"
	"text/template"
)

type Data struct {
	Categories      []models.Category
	IsAuthenticated bool
	Username        string
	Email           string
	ProfilePicture  string
	Posts           []dto.PostDTO
	Comments        []dto.CommentDTO
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		RenderErrorPage(http.StatusNotFound, w)
		return
	}
	categories, _ := repositories.CategRepo.GetCategories()

	var data Data
	tokenData, err := authService.VerifyToken(r)
	data.Categories = categories

	if err != nil {
		tmpl, err := template.ParseFiles(("./frontend/index.html"))
		if err != nil {
			RenderErrorPage(http.StatusInternalServerError, w)
			return
		}
		err = tmpl.Execute(w, data)
		fmt.Println(err)
		return
	}

	user, _ := repositories.UserRepo.GetUserById(tokenData.UserId)
	data.Email = user.Email
	data.Username = user.Username
	data.ProfilePicture = user.AvatarUrl
	data.IsAuthenticated = true
	tml, err := template.ParseFiles("templates/main.html", "templates/post_layout.html", "templates/profile.html")
	if err != nil {
		RenderErrorPage(http.StatusInternalServerError, w)
		return
	}
	tmpl := template.Must(tml, err)
	err = tmpl.ExecuteTemplate(w, "main", data)
	fmt.Println(err)
}

func StaticHandler(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	println(path)
	http.ServeFile(w, r, "./frontend/"+path)
}

func ServeAppHandler(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	println(path)
	http.ServeFile(w, r, "./frontend/"+path)
}

func Authorization(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if _, err := authService.VerifyToken(r); err != nil {
			// w.WriteHeader(http.StatusPermanentRedirect)
			if r.URL.Path == "/post/create" {
				RenderErrorPage(http.StatusUnauthorized, w)
				return
			}
			w.Header().Add("HX-Redirect", "/auth/signin")
			return
		}

		next.ServeHTTP(w, r)

	}
}

func RenderErrorPage(errorCode int, w http.ResponseWriter) {
	tmpl, err := template.ParseFiles("templates/error.html")
	if err != nil {
		w.WriteHeader(errorCode)
		fmt.Fprintln(w, ErrorMsgMap[errorCode])
		return
	}
	w.WriteHeader(errorCode)
	err = tmpl.Execute(w, ErrorMsgMap[errorCode])
	fmt.Println(err)
}

var ErrorMsgMap = map[int]Error{
	http.StatusBadRequest:          {http.StatusBadRequest, "Bad Request"},
	http.StatusNotFound:            {http.StatusNotFound, "Not Found"},
	http.StatusInternalServerError: {http.StatusInternalServerError, "Internal Server Error"},
	http.StatusMethodNotAllowed:    {http.StatusMethodNotAllowed, "Method Not Allowed"},
	http.StatusUnauthorized:        {http.StatusUnauthorized, "Unauthorized"},
}

type Error struct {
	Code int
	Msg  string
}

func RespondWithJSON(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}
}

func RespondWithError(w http.ResponseWriter, code int, msg string) {
	RespondWithJSON(w, code, map[string]string{"message": msg})
}
