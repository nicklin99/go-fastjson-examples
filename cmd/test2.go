package main

import (
	"log"

	gofastjsonexamples "github.com/nicklin99/go-fastjson-examples"
	"github.com/valyala/fastjson"
)

func main() {
	nodes := make([]gofastjsonexamples.JsonPathNode, 0)
	nodes = append(nodes, gofastjsonexamples.JsonPathNode{
		Name:  "taskList",
		Type:  "Array",
		Value: "",
	})
	nodes = append(nodes, gofastjsonexamples.JsonPathNode{
		Name:  "taskList.taskId",
		Type:  "String",
		Value: "fca2fd8daba74f00b31cbcbd306c9289",
	})
	// if jsonRoot type array replace {} to []
	// 如果jsonRoot 类型是数组，替换 {} 为 []
	jsonRoot := fastjson.MustParse(`{}`)
	gofastjsonexamples.JsonPathNodeToJson(nodes, jsonRoot)
	log.Printf("%s", jsonRoot)
}
