package handler

import "net/http"

func CreateChatHandler(w http.ResponseWriter, r *http.Request) {
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
		

}