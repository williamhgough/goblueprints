package meander

import "strings"

type journey struct {
	Name       string
	PlaceTypes []string
}

// Journeys is a simple representation of the exposed data
var Journeys = []interface{}{
	&journey{Name: "Romantic", PlaceTypes: []string{"park", "bar", "movie_theater", "restaurant", "florist", "taxi_stand"}},
	&journey{Name: "Shopping", PlaceTypes: []string{"department_store", "cafe", "clothing_store", "jewelry_store", "shoe_store"}},
	&journey{Name: "Night Out", PlaceTypes: []string{"bar", "casino", "food", "bar", "night_club", "bar", "bar", "hospital"}},
	&journey{Name: "Culture", PlaceTypes: []string{"museum", "cafe", "cemetery", "library", "art_gallery"}},
	&journey{Name: "Pamper", PlaceTypes: []string{"hair_care", "beauty_salon", "cafe", "spa"}},
}

func (j *journey) Public() interface{} {
	return map[string]interface{}{
		"name":    j.Name,
		"journey": strings.Join(j.PlaceTypes, "|"),
	}
}
