package resolver

type Cents struct {
	Total     int `json:"total"`
	Deposited int `json:"deposited"`
	Earned    int `json:"earned"`
	Given     int `json:"given"`
}
