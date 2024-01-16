package handler

import (
	"database/sql"
	"encoding/json"
	"errors"
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
	"time"

	"github.com/gofrs/uuid/v5"
)

func CreatePostHandler(w http.ResponseWriter, r *http.Request) {
	tokenData, err := service.AuthSrvice.VerifyToken(r)
	if err != nil {
		RespondWithError(w, http.StatusUnauthorized, "Unauthorized")
		return
	}

	switch r.Method {
	case http.MethodPost:
		err := r.ParseForm()
		if err != nil {
			fmt.Println(err)
			RespondWithError(w, http.StatusInternalServerError, "Internal Server Error")
			return
		}

		cats := strings.Split(r.URL.Query().Get("categories"), ",")
		catsInt, err := utils.ParseArrayInt(cats)
		if err != nil {
			RespondWithError(w, http.StatusBadRequest, "Bad Request")
			return
		}

		var post models.Post
		post.Username = tokenData.Username

		err = json.NewDecoder(r.Body).Decode(&post)
		if err != nil {
			fmt.Println(err)
			RespondWithError(w, http.StatusInternalServerError, "Internal Server Error")
			return
		}

		post.Body = strings.ReplaceAll(post.Body, "\"", `&quot;`)
		post.Title = strings.ReplaceAll(post.Title, "\"", `&quot;`)

		if strings.TrimSpace(post.Body) == "" || strings.TrimSpace(post.Title) == "" {
			RespondWithError(w, http.StatusBadRequest, "Bad Request")
			return
		}

		post.UserId, err = uuid.FromString(tokenData.UserId)
		if err != nil {
			fmt.Println(err)
			RespondWithError(w, http.StatusInternalServerError, "Internal Server Error")
			return
		}

		err = service.PostSrvice.NewPost(post, catsInt)
		if err != nil {
			fmt.Println(err)
			RespondWithError(w, http.StatusInternalServerError, "Internal Server Error")
			return
		}

		RespondWithJSON(w, http.StatusOK, post)

	default:
		RespondWithError(w, http.StatusMethodNotAllowed, "Method Not Allowed")
	}
}
func EditPostHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPut:
		tokenData, err := authService.VerifyToken(r)
		if err != nil {
			RespondWithError(w, http.StatusUnauthorized, "Unauthorized")
			return
		}

		postId := r.URL.Query().Get("postid")
		if postId == "" {
			RespondWithError(w, http.StatusBadRequest, "Bad Request")
			return
		}

		body := r.Body
		if _, err := io.ReadAll(body); err != nil {
			if err == io.EOF {
				RespondWithError(w, http.StatusBadRequest, "Bad Request")
				return
			} else {
				fmt.Println(err)
				RespondWithError(w, http.StatusInternalServerError, "Internal Server Error")
				return
			}
		}

		var updatedPost models.Post
		err = json.NewDecoder(body).Decode(&updatedPost)
		if err != nil {
			fmt.Println(err)
			RespondWithError(w, http.StatusInternalServerError, "Internal Server Error")
			return
		}

		// Check if the user is the owner of the post
		post, err := service.PostSrvice.GetPost(postId)
		if err != nil {
			RespondWithError(w, http.StatusInternalServerError, "Internal Server Error")
			return
		}
		if post.UserId != uuid.FromStringOrNil(tokenData.UserId) {
			RespondWithError(w, http.StatusUnauthorized, "Unauthorized")
			return
		}

		// Update the post
		updatedPost.PostId = post.PostId
		err = service.PostSrvice.UpdatePost(updatedPost)
		if err != nil {
			fmt.Println(err)
			RespondWithError(w, http.StatusInternalServerError, "Internal Server Error")
			return
		}

		// Return the updated post as response
		RespondWithJSON(w, http.StatusOK, updatedPost)

	default:
		RespondWithError(w, http.StatusMethodNotAllowed, "Method Not Allowed")
	}
}

func PostReactHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		tokenData, err := authService.VerifyToken(r)
		if err != nil {
			RespondWithError(w, http.StatusUnauthorized, "Unauthorized")
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
			RespondWithError(w, http.StatusInternalServerError, "Internal Server Error")
			return
		}
		var votes int
		if react == "LIKE" || react == "DISLIKE" {
			reaction, err := service.PostSrvice.GetUserPostReact(tokenData.UserId, postId)
			if err != nil {
				if errors.Is(err, sql.ErrNoRows) {
					err := service.PostSrvice.SavePostReaction(post, react, tokenData.UserId)
					if err != nil {
						RespondWithError(w, http.StatusInternalServerError, "Internal Server Error")
						return
					}
				}
			} else if reaction.Reactions == react {
				err := repositories.ReactRepo.DeleteReaction(reaction.ReactId.String())
				if err != nil {
					RespondWithError(w, http.StatusInternalServerError, "Internal Server Error")
					return
				}
			} else {
				reaction.Reactions = react
				err := repositories.ReactRepo.UpdateReaction(reaction)
				if err != nil {
					RespondWithError(w, http.StatusInternalServerError, "Internal Server Error")
					return
				}
			}
			votes, err = service.PostSrvice.GetPostVotes(postId)
			if err != nil {
				RespondWithError(w, http.StatusInternalServerError, "Internal Server Error")
				return
			}
			RespondWithJSON(w, http.StatusOK, votes)
		} else {
			RespondWithError(w, http.StatusBadRequest, "Bad Request")
			return
		}
	default:
		RespondWithError(w, http.StatusMethodNotAllowed, "Method Not Allowed")
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

		tokenData, err := authService.VerifyToken(r)
		if err != nil {
			RespondWithError(w, http.StatusUnauthorized, "Unauthorized")
			return
		}

		posts, err := service.PostSrvice.GetAllPosts(tokenData, options)
		if err != nil {
			RespondWithError(w, http.StatusInternalServerError, "Internal Server Error")
			return
		}

		response := struct {
			Posts []dto.PostDTO `json:"posts"`
		}{
			Posts: posts,
		}

		RespondWithJSON(w, http.StatusOK, response)
		w.Header().Set("Access-Control-Allow-Origin", "*")

	default:
		RespondWithError(w, http.StatusMethodNotAllowed, "Method Not Allowed")
	}
}

func GetPostHandler(w http.ResponseWriter, r *http.Request) {
	postID := r.URL.Query().Get("postid")
	tokenData, err := authService.VerifyToken(r)
	if err != nil {
		RespondWithError(w, http.StatusUnauthorized, "Unauthorized")
		return
	}
	if postID == "" {
		RespondWithError(w, http.StatusBadRequest, "Bad Request")
		return
	}

	post, err := service.PostSrvice.GetPost(postID)
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}

	creationDate, _ := time.Parse(config.Get("TIME_FORMAT").ToString(), post.CreatedAt)
	now, _ := time.Parse(config.Get("TIME_FORMAT").ToString(), time.Now().Format(config.Get("TIME_FORMAT").ToString()))
	age := utils.FormatDuration(now.Sub(creationDate))
	commCount, _ := service.ComSrvice.CommentRepo.GetCommentsCount(postID)
	react, _ := service.PostSrvice.GetUserPostReact(tokenData.UserId, post.PostId.String())
	votes, _ := service.PostSrvice.GetPostVotes(postID)
	categories, _ := service.PostSrvice.PostRepo.GetPostCategories(post.PostId.String())

	postDTO := dto.PostDTO{
		Post:          post,
		CommentsCount: commCount,
		Age:           age,
		UserReact:     react.Reactions,
		Votes:         votes,
		Categories:    categories,
	}

	RespondWithJSON(w, http.StatusOK, postDTO)
	w.Header().Set("Access-Control-Allow-Origin", "*")

}
