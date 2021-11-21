package util

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

func UnmarshalJson(context *gin.Context) map[string]interface{} {
	var jsonData map[string]interface{}

	data, err := ioutil.ReadAll(context.Request.Body)
	if err != nil {
		context.JSON(http.StatusInternalServerError, err)
		return nil
	}

	err = json.Unmarshal(data, &jsonData)
	if err != nil {
		context.JSON(http.StatusInternalServerError, err)
		return nil
	}

	return jsonData
}
