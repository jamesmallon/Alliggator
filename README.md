# Alliggator

Do you want to use mongodb aggregate syntax on your rest requests? Now you can do it! Alliggator provides a string/bson.M convertion. There is no need to stablish a pattern to you json string requests, use the mongodb aggregate syntax and Alliggator will convert it to bson!

```
go get https://github.com/johnthegreenobrien/Alliggator.git
```

```go
import "github.com/johnthegreenobrien/Alliggator"
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
