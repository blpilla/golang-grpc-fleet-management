package vehicle

type Vehicle struct {
	ID      int    `json:"id"`
	Make    string `json:"make"`
	Model   string `json:"model"`
	Year    int    `json:"year"`
	License string `json:"license"`
}
