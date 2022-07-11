package movies

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/RutvikS-crest/movies-go-client/client"
)

const movieSelfRequiredCount = 1

var resourceMovieTest = map[string]interface{}{
	"title": map[string]interface{}{
		"valid":           []interface{}{"reb9mfgujf", "rotf377xld", "pigjrkl2i7", "qhlphte62c"},
		"invalid":         []interface{}{10, 12.43},
		"multiple_valids": []interface{}{"kg10fbwcz7", "k1kz1au15a", "2ki0e88t4w", "jd9nakr11f", "bewjgqjo3q", "kpofubatve", "3fo17tajnv", "rsuk6jckg4", "rzl0x2uf6r", "o4qy52ueri", "iydt4lult7", "i8o95dz89b", "m5efm3d33g", "gqr7zlix8h", "ea05g1oocv"},
	},

	"isbn": map[string]interface{}{
		"valid":           []interface{}{"vu6t5rb3qx", "1n3yo1gvcr", "6sbganzn2p", "wblhz74mcp"},
		"invalid":         []interface{}{10, 12.43},
		"multiple_valids": []interface{}{"9d6is9p00i", "dn8ybvbs0b", "5a83o5swuw", "f2mwe83rw3", "0ytgykwnmh", "aegb5ka3o4", "mgjhlvqe2r", "9qh6k1zoss", "4uvy25sqrj", "hqszr06926", "q4p2r3ozcd", "rvhsezo8im", "yc008r2gxg", "758z3xzdc3", "kwhu5f9z9b"},
	},

	"genre": map[string]interface{}{
		"valid":           []interface{}{"thriller", "action", "horror", "fiction", "comedy"},
		"invalid":         []interface{}{"u10gyzoswm"},
		"multiple_valids": []interface{}{"thriller", "action", "horror", "fiction", "comedy"},
	},

	"director": map[string]interface{}{
		"firstname": map[string]interface{}{
			"valid":           []interface{}{"y8fbb197bv", "s5cawisxjs", "e21is5fy10", "2tkj8myojf"},
			"invalid":         []interface{}{10, 12.43},
			"multiple_valids": []interface{}{"qdaiwihjj8", "zcs4ccpv1c", "q6vcj9xzkg", "irddgz2emm", "gbqthyjni6", "dt7lz1w6wr", "fvsylqhxgc", "mqd4ni2f8b", "8wd0oo9ske", "eai9aceltj", "l6iufp0fg8", "m1xlst9r6m", "dnxwmzc1l3", "avwjtdkl0l", "cptpne6zgk"},
		},

		"lastname": map[string]interface{}{
			"valid":           []interface{}{"9cfkeac00o", "sm9jlwkl0o", "l0q337pqsq", "i0aql3afvq"},
			"invalid":         []interface{}{10, 12.43},
			"multiple_valids": []interface{}{"mlex79t71s", "fi7geoq37m", "2qqy0e8wo0", "3wwgaz0pnh", "syh8xmyje8", "qh51boe7gk", "7lk5nqro44", "dziei9tme3", "807yse8myd", "6fm5u3akp0", "1loclap0wh", "d32s6a19ip", "dinxpsuhg5", "o9uumlhb2b", "7v23dgzv53"},
		},
	},

	"rating": map[string]interface{}{
		"rater": map[string]interface{}{
			"valid":           []interface{}{"j941xenc9g", "0clp8bilj4", "k8nc5mdrm9", "0zjd4itmq2"},
			"invalid":         []interface{}{10, 12.43},
			"multiple_valids": []interface{}{"qu03wzno10", "afxdgb0b34", "8dxy78k9wp", "iua9krjv57", "9gyqp3lx58", "jdh0qfd4yx", "z37cp0jihw", "ljplqz1s77", "owqmlpqlvb", "9wmh36ucg6", "zjibh46nw1", "byhhe6vmx9", "m4dzpv11tn", "g2un4gqn1q", "tp4t74yyu1"},
		},

		"rating": map[string]interface{}{
			"valid":           []interface{}{1, 10, 5.5, 5.327493256150269},
			"invalid":         []interface{}{0, 11},
			"multiple_valids": []interface{}{1, 10, 5.5, 6.3077942796079425, 3.8319217250341877, 7.284504431375421, 3.938417818790897, 4.803870162269427, 4.125968612341213, 5.012689864396442, 6.666219282869619, 9.076894442102645, 1.0244780871351846, 6.451452026875719, 2.2539842447992426},
		},
	},

	"casts": map[string]interface{}{
		"valid":           []interface{}{"ao08fjphu0", "tj7oh8qu2p", "fs2hzniqj6", "ijcsi2kocb"},
		"invalid":         []interface{}{10, 12.43},
		"multiple_valids": []interface{}{"og35f0mipd", "vm5qe7ezh3", "0a107dr14y", "lofl9pcn3p", "x21nq0p2dg", "m7wyp6z58z", "8n2xfwcrod", "z39ir85570", "j6r31a7svi", "2es32rzq3g", "aptehn6o3k", "kc6x7umvk7", "ijjqk3dvbs", "kppufs3o2j", "rsu1fnbdh9"},
	},
}

