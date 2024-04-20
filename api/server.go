package main

import (
	"encoding/json"
	"github.com/kcbojanowski/aws-iam-policy-verifier/pkg/validator"
	"log"
	"net/http"
)

type PolicyRequest struct {
	Policy validator.IAMPolicy `json:"policy"`
}

func main() {
	http.HandleFunc("/validate", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
			return
		}

		var req PolicyRequest
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			http.Error(w, "Error reading rq body (can be missing data or incorrect types)", http.StatusBadRequest)
			return
		}

		err = validator.ValidateIAMPolicy(req.Policy)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusOK)
		write, err := w.Write([]byte("Policy is valid"))
		if err != nil {
			return
		}

		log.Printf("Wrote %d bytes to the response\n", write)
	})

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		return
	}
}
