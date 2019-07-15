package controllers

import (

	"net/http"
	// "fmt"
	// "encoding/json"

	"io/ioutil"
	"log"
	"strconv"
	"strings"
	"time"

	// "github.com/davecgh/go-spew/spew"	
	"github.com/gorilla/mux"
	"utils"
)

func FetchOpeningPrice() http.HandlerFunc {

	return func (w http.ResponseWriter, r *http.Request) {

		allParams := mux.Vars(r)

		param := allParams["symbol"]

		url := []string{"https://cloud.iexapis.com/beta/stock/", param, "/quote/open?token=sk_5c1a1ec78f534b179da588b787245fe6"}

		openingPriceEndpoint := http.Client{

			Timeout: time.Second * 20, // Maximum of 2 secs
		}

		req, err := http.NewRequest(http.MethodGet, strings.Join(url, ""), nil)

		if err != nil {
			log.Fatal(err)
		}

		req.Header.Set("User-Agent", "TTP-FP")

		res, getErr := openingPriceEndpoint.Do(req)

		if getErr != nil {
			log.Fatal(getErr)
		}

		body, readErr := ioutil.ReadAll(res.Body)

		if readErr != nil {
			log.Fatal(readErr)
		}

		formatted, _ := strconv.ParseFloat(string(body), 32)

		returnJason := make(map[string]float64) 

		returnJason[param] = formatted

		utils.ResponseJSON(w, returnJason) 
	}
}	
