package cache

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gomodule/redigo/redis"
	"github.com/gorilla/mux"
)

type Cache struct{}

func (C Cache) SetCache(c redis.Conn) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		allParams := mux.Vars(r)

		param := allParams["key"]

		fmt.Println(param)

		_, err := c.Do("SET", param, "DigereeDOO")

		if err != nil {
			fmt.Println(err)
		}
	}
}

func (C Cache) GetCache(c redis.Conn) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		allParams := mux.Vars(r)

		param := allParams["key"]

		s, err := redis.String(c.Do("GET", param))

		returnJSON := make(map[string]string)

		if err == redis.ErrNil {

			fmt.Printf("%s does not exist\n", param)

		} else if err != nil {

			fmt.Println(err)
		} else {

			returnJSON[param] = s
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(returnJSON)
		}
	}
}
