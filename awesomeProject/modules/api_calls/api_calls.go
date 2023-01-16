package api_calls

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func Get() string {
	url := "https://dog.ceo/api/breeds/image/random"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return err.Error()
	}
	req.Header.Add("Cookie", "AB_forceLoginOnHome=0; PHPSESSID=45caa82a789f30ae93a05e6b1352d1bd; lamao=2147483647")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return err.Error()
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return err.Error()
	}
	fmt.Println(body, "\n")

	//////// Way to convert to object the json response ////////
	var jsonMap map[string]interface{}
	json.Unmarshal([]byte(string(body)), &jsonMap)

	fmt.Println(jsonMap)
	////////////////////////////////////////////////////////////

	return string(body)
}

func Post() string {
	url := "http://sandbox.oauth.undostres.com.mx/api/oauth/login"
	method := "POST"

	payload := strings.NewReader(`{
		"email":"oscar.millan@undostres.com.mx",
		"password":"Deaththekid1#"
	}`)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		return err.Error()
	}
	req.Header.Add("SESSID", "9292aa9bfebf1ce000880325f439588c")
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		return err.Error()
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err.Error()
	}
	fmt.Println(body, "\n")

	//////// Way to convert to object the json response ////////
	var jsonMap map[string]interface{}
	json.Unmarshal([]byte(string(body)), &jsonMap)

	fmt.Println(jsonMap, "\n")
	////////////////////////////////////////////////////////////

	return string(body)
}
