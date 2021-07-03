package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type Result struct {
	Request_hash          string      `json:"request_hash"`
	Request_cached        bool        `json:"request_cached"`
	Request_cache_expired int         `json:"request_cache_expired"`
	Anime                 []AnimeItem `json:"anime"`
}

type AnimeItem struct {
	Mal_id            int           `json:"mal_id"`
	Title             string        `json:"title"`
	Video_url         string        `json:"video_url"`
	Url               string        `json:"url"`
	Image_url         string        `json:"image_url"`
	Type              string        `json:"type"`
	Watching_status   int           `json:"watching_status"`
	Score             int           `json:"score"`
	Watched_episodes  int           `json:"watched_episodes"`
	Total_episodes    int           `json:"total_episodes"`
	Airing_status     int           `json:"airing_status"`
	Season_name       string        `json:"season_name"`
	Season_year       string        `json:"season_year"`
	Has_episode_video bool          `json:"has_episode_video"`
	Has_promo_video   bool          `json:"has_promo_video"`
	Has_video         bool          `json:"has_video"`
	Is_rewatching     bool          `json:"is_rewatching"`
	Tags              string        `json:"tags"`
	Rating            string        `json:"rating"`
	Start_date        string        `json:"start_date"`
	End_date          string        `json:"end_date"`
	Watch_start_date  string        `json:"watch_start_date"`
	Watch_end_date    string        `json:"watch_end_date"`
	Days              string        `json:"days"`
	Storage           string        `json:"storage"`
	Priority          string        `json:"priority"`
	Added_to_list     bool          `json:"added_to_list"`
	Studios           []interface{} `json:"studios"`
	Licensors         []interface{} `json:"licensors"`
}

func main() {
	user := os.Args[1]
	openMessage := startMessage(user)
	fmt.Println(openMessage)
	makeRequest(user)
}

func createHttpString(uname string) string {
	return fmt.Sprintf("https://api.jikan.moe/v3/user/%s/animelist", uname)
}

func makeRequest(uname string) {
	resp, err := http.Get(createHttpString(uname))
	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatalln(err)
	}

	// sb := string(body)
	// fmt.Println(sb)

	var result Result
	json.Unmarshal(body, &result)

	// fmt.Println(result.Anime)

	for i, val := range result.Anime {
		fmt.Printf("%d:  %s\n", i, val.Title)
	}

	// var animeList []AnimeItem

	// if err := json.Unmarshal([]byte(string(result.Anime)), &animeList); err != nil {
	// 	panic(err)
	// }

	// for _, val := range animeList {
	// 	fmt.Println(val)
	// }

}

func startMessage(uname string) string {
	return "Searching for user " + uname
}
