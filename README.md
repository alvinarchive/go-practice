How to run app

```
go run ./cmd/api 
```

or you can set port or/and set environment 
```
go run ./cmd/api -port="3001" -env="development"
```


### current api

/v1/healthcheck

will return health and selected port and environment
```
{
	"environment": "development",
	"status": "available",
	"version": "1.0.0"
}
```

/v1/movies/:id 

id must be integer > 0 

all return are still hard coded
```
{
	"movie": {
		"id": 32,
		"title": "Batman",
		"runtime": "102 mins",
		"genres": [
			"vengeance",
			"suspense",
			"superhero"
		],
		"version": 1
	}
}
```
