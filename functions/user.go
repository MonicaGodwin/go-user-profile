package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type User struct{
	Results []Result `json:"results"` 
}

type Name struct{
	First string `json:"first"`
	Last string `json:"last"`
}
type Picture struct{
	Large string `json:"large"`
}

type Result struct{
	Name Name `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
	Picture Picture `json:"picture"`
}

func UserProfile() (Result, error) {
	resp, err := http.Get("https://randomuser.me/api/")
	if err != nil {
		return Result{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println(resp.Status)
		return Result{}, fmt.Errorf("status not OK have %v", resp.StatusCode)
	}

	var APIResponse User

	err = json.NewDecoder(resp.Body).Decode(&APIResponse)
	if err != nil {
		return Result{}, err
	}
	return APIResponse.Results[0], nil
}
