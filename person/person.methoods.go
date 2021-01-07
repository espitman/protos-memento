package Movie

import "encoding/json"

func (t *ResponseDetails) Unmarshal(entry interface{}) ResponseDetails {
	obj, _ := json.Marshal(entry)
	response := ResponseDetails{}
	_ = json.Unmarshal(obj, &response)
	return response
}

func (t *Movie) Unmarshal(entry interface{}) Movie {
	obj, _ := json.Marshal(entry)
	response := Movie{}
	_ = json.Unmarshal(obj, &response)
	return response
}
