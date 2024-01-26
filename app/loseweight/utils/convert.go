package utils

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

var APPID, APPKEY, ENDPOINT, PATH string 

func init() {
	fmt.Println("Initializing variables")
	// load .env file
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	APPID = os.Getenv("APPID")
	APPKEY = os.Getenv("APPKEY")

	ENDPOINT = "http://api.fanyi.baidu.com"
	PATH = "/api/trans/vip/translate"
}


func Convert(query, fromLang, toLang string) string {
	salt := strconv.Itoa(randInt(32768, 65536))
	sign := makeMd5(APPID + query + salt +APPKEY)

	data := url.Values{
		"appid":  {APPID},
		"q":      {query},	
		"from":   {fromLang},
		"to":     {toLang},
		"salt":   {salt},
		"sign":   {sign},
	}

	response, err := http.PostForm(ENDPOINT+PATH, data)
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
	if result == nil {
		log.Fatal("The original extract field is nil")
	}
	temp_result := result["trans_result"].([]interface{})
	trans_result := temp_result[0].(map[string]interface{})
	dst := trans_result["dst"].(string)
	return dst
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