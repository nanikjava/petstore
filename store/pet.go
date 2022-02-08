//Handle all Pet related operation
package store

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

type Pet struct{}

type category struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type tags struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type petResponse struct {
	ID       int      `json:"id"`
	Category category `json:"category"`
	Name     string   `json:"name"`
	Tags     []tags   `json:"tags"`
	Status   string   `json:"status"`
}

//ConfigureRoute configure route for Pet
func (p Pet) ConfigureRoute(r *mux.Router) {
	r.HandleFunc("/products", p.FindByStatus).Queries("status", "{status}").Methods("GET")

}

//FindByStatus find pet based on the 'status' parameter
func (p Pet) FindByStatus(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	s, ok := vars["status"]
	if !ok {
		respondWithJSON(w, http.StatusBadRequest, "Invalid status")
		return
	}

	resp := preparePetResponse(s)

	if resp != nil {
		respondWithJSON(w, http.StatusOK, preparePetResponse(s))
		return
	}
	respondWithJSON(w, http.StatusBadRequest, "invalid parameter")

}

//preparePetResponse prepare response for Pet
func preparePetResponse(s string) []petResponse {
	if s == "pending" {
		return preparePetResponsePending(s)
	}
	if s == "sold" {
		return preparePetResponseSold(s)
	}
	if s == "available" {
		return preparePetResponseAvailable(s)
	}

	return nil
}

//preparePetResponseAvailable prepare response for available pet
func preparePetResponseAvailable(s string) []petResponse {
	var pets []petResponse
	var pet petResponse

	pet.ID = 10001
	pet.Category = category{
		ID:   1,
		Name: "c",
	}
	pet.Name = "doggy"
	pet.Tags =
		[]tags{
			{
				ID:   1,
				Name: "dt",
			},
		}
	pet.Status = "available"

	pets = append(pets, pet)
	return pets
}

//preparePetResponseSold prepare response for sold pet
func preparePetResponseSold(s string) []petResponse {
	var pets []petResponse
	var pet petResponse

	pet.ID = 10001
	pet.Category = category{
		ID:   1,
		Name: "c",
	}
	pet.Name = "doggy"
	pet.Tags =
		[]tags{
			{
				ID:   1,
				Name: "dt",
			},
		}
	pet.Status = "sold"

	pets = append(pets, pet)
	return pets
}

//preparePetResponseSold prepare response for pending pet
func preparePetResponsePending(s string) []petResponse {
	var pets []petResponse
	var pet petResponse

	pet.ID = 10001
	pet.Category = category{
		ID:   1,
		Name: "c",
	}
	pet.Name = "doggy"
	pet.Tags =
		[]tags{
			{
				ID:   1,
				Name: "dt",
			},
		}
	pet.Status = "pending"

	pets = append(pets, pet)
	return pets
}

//respondWithJSON to convert to json string
func respondWithJSON(w http.ResponseWriter, c int, p interface{}) {
	response, _ := json.Marshal(p)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(c)
	w.Write(response)
}
