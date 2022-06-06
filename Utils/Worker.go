package Utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"bitcoin-checker/Models"

	"github.com/jasonlvhit/gocron"
	_ "github.com/mattn/go-sqlite3"
)

type Response struct {
	Bitcoin struct {
		Usd int `json:"usd"`
	}
}

func fetchBitcoinRate() {

	base, err := url.Parse("https://api.coingecko.com/api/v3/simple/price")
	if err != nil {
		return
	}

	//add query params to the coingecko api
	params := url.Values{}
	params.Add("ids", "bitcoin")
	params.Add("vs_currencies", "usd")
	base.RawQuery = params.Encode()

	client := &http.Client{}
	req, err := http.NewRequest("GET", base.String(), nil)
	if err != nil {
		fmt.Print(err.Error())
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		fmt.Print(err.Error())
	}
	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Print(err.Error())
		return
	}

	var BitcoinInfo Response
	json.Unmarshal(bodyBytes, &BitcoinInfo)
	prepareForInsert(BitcoinInfo)
	prepareForNotification(BitcoinInfo)
}

func StartPolling() {
	// cron will call wth worker function for every 30 seconds
	gocron.Every(30).Second().Do(fetchBitcoinRate)
	<-gocron.Start()
}

func prepareForInsert(BitcoinInfo Response) {
	// insert data to bitcoin_tracker table
	Models.AddBitcoinInfo("BITCOIN", BitcoinInfo.Bitcoin.Usd, getCurrentTime())
}

/*
	Notify the user when price crosses the below min or above the max value
*/
func prepareForNotification(BitcoinInfo Response) {
	message := ""
	if BitcoinInfo.Bitcoin.Usd > StringToInt(getEnvVariables("MAX_PRICE")) {
		message = "Bitcoin price is greater than " + getEnvVariables("MAX_PRICE")
	} else if BitcoinInfo.Bitcoin.Usd < StringToInt(getEnvVariables("MIN_PRICE")) {
		message = "Bitcoin price is less than " + getEnvVariables("MIN_PRICE")
	} else {
		return
	}

	email := getEnvVariables("TO_MAIL")
	sendMail(message, email)
}
