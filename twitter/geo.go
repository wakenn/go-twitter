package twitter

import (
	"net/http"

	"github.com/dghubble/sling"
)

// Result from reverse geo lookup
type Geo struct {
	Result struct {
		Places []*Place `json:"places"`
	} `json:"result"`
}

// GeoParams are the parameters for SearchService.Tweets
type GeoParams struct {
	Lat   string `url:"lat,omitempty"`
	Long  string `url:"long,omitempty"`
	Query string `url:"query,omitempty"`
}

// GeoService provides methods for accessing Twitter geo API endpoints.
type GeoService struct {
	sling *sling.Sling
}

// newGeoService returns a new SearchService.
func newGeoService(sling *sling.Sling) *GeoService {
	return &GeoService{
		sling: sling.Path("geo/"),
	}
}

// Tweets returns a collection of Tweets matching a search query.
// https://dev.twitter.com/rest/reference/get/search/tweets
func (g *GeoService) Reverse(params *GeoParams) (*Geo, *http.Response, error) {
	geo := new(Geo)
	apiError := new(APIError)
	resp, err := g.sling.New().Get("reverse_geocode.json").QueryStruct(params).Receive(geo, apiError)
	return geo, resp, relevantError(err, *apiError)
}

func (g *GeoService) Search(params *GeoParams) (*Geo, *http.Response, error) {
	geo := new(Geo)
	apiError := new(APIError)
	resp, err := g.sling.New().Get("search.json").QueryStruct(params).Receive(geo, apiError)
	return geo, resp, relevantError(err, *apiError)
}
