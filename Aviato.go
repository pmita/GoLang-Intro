package main

import ("fmt"
		"net/http"
		"encoding/json"
		"time"
)

type WeatherData struct{
	Timezone 		string	`json:"timezone"`
	Currently struct{
		Summary		string	`json:"summary"`
		Temperature	float64	`json:"temperature"`
	} `json:"currently"`
}

type FlightData struct{
	FlightStatuses []struct{
		FlightID				int		`json:"flightID"`
		FlightNumber 			int 	`json:"flightNumber"`
		DepartureAirportFsCode	string	`json:"departureAirportFsCode"`
	}`json:"flightStatuses"`
}

func weatherApp() []string {
	CitiesToVisit := [][] string {
		{"LHR", "51.5074", "0.1278"},
		{"WAW","37.9838","23.7275"},
		{"MXP", "45.4642", "9.1900"},
		{"WAW", "41.3851", "2.1734"},
		{"LIS","38.7223","9.1393"},
		{"FRA", "50.1109","8.6821"},
		{"CDG","48.8566","2.3522"},
		{"VIE","48.2082","16.3738"},
		{"PRG","50.0755","14.4378"},
		{"BUD","47.4979","19.0402"},
		{"SVO","55.7558","37.6173"},
		{"RIX","56.9496","24.1052"},
	} 
	var data WeatherData
	var perfectCitiesToVisit []string
	
	for i:=0; i<12; i++{
			var urlAPI = fmt.Sprintf("https://api.darksky.net/forecast/58ffcf3ca186b5068bc9918ad2c16d8e/%s,%s", CitiesToVisit[i][1], CitiesToVisit[i][2])
			response, err :=http.Get(urlAPI)
			if err != nil{
				fmt.Println("Failed HTTP request %s\n", err)
			}

			defer response.Body.Close()

			if err := json.NewDecoder(response.Body).Decode(&data); err !=nil{
				fmt.Println( WeatherData{}, err)
			} else {
				if (data.Currently.Summary == "Clear") {
					perfectCitiesToVisit = append(perfectCitiesToVisit, CitiesToVisit[i][0])
				}
			}
		}

		return (perfectCitiesToVisit)
}

func flightApp() FlightData{
	//currentData= time.Now()Format("2006/01/02")
	var Data FlightData
	var citiesToVisitURLs = fmt.Sprintf("http://www.json-generator.com/api/json/get/coVHwkCWaa?indent=2")
	response, err :=http.Get(citiesToVisitURLs)
	if err != nil{
		fmt.Println("Failed HTTP request %s\n", err)
	}

	defer response.Body.Close()

	if err := json.NewDecoder(response.Body).Decode(&Data); err !=nil{
		fmt.Println( FlightData{}, err)
	} else {
		//fmt.Println(Data)
		//counter := len(Data.FlightStatuses)
		//fmt.Println(counter)
	}

	return (Data)
}

func main(){

	perfectCitiesToVisit:= weatherApp()
	totalPerfectCitiesToVisit := len(perfectCitiesToVisit)
	if  totalPerfectCitiesToVisit!= 0{
		fmt.Println("Your ideal weather conditions exists in:")
		fmt.Println(perfectCitiesToVisit)
	}else{
		fmt.Println("Apologies but there are no cities with good weather!")
	}

	Data := flightApp()
	counter := len(Data.FlightStatuses)


	for m:=0; m< totalPerfectCitiesToVisit; m++{
		for p:=0; p<counter; p++{
			if Data.FlightStatuses[p].DepartureAirportFsCode == perfectCitiesToVisit[m]{
				fmt.Printf("Your destination for %s has flight code:%d and flight number: %d\n", perfectCitiesToVisit[m], Data.FlightStatuses[p].FlightID, Data.FlightStatuses[p].FlightNumber )
			}
		}
	}
}