func TestAccMoviesMovie_Basic(t *testing.T) {
	var movie_default models.Movie
	var movie_updated models.Movie
	resourceName := "movies_movie.test"

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		CheckDestroy:      testAccCheckMoviesMovieDestroy,
		Steps: append([]resource.TestStep{
			{
				Config:      CreateAccMovieWithoutTitle(),
				ExpectError: regexp.MustCompile(`Missing required argument`),
			},
			{
				Config: CreateAccMovieConfig(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMoviesMovieExists(resourceName, &movie_default),
					resource.TestCheckResourceAttr(resourceName, "title", fmt.Sprintf("%v", searchInObject(resourceMovieTest, "title.valid.0"))),

					resource.TestCheckResourceAttr(resourceName, "isbn", ""),

					resource.TestCheckResourceAttr(resourceName, "genre", "thriller"),

					resource.TestCheckResourceAttr(resourceName, "director.#", "0"),

					resource.TestCheckResourceAttr(resourceName, "rating.#", "0"),

					resource.TestCheckResourceAttr(resourceName, "casts.#", "0"),
				),
			},
			{
				Config: CreateAccMovieConfigWithOptional(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMoviesMovieExists(resourceName, &movie_updated),
					resource.TestCheckResourceAttr(resourceName, "title", fmt.Sprintf("%v", searchInObject(resourceMovieTest, "title.valid.0"))),
					resource.TestCheckResourceAttr(resourceName, "isbn", fmt.Sprintf("%v", searchInObject(resourceMovieTest, "isbn.valid.0"))),
					resource.TestCheckResourceAttr(resourceName, "genre", fmt.Sprintf("%v", searchInObject(resourceMovieTest, "genre.valid.0"))),
					resource.TestCheckResourceAttr(resourceName, "director.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "director.0.firstname", fmt.Sprintf("%v", searchInObject(resourceMovieTest, "director.firstname.valid.0"))),
					resource.TestCheckResourceAttr(resourceName, "director.0.lastname", fmt.Sprintf("%v", searchInObject(resourceMovieTest, "director.lastname.valid.0"))),

					resource.TestCheckResourceAttr(resourceName, "rating.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "rating.0.rater", fmt.Sprintf("%v", searchInObject(resourceMovieTest, "rating.rater.valid.0"))),
					resource.TestCheckResourceAttr(resourceName, "rating.0.rating", fmt.Sprintf("%v", searchInObject(resourceMovieTest, "rating.rating.valid.0"))),

					resource.TestCheckResourceAttr(resourceName, "casts.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "casts.0", fmt.Sprintf("%v", searchInObject(resourceMovieTest, "casts.valid.0"))),
					resource.TestCheckResourceAttr(resourceName, "casts.1", fmt.Sprintf("%v", searchInObject(resourceMovieTest, "casts.valid.1"))),
					testAccCheckMoviesMovieIdEqual(&movie_default, &movie_updated),
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
		}, generateStepForUpdatedRequiredAttr(resourceName, &movie_default, &movie_updated)...),
	})
}

func TestAccMoviesMovie_Update(t *testing.T) {
	var movie_default models.Movie
	var movie_updated models.Movie
	resourceName := "movies_movie.test"

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

func TestAccMoviesMovie_NegativeCases(t *testing.T) {
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
			},
		}, generateNegativeSteps(resourceName)...),
	})
}

