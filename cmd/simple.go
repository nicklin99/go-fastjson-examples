package main

import (
	"fmt"
	"log"

	gofastjsonexamples "github.com/nicklin99/go-fastjson-examples"
	"github.com/valyala/fastjson"
)

func main() {
	s := []byte(`{"foo": [123, "bar"]}`)
	fmt.Printf("foo.0=%d\n", fastjson.GetInt(s, "foo", "0"))

	jsonRoot := fastjson.MustParse(`{}`)
	gofastjsonexamples.JsonPathNodeToJson(gofastjsonexamples.JsonPathNodes, jsonRoot)
	log.Printf("%s", jsonRoot)
}
