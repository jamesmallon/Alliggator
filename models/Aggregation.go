package models

type Aggregation struct {
  Match   interface{} `jaon:"match"`
  Project interface{} `json:"project`
  Sort    interface{} `json:"sort"`
  Limit   int         `json:"limit"`
  Skip    int         `json:"skip"`
}
