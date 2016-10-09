package igdb

const (
	defaultBaseURL = "https://igdbcom-internet-game-database-v1.p.mashape.com/"
	headerMashape  = "X-Mashape-Key"
)

//Client struct
type Client struct {
	apiKey         string
	defaultBaseURL string
	headerMashape  string
}

//NewClient create new client
func NewClient(key string) *Client {
	return &Client{apiKey: key, defaultBaseURL: defaultBaseURL, headerMashape: headerMashape}
}
