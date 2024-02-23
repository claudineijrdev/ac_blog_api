package main

import (
	"ac_blog_api/infra/db"
	"ac_blog_api/internal/api"
	"net/http"
)

func main() {
	mongodb := db.NewDB("mongodb://root:root@localhost:27017")
	mongoRepository := db.NewMongoRepository(mongodb)
	router := api.Router(mongoRepository)
	err := http.ListenAndServe(":3000", router)
	if err != nil {
		panic(err)
	}
}
