package twitter

import (
	"fmt"
	"github.com/dghubble/oauth1"
	"io/ioutil"
	"net/url"
	"os"
)

func GetWebhook(){
	consumerKey := os.Getenv("CONSUMER_KEY")
	consumerSecret := os.Getenv("CONSUMER_SECRET")
	accessToken := os.Getenv("ACCESS_TOKEN")
	accessSecret := os.Getenv("ACCESS_SECRET")
	//fmt.Println(consumerSecret)
	if consumerKey == "" || consumerSecret == "" || accessToken == "" || accessSecret == "" {
		panic("Missing required environment variable")
	}

	config := oauth1.NewConfig(consumerKey, consumerSecret)
	token := oauth1.NewToken(accessToken, accessSecret)

	// httpClient will automatically authorize http.Request's
	httpClient := config.Client(oauth1.NoContext, token)

	path := "https://api.twitter.com/1.1/account_activity/all/dev/webhooks.json"
	values := url.Values{}
	values.Set("url","http://3.18.125.1/twitter/webhook")
	resp, _ := httpClient.PostForm(path,values)
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Printf("Raw Response Body:\n%v\n", string(body))
}