package provider

import (
	. "C-SquaredService/model"
	"encoding/json"
)

type ResponseData struct {
	Id   string
	Data int
}

func GetTest(id string) []byte {
	response := ApiResponse{ResultCode: "200", ResultMessage: ResponseData{Id: id, Data: 123}}

	responseData, _ := json.Marshal(response)

	return responseData
}
