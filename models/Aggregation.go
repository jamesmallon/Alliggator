package models

type Aggregation struct {
	Match   interface{} `bson:"match" json:"match"`
	Project interface{} `bson:"project" json:"project"`
	Sort    interface{} `bson:"sort" json:"sort"`
	Limit   int         `bson:"limit" json:"limit"`
	Skip    int         `bson:"skip" json:"skip"`
	Count   string      `bson:"count" json:"count"`
}
