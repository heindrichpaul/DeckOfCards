package deckOfCards

import jsoniter "github.com/json-iterator/go"

func UnmarshalDraw(data []byte) (*Draw, error) {
	var r *Draw
	json := jsoniter.ConfigCompatibleWithStandardLibrary
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Draw) Marshal() ([]byte, error) {
	json := jsoniter.ConfigCompatibleWithStandardLibrary
	return json.Marshal(r)
}

type Draw struct {
	Success   bool    `json:"success"`
	Cards     []*Card `json:"cards"`
	Remaining int     `json:"remaining"`
}
