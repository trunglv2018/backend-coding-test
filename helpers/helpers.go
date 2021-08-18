package helpers

import (
	"encoding/json"
	"io"
	"io/ioutil"
)

// UnmarshalResponseBody parses the JSON-encoded data and stores the result
// in the value pointed to by v.
func UnmarshalResponseBody(r io.Reader, v interface{}) error {
	body, err := ioutil.ReadAll(r)
	if err != nil {
		return err
	}
	return json.Unmarshal(body, &v)
}
