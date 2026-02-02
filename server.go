package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

type APIResponse struct {
	USDBRL struct {
		Bid string `json:"bid"`
	} `json:"USDBRL"`
}

func getExchangeRate() (*APIResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 200*time.Millisecond)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", "https://economia.awesomeapi.com.br/json/last/USD-BRL", nil)
	if err != nil {
		return nil, err
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var data APIResponse
	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}

func saveExchangeRate(db *sql.DB, bid string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Millisecond)
	defer cancel()

	query := `INSERT INTO exchange_rate (bid, created_at) VALUES (?, ?)`
	_, err := db.ExecContext(ctx, query, bid, time.Now())

	return err
}

func main() {
	db, err := sql.Open("sqlite3", "./quotations.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	stmt := `CREATE TABLE IF NOT EXISTS exchange_rate (
    		 id INTEGER PRIMARY KEY AUTOINCREMENT,
    		 bid TEXT,
    		 created_at DATETIME
);`
	_, err = db.Exec(stmt)
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/cotacao", func(w http.ResponseWriter, r *http.Request) {
		quotation, err := getExchangeRate()
		if err != nil {
			log.Printf("Error fetching exchange rate: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		err = saveExchangeRate(db, quotation.USDBRL.Bid)
		if err != nil {
			log.Printf("Error saving to database: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(quotation.USDBRL)
	})

	log.Printf("server started at port 8080")
	http.ListenAndServe(":8080", nil)
}
