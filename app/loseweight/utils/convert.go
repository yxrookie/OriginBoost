package utils

import (
	"OriginBoost/pkg/config"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"net/url"

	"strconv"
	"time"
)
 

func Convert(query, fromLang, toLang string) string {
	tmep := config.Get("appid")
	if tmep != "" {
		fmt.Println(tmep)
	} else {
		fmt.Println("无法取出")
	}

	salt := strconv.Itoa(randInt(32768, 65536))
	sign := makeMd5(config.Get("appid") + query + salt +config.Get("appkey"))

	data := url.Values{
		"appid":  {config.Get("appid")},
		"q":      {query},	
		"from":   {fromLang},
		"to":     {toLang},
		"salt":   {salt},
		"sign":   {sign},
	}
	response, err := http.PostForm(config.Get("endpoint")+config.Get("path"), data)
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