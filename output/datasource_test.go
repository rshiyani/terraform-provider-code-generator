package movies

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func TestAccMoviesMovieDataSource_Basic(t *testing.T) {
	resourceName := "movies_movie.test"
	dataSourceName := "data.movies_movie.test"
	randomParameter := acctest.RandStringFromCharSet(5, "abcdefghijklmnopqrstuvwxyz")
	randomValue := acctest.RandString(5)
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		CheckDestroy:      testAccCheckMoviesMovieDestroy,
		Steps: append([]resource.TestStep{

			{
				Config:      CreateAccMovieDataSourceWithoutId(),
				ExpectError: regexp.MustCompile(`Missing required argument`),
			},

			{
				Config: CreateAccMovieDataSourceConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair(dataSourceName, "title", resourceName, "title"),
					resource.TestCheckResourceAttrPair(dataSourceName, "isbn", resourceName, "isbn"),
					resource.TestCheckResourceAttrPair(dataSourceName, "genre", resourceName, "genre"),
					resource.TestCheckResourceAttrPair(dataSourceName, "director.#", resourceName, "director.#"),
					resource.TestCheckResourceAttrPair(dataSourceName, "rating.#", resourceName, "rating.#"),
					resource.TestCheckResourceAttrPair(dataSourceName, "casts.#", resourceName, "casts.#"),
				),
			},
			{
				Config:      CreateAccMovieUpdatedConfigDataSourceRandomAttr(randomParameter, randomValue),
				ExpectError: regexp.MustCompile(`An argument named (.)+ is not expected here.`),
			},
			{
				Config:      CreateAccMovieDataSourceWithInvalidId(),
				ExpectError: regexp.MustCompile(""), // `(.)+ Object may not exists`
			},
		}, generateStepForDataSourceUpdatedOptionalAttr(dataSourceName, resourceName)...),
	})
}
func CreateAccMovieDataSourceWithoutId() string {
	resource := CreateAccMovieConfigWithOptional()
	resource += fmt.Sprintf(`
			data "movies_movie" "test" {
			}
			`)
	return resource
}
func CreateAccMovieDataSourceConfig() string {
	resource := CreateAccMovieConfigWithOptional()
	resource += fmt.Sprintf(`
	data "movies_movie" "test" {

					id = movies_movie.test.id
	}
	`)
	return resource
}
func CreateAccMovieUpdatedConfigDataSourceRandomAttr(key, value string) string {
	resource := CreateAccMovieConfigWithOptional()
	resource += fmt.Sprintf(`
	data "movies_movie" "test" {

					id = movies_movie.test.id
		%s = "%s"
	}
	`, key, value)
	return resource
}

func CreateAccMovieDataSourceWithInvalidId() string {
	var resource string
	parentResources := getParentMovie()
	parentResources = parentResources[:len(parentResources)-1]
	resource += createMovieConfig(parentResources)
	resource += fmt.Sprintf(`
			data "movies_movie" "test" {
							
							id = "%v"


							title = "%v"

							isbn = "%v"

							genre = "%v"

						director {
    
						                        
                            firstname = "%v"
                        
                            lastname = "%v"

						}

						rating {
    
						                        
                            rater = "%v"
                        
                            rating = %v

						}

							casts = ["%v", "%v"]
		}
	`, "abcd",
		searchInObject(resourceMovieTest, "title.valid.0"),
		searchInObject(resourceMovieTest, "isbn.valid.0"),
		searchInObject(resourceMovieTest, "genre.valid.0"),
		searchInObject(resourceMovieTest, "director.firstname.valid.0"),
		searchInObject(resourceMovieTest, "director.lastname.valid.0"),
		searchInObject(resourceMovieTest, "rating.rater.valid.0"),
		searchInObject(resourceMovieTest, "rating.rating.valid.0"),
		searchInObject(resourceMovieTest, "casts.valid.0"),
		searchInObject(resourceMovieTest, "casts.valid.1"))
	return resource
}

func generateStepForDataSourceUpdatedOptionalAttr(dataSourceName, resourceName string) []resource.TestStep {
	testSteps := make([]resource.TestStep, 0, 1)
	var valid interface{}
	valid = searchInObject(resourceMovieTest, "title.valid.1")
	testSteps = append(testSteps, resource.TestStep{
		Config: CreateAccMovieDataSourceUpdatedOptionalAttrTitle(valid),
		Check: resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttrPair(dataSourceName, "title", resourceName, "title"),
		),
	})
	valid = searchInObject(resourceMovieTest, "isbn.valid.1")
	testSteps = append(testSteps, resource.TestStep{
		Config: CreateAccMovieDataSourceUpdatedOptionalAttrIsbn(valid),
		Check: resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttrPair(dataSourceName, "isbn", resourceName, "isbn"),
		),
	})
	valid = searchInObject(resourceMovieTest, "genre.valid.1")
	testSteps = append(testSteps, resource.TestStep{
		Config: CreateAccMovieDataSourceUpdatedOptionalAttrGenre(valid),
		Check: resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttrPair(dataSourceName, "genre", resourceName, "genre"),
		),
	})
	return testSteps
}

