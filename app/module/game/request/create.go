package request

type Create struct {
	Name       string  `json:"name" example:"test" validate:"required"`
	MaxPlayers int     `json:"max_players" validate:"required,min=3,max=6" example:"5"`
	Password   *string `json:"password,omitempty" example:"123456"`
}
