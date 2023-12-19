package handler

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"forum/config"
	"forum/dto"
	"forum/models"
	"forum/server/repositories"
	"forum/server/service"
	"forum/utils"
	"io"
	"net/http"
	"strings"
	"text/template"
	"time"

	"github.com/gofrs/uuid/v5"
)

func CreatePostHandler(w http.ResponseWriter, r *http.Request) {
	tokenData, err := service.AuthSrvice.VerifyToken(r)
	if err != nil {
		RenderErrorPage(http.StatusUnauthorized, w)
		return
	}
	switch r.Method {
	case http.MethodGet:
		tml, err := template.ParseFiles("./templates/createpost.html", "./templates/main.html", "./templates/profile.html")
		if err != nil {
			RenderErrorPage(http.StatusInternalServerError, w)
			return
		}
		tmpl := template.Must(tml, err)
		categories, _ := repositories.CategRepo.GetCategories()
		data := Data{
			Categories:      categories,
			IsAuthenticated: true,
			Username:        tokenData.Username,
		}
		err = tmpl.ExecuteTemplate(w, "main", data)
		fmt.Println(err)
	case http.MethodPost:
		err := r.ParseForm()
		if err != nil {
			fmt.Println(err)
			RenderErrorPage(http.StatusInternalServerError, w)
			return
		}
		cats := strings.Split(r.URL.Query().Get("categories"), ",")
		catsInt, err := utils.ParseArrayInt(cats)
		if err != nil {
			RenderErrorPage(http.StatusBadRequest, w)
			return
		}
		body := r.Body
		if err != nil {
			if err == io.EOF {
				RenderErrorPage(http.StatusBadRequest, w)
				return
			} else {
				fmt.Println(err)
				RenderErrorPage(http.StatusInternalServerError, w)
				return
			}
		}
		var post models.Post
		post.Username = tokenData.Username

		err = json.NewDecoder(body).Decode(&post)
		post.Body = strings.ReplaceAll(post.Body, "\"", `&quot;`)
		post.Title = strings.ReplaceAll(post.Title, "\"", `&quot;`)

		if strings.TrimSpace(post.Body) == "" || strings.TrimSpace(post.Title) == "" {
			RenderErrorPage(http.StatusBadRequest, w)
			return
		}

		if err != nil {
			fmt.Println(err)
			RenderErrorPage(http.StatusInternalServerError, w)
			return
		}
		post.UserId, err = uuid.FromString(tokenData.UserId)
		if err != nil {
			fmt.Println(err)
			RenderErrorPage(http.StatusInternalServerError, w)
			return
		}
		err = service.PostSrvice.NewPost(post, catsInt)
		if err != nil {
			fmt.Println(err)
			RenderErrorPage(http.StatusInternalServerError, w)
			return
		} else {
			http.Redirect(w, r, "/", http.StatusPermanentRedirect)
			return
		}
	}
}

func EditPostHandler(w http.ResponseWriter, r *http.Request) {
	tokenData, _ := service.AuthSrvice.VerifyToken(r)

	switch r.Method {
	case http.MethodGet:
		tml, err := template.ParseFiles("./templates/edit.post", "./templates/main.html")
		if err != nil {
			RenderErrorPage(http.StatusInternalServerError, w)
			return
		}
		tmpl := template.Must(tml, err)
		categories, _ := repositories.CategRepo.GetCategories()
		data := Data{
			Categories:      categories,
			IsAuthenticated: true,
			Username:        tokenData.Username,
		}
		err = tmpl.ExecuteTemplate(w, "main", data)
		fmt.Println(err)
	case http.MethodPost:
		cats := strings.Split(r.URL.Query().Get("categories"), ",")
		catsInt, err := utils.ParseArrayInt(cats)
		if err != nil {
			RenderErrorPage(http.StatusBadRequest, w)
			return
		}
		body := r.Body
		if _, err := io.ReadAll(body); err != nil {
			if err == io.EOF {
				RenderErrorPage(http.StatusBadRequest, w)
				return
			} else {
				fmt.Println(err)
				RenderErrorPage(http.StatusInternalServerError, w)
				return
			}
		}
		var post models.Post

		err = json.NewDecoder(body).Decode(&post)
		if err != nil {
			fmt.Println(err)
			RenderErrorPage(http.StatusInternalServerError, w)
			return
		}
		post.UserId, err = uuid.FromString(tokenData.UserId)
		if err != nil {
			fmt.Println(err)
			RenderErrorPage(http.StatusInternalServerError, w)
			return
		}
		err = service.PostSrvice.NewPost(post, catsInt)
		if err != nil {
			fmt.Println(err)
			RenderErrorPage(http.StatusInternalServerError, w)
			return
		}
		http.Redirect(w, r, "/HELLO", http.StatusPermanentRedirect)
		return
	}
}

