package gofastjsonexamples

import (
	"fmt"
	"strings"

	"github.com/valyala/fastjson"
)

type JsonPathNode struct {
	Name          string            `json:"name"`
	Type          string            `json:"type"`
	Value         string            `json:"value"`
	FastJsonValue *fastjson.Value   `json:"jsonValue"`
	Items         []*fastjson.Value `json:"items"`
}

func (n JsonPathNode) GetParentID() string {
	paths := strings.Split(n.Name, ".")
	if len(paths)-1 < 0 {
		return "0"
	}
	return strings.Join(paths[:len(paths)-1], ".")
}

func (n JsonPathNode) GetID() string {
	return n.Name
}

func (n JsonPathNode) GetKey() string {
	paths := strings.Split(n.Name, ".")
	return paths[len(paths)-1]
}

var JsonPathMap = []map[string]interface{}{
	{"name": "taskList", "type": "Array", "value": ""},
	{"name": "taskList.taskId", "type": "String", "value": "fca2fd8daba74f00b31cbcbd306c9289"},
	{"name": "taskList.taskInfo", "type": "Object", "value": ""},
	{"name": "taskList.taskInfo.priority", "type": "String", "value": ""},
	{"name": "taskList.taskInfo.destination", "type": "String", "value": ""},
	{"name": "taskList.taskInfo.taskType", "type": "Number", "value": "123"},
	{"name": "taskList.taskInfo.imageInfos", "type": "Array", "value": ""},
	{"name": "taskList.taskInfo.imageInfos.id", "type": "Number", "value": ""},
	{"name": "taskList.taskInfo.imageInfos.imageType", "type": "Number", "value": "jpg"},
	{"name": "taskList.taskInfo.imageInfos.data", "type": "String", "value": "http://10.66.88.1/xxxx.jpg"},
	{"name": "taskList.taskInfo.imageInfos.dataType", "type": "Number", "value": "1"},
	{"name": "taskList.taskInfo.imageInfos.ruleInfo", "type": "String", "value": "xxx"},
	{"name": "taskList.taskInfo.imageInfos.descriptionData", "type": "String", "value": "xxxx"},
	{"name": "taskList.taskInfo.streamInfo", "type": "Object", "value": ""},
	{"name": "taskList.taskInfo.streamInfo.streamType", "type": "Number", "value": ""},
	{"name": "taskList.taskInfo.streamInfo.url", "type": "String", "value": ""},
	{"name": "taskList.taskInfo.streamInfo.username", "type": "String", "value": ""},
	{"name": "taskList.taskInfo.streamInfo.password", "type": "String", "value": ""},
	{"name": "taskList.taskInfo.streamInfo.begTime", "type": "String", "value": ""},
	{"name": "taskList.taskInfo.streamInfo.endTime", "type": "String", "value": ""},
	{"name": "taskList.taskInfo.timeZone", "type": "String", "value": "28800"},
	{"name": "taskList.taskInfo.valid", "type": "Boolean", "value": "true"},
	{"name": "taskList.taskInfo.dstDetail", "type": "String", "value": ""},
	{"name": "taskList.taskInfo.ruleInfo", "type": "String", "value": ""},
	{"name": "taskList.taskInfo.descriptionData", "type": "String", "value": ""},
}

var JsonPathNodes []JsonPathNode

func init() {
	for _, node := range JsonPathMap {
		JsonPathNodes = append(JsonPathNodes, JsonPathNode{
			Name:  node["name"].(string),
			Type:  node["type"].(string),
			Value: node["value"].(string),
		})
	}
}

func JsonPathNodeToJson(nodes []JsonPathNode, jsonRoot *fastjson.Value) {
	rootMap := NewTreeManager(nodes)
	roots := rootMap.GetTree()
	var callback func(int, *Node[JsonPathNode])
	callback = func(i int, node *Node[JsonPathNode]) {
		nValue := new(fastjson.Value)
		if node.Parent == nil {
			nValue = jsonRoot
		} else {
			nValue = node.Parent.Data.FastJsonValue
			if node.Parent.Data.Type == "Array" {
				items, _ := nValue.Array()
				if len(items) == 0 {
					// 默认
					subV := fastjson.MustParse(`{}`)
					nValue.SetArrayItem(0, subV)
				}
				items, _ = nValue.Array()
				nValue = items[0]
			}
		}

		v := new(fastjson.Value)
		key := node.Data.GetKey()
		switch node.Data.Type {
		case "Array":
			v = fastjson.MustParse(`[]`)
			node.Data.FastJsonValue = v
			nValue.Set(key, v)
		case "Object":
			v = fastjson.MustParse(`{}`)
			node.Data.FastJsonValue = v
			nValue.Set(key, v)
		case "String":
			v = fastjson.MustParse(fmt.Sprintf(`"%s"`, node.Data.Value))
			nValue.Set(key, v)
		default:
			v, _ = fastjson.Parse(node.Data.Value)
			nValue.Set(key, v)
		}

		if len(node.Children) > 0 {
			MapNode(node.Children, callback)
		}
	}
	MapNode(roots, callback)
}
