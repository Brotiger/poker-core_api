package response

type GameList struct {
	Total int64  `json:"total" example:"100"`
	Games []Game `json:"games"`
}
