package shared

type FilterByQueryParams struct {
	TimeFrom    string            `json:"time_from"`
	TimeTo      string            `json:"time_to"`
	UserIDs     []string          `json:"user_ids"`
	StatusCodes []int             `json:"status_codes"`
	Endpoints   []string          `json:"endpoints"`
	Meta        map[string]string `json:"meta"`
}
