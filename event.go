package main

import (
  "time"
)

type Event struct {
  Name     string                   `json:"name" bson:"n"`
  Stream   string                   `json:"stream" bson:",omitempty"`
  Data     interface{}              `json:"data" bson:"d"`
  Time     time.Time                `json:"time" bson:"t"`
}
