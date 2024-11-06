package xjson

import jsoniter "github.com/json-iterator/go"

var json = jsoniter.ConfigCompatibleWithStandardLibrary

func Marshal(v interface{}) (string, error) {
	return json.MarshalToString(v)
}
func MarshalByte(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}
func MustMarshal(v interface{}) string {
	d, err := Marshal(v)
	if err != nil {
		return ""
	}
	return d
}
func UnmarshalByString(str string, v interface{}) error {
	return json.UnmarshalFromString(str, v)
}
func Unmarshal(bytes []byte, v interface{}) error {
	return json.Unmarshal(bytes, v)
}
func Get(str string, path ...interface{}) jsoniter.Any {
	return json.Get([]byte(str), path...)
}

func T() jsoniter.API {
	return json
}
