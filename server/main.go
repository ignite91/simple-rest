package main

import (
	"encoding/json"
	"fmt"
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

var model = Model{
	Campo1:  99999999.22,
	Campo2:  99999999.22,
	Campo3:  99999999.22,
	Campo4:  99999999.22,
	Campo5:  99999999.22,
	Campo6:  99999999.22,
	Campo7:  99999999.22,
	Campo8:  99999999.22,
	Campo9:  99999999.22,
	Campo10: 99999999.22,
	Campo11: 99999999.22,
	Campo12: 99999999.22,
	Campo13: 99999999.22,
	Campo14: 99999999.22,
	Campo15: 99999999.22,
	Campo16: 99999999.22,
	Campo17: 99999999.22,
	Campo18: 99999999.22,
	Campo19: 99999999.22,
	Campo20: 99999999.22,
	Campo21: 99999999.22,
	Campo22: 99999999.22,
	Campo23: 99999999.22,
	Campo24: 99999999.22,
	Campo25: 99999999.22,
	Campo26: 99999999.22,
	Campo27: 99999999.22,
	Campo28: 99999999.22,
	Campo29: 99999999.22,
	Campo30: 99999999.22,
	Campo31: 99999999.22,
	Campo32: 99999999.22,
	Campo33: 99999999.22,
	Campo34: 99999999.22,
	Campo35: 99999999.22,
	Campo36: 99999999.22,
	Campo37: 99999999.22,
	Campo38: 99999999.22,
	Campo39: 99999999.22,
	Campo40: 99999999.22,
	Campo41: 99999999.22,
	Campo42: 99999999.22,
	Campo43: 99999999.22,
	Campo44: 99999999.22,
	Campo45: 99999999.22,
	Campo46: 99999999.22,
	Campo47: 99999999.22,
	Campo48: 99999999.22,
	Campo49: 99999999.22,
	Campo50: 99999999.22,
	Campo51: 99999999.22,
	Campo52: 99999999.22,
	Campo53: 99999999.22,
	Campo54: 99999999.22,
	Campo55: 99999999.22,
	Campo56: 99999999.22,
	Campo57: 99999999.22,
	Campo58: 99999999.22,
	Campo59: 99999999.22,
	Campo60: 99999999.22,
}

func main() {
	fmt.Println("Server Starting at :8080 . . .")
	http.HandleFunc("/getall", getAll)
	http.HandleFunc("/getone", getOne)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}

func getAll(w http.ResponseWriter, r *http.Request) {
	t := time.Now()
	var models []Model
	for i := 0; i < 20000; i++ {
		models = append(models, model)
	}
	data, err := json.Marshal(models)
	if err != nil {
		panic(err)
	}
	w.Write(data)
	fmt.Println("GetAll from server Elapsed: ", time.Since(t).Seconds())
}

func getOne(w http.ResponseWriter, r *http.Request) {
	//	t := time.Now()
	data, err := json.Marshal(model)
	if err != nil {
		panic(err)
	}
	w.Write(data)
	//fmt.Println("GetOne from server Elapsed: ", time.Since(t).Seconds())
}
