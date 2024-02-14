package handlers

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
	_ "github.com/go-sql-driver/mysql"
)

// Type 1: Getting data from sql and putting data into a map, unstructured data approach
func (h *Handlers) GetAllData(w http.ResponseWriter, r *http.Request) {
	h.infoLog.Println("hit the nasa data")

	//Set up connection
	db, err := sql.Open("mysql", "testUser:testPassword@/testGoDb")
	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	defer db.Close()

	// Execute the query
	rows, err := db.Query("SELECT * FROM `NasaData` LIMIT 50")
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	defer rows.Close()

	// Get the column names from the query
	var columns []string
	// Columns: ["column name 1", "column name 2", ...]
	columns, err = rows.Columns()
	checkErr(err)

	colNum := len(columns)

	var results []map[string]interface{}

	for rows.Next() {
		// Prepare to read row using Scan
		r := make([]interface{}, colNum)
		//r: [<nil>, <nil>, ...(amount of columns total)]

		for i := range r {
			r[i] = &r[i]
		}
		//r: [0xc00020c1e0, 0xc00020c1f0, ...(memory locations of interfaces)]
		// This is done because the scan below takes in a list of memory addresses when scanning in data

		// Read rows using Scan
		err = rows.Scan(r...)
		checkErr(err)
		//Format for this example
		//r: [[51 55 57][65 99 104 105 114 97 115]...(byte array representations for data in the database row)]
		//r: [[id as byte array] [name as byte array] ...[other attributes for database row as byte array]]

		// Create a row map to store row's data, and converts from byte array to string
		var row = map[string]interface{}{}
		for i := range r {
			row[columns[i]] = (string(r[i].([]byte)))
		}
		// Creates a map for the data in the database row, similar to a python dict or javascript object
		// row: map[id: 370 name:Achiras ...(other attributes for database row as a map and strings]

		// Append to the final results slice
		results = append(results, row)
	}

	// results: [
	// map[id: 370 name:Achiras ...(other attributes for database row as a map and strings],
	// map[id: 1 name:Aachen ...(other attributes for database row as a map and strings]
	// ]
	// From here we can JSON marshal this data and send back to user as JSON data

	// Structure for response
	var payload jsonResponse
	payload.Error = false
	payload.Data = results
	h.writeJSON(w, http.StatusOK, payload)

}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// Type 2: Getting data from sql and putting data into structured data approach

// Create a struct that contains all fields for database row
type NasaData struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	NameType    string `json:"nameType"`
	Recclass    string `json:"recclass"`
	Mass        string `json:"mass"`
	Fall        string `json:"fall"`
	Year        string `json:"year"`
	Reclat      string `json:"reclat"`
	Reclong     string `json:"reclong"`
	Geolocation string `json:"geolocation"`
}

func (h *Handlers) GetNasaDataForId(w http.ResponseWriter, r *http.Request) {
	// How you access url parameters
	nasaDataId := chi.URLParam(r, "nasaDataId")
	h.infoLog.Println("hit the nasa data for an id", nasaDataId)

	//Set up connection
	db, err := sql.Open("mysql", "testUser:testPassword@/testGoDb")
	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	defer db.Close()

	// Execute the query
	rows, err := db.Query("SELECT * FROM `NasaData` WHERE id=" + nasaDataId + " LIMIT 50")
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	defer rows.Close()

	// Create a list of structs
	var nasaDataPoints []NasaData
	for rows.Next() {
		// Create a struct for single database row
		var nasaDataPoint NasaData
		// Deserialize database row into struct
		if err := rows.Scan(&nasaDataPoint.Id, &nasaDataPoint.Name, &nasaDataPoint.NameType, &nasaDataPoint.Recclass, &nasaDataPoint.Mass, &nasaDataPoint.Fall, &nasaDataPoint.Year, &nasaDataPoint.Reclat, &nasaDataPoint.Reclong, &nasaDataPoint.Geolocation); err != nil {
			fmt.Println("failed at id" + nasaDataId)
		}
		// Data Format nasaDataPoint: {id name nameType recclass mass fall year reclat reclong geolocation}
		// Example nasaDataPoint: {1 Aachen Valid L5 21 Fell 1880-01-01 00:00:00 50.775000 6.083330 {"coordinates":[6.08333,50.775],"type":"Point"}}
		nasaDataPoints = append(nasaDataPoints, nasaDataPoint)
	}

	// In this example the list should always be length 1 because we are seraching by id, but it shows a more generic way

	// Structure for response
	var payload jsonResponse
	payload.Error = false
	payload.Data = nasaDataPoints
	h.writeJSON(w, http.StatusOK, payload)

}

