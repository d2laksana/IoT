package request

type SuhuCreateRequest struct {
	Temp float64 `json:"temp" validate:"required"`
	Hum  float64 `json:"hum" validate:"required"`
}

type SuhuUpdateRequest struct {
	Temp float64 `json:"temp"`
	Hum  float64 `json:"hum"`
}
