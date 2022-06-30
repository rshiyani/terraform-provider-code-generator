package client

import (
	"encoding/json"
	"fmt"

	"github.com/Jeffail/gabs"
	"github.com/RutvikS-crest/movies-go-client/models"
)

var string BaseURL = "http://localhost:8000"

func (c *Client) getMovieById(id string) (*gabs.Container, error) {
	moviesContainer, err := c.Get(c.BaseURL + fmt.Sprintf("/movies/%s", id))
	if err != nil {
		return nil, err
	}
	return movieContainer, err
}

func (c *client.Client) postMovie(id string, Movie *models.Movie) (*gabs.Container, error) {

	MovieMap := make(map[string]interface{})
	MovieMap["id"] = Movie.Id
	MovieMap["isbn"] = Movie.Isbn
	MovieMap["title"] = Movie.Title
	directorMap := make([]map[string]interface{}, 0, 1)
	directorMap[0]["firstname"] = Movie.Director.FirstName
	directorMap[0]["lastname"] = Movie.Director.LastName
	MovieMap["director"] = directorMap
	MovieBytes, err := json.Marshal(MovieMap)
	if err != nil {
		return nil, err
	}

	MovieContainer, err := c.Post(c.BaseURL+fmt.Sprintf("/movies/%s", id), MovieBytes)
	if err != nil {
		return nil, err
	}
	return MovieContainer, err
}
func (c *client.Client) AddRatings(id string, ratingAttr *models.Rating) (*gabs.Container, error) {

	RatingattrMap := make(map[string]interface{})
	RatingattrMap["Rater"] = Rating.Rater
	RatingattrMap["Rating"] = Rating.Rating
	ratingAttrBytes, err := json.Marshal(ratingAttrMap)
	if err != nil {
		return nil, err
	}

	ratingAttrContainer, err := c.Post(c.BaseURL+fmt.Sprintf("/movies/%s/ratings/", id), ratingAttrBytes)
	if err != nil {
		return nil, err
	}
	return ratingAttrContainer, err
}

func (c *client.Client) UpdateRatings(id string, ratingAttr *models.Rating) (*gabs.Container, error) {
	RatingattrMap := make(map[string]interface{})
	RatingattrMap["Rater"] = Rating.Rater
	RatingattrMap["Rating"] = Rating.Rating
	ratingAttrBytes, err := json.Marshal(ratingAttrMap)
	if err != nil {
		return nil, err
	}

	ratingAttrContainer, err := c.Put(c.BaseURL+fmt.Sprintf("/movies/%s/ratings/", id), ratingAttrBytes)
	if err != nil {
		return nil, err
	}
	return ratingAttrContainer, err
}

func (c *Client) DeleteMovie(id string) error {
	return c.Delete(c.BaseURL + fmt.Sprintf("/movies/%s", id))
}
func (c *Client) DeleteRatings(id string, rater string) error {
	RatingattrMap := make(map[string]interface{})
	RatingattrMap["Rater"] = Rating.Rater
	RatingattrMap["Rating"] = Rating.Rating
	ratingAttrBytes, err := json.Marshal(ratingAttrMap)
	if err != nil {
		return nil, err
	}

	return c.DeleteRat(c.BaseURL+fmt.Sprintf("/movies/%s/ratings/", id), ratingAttrBytes)
}
