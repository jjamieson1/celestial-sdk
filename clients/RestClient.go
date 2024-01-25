package clients

import (
	"bytes"
	"io/ioutil"
	"net/http"
)

func CallRestEndPoint(url string, method string, tenantId string, b []byte) ([]byte, int, error) {
	client := &http.Client{}
	req, err := http.NewRequest(method, url, bytes.NewReader(b))
	if err != nil {
		return nil, 500, err
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("tenantId", tenantId)

	res, err := client.Do(req)
	if err != nil {
		return nil, 500, err
	}

	body, err := ioutil.ReadAll(res.Body)
	defer res.Body.Close()
	return body, res.StatusCode, err
}
