package Model

import "time"

type USDVND struct {
	Time     time.Time `json: "time"`
	Exchange uint      `json: "exchange"`
}
