package gofastjsonexamples

import (
	"strings"

	"github.com/valyala/fastjson"
)

type JsonPathNode struct {
	Name  string `json:"name"`
	Type  string `json:"type"`
	Value string `json:"value"`
	FastJsonValue *fastjson.Value `json:"_"`
}

func (n JsonPathNode) GetParentID() string{
	paths := strings.Split(n.Name, ".")
	if len(paths) - 1 < 0 {
		return "0"
	}
	return strings.Join(paths[:len(paths)-1], ".")
}

func (n JsonPathNode) GetID() string{
	return n.Name
}

func (n JsonPathNode) GetKey() string{
	paths := strings.Split(n.Name, ".")
	return paths[len(paths) -1]
}

var JsonPathMap = []map[string]interface{}{
	{"name": "taskList", "type": "Array", "value": ""},
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
	{"name": "taskList.taskInfo.streamInfo", "type": "String", "value": ""},
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

func init()  {
	for _, node := range JsonPathMap {
		JsonPathNodes = append(JsonPathNodes, JsonPathNode{
			Name: node["name"].(string),
			Type: node["type"].(string),
			Value: node["value"].(string),
		})
	}
}