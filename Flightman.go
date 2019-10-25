package main

import ("fmt"
		//"strings"
		//"io/ioutil"
		"net/http"
		"encoding/json"
		"net/url"
		//"log"

)

type FlightData struct{
	ScheduledFlights []struct{
		CarrierFsCode			string	`json:"carrierFsCode"`
		FlightNumber			string	`json:"flightNumber"`
		//DepartureTerminal		string	`json:"departureTerminal"`
		//DepartureTime			string	`json:"departureTime"`
		//Stops					int		`json:"stops"`
		//ArrivalTime 			string	`json:"arrivalTime"`
		//Codeshares	[]struct{
			//CarrierFsCode		string	`json:"carrierFsCode"`
			//FlightNumber 		string	`json:"flightNumber"`
			//ReferenceCode		int		`json:"referenceCode"`	
		//}`json:"codeshares"`
	}`json:"scheduledFlights"`
}

//var airport string

func main(){
	
	var d FlightData
	citiesToVisit := [3]string{"ATH", "FRA", "MLN"}
	var citiesToVisitURLs = fmt.Sprintf("https://api.flightstats.com/flex/schedules/rest/v1/json/from/LHR/to/%s/departing/2020/03/24?appId=605d73d7&appKey=78a37d155d4f07e5f891b8a609f50d59&codeType=IATA", citiesToVisit[1])
	//fmt.Println(example)
	base, err := url.Parse("https://api.flightstats.com/flex/schedules/rest/v1/json/from/LHR/to/")
	if err != nil {
			return
		}
	/*The base.Parse code isn't vaiable since the character ? translates to %3F when combining the paths. Need more debugging*/
	base.Path += citiesToVisit[1]
	base.Path += "/departing/2020/03/24/"
	base.Path += "?"
	base.Path +="appId=605d73d7&appKey=78a37d155d4f07e5f891b8a609f50d59&codeType=IATA"

	//params := url.Values{}
	//base.RawQuery = params.Encode()
	fmt.Println(base.String())
	//a :=base.String()
	//fmt.Println(a)
	//fmt.Println(con)
	//aa := "https://api.flightstats.com/flex/schedules/rest/v1/json/from/LHR/to/ATH/departing/2020/03/24?appId=605d73d7&appKey=78a37d155d4f07e5f891b8a609f50d59&codeType=IATA"
	response, err :=http.Get(citiesToVisitURLs)
	if err != nil{
		fmt.Println("Failed HTTP request %s\n", err)
	}

	defer response.Body.Close()
	
	if err := json.NewDecoder(response.Body).Decode(&d); err !=nil{
		fmt.Println( FlightData{}, err)
	} else {
		fmt.Println(d)
	}
	
}
	

