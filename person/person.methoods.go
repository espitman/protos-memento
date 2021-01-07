package person

import "encoding/json"

func (t *ResponseDetails) Unmarshal(entry interface{}) ResponseDetails {
	obj, _ := json.Marshal(entry)
	response := ResponseDetails{}
	_ = json.Unmarshal(obj, &response)
	return response
}

func (t *Person) Unmarshal(entry interface{}) Person {
	obj, _ := json.Marshal(entry)
	response := Person{}
	_ = json.Unmarshal(obj, &response)
	return response
}
