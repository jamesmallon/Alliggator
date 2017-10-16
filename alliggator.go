package alliggator

import (
	"Alliggator/models"
	"bytes"
	"encoding/json"
	// mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
	"reflect"
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

// TreatDollarSign
func (all *alliggator) TreatDollarSign(jsonString string) string {
	var reg = regexp.MustCompile(`\$`)
	jsonStream := reg.ReplaceAllString(string(jsonString), ``)
	// log.Println(jsonStream)
	return jsonStream
}

// GetJsonField
func (all *alliggator) GetJsonField(agg models.Aggregation) {
	val := reflect.ValueOf(agg)
	for i := 0; i < val.Type().NumField(); i++ {
		log.Println(val.Type().Field(i).Tag.Get("json"))
	}
}

// GetBsonField
func (all *alliggator) GetBsonField(agg models.Aggregation) {
	val := reflect.ValueOf(agg)
	for i := 0; i < val.Type().NumField(); i++ {
		log.Println(val.Type().Field(i).Tag.Get("bson"))
		log.Println(val.Type().Field(i))
	}
}

// CreateBsonObj
func (all *alliggator) CreateBsonObj() []bson.M {
	query := []bson.M{}
	query = append(query, bson.M{"$match": bson.M{"domain": "carrierexpress.com.br", "ipPort": "55.131.31.42:37020"}})
	query = append(query, bson.M{"$project": bson.M{"_id": 1, "domain": 1, "ipPort": 1, "available": 1}})
	return query
}

// BuildPipeline
func (all *alliggator) ChargePipeline(result []models.Aggregation) []bson.M {
	query := []bson.M{}
	for _, v := range result {
		// all.GetJsonField(v)
		// all.GetBsonField(v)
		val := reflect.ValueOf(v)
		for i := 0; i < val.Type().NumField(); i++ {
			log.Println(val.Type().Field(i).Tag.Get("bson"))
			// log.Println(val.Type().Field(i))
		}
	}
	return query
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
