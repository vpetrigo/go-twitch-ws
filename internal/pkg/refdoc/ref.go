package refdoc

import (
	"net/http"

	"golang.org/x/net/html"
)

func GetReferenceDocPage(referenceDocURL string) (*html.Node, error) {
	resp, err := http.Get(referenceDocURL)

	if err != nil {
		return nil, err
	}
	defer func() {
		_ = resp.Body.Close()
	}()

	body, err := html.Parse(resp.Body)

	if err != nil {
		return nil, err
	}

	return body, nil
}
