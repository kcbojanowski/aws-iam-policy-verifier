package api

import (
	"encoding/json"
	"github.com/kcbojanowski/aws-iam-policy-verifier/pkg/validator"
	"log"
	"net/http"
)

type PolicyResponse struct {
	IsValid bool   `json:"is_valid"`
	Error   string `json:"error,omitempty"`
}

func main() {
	http.HandleFunc("/validate", ValidateIAMPolicyHandler)
	log.Println("Server starting on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func ValidateIAMPolicyHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	defer r.Body.Close()

	var policy validator.IAMPolicy
	if err := json.NewDecoder(r.Body).Decode(&policy); err != nil {
		respondWithError(w, http.StatusBadRequest, "Bad Request: Error decoding JSON")
		return
	}

	if err := validator.ValidateIAMPolicy(policy); err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, PolicyResponse{IsValid: true})
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, PolicyResponse{IsValid: false, Error: message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload PolicyResponse) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	if err := json.NewEncoder(w).Encode(payload); err != nil {
		log.Printf("Error sending response: %v", err)
	}
}
