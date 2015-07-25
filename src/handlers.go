package main

import (
  "encoding/json"
  "fmt"
  "net/http"
  "os"
  "io"
  "io/ioutil"
  "time"
  "gopkg.in/mgo.v2"
  "gopkg.in/mgo.v2/bson"
)

func Index(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintln(w, "{\"message\":\"Welcome\"}")
}

func DataCreate(w http.ResponseWriter, r *http.Request) {
  var event Event
  var payload Payload
  body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
  if err != nil {
    panic(err)
  }
  json.Unmarshal(body, &payload)
  json.Unmarshal(body, &event)
  event.Time = time.Now()
  mongo, err := mgo.Dial(os.Getenv("MONGODB_URL"))
  collection := mongo.DB(os.Getenv("MONGODB_DB")).C(fmt.Sprintf("events-%s", payload.Stream))
  collection.Insert(event)
  w.Header().Set("Content-Type", "application/json; charset=UTF-8")
  json.NewEncoder(w).Encode(event)
}

func DataIndex(w http.ResponseWriter, r *http.Request) {
  params := r.URL.Query()
  stream := params.Get("stream")
  mongo, err := mgo.Dial(os.Getenv("MONGODB_URL"))
  if err != nil {
    panic(err)
  }
  collection := mongo.DB(os.Getenv("MONGODB_DB")).C(fmt.Sprintf("events-%s", stream))
  var events []Event
  count, err := collection.Count()
  collection.Find(bson.M{}).Sort("-1").All(&events)
  json.NewEncoder(w).Encode(Stream{Name: stream, Count: count, Data: events})
}
