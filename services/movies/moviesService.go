package movies

import (
	"sort"
	"strings"

	// testModel "backend_proj_os/models/test"
	movieModel "backend_proj_os/models/movie"
)

// var data = []movieModel.MovieBasicDetailResponse{

// 	movieModel.MovieBasicDetailResponse{
// 		TrailerLink: "",
// 	}
// 	// testModel.TestModelResponse{Code: "20000", Message: "Everything Alright"},
// 	// testModel.TestModelResponse{Code: apiErr.InternalServerError.Code(), Message: apiErr.InternalServerError.EnMessage()},
// }
// var data = movieModel.Sample_data{
// 	// movieModel.Sample_data
// 	for i := 0; i < count; i++ {

// 	}
// }

func GetTopMovies(count int) ([]movieModel.MovieBasicDetailResponse, error) {
	// var tops [count]movieModel.MovieBasicDetailResponse
	tops := make([]movieModel.MovieBasicDetailResponse, 5)
	var all = movieModel.Sample_data

	// * not good behavior to sort all data without permission
	// sort.SliceStable(all, func(i, j int) bool {
	// 	return all[i].Rating < all[j].Rating
	// })

	// for i := 0; i < count; i++ {

	// movieModel
	var maxRating int = -999
	var j int = 0
	for i := 0; i < len(all); i++ {
		// movieModel.Sample_data
		if all[i].Rating > maxRating {
			var cur = movieModel.MovieBasicDetailResponse{
				ID:           all[i].ID,
				Name:         all[i].Name,
				Description:  all[i].Description,
				Rating:       all[i].Rating,
				DirectorName: all[i].DirectorName,
				Category:     all[i].Category,
				Views:        all[i].Views,
				TrailerLink:  all[i].TrailerLink,
			}

			tops[j] = cur

			j++
			j %= 5
		}

	}

	sort.SliceStable(tops, func(i, j int) bool {
		return tops[i].Rating < tops[j].Rating
	})

	// return tops, nill
	return tops, nil
}

func GetMoviesByCategory(categoryName string) ([]movieModel.MovieBasicDetailResponse, error) {
	var response []movieModel.MovieBasicDetailResponse
	var all = movieModel.Sample_data

	for i := 0; i < len(all); i++ {

		var foundCat bool = false
		for j := 0; j < len(all[i].Category); j++ {
			// if strings.ToLower(all[i].Category[j].Name) == strings.ToLower(categoryName) {
			if strings.EqualFold(strings.ToLower(all[i].Category[j].Name),
				strings.ToLower(categoryName)) {
				foundCat = true
			}
		}
		if foundCat {
			var cur = movieModel.MovieBasicDetailResponse{
				ID:           all[i].ID,
				Name:         all[i].Name,
				Description:  all[i].Description,
				Rating:       all[i].Rating,
				DirectorName: all[i].DirectorName,
				Category:     all[i].Category,
				Views:        all[i].Views,
				TrailerLink:  all[i].TrailerLink,
			}
			response = append(response, cur)
		}

	}

	return response, nil
}

func GetMoviesByDirector(directorName string) ([]movieModel.MovieBasicDetailResponse, error) {
	var response []movieModel.MovieBasicDetailResponse
	var all = movieModel.Sample_data

	for i := 0; i < len(all); i++ {

		var found bool = false
		if strings.EqualFold(strings.ToLower(all[i].DirectorName),
			strings.ToLower(directorName)) {
			found = true
		}

		if found {
			var cur = movieModel.MovieBasicDetailResponse{
				ID:           all[i].ID,
				Name:         all[i].Name,
				Description:  all[i].Description,
				Rating:       all[i].Rating,
				DirectorName: all[i].DirectorName,
				Category:     all[i].Category,
				Views:        all[i].Views,
				TrailerLink:  all[i].TrailerLink,
			}
			response = append(response, cur)
		}
	}

	return response, nil

}
