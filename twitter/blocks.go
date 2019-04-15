package twitter

import (
	"net/http"

	"github.com/dghubble/sling"
)

// BlockService provides methods for accessing Twitter block API
// endpoints.
type BlockService struct {
	sling *sling.Sling
}

// newBlockService returns a new BlockService.
func newBlockService(sling *sling.Sling) *BlockService {
	return &BlockService{
		sling: sling.Path("blocks/"),
	}
}

// BlockCreateParams are parameters for BlockService.Create
type BlockCreateParams struct {
	ScreenName string `url:"screen_name,omitempty"`
	UserID     int64  `url:"user_id,omitempty"`
}

type BlockedUser struct {
	ContributorsEnabled bool   `json:"contributors_enabled"`
	CreatedAt           string `json:"created_at"`
	DefaultProfile      bool   `json:"default_profile"`
	DefaultProfileImage bool   `json:"default_profile_image"`
	Description         string `json:"description"`
	Entities            struct {
		Description struct {
			Urls []interface{} `json:"urls"`
		} `json:"description"`
	} `json:"entities"`
	FavouritesCount                int         `json:"favourites_count"`
	FollowRequestSent              bool        `json:"follow_request_sent"`
	FollowersCount                 int         `json:"followers_count"`
	Following                      bool        `json:"following"`
	FriendsCount                   int         `json:"friends_count"`
	GeoEnabled                     bool        `json:"geo_enabled"`
	ID                             int         `json:"id"`
	IDStr                          string      `json:"id_str"`
	IsTranslator                   bool        `json:"is_translator"`
	Lang                           string      `json:"lang"`
	ListedCount                    int         `json:"listed_count"`
	Location                       string      `json:"location"`
	Name                           string      `json:"name"`
	Notifications                  bool        `json:"notifications"`
	ProfileBackgroundColor         string      `json:"profile_background_color"`
	ProfileBackgroundImageURL      string      `json:"profile_background_image_url"`
	ProfileBackgroundImageURLHTTPS string      `json:"profile_background_image_url_https"`
	ProfileBackgroundTile          bool        `json:"profile_background_tile"`
	ProfileImageURL                string      `json:"profile_image_url"`
	ProfileImageURLHTTPS           string      `json:"profile_image_url_https"`
	ProfileLinkColor               string      `json:"profile_link_color"`
	ProfileSidebarBorderColor      string      `json:"profile_sidebar_border_color"`
	ProfileSidebarFillColor        string      `json:"profile_sidebar_fill_color"`
	ProfileTextColor               string      `json:"profile_text_color"`
	ProfileUseBackgroundImage      bool        `json:"profile_use_background_image"`
	Protected                      bool        `json:"protected"`
	ScreenName                     string      `json:"screen_name"`
	ShowAllInlineMedia             bool        `json:"show_all_inline_media"`
	StatusesCount                  int         `json:"statuses_count"`
	TimeZone                       string      `json:"time_zone"`
	URL                            interface{} `json:"url"`
	UtcOffset                      int         `json:"utc_offset"`
	Verified                       bool        `json:"verified"`
}

func (s *BlockService) Create(params *BlockCreateParams) (*BlockedUser, *http.Response, error) {
	user := new(BlockedUser)
	apiError := new(APIError)
	resp, err := s.sling.New().Post("create.json").QueryStruct(params).Receive(user, apiError)
	return user, resp, relevantError(err, *apiError)
}
