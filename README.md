# Alliggator

Alliggator provides a string/bson.M convertion, to be used with the mongodb native driver for Golang. There is no need to stablish a pattern to you json string requests, use the consolidate mongodb aggregate syntax and Alliggator will convert your string it to mgo.bson.M!

```
Install: go get https://github.com/jamesmallon/Alliggator.git (better use it with golang Dep package manager)
```

```go
import "github.com/jamesmallon/Alliggator"
```

```go
jsonStr:= `[{"$match":
                {"domain": "carrierexpress.com.br"}
            },
            {"$project": {
                "_id": 0,
                "domain": 1,
                "ipPort": 1}
            },
            {"$sort": {"ipPort": 1}},
            {"$skip": 1},
            {"$limit": 10}]`

c := session.DB("test").C("people")
err := c.Pipe(alliggator.Piperize(jsonStr)).One(&result)
if err != nil {
	fmt.Println(err)
}
```

> Author: Thiago Mallon <thiagomallon@gmail.com>
