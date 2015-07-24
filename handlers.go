package main

import (
  "encoding/json"
  "fmt"
  "net/http"
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
  mongo, err := mgo.Dial("mongodb://127.0.0.1:27017/hoard-development")
  collection := mongo.DB("hoard-development").C(fmt.Sprintf("events-%s", payload.Stream))
  collection.Insert(event)
  w.Header().Set("Content-Type", "application/json; charset=UTF-8")
  json.NewEncoder(w).Encode(event)
}

func DataIndex(w http.ResponseWriter, r *http.Request) {
  params := r.URL.Query()
  stream := params.Get("stream")
  mongo, err := mgo.Dial("mongodb://127.0.0.1:27017/hoard-development")
  if err != nil {
    panic(err)
  }
  collection := mongo.DB("hoard-development").C(fmt.Sprintf("events-%s", stream))
  // event := Event{Stream: params.Get("stream")}
  var events []Event
  count, err := collection.Count()
  collection.Find(bson.M{}).Sort("-1").All(&events)
  // json.Marshel(results, &events)
  // events[0] = event
  json.NewEncoder(w).Encode(Stream{Name: stream, Count: count, Data: events})
  // mongo, err := mgo.Dial("mongodb://127.0.0.1:27017/hoard-development")
  // collection := mongo.DB("hoard-development").C(fmt.Sprintf("events-%s", event.Stream))

}
