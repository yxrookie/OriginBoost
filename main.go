package main

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)




func main() {
	// load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	appid := os.Getenv("APPID")
	appkey := os.Getenv("APPKEY")
	
	fromLang := "en"
	toLang := "zh"
	endpoint := "http://api.fanyi.baidu.com"
	path := "/api/trans/vip/translate"
	query := "Hello World!"

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
	//fmt.Println(result)
	extractField(result)
	//resultJSON, _ := json.MarshalIndent(result, "", "    ")
	//fmt.Println(string(resultJSON))
}

func extractField(result map[string]interface{}) {
	if result == nil {
		log.Fatal("The original extract field is nil")
	}
	from := result["from"].(string)
	to := result["to"].(string)
	temp_result := result["trans_result"].([]interface{})
	trans_result := temp_result[0].(map[string]interface{})
	dst := trans_result["dst"].(string)
	src := trans_result["src"].(string)
	fmt.Println(from, to, dst, src)
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
