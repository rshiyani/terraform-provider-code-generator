package movies


var resourceMovieTest = map[string]interface{}{
			"title" : map[string]interface{}{
	"valid": []interface{}{ "fgxpqzrfed", "v004mg74n4", "afecl88eca", "6g2mv0fqz8" },
	"invalid": []interface{}{ 10, 12.43 },
	"multiple_valids": []interface{}{ "vriqwy63tr", "2wbyd53l2m", "0qafzrvcwu", "iq5ryi7jx4", "bp6zozq4f5", "lyku3sqim7", "o258l1nj0m", "gxb3av4m2p", "208f3vbrob", "nl5d4iw6ga", "sso7mfgx2r", "l7udbyrjnb", "348la6pic8", "1kyqe6a1x5", "7cis1d7ywc" },
},

			"isbn" : map[string]interface{}{
	"valid": []interface{}{ "0txp0u7lvc", "romj7abhkq", "hhlo68c83j", "b0hy5b5svb" },
	"invalid": []interface{}{ 10, 12.43 },
	"multiple_valids": []interface{}{ "xpk4vhwf5r", "9l0q7b91dw", "s8x2q50oit", "1qmo18o37l", "m2gc5fk9t1", "22q74jqy3a", "4zvn5a0t2l", "7j7w4ezr4l", "9of5amrv2a", "cfy3zjen74", "a73q2yf8rs", "px9ewygk60", "s2m9prtr8c", "ikjciim2pl", "ey70vp3zgn" },
},

			"genre" : map[string]interface{}{
	"valid": []interface{}{ "thriller", "action", "horror", "fiction", "comedy" },
	"invalid": []interface{}{ "37xaa7khob" },
	"multiple_valids": []interface{}{ "thriller", "action", "horror", "fiction", "comedy" },
},

				"director" : map[string]interface{}{
			"firstname" : map[string]interface{}{
	"valid": []interface{}{ "i9q9x0474n", "q6bym76vrn", "3aykyi1f61", "trhugpaxb1" },
	"invalid": []interface{}{ 10, 12.43 },
	"multiple_valids": []interface{}{ "mosknpuezg", "1au17yzeo5", "wna9yiwv7m", "du8hdjwi2h", "3gswxsyybx", "k1eeqanb0x", "pormxun4y4", "hpgdtdx69d", "ama87qj203", "7rdgqn76mo", "cy3doe0v7z", "kcixk9z5o9", "966r6oto5k", "z2p9pj7o4t", "6t1dxrg1t5" },
},

			"lastname" : map[string]interface{}{
	"valid": []interface{}{ "h2hk7slyam", "1vhc3jk7oq", "91s5gm47px", "orpspigkso" },
	"invalid": []interface{}{ 10, 12.43 },
	"multiple_valids": []interface{}{ "16m3kzvwjg", "6lu1e2dllr", "2y46cqn5i4", "83dp9jnzgz", "6zdq9jxuho", "mtw3o65ojc", "ld7xgx1ofk", "4thclb02n0", "pbzizzjcnv", "dks277hddc", "1u0ums271d", "6uilv2fadj", "67ng6g2so8", "kzq8c7q5l7", "zqr790m5pr" },
},

},

				"rating" : map[string]interface{}{
			"rater" : map[string]interface{}{
	"valid": []interface{}{ "91bs56qalz", "objx6x9wfp", "y6nx7ahor4", "e9wvsixfxr" },
	"invalid": []interface{}{ 10, 12.43 },
	"multiple_valids": []interface{}{ "v9b385o6qr", "wi1371lix8", "rom4sfgixh", "111a47hwuo", "9f2rz7frbv", "raf3gs4ufz", "1urvkwvd3e", "shb1ue2adn", "m4xhb0pkrd", "2ekirf76tf", "agvsz31arj", "57epz0rn9a", "lju27ruedu", "z1owargelr", "id77gncis8" },
},

			"rating" : map[string]interface{}{
	"valid": []interface{}{ 1, 10, 5.5, 7.237213359750473 },
	"invalid": []interface{}{ 0, 11 },
	"multiple_valids": []interface{}{ 1, 10, 5.5, 1.2027848540318726, 0.3050486120495358, 8.247855008448516, 2.4182360306341715, 0.18915413643883738, 4.388729606962679, 5.4314244970199885, 5.082857689879004, 2.5444105102010495, 0.4318683728142736, 1.051587031051524, 0.22144530099498627, 0.7505274142672675 },
},

},

				"casts" : map[string]interface{}{
	"valid": []interface{}{ "0dw5f5m5oe", "g56zr2lauy", "u8qlmvofgk", "u184pvwef5" },
	"invalid": []interface{}{ 10, 12.43 },
	"multiple_valids": []interface{}{ "fkvc673p91", "e1s6xb7s9s", "h03sb6yurk", "v3s5vjbx3e", "nedzpf0jhv", "9lvydulhpe", "det9s0keke", "hnk1gkvzzn", "ylzgro5tqa", "ulvmgb316f", "d3smyamufk", "ripfhcji8s", "jhuaw8zyod", "pwsh892i9u", "cpvs6s2aqp" },
},

}

