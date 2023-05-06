package jxon

import "encoding/json"

//Must mean panic when err

func MustMarshal(obj interface{}) []byte {
	content, err := json.Marshal(obj)
	panicIfErr(err)
	return content
}

func MustMarshalIndent(obj interface{}) []byte {
	content, err := json.MarshalIndent(obj, "", "\t")
	panicIfErr(err)
	return content
}

func MustMarshalString(obj interface{}) string {
	return string(MustMarshal(obj))
}

func MustMarshalIndentString(obj interface{}) string {
	return string(MustMarshalIndent(obj))
}

func MustUnmarshal(content []byte, obj interface{}) {
	err := json.Unmarshal(content, obj)
	panicIfErr(err)
}

func MustUnmarshalString(str string, obj interface{}) {
	MustUnmarshal([]byte(str), obj)
}
