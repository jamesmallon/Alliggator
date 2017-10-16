package alliggator

import (
	"Alliggator/models"
	"bytes"
	"encoding/json"
	// mgo "gopkg.in/mgo.v2"
	// "gopkg.in/mgo.v2/bson"
	"log"
	"regexp"
)

// alliggator
type alliggator struct {
	// Data bson.M `json:"data"`
}

// New
func New() *alliggator {
	return &alliggator{}
}

func (all *alliggator) TreatDollarSign(jsonString string) string {
	var reg = regexp.MustCompile(`\$`)
	jsonStream := reg.ReplaceAllString(string(jsonString), ``)
	// log.Println(jsonStream)
	return jsonStream
}

// BuildPipeline
func (all *alliggator) ChargePipeline(result []models.Aggregation) {
	log.Println(result)
	log.Println(result[0].Match)
	// total := 0
	for _, v := range result {
		// for key, val := range v {
		// 	log.Println(key, ":", val)
		// }
		log.Println(v)
	}
}

// FromString
func (all *alliggator) FromString(jsonString string) {
	jsonStream := all.TreatDollarSign(jsonString)
	result := make([]models.Aggregation, 0)
	decoder := json.NewDecoder(bytes.NewBufferString(jsonStream))
	err := decoder.Decode(&result)
	if err != nil {
		panic(err)
	}
	all.ChargePipeline(result)
}
