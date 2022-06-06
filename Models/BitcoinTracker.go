package Models

import (
	"fmt"
	"log"

	"bitcoin-checker/Config"
)

type Bitcoin struct {
	Id        uint   `json:"id"`
	CoinType  string `json:"coin_type"`
	Price     int    `json:"price"`
	Timestamp string `json:"timestamp"`
}
type BitcoinData struct {
	Url   string    `json:"url"`
	Next  string    `json:"next"`
	Count int       `json:"count"`
	Data  []Bitcoin `json:"data"`
}

func AddBitcoinInfo(CoinType string, Price int, Timestamp string) {
	db := Config.DBConnect()
	defer db.Close()
	records := `INSERT INTO bitcoin_tracker(CoinType, Price, Timestamp) VALUES (?, ?, ?)`
	query, err := db.Prepare(records)
	if err != nil {
		log.Fatal(err)
	}
	_, err = query.Exec(CoinType, Price, Timestamp)
	if err != nil {
		log.Fatal(err)
	}
}

func FetchBitcoinInfo(date string, limit, offset int) (BitcoinData, error) {

	db := Config.DBConnect()
	defer db.Close()
	var sqlWhere string
	if date != "" {
		sqlWhere = "WHERE strftime('%d-%m-%Y', Timestamp) = '" + date + "'"

	}
	queryString := fmt.Sprintf("SELECT * FROM bitcoin_tracker %s ORDER BY id DESC LIMIT %d OFFSET %d ", sqlWhere, limit, offset)
	record, err := db.Query(queryString)
	if err != nil {
		log.Fatal(err)
	}
	defer record.Close()

	bitcoin := make([]Bitcoin, 0)
	bitcoinData := BitcoinData{}

	//create bitcoin array by looping the query data
	for record.Next() {
		item := Bitcoin{}
		err = record.Scan(&item.Id, &item.CoinType, &item.Price, &item.Timestamp)

		if err != nil {
			return bitcoinData, err
		}

		bitcoin = append(bitcoin, item)
	}

	bitcoinData.Count = FetchBitcoinCount()
	bitcoinData.Data = bitcoin

	err = record.Err()
	if err != nil {
		return bitcoinData, err
	}

	return bitcoinData, err
}

/*
	fetch total number of records from bitcoin_tracker table
*/
func FetchBitcoinCount() int {
	db := Config.DBConnect()
	rows, err := db.Query("SELECT COUNT(*) FROM bitcoin_tracker")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var count int

	for rows.Next() {
		if err := rows.Scan(&count); err != nil {
			log.Fatal(err)
		}
	}
	return count
}
