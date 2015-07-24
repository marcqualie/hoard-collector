package main

import (
  "time"
)

type Event struct {
  Stream   string             `json:"stream"`
  Data     map[string]string  `json:"data"`
  Time     time.Time          `json:"time"`
}
