package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type Model struct {
	Campo1  float64 `json:"campo1"`
	Campo2  float64 `json:"campo2"`
	Campo3  float64 `json:"campo3"`
	Campo4  float64 `json:"campo4"`
	Campo5  float64 `json:"campo5"`
	Campo6  float64 `json:"campo6"`
	Campo7  float64 `json:"campo7"`
	Campo8  float64 `json:"campo8"`
	Campo9  float64 `json:"campo9"`
	Campo10 float64 `json:"campo10"`
	Campo11 float64 `json:"campo11"`
	Campo12 float64 `json:"campo12"`
	Campo13 float64 `json:"campo13"`
	Campo14 float64 `json:"campo14"`
	Campo15 float64 `json:"campo15"`
	Campo16 float64 `json:"campo16"`
	Campo17 float64 `json:"campo17"`
	Campo18 float64 `json:"campo18"`
	Campo19 float64 `json:"campo19"`
	Campo20 float64 `json:"campo20"`
	Campo21 float64 `json:"campo21"`
	Campo22 float64 `json:"campo22"`
	Campo23 float64 `json:"campo23"`
	Campo24 float64 `json:"campo24"`
	Campo25 float64 `json:"campo25"`
	Campo26 float64 `json:"campo26"`
	Campo27 float64 `json:"campo27"`
	Campo28 float64 `json:"campo28"`
	Campo29 float64 `json:"campo29"`
	Campo30 float64 `json:"campo30"`
	Campo31 float64 `json:"campo31"`
	Campo32 float64 `json:"campo32"`
	Campo33 float64 `json:"campo33"`
	Campo34 float64 `json:"campo34"`
	Campo35 float64 `json:"campo35"`
	Campo36 float64 `json:"campo36"`
	Campo37 float64 `json:"campo37"`
	Campo38 float64 `json:"campo38"`
	Campo39 float64 `json:"campo39"`
	Campo40 float64 `json:"campo40"`
	Campo41 float64 `json:"campo41"`
	Campo42 float64 `json:"campo42"`
	Campo43 float64 `json:"campo43"`
	Campo44 float64 `json:"campo44"`
	Campo45 float64 `json:"campo45"`
	Campo46 float64 `json:"campo46"`
	Campo47 float64 `json:"campo47"`
	Campo48 float64 `json:"campo48"`
	Campo49 float64 `json:"campo49"`
	Campo50 float64 `json:"campo50"`
	Campo51 float64 `json:"campo51"`
	Campo52 float64 `json:"campo52"`
	Campo53 float64 `json:"campo53"`
	Campo54 float64 `json:"campo54"`
	Campo55 float64 `json:"campo55"`
	Campo56 float64 `json:"campo56"`
	Campo57 float64 `json:"campo57"`
	Campo58 float64 `json:"campo58"`
	Campo59 float64 `json:"campo59"`
	Campo60 float64 `json:"campo60"`
}

func main() {
	t := time.Now()
	//GetAll
	r, err := http.Get("http://0.0.0.0:8080/getall")
	if err != nil {
		panic(err)
	}
	body, err := io.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	var model []Model
	err = json.Unmarshal(body, &model)
	if err != nil {
		panic(err)
	}
	//GetOne
	/* 	for i := 0; i < 20000; i++ {
		r, err := http.Get("http://0.0.0.0:8080/getone")
		if err != nil {
			panic(err)
		}
		body, err := io.ReadAll(r.Body)
		if err != nil {
			panic(err)
		}
		var model Model
		err = json.Unmarshal(body, &model)
		if err != nil {
			panic(err)
		}
	} */
	fmt.Println("Elapsed from client: ", time.Since(t).Seconds())
}
