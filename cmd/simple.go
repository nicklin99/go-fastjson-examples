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
	
	rootMap := gofastjsonexamples.NewTreeManager(gofastjsonexamples.JsonPathNodes)
	roots := rootMap.GetTree()
	// if jsonRoot type array replace {} to []
	// 如果jsonRoot 类型是数组，替换 {} 为 []
	jsonRoot := fastjson.MustParse(`{}`) 
	var callback func (int, *gofastjsonexamples.Node[gofastjsonexamples.JsonPathNode]) 
	callback = func (i int, node *gofastjsonexamples.Node[gofastjsonexamples.JsonPathNode])  {
		nValue := new(fastjson.Value)
		if node.Parent == nil {
			nValue = jsonRoot
		} else {
			nValue = node.Parent.Data.FastJsonValue
		}
		v := new(fastjson.Value)
		switch node.Data.Type {
		case "Array":
			v = fastjson.MustParse(`[]`)
			node.Data.FastJsonValue = v
			if node.Parent!= nil && node.Parent.Data.Type == "Array" {
				node.Parent.Data.FastJsonValue.SetArrayItem(0, v)
			}
			nValue.Set(node.Data.GetKey(), v)	
		case "Object":
			v = fastjson.MustParse(`{}`)
			node.Data.FastJsonValue = v
			if node.Parent!= nil && node.Parent.Data.Type == "Array" {
				node.Parent.Data.FastJsonValue.SetArrayItem(0, v)
			}
			nValue.Set(node.Data.GetKey(), v)
		default:
			v, _ = fastjson.Parse(node.Data.Value)
			nValue.Set(node.Data.GetKey(), v)
		}
		
		if len(node.Children) > 0 {
			gofastjsonexamples.MapNode(node.Children, callback)
		}
	}
	gofastjsonexamples.MapNode(roots, callback)
	log.Printf("%s", jsonRoot)
}