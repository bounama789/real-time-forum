package handler

import (
	"fmt"
	"forum/backend/dto"
	"forum/backend/models"
	"forum/backend/server/cors"
	"net/http"
	"os"
	"path"
	"regexp"
	"strings"
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
	cors.SetCors(&w)

	re := regexp.MustCompile(`.*\.(js|css|ico|svg)$`)
	p := r.URL.Path

	if re.MatchString(p) {
		p2,_ := strings.CutPrefix(p,"/src")
		http.ServeFile(w, r, path.Join("./frontend/src", p2))
		return
	}

	// w.WriteHeader(200)
	http.ServeFile(w, r, "./frontend/index.html")

}

func StaticHandler(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	println(path)
	http.ServeFile(w, r, "./frontend/"+path)
}

// func ServeAppHandler(w http.ResponseWriter, r *http.Request) {
// 	path := r.URL.Path
// 	println(path)
// 	http.ServeFile(w, r, "./frontend/"+path)
// }

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

func ServeApp(w http.ResponseWriter, r *http.Request) {
	// return func(w http.ResponseWriter, r *http.Request) {
	infos, err := os.Stat("./frontend/")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(infos)
	}
	http.ServeFile(w, r, "./frontend/")

	// }
}
