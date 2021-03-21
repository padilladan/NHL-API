package app

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type teamsReceived struct {
	Teams []struct {
		ID    int    `json:"id"`
		Name  string `json:"name"`
		Venue struct {
			City string `json:"city"`
		}
	}
}

func GetAllTeams() {

	listofTeams := teamsReceived{}

	url := "https://statsapi.web.nhl.com/api/v1/teams"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
	}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}

	err = json.Unmarshal(body, &listofTeams)
	if err != nil {
		fmt.Printf("There was an error unmarshaling data received from API, %v", err)
	}
	fmt.Println(listofTeams)
}
