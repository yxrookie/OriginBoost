package main

import (
	
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

func main() {
	appid := "20240125001951130"
	appkey := "qVxWQ5eDioMWMTr5_YOh"
	fromLang := "en"
	toLang := "zh"
	endpoint := "http://api.fanyi.baidu.com"
	path := "/api/trans/vip/translate"
	query := "Hello World! This is 1st paragraph.\nThis is 2nd paragraph."

	salt := strconv.Itoa(randInt(32768, 65536))
	sign := makeMd5(appid + query + salt + appkey)

	data := url.Values{
		"appid":  {appid},
		"q":      {query},
		"from":   {fromLang},
		"to":     {toLang},
		"salt":   {salt},
		"sign":   {sign},
	}

	response, err := http.PostForm(endpoint+path, data)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	var result map[string]interface{}
	json.Unmarshal(body, &result)

	resultJSON, _ := json.MarshalIndent(result, "", "    ")
	fmt.Println(string(resultJSON))
}

func makeMd5(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}

func randInt(min int, max int) int {
	rand.Seed(time.Now().UnixNano())
	return min + rand.Intn(max-min)
}
