package alfred

import "encoding/json"

type Item struct {
	Title    string `json:"title"`
	Valid    bool   `json:"valid"`
	Subtitle string `json:"subtitle"`
	Arg      string `json:"arg"`
}

type Response struct {
	Items []Item `json:"items"`
}

func MakeResponse() Response {
	items := []Item{}
	ret := Response{items}
	return ret
}

func ToJson(response Response) string {
	ret, _ := json.Marshal(response)
	return string(ret)
}
