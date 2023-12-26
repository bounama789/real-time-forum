package handler

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"forum/models"
	repo "forum/server/repositories"
	"forum/server/service"
	"io"
	"net/http"
	"strings"

	"github.com/gofrs/uuid/v5"
)

func CreateCommentHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		// Get token data from authentication cookie
		cookie, err := r.Cookie("auth-cookie")
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		tokenData, err := service.AuthSrvice.GetTokenData(cookie.Value)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		// Read request body
		body := r.Body
		content, err := io.ReadAll(body)
		if err != nil {
			if err == io.EOF {
				w.WriteHeader(http.StatusBadRequest)
				return
			} else {
				fmt.Println(err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
		}

		// Unmarshal request body into comment struct
		var comment models.Comment
		err = json.Unmarshal(content, &comment)
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		// Set user ID and username from token data
		comment.UserId, err = uuid.FromString(tokenData.UserId)
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		comment.Username = tokenData.Username

		// Validate comment body
		if strings.TrimSpace(comment.Body) == "" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		// Create comment in the database
		err = service.ComSrvice.NewComment(comment)
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		} else {
			w.WriteHeader(http.StatusCreated)
		}
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func EditCommentHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPut:
		// Get comment ID from URL parameter
		commentID := r.URL.Query().Get("commentid")
		if commentID == "" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		// Get token data from authentication cookie
		cookie, err := r.Cookie("auth-cookie")
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		tokenData, err := service.AuthSrvice.GetTokenData(cookie.Value)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		// Read request body
		body := r.Body
		content, err := io.ReadAll(body)
		if err != nil {
			if err == io.EOF {
				w.WriteHeader(http.StatusBadRequest)
				return
			} else {
				fmt.Println(err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
		}

		// Unmarshal request body into updatedComment struct
		var updatedComment models.Comment
		err = json.Unmarshal(content, &updatedComment)
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		// Set user ID from token data
		updatedComment.UserId, err = uuid.FromString(tokenData.UserId)
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		// Update comment in the database
		err = service.ComSrvice.EditComment(updatedComment)
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		} else {
			w.WriteHeader(http.StatusOK)
		}
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func DeleteCommentHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodDelete:
		commentID := r.URL.Query().Get("commentid")
		if commentID == "" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		err := service.ComSrvice.DeleteComment(commentID)
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func GetCommentsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		postId := r.URL.Query().Get("postid")
		tokenData, _ := authService.VerifyToken(r)
		comments, _ := service.ComSrvice.GetCommentsByPostId(postId, tokenData)
		data := Data{
			Username: tokenData.Username,
			Comments: comments,
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(data)
	}
}

func CommReactHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getCommReactHandler(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func getCommReactHandler(w http.ResponseWriter, r *http.Request) {
	tokenData, err := authService.VerifyToken(r)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	react := r.URL.Query().Get("react")
	commentId := r.URL.Query().Get("commentid")
	if react == "" || commentId == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	comment, err := service.ComSrvice.GetComment(commentId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var votes int
	if react == "LIKE" || react == "DISLIKE" {
		reaction, err := service.ComSrvice.GetUserCommReact(tokenData.UserId, commentId)
		if err != nil {
			if err == sql.ErrNoRows {
				service.ComSrvice.SaveCommReaction(comment, react, tokenData.UserId)
			}
		} else if reaction.Reactions == react {
			repo.ReactRepo.DeleteReaction(reaction.ReactId.String())
		} else {
			reaction.Reactions = react
			repo.ReactRepo.UpdateReaction(reaction)
		}

		votes, err = service.ComSrvice.GetCommentVotes(commentId)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(votes)
	} else {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}
