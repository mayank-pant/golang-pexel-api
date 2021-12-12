package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

func handleRequests() {
	http.HandleFunc("/SearchPhotos", SearchPhotoshandler)
	http.HandleFunc("/CuratedPhotos", CuratedPhotoshandler)
	http.HandleFunc("/GetPhoto", GetPhotohandler)
	http.HandleFunc("/GetRandomPhoto", GetRandomPhotohandler)
	http.HandleFunc("/SearchVideo", SearchVideohandler)
	http.HandleFunc("/PopularVideo", PopularVideohandler)
	http.HandleFunc("/GetRandomVideo", GetRandomVideohandler)
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func SearchPhotoshandler(w http.ResponseWriter, r *http.Request) {

	os.Setenv("PexelsToken", "563492ad6f9170000100000199336114cf7349b4bbce9eee7cb15519")
	var TOKEN = os.Getenv("PexelsToken")
	var c = NewClient(TOKEN)

	result, err := c.SearchPhotos("waves", 15, 1)
	//result, err := c.GetRandomVideo()

	if err != nil {
		fmt.Errorf("Search error: %v", err)
	}
	if result.Page == 0 {
		fmt.Errorf("search result wrong")
	}

	json.NewEncoder(w).Encode(result)

}

func CuratedPhotoshandler(w http.ResponseWriter, r *http.Request) {

	os.Setenv("PexelsToken", "563492ad6f9170000100000199336114cf7349b4bbce9eee7cb15519")
	var TOKEN = os.Getenv("PexelsToken")
	var c = NewClient(TOKEN)

	result, err := c.CuratedPhotos(15, 1)
	//result, err := c.GetRandomVideo()

	if err != nil {
		fmt.Errorf("Search error: %v", err)
	}
	if result.Page == 0 {
		fmt.Errorf("search result wrong")
	}

	json.NewEncoder(w).Encode(result)

}

func GetPhotohandler(w http.ResponseWriter, r *http.Request) {

	os.Setenv("PexelsToken", "563492ad6f9170000100000199336114cf7349b4bbce9eee7cb15519")
	var TOKEN = os.Getenv("PexelsToken")
	var c = NewClient(TOKEN)

	result, err := c.GetPhoto(1)
	//result, err := c.GetRandomVideo()

	if err != nil {
		fmt.Errorf("Search error: %v", err)
	}
	if result.Id == 0 {
		fmt.Errorf("search result wrong")
	}

	json.NewEncoder(w).Encode(result)

}

func GetRandomPhotohandler(w http.ResponseWriter, r *http.Request) {

	os.Setenv("PexelsToken", "563492ad6f9170000100000199336114cf7349b4bbce9eee7cb15519")
	var TOKEN = os.Getenv("PexelsToken")
	var c = NewClient(TOKEN)

	result, err := c.GetRandomPhoto()
	//result, err := c.GetRandomVideo()

	if err != nil {
		fmt.Errorf("Search error: %v", err)
	}
	if result.Id == 0 {
		fmt.Errorf("search result wrong")
	}

	json.NewEncoder(w).Encode(result)

}

func SearchVideohandler(w http.ResponseWriter, r *http.Request) {

	os.Setenv("PexelsToken", "563492ad6f9170000100000199336114cf7349b4bbce9eee7cb15519")
	var TOKEN = os.Getenv("PexelsToken")
	var c = NewClient(TOKEN)

	result, err := c.SearchVideo("waves", 15, 1)
	//result, err := c.GetRandomVideo()

	if err != nil {
		fmt.Errorf("Search error: %v", err)
	}
	if result.Page == 0 {
		fmt.Errorf("search result wrong")
	}

	json.NewEncoder(w).Encode(result)

}

func PopularVideohandler(w http.ResponseWriter, r *http.Request) {

	os.Setenv("PexelsToken", "563492ad6f9170000100000199336114cf7349b4bbce9eee7cb15519")
	var TOKEN = os.Getenv("PexelsToken")
	var c = NewClient(TOKEN)

	result, err := c.PopularVideo(15, 1)
	//result, err := c.GetRandomVideo()

	if err != nil {
		fmt.Errorf("Search error: %v", err)
	}
	if result.Page == 0 {
		fmt.Errorf("search result wrong")
	}

	json.NewEncoder(w).Encode(result)

}

func GetRandomVideohandler(w http.ResponseWriter, r *http.Request) {

	os.Setenv("PexelsToken", "563492ad6f9170000100000199336114cf7349b4bbce9eee7cb15519")
	var TOKEN = os.Getenv("PexelsToken")

	var c = NewClient(TOKEN)

	result, err := c.GetRandomVideo()

	if err != nil {
		fmt.Errorf("Search error: %v", err)
	}
	if result.ID == 0 {
		fmt.Errorf("search result wrong")
	}

	json.NewEncoder(w).Encode(result)

}
