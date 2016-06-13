package api

import (
	"errors"
	"io/ioutil"
	"net/http"
)

func GetResults(URL string) ([]byte, error) {
	res, err := http.Get(URL)

	if err != nil {
		return nil, errors.New("HTTP: Could not open a connection to the Osu! API server.")
	}

	content, err := ioutil.ReadAll(res.Body) // Change to io.Copy later
	res.Body.Close()

	if err != nil {
		return nil, errors.New("HTTP: Could not read the results.")
	}

	return content, err
}
