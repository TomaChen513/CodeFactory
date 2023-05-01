package lib

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
)

//将body中=号格式的字符串转为map
func ConvertToMap(str string) map[string]string {
	var resultMap = make(map[string]string)
	values := strings.Split(str, "&")
	for _, value := range values {
		vs := strings.Split(value, "=")
		resultMap[vs[0]] = vs[1]
	}
	return resultMap
}

func ParseResponse(response *http.Response) (map[string]interface{}, error){
	var result map[string]interface{}
    body,err := ioutil.ReadAll(response.Body)
    if err == nil {
        err = json.Unmarshal(body, &result)
    }
 
    return result,err
}
