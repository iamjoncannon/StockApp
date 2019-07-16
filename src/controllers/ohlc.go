// #

package controllers

import (
	"net/http"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
	"time"
	"os"
	"fmt"

	"github.com/gorilla/mux"
	"utils"
)

func FetchOpeningPrice() http.HandlerFunc {

	return func (w http.ResponseWriter, r *http.Request) {

		allParams := mux.Vars(r)

		param := allParams["symbol"]

		url := []string{"https://cloud.iexapis.com/beta/stock/", param, "/quote/open?token=", os.Getenv("IEX_API_KEY")}

		fmt.Println(strings.Join(url, ""))

		openingPriceEndpoint := http.Client{

			Timeout: time.Second * 20,
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