func PostReactHandler(w http.ResponseWriter, r *http.Request) {
	tokenData, err := authService.VerifyToken(r)
	if err != nil {
		RenderErrorPage(http.StatusUnauthorized, w)
		return
	}
	react := r.URL.Query().Get("react")
	postId := r.URL.Query().Get("postid")
	if react == "" || postId == "" {
		RenderErrorPage(http.StatusBadRequest, w)
		return
	}
	post, err := service.PostSrvice.GetPost(postId)
	if err != nil {
		RenderErrorPage(http.StatusNotFound, w)
		return
	}
	switch r.Method {
	case http.MethodGet:
		var votes int
		if react == "LIKE" || react == "DISLIKE" {
			reaction, err := service.PostSrvice.GetUserPostReact(tokenData.UserId, postId)
			if err != nil {
				if err == sql.ErrNoRows {
					service.PostSrvice.SavePostReaction(post, react, tokenData.UserId)
				}
			} else if reaction.Reactions == react {
				repositories.ReactRepo.DeleteReaction(reaction.ReactId.String())
			} else {
				reaction.Reactions = react
				repositories.ReactRepo.UpdateReaction(reaction)

			}
			votes, err = service.PostSrvice.GetPostVotes(postId)
			if err != nil {
				RenderErrorPage(http.StatusInternalServerError, w)
			}
			json.NewEncoder(w).Encode(votes)
		} else {
			RenderErrorPage(http.StatusBadRequest, w)
			return
		}
	}
}

func GetAllPostHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		liked := r.URL.Query().Get("liked")
		order := r.URL.Query().Get("order")
		category := r.URL.Query().Get("category")
		created := r.URL.Query().Get("created")
		commented := r.URL.Query().Get("commented")
		var options = map[string]string{
			"liked":     liked,
			"order":     order,
			"category":  category,
			"created":   created,
			"commented": commented,
		}

		var data Data
		tokenData, err := authService.VerifyToken(r)
		var posts []dto.PostDTO
		if err != nil {
			posts, _ = service.PostSrvice.GetAllPosts(tokenData, options)
			data.Posts = posts
			tml, err := template.ParseFiles("templates/post.html")
			if err != nil {
				RenderErrorPage(http.StatusInternalServerError, w)
				return
			}
			tmpl := template.Must(tml, err)
			err = tmpl.Execute(w, data)
			if err != nil {
				fmt.Println(err)
			}
			return
		}

		posts, err = service.PostSrvice.GetAllPosts(tokenData, options)
		if err != nil {
			fmt.Println(err)
		}
		data.IsAuthenticated = true
		data.Posts = posts
		tml, err := template.ParseFiles("templates/post.html")
		if err != nil {
			RenderErrorPage(http.StatusInternalServerError, w)
			return
		}
		tmpl := template.Must(tml, err)
		err = tmpl.Execute(w, data)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func GetPostHandler(w http.ResponseWriter, r *http.Request) {
	var data Data
	tokenData, _ := authService.VerifyToken(r)
	postId := r.URL.Query().Get("postid")
	if id := uuid.FromStringOrNil(postId); id == uuid.Nil {
		RenderErrorPage(http.StatusBadRequest, w)
		return
	}
	post, err := service.PostSrvice.GetPost(postId)
	if err != nil {
		RenderErrorPage(http.StatusNotFound, w)
		return
	}
	creationDate, _ := time.Parse(config.Get("TIME_FORMAT").ToString(), post.CreatedAt)
	now, _ := time.Parse(config.Get("TIME_FORMAT").ToString(), time.Now().Format(config.Get("TIME_FORMAT").ToString()))
	age := utils.FormatDuration(now.Sub(creationDate))
	commCout, _ := service.ComSrvice.CommentRepo.GetCommentsCount(postId)
	react, _ := service.PostSrvice.GetUserPostReact(tokenData.UserId, post.PostId.String())
	votes, _ := service.PostSrvice.GetPostVotes(postId)
	categories, _ := service.PostSrvice.PostRepo.GetPostCategories(post.PostId.String())

	data.Posts = []dto.PostDTO{{Post: post, CommentsCount: commCout, Age: age, UserReact: react.Reactions, Votes: votes, Categories: categories}}
	tml, err := template.ParseFiles("templates/post.html")
	if err != nil {
		RenderErrorPage(http.StatusInternalServerError, w)
		return
	}
	tmpl := template.Must(tml, err)
	err = tmpl.Execute(w, data)
	if err != nil {
		fmt.Println(err)
	}
}