func TestAccMoviesMovie_Basic(t *testing.T) {
	var movie_default models.Movie
	var movie_updated models.Movie
	resourceName := "movies_movie.test"
	
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		CheckDestroy:      testAccCheckMoviesMovieDestroy,
		Steps: []resource.TestStep{
					{
			Config: CreateAccMovieWithoutTitle(),
			ExpectError: regexp.MustCompile(`Missing required argument`),
		},
	{
		Config: CreateAccMovieConfig(),
		Check: resource.ComposeTestCheckFunc(
			testAccCheckMoviesMovieExists(resourceName, &movie_default),
							resource.TestCheckResourceAttr(resourceName, "title", fmt.Sprintf("%v", searchInObject(resourceMovieTest, "title.valid.0"))),
	
							resource.TestCheckResourceAttr(resourceName, "isbn", ""),
	
							resource.TestCheckResourceAttr(resourceName, "genre", "thriller"),
	
							resource.TestCheckResourceAttr(resourceName,"director.#", "0"),
	
							resource.TestCheckResourceAttr(resourceName,"rating.#", "0"),
	
							resource.TestCheckResourceAttr(resourceName,"casts.#", "0"),
	
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
			}
		},
	})
}

func TestAccMoviesMovie_Update(t *testing.T) {
	var movie_default models.Movie
	var movie_updated models.Movie
	resourceName := "movies_movie.test"
	
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		CheckDestroy:      testAccCheckMoviesMovieDestroy,
		Steps: append([]resource.TestStep{
			{
				Config: CreateAccMovieConfig(),
				Check: testAccCheckMoviesMovieExists(resourceName, &movie_default),
			},
		},generateStepForUpdatedAttr(resourceName, &movie_default, &movie_updated)...),
	})
}

func TestAccMoviesMovie_NegativeCases(t *testing.T) {
	resourceName := "movies_movie.test"
	
	// [TODO]: Add makeTestVariable() to utils.go file
	// rName := makeTestVariable(acctest.RandString(5))
	// rOther := makeTestVariable(acctest.RandString(5))
	
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		CheckDestroy:      testAccCheckMoviesMovieDestroy,
		Steps: append([]resource.TestStep{
			{
				Config: CreateAccMovieConfig(),
			},
		},generateNegativeSteps(resourceName)...),
	})
}

				



func TestAccMoviesMovie_MultipleCreateDelete(t *testing.T) {
	resourceName := "movies_movie.test"
	
	// [TODO]: Add makeTestVariable() to utils.go file
	// rName := makeTestVariable(acctest.RandString(5))
	// rOther := makeTestVariable(acctest.RandString(5))
	
	resource.ParallelTest(t, resource.TestCase{
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
			`,searchInObject(resourceMovieTest, "isbn.valid.0"),
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
	`,searchInObject(resourceMovieTest, "title.valid.0"))
	return resource
}

func CreateAccMovieConfigWithOptional() string {
	resource := createMovieConfig(getParentMovie())
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
	`,searchInObject(resourceMovieTest, "title.valid.0"),
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
	`,searchInObject(resourceMovieTest, "title.valid.0"),
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
	`,searchInObject(resourceMovieTest, "title.valid.0"),
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
	`,searchInObject(resourceMovieTest, "title.valid.0"),
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
	`,searchInObject(resourceMovieTest, "title.valid.0"),
searchInObject(resourceMovieTest, "isbn.valid.0"),
searchInObject(resourceMovieTest, "genre.valid.0"),
searchInObject(resourceMovieTest, "director.firstname.valid.0"),
searchInObject(resourceMovieTest, "director.lastname.valid.0"),
searchInObject(resourceMovieTest, "rating.rater.valid.0"),
searchInObject(resourceMovieTest, "rating.rating.valid.0"),
value)
			return resource
		}

	func generateStepForUpdatedAttr(resourceName string,movie_default,movie_updated *models.Movie) []resource.TestStep{
		testSteps := make([]resource.TestStep, 0, 1)
		var valid []interface{} 
				valid = searchInObject(resourceMovieTest, "isbn.valid").([]interface{})
				for _, value := range valid {
					v := fmt.Sprintf("%v", value)
					testSteps = append(testSteps,resource.TestStep{
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
					testSteps = append(testSteps,resource.TestStep{
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
				testSteps = append(testSteps,resource.TestStep{
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
				testSteps = append(testSteps,resource.TestStep{
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
						testSteps = append(testSteps,resource.TestStep{
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


	func generateNegativeSteps(resourceName string) []resource.TestStep{
		//Use Update Config Function with false value 
		testSteps := make([]resource.TestStep, 0, 1)
		var invalid []interface{} 
				invalid = searchInObject(resourceMovieTest, "genre.invalid").([]interface{})
				for _, value := range invalid {
					testSteps = append(testSteps,resource.TestStep{
						Config: CreateAccMovieUpdatedAttrGenre(value),
						ExpectError: regexp.MustCompile(expectErrorMap["StringInSlice"]),
 
					})
				} 
					
								invalid = searchInObject(resourceMovieTest, "rating.rating.invalid").([]interface{})
			for _, value := range invalid {
				testSteps = append(testSteps,resource.TestStep{
					Config: CreateAccMovieUpdatedAttrRatingRating(value),
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
		for i,val := range multipleValues {
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


	func getParentMovie()  []string{
		t := []string{}
		t = append(t, movieBlock())
		return t
	}

	func movieBlock() string{
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
	`,searchInObject(resourceMovieTest, "title.valid.0"),
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
