package studiogbl

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const baseURL string = "https://ghibliapi.herokuapp.com"

//Species data type for the species and their fims
type Species struct {
	ID    string   `json:"id"`
	Name  string   `json:"name"`
	Films []string `json:"films"`
}

//GetSpecies returns the species
func GetSpecies(w http.ResponseWriter, r *http.Request) {
	v := r.URL.Query()
	queryParam := v.Get("species")

	speciesRes, err := http.Get(baseURL + "/species/" + queryParam)
	if err != nil || speciesRes.StatusCode != 200 {
		handleError(w, speciesRes.StatusCode)
		return
	}

	defer speciesRes.Body.Close()

	speciesBody, byteErr := ioutil.ReadAll(speciesRes.Body)
	if byteErr != nil {
		handleError(w, 500)
		return
	}

	var species Species

	jsonError := json.Unmarshal(speciesBody, &species)
	if jsonError != nil {
		handleError(w, 500)
		return
	}

	var speciesFilms []Film
	var film Film

	for _, filmURL := range species.Films {
		film, err = GetFilms(filmURL)
		speciesFilms = append(speciesFilms, film)
	}

	prettyJSON, err := json.MarshalIndent(speciesFilms, "", "    ")
	if err != nil {
		handleError(w, 500)
		return
	}
	fmt.Printf("%s\n", prettyJSON)

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(speciesFilms)
}

//handleError handles the error cases
func handleError(w http.ResponseWriter, code int) {
	switch code {
	case 400:
		log.Printf("Bad Request")
		http.Error(w, "Bad Request", http.StatusBadRequest)
	case 404:
		log.Printf("Not Found")
		http.Error(w, "Not Found", http.StatusNotFound)
	case 500:
		log.Printf("Server error")
		http.Error(w, "Server error", http.StatusInternalServerError)
	}
}
