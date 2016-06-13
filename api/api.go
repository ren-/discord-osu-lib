package api

import (
	"encoding/json"
	"errors"
	"strconv"
)

// Game types
const (
	OSU   = "0"
	TAIKO = "1"
	CTB   = "2"
	MANIA = "3"
)

// Beatmap Identifier Types for use with GetBeatmaps
const (
	BEATMAPSET = "s"
	BEATMAPID  = "b"
	USERID     = "u"
)

// Modlist
const (
	None   = 0
	NoFail = 1 << (iota - 1)
	Easy
	NoVideo
	Hidden
	HardRock
	SuddenDeath
	DoubleTime
	Relax
	HalfTime
	Nightcore
	Flashlight
	Autoplay
	SpunOut
	Relax2
	Perfect
	Key4
	Key5
	Key6
	Key7
	Key8
	FadeIn
	Random
	LastMod
)

var (
	API_URL           string = "https://osu.ppy.sh/api/"
	API_RECENT_PLAYS  string = "get_user_recent"
	API_GET_BEATMAPS  string = "get_beatmaps"
	API_GET_USER      string = "get_user"
	API_GET_SCORES    string = "get_scores"
	API_GET_USER_BEST string = "get_user_best"
)

func (c *Config) SetAPIKey(API_KEY string) error {
	var err error
	if len(API_KEY) <= 1 {
		err = errors.New("API Key: invalid/non-existant API key.")
		return err
	} else {
		c.API_KEY = API_KEY
	}
	return err
}

func (c Config) BuildRecentURL(USER_ID string, GAME_TYPE string, LIMIT int) string {
	return API_URL + API_RECENT_PLAYS + "?k=" + c.API_KEY + "&u=" + USER_ID + "&m=" + GAME_TYPE + "&limit=" + strconv.Itoa(LIMIT)
}

func (c Config) BuildBeatmapURL(ID string, TYPE string) string {
	return API_URL + API_GET_BEATMAPS + "?k=" + c.API_KEY + "&" + TYPE + "=" + ID
}

func (c Config) BuildUserURL(USER_ID string, GAME_TYPE string, DAYS string) string {
	return API_URL + API_GET_USER + "?k=" + c.API_KEY + "&u=" + USER_ID + "&m=" + GAME_TYPE + "&event_days=" + DAYS
}

func (c Config) BuildUserBestURL(USER_ID string, GAME_TYPE string, NUM_SCORES int) string {
	return API_URL + API_GET_USER_BEST + "?k=" + c.API_KEY + "&u=" + USER_ID + "&m=" + GAME_TYPE + "&limit=" + strconv.Itoa(NUM_SCORES)
}

func (c Config) BuildScoreURL(BEATMAP_ID string, USER_ID string, GAME_TYPE string) string {
	return API_URL + API_GET_SCORES + "?k=" + c.API_KEY + "&b=" + BEATMAP_ID + "&m=" + GAME_TYPE + "&u=" + USER_ID
}

func (c Config) GetUserBest(USER_ID string, GAME_TYPE string, NUM_SCORES int) ([]TopScore, error) {
	var songs []TopScore
	url := c.BuildUserBestURL(USER_ID, GAME_TYPE, NUM_SCORES)
	content, err := GetResults(url)

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(content, &songs)

	if err != nil {
		return nil, errors.New("JSON: Couldn't process the results into JSON. Possibly wrong page/wrong API key/offline server.")
	}

	return songs, err
}
func (c Config) GetRecentPlays(USER_ID string, GAME_TYPE string, LIMIT int) ([]Song, error) {
	var songs []Song
	url := c.BuildRecentURL(USER_ID, GAME_TYPE, LIMIT)
	html, err := GetResults(url)

	if err != nil {
		return nil, err
	}
	//fmt.Println(string(html))
	err = json.Unmarshal(html, &songs)

	if err != nil {
		//fmt.Println(err)
		//fmt.Println(string(html))
		return nil, errors.New("JSON(GetRecentPlays): The response was not in JSON format, probably server error.")
	}

	return songs, err
}
