package main

import (
  "os"
  "time"
  "github.com/PyramidSystemsInc/go/aws"
  "github.com/PyramidSystemsInc/go/aws/cloudfront"
)

func main() {
  if len(os.Args) > 1 {
    waitMinutes(45)
    region := "us-east-2"
    awsSession := aws.CreateAwsSession(region)
    for _, alias := range os.Args[1:] {
      if alias != "--" {
        cloudfront.DeleteDistribution(alias, awsSession)
      }
    }
  }
}

func waitMinutes(minutes time.Duration) {
  time.Sleep(minutes * 60 * 1000 * time.Millisecond)
}
