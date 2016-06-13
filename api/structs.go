package api

type Config struct {
	API_KEY string
}

type Beatmap struct {
	Beatmapset_ID     string
	Beatmap_ID        string
	Approved          string
	Total_Length      string
	Hit_Length        string
	Version           string
	File_MD5          string
	Diff_Size         string
	Diff_Overall      string
	Diff_Approach     string
	Diff_Drain        string
	Mode              string
	Approved_Date     string
	Last_Update       string
	Artist            string
	Title             string
	Creator           string
	Bpm               string
	Source            string
	Tags              string
	Genre_ID          string
	Language_ID       string
	Favourite_Count   string
	PlayCount         string
	PassCount         string
	Max_Combo         string
	Difficulty_Rating string
}

type Song struct {
	BeatmapID   int    `json:"beatmap_id,string" db:"beatmap_id"`
	Score       int    `json:"score,string" db:"score"`
	MaxCombo    int    `json:"maxcombo,string" db:"max_combo"`
	Count50     int    `json:"count50,string" db:"count50"`
	Count100    int    `json:"count100,string" db:"count100"`
	Count300    int    `json:"count300,string" db:"count300"`
	CountMiss   int    `json:"countmiss,string" db:"count_miss"`
	CountKatu   int    `json:"countkatu,string" db:"count_katu"`
	CountGeki   int    `json:"countgeki,string" db:"count_geki"`
	Perfect     int    `json:"perfect,string" db:"perfect"`
	EnabledMods int    `json:"enabled_mods,string" db:"enabled_mods"`
	UserID      int    `json:"user_id,string" db:"user_id"`
	Date        string `json:"date" db:"date"`
	Rank        string `json:"rank" db:"rank"`
	Username    string `json:"-" db:"username"`
}

type User struct {
	User_ID         string
	Username        string
	Count300        string
	Count100        string
	Count50         string
	PlayCount       string
	Ranked_Score    string
	Total_Score     string
	PP_Rank         string
	Level           string
	PP_Raw          string
	Accuracy        string
	Count_Rank_SS   string
	Count_Rank_S    string
	Count_Rank_A    string
	Country         string
	PP_Country_Rank string
	Events          []Event
}
type Event struct {
	Display_HTML  string
	Beatmap_ID    string
	Beatmapset_ID string
	Date          string
	EpicFactor    string
}

type Score struct {
	Score        string
	Username     string
	MaxCombo     string
	Count50      string
	Count100     string
	Count300     string
	CountMiss    string
	CountKatu    string
	CountGeki    string
	Perfect      string
	Enabled_Mods string
	User_ID      string
	Date         string
	Rank         string
	PP           string
}

type TopScore struct {
	Beatmap_ID   string
	Score        string
	MaxCombo     string
	Count50      string
	Count100     string
	Count300     string
	CountMiss    string
	CountKatu    string
	CountGeki    string
	Perfect      string
	Enabled_Mods string
	User_ID      string
	Date         string
	Rank         string
	PP           float64 `json:",string"`
}
