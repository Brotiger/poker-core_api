package response

type GameList struct {
	Total int64  `json:"size"`
	Games []Game `json:"games"`
}
