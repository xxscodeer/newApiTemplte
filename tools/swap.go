package tools


import jsoniter "github.com/json-iterator/go"

//通过json tag 进行结构体赋值
func SwapTo(request, category interface{}) (err error)  {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	dataByte, err := json.Marshal(request)
	if err != nil {
		return
	}
	err = json.Unmarshal(dataByte, category)
	return
}
