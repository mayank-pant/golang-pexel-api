package main

type SearchResult struct {
	Page         int32   `json:"page"`
	PerPage      int32   `json:"per_page"`
	TotalResults int32   `json:"total_Results"`
	NextPage     string  `json:"next_page"`
	Photos       []Photo `json:"photos"`
}

type Photo struct {
	Id              int32        `json:"id"`
	Width           int32        `json:"width"`
	Height          int32        `json:"height"`
	Url             string       `json:"url"`
	Photographer    string       `json:"photographer"`
	PhotographerUrl string       `json:"photographer_url"`
	Src             PhotosSource `json:"src"`
}

type PhotosSource struct {
	Original  string `json:"original"`
	Large     string `json:"large"`
	Large2x   string `json:"large2x"`
	Medium    string `json:"medium"`
	Small     string `json:"small"`
	Portrait  string `json:"portrait"`
	Square    string `json:"square"`
	Landscape string `json:"landscape"`
	Tiny      string `json:"tiny"`
}

type CuratedResult struct {
	Page     int     `json:"page"`
	PerPage  int     `json:"per_page"`
	Photos   []Photo `json:"photos"`
	NextPage string  `json:"next_page"`
}

type VideoSearchResult struct {
	Page         int     `json:"page"`
	PerPage      int     `json:"per_page"`
	TotalResults int     `json:"total_results"`
	NextPage     string  `json:"next_page"`
	Videos       []Video `json:"videos"`
}

type PopularVideos struct {
	Page         int     `json:"page"`
	PerPage      int     `json:"per_page"`
	TotalResults int     `json:"total_results"`
	Url          string  `json:"url"`
	Videos       []Video `json:"videos"`
}

type Video struct {
	ID            int             `json:"id"`
	Width         int             `json:"width"`
	Height        int             `json:"height"`
	URL           string          `json:"url"`
	Image         string          `json:"image"`
	Duration      float64         `json:"duration"`
	FullRes       interface{}     `json:"full_res"`
	VideoFiles    []VideoFiles    `json:"video_files"`
	VideoPictures []VideoPictures `json:"video_pictures"`
}

type VideoFiles struct {
	ID       int    `json:"id"`
	Quality  string `json:"quality"`
	FileType string `json:"file_type"`
	Width    int    `json:"width"`
	Height   int    `json:"height"`
	Link     string `json:"link"`
}

type VideoPictures struct {
	ID      int    `json:"id"`
	Picture string `json:"picture"`
	Nr      int    `json:"nr"`
}
