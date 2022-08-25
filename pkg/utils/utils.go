package utils

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)


func ParseBody(r *http.Request, x interface{})  {
  if body, err := ioutil.ReadAll(r.Body); err == nil {
    if err := json.Unmarshal([]byte(body), x); err == nil {
      return
    }
  }

 //  body, err := ioutil.ReadAll(r.Body)
	// if err != nil {
	// 	fmt.Printf("Get the following error: %v", err)
	// }
	//
	// err = json.Unmarshal([]byte(body), x)
	// if err != nil {
	// 	fmt.Printf("Error at unmarshalling %v", err)
	// }
}
