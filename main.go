package main

import (
  "time"
  "os"
)

func main() {
  FORTY_FIVE_MINUTES = 45 * 60 * 1000 * time.Millisecond
  time.Sleep(FORTY_FIVE_MINUTES)
  os.Mkdir("/home/jeff/Desktop/I-Worked", 0755)
}
