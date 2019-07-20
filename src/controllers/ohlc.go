package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
	"utils"

	"github.com/gomodule/redigo/redis"
	"github.com/gorilla/mux"
)

func FetchOpeningPrice(c redis.Conn) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		allParams := mux.Vars(r)
		param := allParams["symbol"]

		s, err := redis.String(c.Do("GET", param))

		if err != redis.ErrNil {

			fmt.Println("hit the cache: ", allParams)

			utils.ResponseJSON(w, s)

		} else if err == redis.ErrNil {

			url := []string{"https://cloud.iexapis.com/beta/stock/", param, "/quote/ohlc?token=", os.Getenv("IEX_API_KEY")}

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

			twelveHours := 60 * 60 * 12

			_, err = c.Do("SET", param, body)
			expire, err := c.Do("EXPIRE", param, twelveHours)

			fmt.Println(expire)

			var f interface{}

			err = json.Unmarshal(body, &f)

			fmt.Println("call the api: ", allParams)

			if err != nil {
				fmt.Println(err)
			}

			utils.ResponseJSON(w, f)
		} else {

			fmt.Println(err)
		}
	}
}
