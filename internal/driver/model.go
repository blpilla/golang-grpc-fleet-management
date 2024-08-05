package driver

type Driver struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	License   string `json:"license"`
}
