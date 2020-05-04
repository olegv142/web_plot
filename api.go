package main

import (
	"fmt"
	"math"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func apiInit(r *mux.Router) {
	api := r.PathPrefix("/api/v1").Subrouter()
	api.HandleFunc("/data/{days}", getData).Methods(http.MethodGet)
}

func getData(w http.ResponseWriter, r *http.Request) {
	pathParams := mux.Vars(r)
	w.Header().Set("Content-Type", "application/json")

	var days int
	var err error
	if val, ok := pathParams["days"]; ok {
		days, err = strconv.Atoi(val)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(`{"message": "need a number of days"}`))
			return
		}
	} else {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message": "need a number of days"}`))
		return
	}

	query := r.URL.Query()
	val := query.Get("interval")
	var interval int
	if interval, err = strconv.Atoi(val); err != nil {
		interval = 1
	}

	w.Write([]byte(`{"data": [`))

	const day_seconds = 3600 * 24
	samples := days * day_seconds / interval
	for i := 0; i < samples; i++ {
		if i > 0 {
			w.Write([]byte(`,`))
		}
		w.Write([]byte(
			fmt.Sprintf(`{"x":%d,"y":%f }`,
				(i-samples)*interval,
				math.Sin(2*math.Pi*float64(i*interval)/3600))))
	}
	w.Write([]byte(`]}`))
}
