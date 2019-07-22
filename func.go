package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/k-ueki/TwitterManager/config"
	"github.com/k-ueki/TwitterManager/db"
	"github.com/k-ueki/TwitterManager/timeline"
	"github.com/k-ueki/TwitterManager/users"
)

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

	var mode string
	if r.ContentLength != 0 {
		mode = GetMode(r)
	}

	var dbh = &db.DBHandler{
		DB: config.SetDB(),
	}

	pathToGetFollowers := baseURL + "followers/list.json"
	pathToGetIds := baseURL + "followers/ids.json"
	bodyF, Ids := ucl.GetFollowersList(pathToGetFollowers, pathToGetIds)

	if mode == "register" {
		_, fromdb := dbh.Select("followers")

		//dbの情報とIdsを比較
		newf, byef := db.FindNewBye(&Ids, fromdb)
		fmt.Println("NEW", newf, "\nBYE", byef) //Ids

		if len(byef.Ids) != 0 {
			body := ucl.ConvertIdsToUsers(byef.Ids)
			fmt.Fprintf(w, string(body))
			return
		}

		//init register
		//if err := dbh.RegisterIds(Ids); err != nil {
		//	fmt.Println("ERR", err)
		//}
		//fmt.Println("OK")

		//-----------new register動作確認済み
		//		if err := dbh.RegisterIds(newf); err != nil {
		//			fmt.Println("ERR", err)
		//		}

		//------------bye dropout動作確認済み
		//if len(byef.Ids) >= 1 {
		//dbh.DropOutByes(byef)
		//}

	}

	//fmt.Println(dbh, Ids)
	fmt.Println(string(bodyF))
	//	fmt.Fprintf(w, string(bodyF))
}

// ---------------------------

// ----------Tweets-----------

// ---------------------------

// ---------timeline----------
func NewTimelineClient() *timeline.Client {
	conf, token, client := config.Set()

	return &timeline.Client{
		Config:     conf,
		Token:      token,
		HttpClient: client,
	}
}

func Timeline(w http.ResponseWriter, r *http.Request) {
	var tcl = NewTimelineClient()

	path := baseURL + "/statuses/home_timeline.json"
	body := tcl.GetTimeline(path)
	w.Write(body)
}

// --------------------------
func GetMode(r *http.Request) string {
	var body = make([]byte, r.ContentLength)
	r.Body.Read(body)
	return Separate(string(body))
}
func Separate(str string) string {
	tmp := strings.Split(str, "=")
	return tmp[1]
}
