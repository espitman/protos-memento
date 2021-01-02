package merchandising

import "encoding/json"

func (t *RequestCreate) Unmarshal(entry interface{}) RequestCreate {
	obj, _ := json.Marshal(entry)
	response := RequestCreate{}
	_ = json.Unmarshal(obj, &response)
	return response
}

func (t *ResponseDetails) Unmarshal(entry interface{}) ResponseDetails {
	obj, _ := json.Marshal(entry)
	response := ResponseDetails{}
	_ = json.Unmarshal(obj, &response)
	return response
}

func (t *ResponseDetailsWithMovies) Unmarshal(entry interface{}) ResponseDetailsWithMovies {
	obj, _ := json.Marshal(entry)
	response := ResponseDetailsWithMovies{}
	_ = json.Unmarshal(obj, &response)
	return response
}

func (t *Movie) Unmarshal(entry interface{}) Movie {
	obj, _ := json.Marshal(entry)
	response := Movie{}
	_ = json.Unmarshal(obj, &response)
	return response
}
