package common

import (
	"crypto/tls"
	"errors"
	"fmt"
	smartisanData "github.com/princeyuaner/smartisan/Data"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strings"
)

func HttpSend(url string, body io.Reader, method string, headers map[string]string) ([]byte, error) {
	if len(method) == 0 {
		method = "GET"
	}
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	for k, v := range headers {
		fmt.Println(k, v)
		req.Header.Set(k, v)
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, errors.New(resp.Status)
	}

	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if len(content) == 0 {
		return nil, errors.New("nil resp")
	}
	return content, nil
}

func HttpSend2(url string, body io.Reader, method string, headers map[string]string) ([]byte, error) {
	if len(method) == 0 {
		method = "GET"
	}
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}

	//req.Header.Set("Content-Type", "application/json")
	for k, v := range headers {
		fmt.Println(k, v)
		req.Header.Set(k, v)
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, errors.New(resp.Status)
	}

	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if len(content) == 0 {
		return nil, errors.New("nil resp")
	}
	return content, nil
}

func HttpFormSend(url string, formBody url.Values, method string, headers map[string]string) ([]byte, error) {
	if len(method) == 0 {
		method = "GET"
	}
	body := strings.NewReader(formBody.Encode())
	fmt.Println("form", url, method, formBody.Encode())
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}

	//req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	for k, v := range headers {
		req.Header.Set(k, v)
	}

	smartisanData.CookieJar, _ = cookiejar.New(nil)
	smartisanData.Tr = &http.Transport{
		TLSClientConfig:    &tls.Config{InsecureSkipVerify: true},
		DisableCompression: true,
	}
	http.DefaultClient.Jar = smartisanData.CookieJar
	http.DefaultClient.Transport = smartisanData.Tr
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, errors.New(resp.Status)
	}

	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if len(content) == 0 {
		return nil, errors.New("nil resp")
	}
	return content, nil
}
