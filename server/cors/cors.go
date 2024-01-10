package cors

import "net/http"

func SetCors(w *http.ResponseWriter) {
   (*w).Header().Set("Access-Control-Allow-Origin", "http://localhost:8000")
   (*w).Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS, PUT, DELETE")
   (*w).Header().Set("Access-Control-Allow-Headers", "*, auth-token")
   (*w).Header().Set("Access-Control-Allow-Credentials", "true")
}