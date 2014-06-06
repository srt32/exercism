package gigasecond

import (
  "time"
)

const (
  Gigasecond = 1000000000000000000
)

func AddGigasecond(t time.Time) time.Time {
  return t.Add(Gigasecond)
}