func TestAccMoviesMovie_MultipleCreateDelete(t *testing.T) {
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
				Config: CreateAccMovieMultipleConfig(),
			},
		},
	})
}

func CreateAccMovieWithoutTitle() string {
	var resource string
	parentResources := getParentMovie()
	parentResources = parentResources[:len(parentResources)-1]
	resource += createMovieConfig(parentResources)
	resource += fmt.Sprintf(`
				resource  "movies_movie" "test" {

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

									casts = ["%v","%v"]
				}
			`, searchInObject(resourceMovieTest, "isbn.valid.0"),
		searchInObject(resourceMovieTest, "genre.valid.0"),
		searchInObject(resourceMovieTest, "director.firstname.valid.0"),
		searchInObject(resourceMovieTest, "director.lastname.valid.0"),
		searchInObject(resourceMovieTest, "rating.rater.valid.0"),
		searchInObject(resourceMovieTest, "rating.rating.valid.0"),
		searchInObject(resourceMovieTest, "casts.valid.0"),
		searchInObject(resourceMovieTest, "casts.valid.1"))
	return resource
}

func CreateAccMovieConfig() string {
	var resource string
	parentResources := getParentMovie()
	parentResources = parentResources[:len(parentResources)-1]
	resource += createMovieConfig(parentResources)
	resource += fmt.Sprintf(`
		resource  "movies_movie" "test" {

							title = "%v"
		}
	`, searchInObject(resourceMovieTest, "title.valid.0"))
	return resource
}

func CreateAccMovieConfigWithOptional() string {
	resource := createMovieConfig(getParentMovie())
	return resource
}

func generateStepForUpdatedRequiredAttr(resourceName string, movie_default, movie_updated *models.Movie) []resource.TestStep {
	testSteps := make([]resource.TestStep, 0, 1)
	var value interface{}
	value = searchInObject(resourceMovieTest, "title.valid.1")
	testSteps = append(testSteps, resource.TestStep{
		Config: CreateAccMovieUpdateRequiredTitle(),
		Check: resource.ComposeTestCheckFunc(
			testAccCheckMoviesMovieExists(resourceName, movie_updated),
			resource.TestCheckResourceAttr(resourceName, "title", fmt.Sprintf("%v", value)),
			testAccCheckMoviesMovieIdNotEqual(movie_default, movie_updated),
		),
	})
	return testSteps
}
func CreateAccMovieUpdateRequiredTitle() string {
	var resource string
	parentResources := getParentMovie()
	parentResources = parentResources[:len(parentResources)-1]
	resource += createMovieConfig(parentResources)
	value := searchInObject(resourceMovieTest, "title.valid.1")
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
	return resource
}

