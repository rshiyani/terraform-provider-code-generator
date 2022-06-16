package movies

import (
	"fmt"
	"regexp"
	"testing"
)

var resourceMovieTest = map[string]interface{}{
	"title": map[string]interface{}{
		"valid":   []interface{}{"Hello", "World"},
		"invalid": []interface{}{234, 987},
	},
	"isbn": map[string]interface{}{
		"valid":   []interface{}{"Hello", "World"},
		"invalid": []interface{}{234, 987},
	},
	"genre": map[string]interface{}{
		"valid":   []interface{}{"thriller", "action", "horror", "fiction", "comedy"},
		"invalid": []interface{}{"6noqck58mm"},
	},
}

func TestAccMoviesMovie_Basic(t *testing.T) {
	var movie_default models.Movie
	var movie_updated models.Movie
	resourceName := "movies_movie.test"

	// [TODO]: Add makeTestVariable() to utils.go file
	// rName := makeTestVariable(acctest.RandString(5))
	// rOther := makeTestVariable(acctest.RandString(5))

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		CheckDestroy:      testAccCheckMoviesMovieDestroy,
		Steps: []resource.TestStep{
			{
				Config:      CreateAccMovieWithoutTitle(),
				ExpectError: regexp.MustCompile(`Missing required argument`),
			},
			{
				Config: CreateAccMovieConfig(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMoviesMovieExists(resourceName, &movie_default),
					resource.TestCheckResourceAttr(resourceName, "title", fmt.Sprintf("%v", resourceMovieTest["title"].(map[string]interface{})["valid"].([]interface{})[0])),

					resource.TestCheckResourceAttr(resourceName, "isbn", ""),

					resource.TestCheckResourceAttr(resourceName, "genre", "thriller"),
				),
			},
			{
				Config: CreateAccMovieConfigWithOptional(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMoviesMovieExists(resourceName, &movie_updated),
					resource.TestCheckResourceAttr(resourceName, "title", fmt.Sprintf("%v", resourceMovieTest["title"].(map[string]interface{})["valid"].([]interface{})[0])),
					resource.TestCheckResourceAttr(resourceName, "isbn", fmt.Sprintf("%v", resourceMovieTest["isbn"].(map[string]interface{})["valid"].([]interface{})[0])),
					resource.TestCheckResourceAttr(resourceName, "genre", fmt.Sprintf("%v", resourceMovieTest["genre"].(map[string]interface{})["valid"].([]interface{})[0])),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: CreateAccMovieConfig(),
			},
		},
	})
}

func generateStepForUpdatedAttr(resourceName string, movie_default, movie_updated *models.Movie) []resource.TestStep {
	testSteps := make([]resource.TestStep, 0, 1)
	var valid []interface{}
	valid = resourceMovieTest["isbn"].(map[string]interface{})["valid"].([]interface{})
	for _, value := range valid {
		v := fmt.Sprintf("%v", value)
		testSteps = append(testSteps, resource.TestStep{
			Config: CreateAccMovieUpdatedAttr("isbn", value),
			Check: resource.ComposeTestCheckFunc(
				testAccCheckMoviesMovieExists(resourceName, movie_updated),
				resource.TestCheckResourceAttr(resourceName, "isbn", v),
				testAccCheckMoviesMovieIdEqual(movie_default, movie_updated),
			),
		})
	}
	valid = resourceMovieTest["genre"].(map[string]interface{})["valid"].([]interface{})
	for _, value := range valid {
		v := fmt.Sprintf("%v", value)
		testSteps = append(testSteps, resource.TestStep{
			Config: CreateAccMovieUpdatedAttr("genre", value),
			Check: resource.ComposeTestCheckFunc(
				testAccCheckMoviesMovieExists(resourceName, movie_updated),
				resource.TestCheckResourceAttr(resourceName, "genre", v),
				testAccCheckMoviesMovieIdEqual(movie_default, movie_updated),
			),
		})
	}
	return testSteps
}

func TestAccMoviesMovie_Update(t *testing.T) {
	var movie_default models.Movie
	var movie_updated models.Movie
	resourceName := "movies_movie.test"

	// [TODO]: Add makeTestVariable() to utils.go file
	// rName := makeTestVariable(acctest.RandString(5))
	// rOther := makeTestVariable(acctest.RandString(5))

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		CheckDestroy:      testAccCheckMoviesMovieDestroy,
		Steps: append([]resource.TestStep{
			{
				Config: CreateAccMovieConfig(),
				Check:  testAccCheckMoviesMovieExists(resourceName, &movie_default),
			},
		}, generateStepForUpdatedAttr(resourceName, &movie_default, &movie_updated)...),
	})
}

func CreateAccMovieWithoutTitle() string {
	return fmt.Sprintf(`
		resource "movies_movie" "test" {
		}
	`)
}

func CreateAccMovieConfig() string {
	var resource string
	resource += fmt.Sprintf(`
		resource  "movies_movie" "test" {
						title = "%v"
						
		}
	`, resourceMovieTest["title"].(map[string]interface{})["valid"].([]interface{})[0])
	return resource
}

func CreateAccMovieConfigWithOptional() string {
	var resource string
	resource += fmt.Sprintf(`
		resource  "movies_movie" "test" {
						title = "%v"
						
						isbn = "%v"
						
						genre = "%v"
						
		}
	`, resourceMovieTest["title"].(map[string]interface{})["valid"].([]interface{})[0],
		resourceMovieTest["isbn"].(map[string]interface{})["valid"].([]interface{})[0],
		resourceMovieTest["genre"].(map[string]interface{})["valid"].([]interface{})[0])
	return resource
}

func CreateAccMovieUpdatedAttr(attr string, value interface{}) string {
	var resource string
	resource += fmt.Sprintf(`
		resource  "movies_movie" "test" {
						title = "%v"
						
						%v = "%v"
		}
	`, resourceMovieTest["title"].(map[string]interface{})["valid"].([]interface{})[0], attr, value)
	return resource
}

func testAccCheckMoviesMovieExists(name string, movie *models.Movie) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		// [TODO]: Write your code here
	}
}

func testAccCheckMoviesMovieDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*client.Client)

	for _, rs := range s.RootModule().Resources {

		if rs.Type == "movies_movie" {
			// [TODO]: Write your code here
		}
	}
	return nil
}

func testAccCheckMoviesMovieIdEqual(movie1, movie2 *models.Movie) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		if movie1.ID != movie2.ID {
			return fmt.Errorf("Movie IDs are not equal")
		}
		return nil
	}
}
