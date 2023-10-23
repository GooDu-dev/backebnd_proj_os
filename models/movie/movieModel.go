package models

type Movie struct {
	ID           int             `json:"id"`
	Name         string          `json:"name"`
	Description  string          `json:"description"`
	Rating       int             `json:"rating"`
	DirectorName string          `json:"directorName"`
	TrailerLink  string          `json:"trailerLink"`
	Category     []MovieCategory `json:"category"`
	Views        int             `json:"views"`
	MovieSeason  []MovieSeason   `json:"movieSeason"`
	CreatedBy    string          `json:"createdBy"`
	CreatedAt    string          `json:"createdAt"`
	UpdatedBy    *string         `json:"updatedBy"`
	UpdatedAt    *string         `json:"updatedAt"`
}

type MovieCategory struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type MovieSeason struct {
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
	Category     []MovieCategory `json:"movieCategory"`
	Views        int             `json:"views"`
	TrailerLink  string          `json:"trailerLink"`
}

var Sample_data = []Movie{
	Movie{
		ID:           1,
		Name:         "",
		Description:  "",
		Rating:       0,
		DirectorName: "kayato",
		Category: []MovieCategory{
			MovieCategory{
				Name:        "porno",
				Description: "",
			},
		},
		TrailerLink: "",
		Views:       0,
		MovieSeason: []MovieSeason{
			MovieSeason{
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