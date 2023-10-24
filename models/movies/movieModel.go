package movies

type Movie struct {
	ID           int             `json:"id"`
	Name         string          `json:"name"`
	Description  string          `json:"description"`
	Rating       int             `json:"rating"`
	DirectorName string          `json:"directorName"`
	TrailerLink  string          `json:"trailerLink"`
	Category     []MovieCategory `json:"category"`
	Views        int             `json:"views"`
	MovieSeasons []MovieSeasons  `json:"movieSeasons"`
	CreatedBy    string          `json:"createdBy"`
	CreatedAt    string          `json:"createdAt"`
	UpdatedBy    *string         `json:"updatedBy"`
	UpdatedAt    *string         `json:"updatedAt"`
}

type MovieCategory struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type MovieSeasons struct {
	SeasonID     int           `json:"id"`
	MovieDetails []MovieDetail `json:"movieDetails"`
}

type MovieDetail struct {
	EpisodeID   int    `json:"epdisodeId"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Video_link  string `json:"video_link"`
}

type MovieBasicDetailResponse struct {
	ID           int             `json:"id"`
	Name         string          `json:"name"`
	Description  string          `json:"description"`
	Rating       int             `json:"rating"`
	DirectorName string          `json:"directorName"`
	Category     []MovieCategory `json:"Category"`
	Views        int             `json:"views"`
	TrailerLink  string          `json:"trailerLink"`
}

var Sample_data = []Movie{
	Movie{
		ID:           1,
		Name:         "batman",
		Description:  "",
		Rating:       9,
		DirectorName: "santana",
		Category: []MovieCategory{
			MovieCategory{
				Name:        "action",
				Description: "",
			},
		},
		TrailerLink: "",
		Views:       9,
		MovieSeasons: []MovieSeasons{
			MovieSeasons{
				SeasonID: 1,
				MovieDetails: []MovieDetail{
					MovieDetail{
						EpisodeID:   1,
						Title:       "",
						Description: "",
						Video_link:  "",
					},
				},
			},
		},
		CreatedBy: "backend_proj_os",
		CreatedAt: "",
		UpdatedBy: nil,
		UpdatedAt: nil,
	},
	Movie{
		ID:           2,
		Name:         "your name",
		Description:  "",
		Rating:       9,
		DirectorName: "yamoto chinkai",
		Category: []MovieCategory{
			MovieCategory{
				Name:        "anime",
				Description: "",
			},
		},
		TrailerLink: "",
		Views:       20,
		MovieSeasons: []MovieSeasons{
			MovieSeasons{
				SeasonID: 1,
				MovieDetails: []MovieDetail{
					MovieDetail{
						EpisodeID:   1,
						Title:       "",
						Description: "",
						Video_link:  "",
					},
				},
			},
		},
		CreatedBy: "backend_proj_os",
		CreatedAt: "",
		UpdatedBy: nil,
		UpdatedAt: nil,
	},
	Movie{
		ID:           3,
		Name:         "bloodhound",
		Description:  "",
		Rating:       5,
		DirectorName: "apex",
		Category: []MovieCategory{
			MovieCategory{
				Name:        "game",
				Description: "",
			},
		},
		TrailerLink: "",
		Views:       25,
		MovieSeasons: []MovieSeasons{
			MovieSeasons{
				SeasonID: 1,
				MovieDetails: []MovieDetail{
					MovieDetail{
						EpisodeID:   1,
						Title:       "",
						Description: "",
						Video_link:  "",
					},
				},
			},
		},
		CreatedBy: "backend_proj_os",
		CreatedAt: "",
		UpdatedBy: nil,
		UpdatedAt: nil,
	},

	Movie{
		ID:           4,
		Name:         "japane in the hole",
		Description:  "",
		Rating:       3,
		DirectorName: "sora aoi",
		Category: []MovieCategory{
			MovieCategory{
				Name:        "porno",
				Description: "",
			},
		},
		TrailerLink: "",
		Views:       0,
		MovieSeasons: []MovieSeasons{
			MovieSeasons{
				SeasonID: 1,
				MovieDetails: []MovieDetail{
					MovieDetail{
						EpisodeID:   1,
						Title:       "",
						Description: "",
						Video_link:  "",
					},
				},
			},
		},
		CreatedBy: "backend_proj_os",
		CreatedAt: "",
		UpdatedBy: nil,
		UpdatedAt: nil,
	},
}
