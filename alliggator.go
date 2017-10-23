package alliggator

import (
	"github.com/johnthegreenobrien/Alliggator/models"
	"bytes"
	"encoding/json"
	"fmt"
	"gopkg.in/mgo.v2/bson"
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

// BuildPipeline
func (all *alliggator) ChargePipeline(result []models.Aggregation) []bson.M {
	query := []bson.M{}

	for _, v := range result {
		val := reflect.ValueOf(v)
		for i := 0; i < val.Type().NumField(); i++ {
			fieldName := val.Type().Field(i).Tag.Get("json")
			fieldValue := val.Field(i)

			if fieldValue.Type().String() == "interface {}" && reflect.ValueOf(fieldValue.Interface()).Kind().String() != "invalid" {
				// log.Println(fieldName,": ", fieldValue.Interface())
				query = append(query, bson.M{fmt.Sprintf("%s%s", "$", fieldName): fieldValue.Interface()})
			} else if fieldValue.Type().String() == "int" && fieldValue.Int() > 0 {
				// log.Println(fieldName,": ", fieldValue.Int())
				query = append(query, bson.M{fmt.Sprintf("%s%s", "$", fieldName): fieldValue.Int()})
			} else if fieldValue.Type().String() == "string" && len(fieldValue.String()) > 0 {
				// log.Println(fieldName, ": ", fieldValue.String())
				query = append(query, bson.M{fmt.Sprintf("%s%s", "$", fieldName): fieldValue.String()})
			}
		}
	}
	return query
}

// FromString
func (all *alliggator) FromString(jsonString string) []bson.M {
	jsonStream := all.TreatDollarSign(jsonString)
	result := make([]models.Aggregation, 0)
	decoder := json.NewDecoder(bytes.NewBufferString(jsonStream))
	err := decoder.Decode(&result)
	if err != nil {
		panic(err)
	}
	return all.ChargePipeline(result)
}
