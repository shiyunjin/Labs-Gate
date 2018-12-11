package device

import (
	"github.com/gin-gonic/gin"
	"github.com/shiyunjin/Labs-Gate/system/e"
	"github.com/shiyunjin/Labs-Gate/system/service/model"
)

type dataType struct {
	State 	string
	Down	int64	`json:"下载流量"`
	Up	 	int64	`json:"上传流量"`
}

func Bandwidth(c *gin.Context){
	code := c.Param("code")

	Channel := c.MustGet("Channel").(serviceModel.Channel)
	resultch := make(chan serviceModel.BandwidthCall)
	defer close(resultch)
	Channel.Bandwidthch <- serviceModel.BandwidthMsg{Type: 1, Device: code, Callback: resultch}
	result := <- resultch
	if result.Err != nil {
		c.JSON(e.SUCCESS, gin.H{
			"status" : e.ERROR,
			"statusText" : result.Err.Error(),
		})
	} else {
		data := []dataType{}
		for k, v := range result.Data.(map[string][]int64) {
			data = append(data, dataType{
				State:k,
				Down:v[0],
				Up:v[1],
			})
		}
		c.JSON(e.SUCCESS, gin.H{
			"status":     e.SUCCESS,
			"statusText": e.GetMsg(e.SUCCESS),
			"data": data,
		})
	}
}
