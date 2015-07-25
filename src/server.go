package main

import (
  "log"
  "fmt"
  "net/http"
  "os"
  "gopkg.in/mgo.v2"
)

func main() {
  router := Router()
  mongo, err := mgo.Dial(os.Getenv("MONGODB_URL"))
  if err != nil {
    panic(err)
  }
  mongoDB := mongo.DB(os.Getenv("MONGODB_DB"))
  buildInfo, err := mongo.BuildInfo()
  collectionNames, err := mongoDB.CollectionNames()
  fmt.Println("Connected to MongoDB: %s", os.Getenv("MONGODB_URL"))
  fmt.Println("  ", buildInfo)
  fmt.Println("  ", collectionNames)
  fmt.Println("  ")
  fmt.Println("Listening for connections on port", os.Getenv("PORT"))
  fmt.Println("  ")
  log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("PORT")), router))
}
