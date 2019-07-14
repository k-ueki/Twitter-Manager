package main

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/k-ueki/TwitterManager/config"
	"github.com/k-ueki/TwitterManager/timeline"
	"github.com/k-ueki/TwitterManager/users"
)

//to Front
// test to get timeline
//func APISet(w http.ResponseWriter, r *http.Request) {
//	var tcl = NewClient()
//
//	path := baseURL + "/statuses/home_timeline.json"
//	body := tcl.GetTimeline(path)
//	fmt.Println("HELLO")
//	w.Write(body)
//}

// -------Followers----------
func NewUsersClient() *users.Client {
	conf, token, client := config.Set()

	return &users.Client{
		Config:     conf,
		Token:      token,
		HttpClient: client,
	}
}

func Followers(w http.ResponseWriter, r *http.Request) {
	var ucl = NewUsersClient()
	method := r.Method

	if method == "GET" {
		path := baseURL + "followers/list.json?count=1"
		body := ucl.GetFollowersList(path)

		fmt.Println(string(body))
		//Sep(string(body), ",")
		fmt.Fprintf(w, string(body))
	} else if method == "POST" {

	}
}

//func (c *config.Client) Followers(w http.ResponseWriter, r *http.Request) {
//	v := url.Values{}
//	v.Set("count", "1000")
//
//	followers, err := c.TwitterApi.GetFriendshipsLookup(v)
//	fmt.Println(followers, err)
//}

// ---------------------------

// ----------Tweets-----------

//func NewClient() (*config.Client, *tweets.Client) {
func NewClient() *timeline.Client {
	conf, token, client := config.Set()

	return &timeline.Client{
		Config:     conf,
		Token:      token,
		HttpClient: client,
	}
}

// ---------------------------

// ---------timeline----------

// --------------------------
// ---------others----------
func Sep(str, separator string) {
	tmp := strings.Split(str, separator)
	fmt.Println("TRMP", tmp)
	fmt.Println("KOKOK", tmp[0])

}

// -------------------------
