package interopb

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/mohfahrur/interop-core/entity"
)

type InteropbAgent interface {
	SendTelegram(data entity.User) (err error)
}

type InteropbDomain struct {
}

func NewinteropbDomain() *InteropbDomain {

	return &InteropbDomain{}
}

func (d *InteropbDomain) SendTelegram(data entity.User) (err error) {
	url := "http://localhost:5001/send-telegram"
	method := "POST"

	payload := strings.NewReader(`{
		"user": "` + data.User + `",
		"email": "` + data.Email + `",
		"item": "` + data.Item + `"
	}`)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
	return
}