func CreateAccMovieDataSourceUpdatedOptionalAttrTitle(value interface{}) string {
	var resource string
	parentResources := getParentMovie()
	parentResources = parentResources[:len(parentResources)-1]
	resource += createMovieConfig(parentResources)
	resource += fmt.Sprintf(`
			resource "movies_movie" "test" {
							
							title = "%v"

							isbn = "%v"

							genre = "%v"

							director {
    
							                        
                            firstname = "%v"
                        
                            lastname = "%v"

							}

							rating {
    
							                        
                            rater = "%v"
                        
                            rating = %v

							}

							casts = ["%v", "%v"]
			}
			`, value,
		searchInObject(resourceMovieTest, "isbn.valid.0"),
		searchInObject(resourceMovieTest, "genre.valid.0"),
		searchInObject(resourceMovieTest, "director.firstname.valid.0"),
		searchInObject(resourceMovieTest, "director.lastname.valid.0"),
		searchInObject(resourceMovieTest, "rating.rater.valid.0"),
		searchInObject(resourceMovieTest, "rating.rating.valid.0"),
		searchInObject(resourceMovieTest, "casts.valid.0"),
		searchInObject(resourceMovieTest, "casts.valid.1"))
	resource += fmt.Sprintf(`
		data "movies_movie" "test" {

						id = movies_movie.test.id
		}
		`)
	return resource
}
func CreateAccMovieDataSourceUpdatedOptionalAttrIsbn(value interface{}) string {
	var resource string
	parentResources := getParentMovie()
	parentResources = parentResources[:len(parentResources)-1]
	resource += createMovieConfig(parentResources)
	resource += fmt.Sprintf(`
			resource "movies_movie" "test" {

							title = "%v"
							
							isbn = "%v"

							genre = "%v"

							director {
    
							                        
                            firstname = "%v"
                        
                            lastname = "%v"

							}

							rating {
    
							                        
                            rater = "%v"
                        
                            rating = %v

							}

							casts = ["%v", "%v"]
			}
			`, searchInObject(resourceMovieTest, "title.valid.0"),
		value,
		searchInObject(resourceMovieTest, "genre.valid.0"),
		searchInObject(resourceMovieTest, "director.firstname.valid.0"),
		searchInObject(resourceMovieTest, "director.lastname.valid.0"),
		searchInObject(resourceMovieTest, "rating.rater.valid.0"),
		searchInObject(resourceMovieTest, "rating.rating.valid.0"),
		searchInObject(resourceMovieTest, "casts.valid.0"),
		searchInObject(resourceMovieTest, "casts.valid.1"))
	resource += fmt.Sprintf(`
		data "movies_movie" "test" {

						id = movies_movie.test.id
		}
		`)
	return resource
}
func CreateAccMovieDataSourceUpdatedOptionalAttrGenre(value interface{}) string {
	var resource string
	parentResources := getParentMovie()
	parentResources = parentResources[:len(parentResources)-1]
	resource += createMovieConfig(parentResources)
	resource += fmt.Sprintf(`
			resource "movies_movie" "test" {

							title = "%v"

							isbn = "%v"
							
							genre = "%v"

							director {
    
							                        
                            firstname = "%v"
                        
                            lastname = "%v"

							}

							rating {
    
							                        
                            rater = "%v"
                        
                            rating = %v

							}

							casts = ["%v", "%v"]
			}
			`, searchInObject(resourceMovieTest, "title.valid.0"),
		searchInObject(resourceMovieTest, "isbn.valid.0"),
		value,
		searchInObject(resourceMovieTest, "director.firstname.valid.0"),
		searchInObject(resourceMovieTest, "director.lastname.valid.0"),
		searchInObject(resourceMovieTest, "rating.rater.valid.0"),
		searchInObject(resourceMovieTest, "rating.rating.valid.0"),
		searchInObject(resourceMovieTest, "casts.valid.0"),
		searchInObject(resourceMovieTest, "casts.valid.1"))
	resource += fmt.Sprintf(`
		data "movies_movie" "test" {

						id = movies_movie.test.id
		}
		`)
	return resource
}
