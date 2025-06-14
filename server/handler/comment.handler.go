package handler

import (
	"database/sql"
	"encoding/json"
	"errors"
	"forum/models"
	repo "forum/server/repositories"
	"forum/server/service"
	"io"
	"log"
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
			RespondWithError(w, http.StatusUnauthorized, "Unauthorized")
			return
		}
		tokenData, err := service.AuthSrvice.GetTokenData(cookie.Value)
		if err != nil {
			RespondWithError(w, http.StatusInternalServerError, "Internal Server Error")
			return
		}

		// Read request body
		body := r.Body
		content, err := io.ReadAll(body)
		if err != nil {
			if err == io.EOF {
				log.Printf("Empty body: %v\n", err)
				RespondWithError(w, http.StatusBadRequest, "Bad Request")
				return
			} else {
				log.Printf("Can't read body: %v\n", err) // Added formatting directive %v
				RespondWithError(w, http.StatusInternalServerError, "Internal Server Error")
				return
			}
		}

		// Unmarshal request body into comment struct
		var comment models.Comment
		err = json.Unmarshal(content, &comment)
		if err != nil {
			log.Printf("Can't Unmarshal comment: %v\n", err)
			RespondWithError(w, http.StatusInternalServerError, "Internal Server Error")
			return
		}

		// Set user ID and username from token data
		comment.UserId, err = uuid.FromString(tokenData.UserId)
		if err != nil {
			log.Printf("Can't recup Comment.UserId: %v\n", err)
			RespondWithError(w, http.StatusInternalServerError, "Internal Server Error")
			return
		}
		comment.Username = tokenData.Username

		// Validate comment body
		if strings.TrimSpace(comment.Body) == "" {
			RespondWithError(w, http.StatusBadRequest, "Bad Request")
			return
		}

		// Create comment in the database
		err = service.ComSrvice.NewComment(comment)
		if err != nil {
			log.Printf("Can't create comment: %v\n", err)
			RespondWithError(w, http.StatusInternalServerError, "Internal Server Error")
			return
		} else {
			RespondWithError(w, http.StatusOK, "OK")
		}
	default:
		RespondWithError(w, http.StatusMethodNotAllowed, "Method Not Allowed")
	}
}

func EditCommentHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPut:
		// Get comment ID from URL parameter
		commentID := r.URL.Query().Get("commentid")
		if commentID == "" {
			RespondWithError(w, http.StatusBadRequest, "Bad Request")
			return
		}

		// Get token data from authentication cookie
		cookie, err := r.Cookie("auth-cookie")
		if err != nil {
			RespondWithError(w, http.StatusUnauthorized, "Unauthorized")
			return
		}
		tokenData, err := service.AuthSrvice.GetTokenData(cookie.Value)
		if err != nil {
			RespondWithError(w, http.StatusInternalServerError, "Internal Server Error")
			return
		}

		// Read request body
		body := r.Body
		content, err := io.ReadAll(body)
		if err != nil {
			if err == io.EOF {
				RespondWithError(w, http.StatusBadRequest, "Bad Request")
				return
			} else {
				log.Printf("Can't read Body: %v\n", err)
				RespondWithError(w, http.StatusInternalServerError, "Internal Server Error")
				return
			}
		}

		// Unmarshal request body into updatedComment struct
		var updatedComment models.Comment
		err = json.Unmarshal(content, &updatedComment)
		if err != nil {
			log.Printf("%v\n", err)
			RespondWithError(w, http.StatusInternalServerError, "Internal Server Error")
			return
		}

		// Set user ID from token data
		updatedComment.UserId, err = uuid.FromString(tokenData.UserId)
		if err != nil {
			log.Printf("Error converting user ID to UUID: %v\n", err)
			RespondWithError(w, http.StatusInternalServerError, "Internal Server Error")
			return
		}

		// Update comment in the database
		if err = service.ComSrvice.EditComment(updatedComment); err != nil {
			log.Printf("Error updating comment: %v\n", err)
			RespondWithError(w, http.StatusInternalServerError, "Internal Server Error")
			return
		} else {
			RespondWithError(w, http.StatusOK, "OK")
		}
	default:
		RespondWithError(w, http.StatusMethodNotAllowed, "Method Not Allowed")
	}
}

func DeleteCommentHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodDelete:
		commentID := r.URL.Query().Get("commentid")
		if commentID == "" {
			RespondWithError(w, http.StatusBadRequest, "Bad Request")
			return
		}

		if err := service.ComSrvice.DeleteComment(commentID); err != nil {
			log.Printf("Error deleting comment: %v\n", err)
			RespondWithError(w, http.StatusInternalServerError, "Internal Server Error")
			return
		}

		RespondWithError(w, http.StatusOK, "OK")
	default:
		RespondWithError(w, http.StatusMethodNotAllowed, "Method Not Allowed")
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
		RespondWithJSON(w, http.StatusOK, data)
	}
}

func CommReactHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getCommReactHandler(w, r)
	default:
		RespondWithError(w, http.StatusMethodNotAllowed, "Method Not Allowed")
	}
}

func getCommReactHandler(w http.ResponseWriter, r *http.Request) {
	tokenData, err := authService.VerifyToken(r)
	if err != nil {
		RespondWithError(w, http.StatusUnauthorized, "Unauthorized")
		return
	}

	react := r.URL.Query().Get("react")
	commentId := r.URL.Query().Get("commentid")
	if react == "" || commentId == "" {
		RespondWithError(w, http.StatusBadRequest, "Bad Request")
		return
	}

	comment, err := service.ComSrvice.GetComment(commentId)
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}

	var votes int
	if react == "LIKE" || react == "DISLIKE" {
		reaction, err := service.ComSrvice.GetUserCommReact(tokenData.UserId, commentId)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				err := service.ComSrvice.SaveCommReaction(comment, react, tokenData.UserId)
				if err != nil {
					RespondWithError(w, http.StatusInternalServerError, "Internal Server Error")
					return
				}
			}
		} else if reaction.Reactions == react {
			err := repo.ReactRepo.DeleteReaction(reaction.ReactId.String())
			if err != nil {
				RespondWithError(w, http.StatusInternalServerError, "Internal Server Error")
				return
			}
		} else {
			reaction.Reactions = react
			err := repo.ReactRepo.UpdateReaction(reaction)
			if err != nil {
				RespondWithError(w, http.StatusInternalServerError, "Internal Server Error")
				return
			}
		}

		votes, err = service.ComSrvice.GetCommentVotes(commentId)
		if err != nil {
			RespondWithError(w, http.StatusInternalServerError, "Internal Server Error")
			return
		}

		RespondWithJSON(w, http.StatusOK, votes)
	} else {
		RespondWithError(w, http.StatusBadRequest, "Bad Request")
		return
	}
}
