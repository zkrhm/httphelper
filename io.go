package httphelper

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/oleiade/reflections"
)

//ReadJSON read passed struct as json.
func ReadJSON(r *http.Request, reqObj interface{}) error {
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}
	defer r.Body.Close()
	json.Unmarshal(b, reqObj)

	return nil
}

//WriteResponse writing response content (struct) to as json
//
func WriteAsJSON(w http.ResponseWriter, resObj interface{}) error {

	w.Header().Set("Content-Type", "application/json")

	hasField, err := reflections.HasField(resObj, "Code")
	if err != nil {
		return err
	}

	if hasField {

		val, err := reflections.GetField(resObj, "Code")
		if err != nil {
			return err
		}
		// fmt.Println("has code field, code value : ", val.(int))

		code := val.(int)
		// fmt.Println("code value : %d, truth check %b", code, code == 0)
		if code == 0 {
			code = http.StatusOK
		}

		w.WriteHeader(code)
	} else {
		w.WriteHeader(http.StatusOK)
	}

	b, err := json.Marshal(resObj)
	if err != nil {
		return err
	}

	w.Write(b)

	return nil
}
