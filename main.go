package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

const (
	DBUsername = "root"
	DBPassword = "1234"
	DBName     = "prediction"
)

var db *sql.DB

type Win struct {
	ID     int    `json:"id"`
	Winner string `json:"Winner Name"`
	Rating string `json:"Team Rating"`
	Result string `json:"Result"`
}

func init() {
	dataSourceName := fmt.Sprintf("%s:%s@tcp(localhost:3306)/%s", DBUsername, DBPassword, DBName)
	var err error
	db, err = sql.Open("mysql", dataSourceName)
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
}

func GenerateResult() {
	var selectedTeams struct {
		Team1 string
		Team2 string
	}

	fmt.Print("Enter Team 1: ")
	fmt.Scan(&selectedTeams.Team1)

	fmt.Print("Enter Team 2: ")
	fmt.Scan(&selectedTeams.Team2)

	// Simulate team data
	rand.Seed(time.Now().UnixNano())
	scoreTeam1 := rand.Intn(101)
	scoreTeam2 := rand.Intn(101)

	var winner string
	if scoreTeam1 > scoreTeam2 {
		winner = selectedTeams.Team1
		randomRating := fmt.Sprintf("%d", rand.Intn(6))

		_, err := db.Exec("INSERT INTO win (`Winner`, `Rating`, `Result`) VALUES (?, ?, ?)", winner, randomRating, scoreTeam1)
		if err != nil {
			log.Println(err)
			return
		}

	} else if scoreTeam2 > scoreTeam1 {
		winner = selectedTeams.Team2
		randomRating := fmt.Sprintf("%d", rand.Intn(6))

		_, err := db.Exec("INSERT INTO win (`Winner`, `Rating`, `Result`) VALUES (?, ?, ?)", winner, randomRating, scoreTeam2)
		if err != nil {
			log.Println(err)
			return
		}

	} else {
		winner = "It's a Tie"
	}

	fmt.Printf("Team 1 Score: %d\n", scoreTeam1)
	fmt.Printf("Team 2 Score: %d\n", scoreTeam2)
	fmt.Printf("Winner: %s\n", winner)
}

// GetWinnersHandler retrieves and prints the database information of winners.
func GetWinnersHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Query the database to retrieve winners.
	rows, err := db.Query("SELECT * FROM win")
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var winners []Win
	for rows.Next() {
		var win Win
		err := rows.Scan(&win.Winner, &win.Rating, &win.Result)
		if err != nil {
			log.Println(err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		winners = append(winners, win)
	}

	// Encode and send the winners as JSON.
	err = json.NewEncoder(w).Encode(winners)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func main() {
	GenerateResult()
	http.HandleFunc("/winners", GetWinnersHandler)

	log.Println("Server started on :8080")
	http.ListenAndServe(":8080", nil)
}
