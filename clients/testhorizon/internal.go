package testhorizon

import (
	"encoding/json"
	"net/http"

	"github.com/test/go/support/errors"
)

func decodeResponse(resp *http.Response, object interface{}) (err error) {
	defer resp.Body.Close()
	decoder := json.NewDecoder(resp.Body)

	if !(resp.StatusCode >= 200 && resp.StatusCode < 300) {
		testhorizonError := &Error{
			Response: resp,
		}
		decodeError := decoder.Decode(&testhorizonError.Problem)
		if decodeError != nil {
			return errors.Wrap(decodeError, "error decoding testhorizon.Problem")
		}
		return testhorizonError
	}

	err = decoder.Decode(&object)
	if err != nil {
		return
	}
	return
}

func loadMemo(p *Payment) error {
	res, err := http.Get(p.Links.Transaction.Href)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	return json.NewDecoder(res.Body).Decode(&p.Memo)
}
