package movie

import "encoding/json"

func (t *TMDBDetails) Unmarshal(entry interface{}) TMDBDetails {
	obj, _ := json.Marshal(entry)
	response := TMDBDetails{}
	_ = json.Unmarshal(obj, &response)
	return response
}

func (t *ResponseDetails) Unmarshal(entry interface{}) ResponseDetails {
	obj, _ := json.Marshal(entry)
	response := ResponseDetails{}
	_ = json.Unmarshal(obj, &response)
	return response
}
