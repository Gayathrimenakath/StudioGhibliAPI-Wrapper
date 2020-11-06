package studiogbl

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

//Film  data type for the response
type Film struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Director    string `json:"director"`
	Producer    string `json:"producer"`
	ReleaseDate string `json:"release_date"`
	RtScore     string `json:"rt_score"`
}

//GetFilms returns the films
func GetFilms(id string) (Film, error) {
	filmRes, err := http.Get(id)
	if err != nil {
		return Film{}, err
	}
	defer filmRes.Body.Close()

	filmBody, _ := ioutil.ReadAll(filmRes.Body)

	var film Film

	jsonErr := json.Unmarshal(filmBody, &film)
	if jsonErr != nil {
		return Film{}, err
	}

	return film, nil
}
