package hello

import (
	"rsc.io/quote/v3"
	"github.com/seyedk/hellogo/api"

)

func Hello() string {

	return quote.HelloV3()
}

func Proverb() string {
	return quote.Concurrency()
}