func (h *Handlers) LoadNasaData(w http.ResponseWriter, r *http.Request) {

	db, err := sql.Open("mysql", "testUser:testPassword@/testGoDb")
	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	defer db.Close()

	// Reads JSON file, assumes server is being ran from server directory
	content, err := os.ReadFile("../data-load/testNasa.json")
	if err != nil {
		log.Fatal("Error when opening file: ", err)
	}

	// Puts JSON file contents into an object/slice
	var result []map[string]any
	json.Unmarshal([]byte(content), &result)

	// First 11 nasa metorite landings in json file
	neededData := result[0:11]

	// Query creation
	query := "INSERT IGNORE INTO `NasaData` (`id`, `name`, `nameType`, `recclass`, `mass`, `fall`, `year`, `reclat`, `reclong`, `geolocation`) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"

	for _, v := range neededData {

		// Pulls all needed fields for insert
		id := v["id"]
		name := v["name"]
		nameType := v["nametype"]
		recclass := v["recclass"]
		mass, _ := strconv.Atoi(v["mass"].(string))
		fall := v["fall"]
		year, _ := time.Parse("2006-01-02T15:04:00.000", v["year"].(string))
		//reclat, _ := strconv.ParseFloat(v["reclat"].(string), 32)
		reclat := v["reclat"]
		reclong := v["reclong"]

		geolocation := v["geolocation"]
		geoLocationJsonString, err := json.Marshal(geolocation)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		geolocationFinal := string(geoLocationJsonString)

		// Inserts into database
		insertResult, err := db.ExecContext(context.Background(), query, id, name, nameType, recclass, mass, fall, year, reclat, reclong, geolocationFinal)
		if err != nil {
			log.Fatalf("impossible insert data: %s", err)
		}
		// Prints what id was inserted, or if failure occured
		idInserted, err := insertResult.LastInsertId()
		if err != nil {
			log.Fatalf("impossible to retrieve last inserted id: %s", err)
		}
		log.Printf("inserted id: %d", idInserted)
	}

	// Structure for response
	var payload jsonResponse
	payload.Error = false
	payload.Data = ""
	h.writeJSON(w, http.StatusOK, payload)
}

// EXTRA CODE FROM OTHER ATTEMPTS
// nasaDataPointAsJson, err := json.Marshal(nasaDataPoints)
// if err != nil {
// 	fmt.Println(err)
// 	return
// }
// fmt.Println(nasaDataPointAsJson)

// // Get the column names from the query
// var columns []string
// columns, err = rows.Columns()
// checkErr(err)

// colNum := len(columns)

// var results []map[string]interface{}

// for rows.Next() {
// 	// Prepare to read row using Scan
// 	r := make([]interface{}, colNum)
// 	for i := range r {
// 		r[i] = &r[i]
// 	}

// 	// Read rows using Scan
// 	err = rows.Scan(r...)
// 	checkErr(err)

// 	// Create a row map to store row's data
// 	var row = map[string]interface{}{}
// 	for i := range r {
// 		row[columns[i]] = (string(r[i].([]byte)))
// 	}

// 	// Append to the final results slice
// 	results = append(results, row)
// }

//fmt.Println(results) // You can then json.Marshal or w/e