func CreateAccMovieUpdatedAttrIsbn(value interface{}) string {
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
	return resource
}
func CreateAccMovieUpdatedAttrGenre(value interface{}) string {
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
	return resource
}
func CreateAccMovieUpdatedAttrDirectorLastname(value interface{}) string {
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
		searchInObject(resourceMovieTest, "genre.valid.0"),
		searchInObject(resourceMovieTest, "director.firstname.valid.0"),
		value,
		searchInObject(resourceMovieTest, "rating.rater.valid.0"),
		searchInObject(resourceMovieTest, "rating.rating.valid.0"),
		searchInObject(resourceMovieTest, "casts.valid.0"),
		searchInObject(resourceMovieTest, "casts.valid.1"))
	return resource
}
func CreateAccMovieUpdatedAttrRatingRating(value interface{}) string {
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
		searchInObject(resourceMovieTest, "genre.valid.0"),
		searchInObject(resourceMovieTest, "director.firstname.valid.0"),
		searchInObject(resourceMovieTest, "director.lastname.valid.0"),
		searchInObject(resourceMovieTest, "rating.rater.valid.0"),
		value,
		searchInObject(resourceMovieTest, "casts.valid.0"),
		searchInObject(resourceMovieTest, "casts.valid.1"))
	return resource
}
func CreateAccMovieUpdatedAttrCasts(value interface{}) string {
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

							casts = ["%v"]
		}
	`, searchInObject(resourceMovieTest, "title.valid.0"),
		searchInObject(resourceMovieTest, "isbn.valid.0"),
		searchInObject(resourceMovieTest, "genre.valid.0"),
		searchInObject(resourceMovieTest, "director.firstname.valid.0"),
		searchInObject(resourceMovieTest, "director.lastname.valid.0"),
		searchInObject(resourceMovieTest, "rating.rater.valid.0"),
		searchInObject(resourceMovieTest, "rating.rating.valid.0"),
		value)
	return resource
}

func generateStepForUpdatedAttr(resourceName string, movie_default, movie_updated *models.Movie) []resource.TestStep {
	testSteps := make([]resource.TestStep, 0, 1)
	var valid []interface{}
	valid = searchInObject(resourceMovieTest, "isbn.valid").([]interface{})
	for _, value := range valid {
		v := fmt.Sprintf("%v", value)
		testSteps = append(testSteps, resource.TestStep{
			Config: CreateAccMovieUpdatedAttrIsbn(value),
			Check: resource.ComposeTestCheckFunc(
				testAccCheckMoviesMovieExists(resourceName, movie_updated),
				resource.TestCheckResourceAttr(resourceName, "isbn", v),
				testAccCheckMoviesMovieIdEqual(movie_default, movie_updated),
			),
		})
	}
	valid = searchInObject(resourceMovieTest, "genre.valid").([]interface{})
	for _, value := range valid {
		v := fmt.Sprintf("%v", value)
		testSteps = append(testSteps, resource.TestStep{
			Config: CreateAccMovieUpdatedAttrGenre(value),
			Check: resource.ComposeTestCheckFunc(
				testAccCheckMoviesMovieExists(resourceName, movie_updated),
				resource.TestCheckResourceAttr(resourceName, "genre", v),
				testAccCheckMoviesMovieIdEqual(movie_default, movie_updated),
			),
		})
	}
	valid = searchInObject(resourceMovieTest, "director.lastname.valid").([]interface{})
	for _, value := range valid {
		v := fmt.Sprintf("%v", value)
		testSteps = append(testSteps, resource.TestStep{
			Config: CreateAccMovieUpdatedAttrDirectorLastname(value),
			Check: resource.ComposeTestCheckFunc(
				testAccCheckMoviesMovieExists(resourceName, movie_updated),
				resource.TestCheckResourceAttr(resourceName, "director.0.lastname", v),
				testAccCheckMoviesMovieIdEqual(movie_default, movie_updated),
			),
		})
	}

	valid = searchInObject(resourceMovieTest, "rating.rating.valid").([]interface{})
	for _, value := range valid {
		v := fmt.Sprintf("%v", value)
		testSteps = append(testSteps, resource.TestStep{
			Config: CreateAccMovieUpdatedAttrRatingRating(value),
			Check: resource.ComposeTestCheckFunc(
				testAccCheckMoviesMovieExists(resourceName, movie_updated),
				resource.TestCheckResourceAttr(resourceName, "rating.0.rating", v),
				testAccCheckMoviesMovieIdEqual(movie_default, movie_updated),
			),
		})
	}

	valid = searchInObject(resourceMovieTest, "casts.valid").([]interface{})
	for _, value := range valid {
		v := fmt.Sprintf("%v", value)
		testSteps = append(testSteps, resource.TestStep{
			Config: CreateAccMovieUpdatedAttrCasts(value),
			Check: resource.ComposeTestCheckFunc(
				testAccCheckMoviesMovieExists(resourceName, movie_updated),
				resource.TestCheckResourceAttr(resourceName, "casts.0", v),
				testAccCheckMoviesMovieIdEqual(movie_default, movie_updated),
			),
		})
	}
	return testSteps
}

func generateNegativeSteps(resourceName string) []resource.TestStep {
	//Use Update Config Function with false value
	testSteps := make([]resource.TestStep, 0, 1)
	var invalid []interface{}
	invalid = searchInObject(resourceMovieTest, "genre.invalid").([]interface{})
	for _, value := range invalid {
		testSteps = append(testSteps, resource.TestStep{
			Config:      CreateAccMovieUpdatedAttrGenre(value),
			ExpectError: regexp.MustCompile(expectErrorMap["StringInSlice"]),
		})
	}

	invalid = searchInObject(resourceMovieTest, "rating.rating.invalid").([]interface{})
	for _, value := range invalid {
		testSteps = append(testSteps, resource.TestStep{
			Config:      CreateAccMovieUpdatedAttrRatingRating(value),
			ExpectError: regexp.MustCompile(expectErrorMap["FloatBetween"]),
		})
	}

	testSteps = append(testSteps, resource.TestStep{
		Config: CreateAccMovieConfig(),
	})
	return testSteps
}

func CreateAccMovieMultipleConfig() string {
	var resource string
	parentResources := getParentMovie()
	parentResources = parentResources[:len(parentResources)-1]
	resource += createMovieConfig(parentResources)
	multipleValues := searchInObject(resourceMovieTest, "title.multiple_valids").([]interface{})
	for i, val := range multipleValues {
		resource += fmt.Sprintf(`
			resource "movies_movie" "test%d" {
							
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
		`, i, val,
			searchInObject(resourceMovieTest, "isbn.valid.0"),
			searchInObject(resourceMovieTest, "genre.valid.0"),
			searchInObject(resourceMovieTest, "director.firstname.valid.0"),
			searchInObject(resourceMovieTest, "director.lastname.valid.0"),
			searchInObject(resourceMovieTest, "rating.rater.valid.0"),
			searchInObject(resourceMovieTest, "rating.rating.valid.0"),
			searchInObject(resourceMovieTest, "casts.valid.0"),
			searchInObject(resourceMovieTest, "casts.valid.1"))
	}
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
		Id1, err := getIdFromMovieModel(movie1)
		if err != nil {
			return err
		}
		Id2, err := getIdFromMovieModel(movie2)
		if err != nil {
			return err
		}
		if Id1 != Id2 {
			return fmt.Errorf("Movie IDs are not equal")
		}
		return nil
	}
}

func testAccCheckMoviesMovieIdNotEqual(movie1, movie2 *models.Movie) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		Id1, err := getIdFromMovieModel(movie1)
		if err != nil {
			return err
		}
		Id2, err := getIdFromMovieModel(movie2)
		if err != nil {
			return err
		}
		if Id1 == Id2 {
			return fmt.Errorf("Movie IDs are equal")
		}
		return nil
	}
}

func getParentMovie() []string {
	t := []string{}
	t = append(t, movieBlock())
	return t
}

func movieBlock() string {
	return fmt.Sprintf(`
		resource  "movies_movie" "test" {

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

				        casts = ["%v","%v"]
		}
	`, searchInObject(resourceMovieTest, "title.valid.0"),
		searchInObject(resourceMovieTest, "isbn.valid.0"),
		searchInObject(resourceMovieTest, "genre.valid.0"),
		searchInObject(resourceMovieTest, "director.firstname.valid.0"),
		searchInObject(resourceMovieTest, "director.lastname.valid.0"),
		searchInObject(resourceMovieTest, "rating.rater.valid.0"),
		searchInObject(resourceMovieTest, "rating.rating.valid.0"),
		searchInObject(resourceMovieTest, "casts.valid.0"),
		searchInObject(resourceMovieTest, "casts.valid.1"))
}

// To eliminate duplicate resource block from slice of resource blocks
func createMovieConfig(configSlice []string) string {
	keys := make(map[string]bool)
	str := ""

	for _, entry := range configSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			str += entry
		}
	}

	return str
}
