package main

import (
  "time"
)

type Payload struct {
  Stream   string                   `json:"stream"`
}

type Event struct {
  Name     string                   `json:"name" bson:"n"`
  Data     interface{}              `json:"data" bson:"d"`
  Time     time.Time                `json:"time" bson:"t"`
}

type Stream struct {
  Name     string                   `json:"stream"`
  Count    int                      `json:"count"`
  Data     []Event                  `json:"data"`
}
