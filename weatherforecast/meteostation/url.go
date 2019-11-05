package meteostation

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
)

func GetUrlCurrent(city string) []byte {
	proxyStr := "http://172.26.43.18:3128"
	proxyURL, err := url.Parse(proxyStr)
	if err != nil {
		fmt.Println(err)
	}

	urlStr := "http://api.openweathermap.org/data/2.5/weather?q=" +
		city +
		"&lang=ru" +
		"&units=metric" +
		"&appid=2c19a8c670afc70f2ae7a81f229fce3d"

	url, err := url.Parse(urlStr)
	if err != nil {
		fmt.Println(err)
	}

	transport := &http.Transport{Proxy: http.ProxyURL(proxyURL)}

	client := &http.Client{Transport: transport}

	request, err := http.NewRequest("GET", url.String(), nil)
	if err != nil {
		fmt.Println(err)
	}

	response, err := client.Do(request)
	if err != nil {
		fmt.Println(err)
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
	}

	// resp, err := http.Get(urlStr)
	// if err != nil {
	// 	fmt.Println("Error")
	// }
	// defer resp.Body.Close()
	// body, err := ioutil.ReadAll(resp.Body)
	// if err != nil {
	// 	fmt.Println("Error")
	// }
	return body
}

func GetUrlDaily(city string, cnt int) []byte {
	proxyStr := "http://172.26.43.18:3128"
	proxyURL, err := url.Parse(proxyStr)
	if err != nil {
		fmt.Println(err)
	}

	urlStr := "http://api.openweathermap.org/data/2.5/find?q=" +
		city +
		"&units=metric" +
		"&lang=ru" +
		"&cnt=" +
		strconv.Itoa(cnt) +
		"&appid=2c19a8c670afc70f2ae7a81f229fce3d"

	url, err := url.Parse(urlStr)
	if err != nil {
		fmt.Println(err)
	}

	transport := &http.Transport{Proxy: http.ProxyURL(proxyURL)}

	client := &http.Client{Transport: transport}

	request, err := http.NewRequest("GET", url.String(), nil)
	if err != nil {
		fmt.Println(err)
	}

	response, err := client.Do(request)
	if err != nil {
		fmt.Println(err)
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
	}
	// resp, err := http.Get(urlStr)
	// if err != nil {
	// 	fmt.Println("Error")
	// }
	// defer resp.Body.Close()
	// body, err := ioutil.ReadAll(resp.Body)
	// if err != nil {
	// 	fmt.Println("Error")
	// }

	return body
}
