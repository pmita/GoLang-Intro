package main

import ("fmt"
		//"strings"
		//"io/ioutil"
		"net/http"
		"encoding/json"
		//"log"

)

type WeatherData struct{
	Timezone 		string	`json:"timezone"`
	Currently struct{
		//Time 		int		`json:"time"`
		Summary		string	`json:"summary"` //Acceptable summarys: Clear, Overcast, Partly Cloudy, etc.
		//Temperature	float64	`json:"temperature"`
		//Visibility	float64	`json:"visibility"` //Viisibility needs to be >= 5 
	} `json:"currently"`
}


func main(){
	citiesToVisit := [][] string {
		{"LHR", "51.5074", "0.1278"},
		{"ATH","37.9838","23.7275"},
		{"MXP", "45.4642", "9.1900"},
		{"BCN", "41.3851", "2.1734"},
		{"LIS","38.7223","9.1393"},
		{"FRA", "50.1109","8.6821"},
		{"CDG","48.8566","2.3522"},
		{"VIE","48.2082","16.3738"},
		{"PRG","50.0755","14.4378"},
		{"BUD","47.4979","19.0402"},
	} 

	var d WeatherData
	var perfectCitiesToVisit [] string
	
	j:=1
	for i:=0; i<10; i++{
			var urlAPI = fmt.Sprintf("https://api.darksky.net/forecast/58ffcf3ca186b5068bc9918ad2c16d8e/%s,%s",citiesToVisit[i][j], citiesToVisit[i][j+1])
			response, err :=http.Get(urlAPI)
			if err != nil{
				fmt.Println("Failed HTTP request %s\n", err)
			}

			defer response.Body.Close()
			
			if err := json.NewDecoder(response.Body).Decode(&d); err !=nil{
				fmt.Println( WeatherData{}, err)
			} else {
				//fmt.Println(d)
				if d.Currently.Summary == "Clear"{
					perfectCitiesToVisit = append(perfectCitiesToVisit, citiesToVisit[i][0])
				}
			}

		}
		fmt.Println(perfectCitiesToVisit)
		fmt.Println(len(perfectCitiesToVisit))
	
}
