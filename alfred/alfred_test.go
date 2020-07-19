package alfred_test

import (
	"testing"

	. "wwei10.com/go-timer/alfred"
)

func TestJson(t *testing.T) {
	var response = MakeResponse()
	response.Items = append(response.Items, Item{"hello", false})
	response.Items = append(response.Items, Item{"world", false})
	ret := ToJson(response)
	if ret != "{\"items\":[{\"title\":\"hello\",\"valid\":false},{\"title\":\"world\",\"valid\":false}]}" {
		t.Errorf(ret)
	}
}
