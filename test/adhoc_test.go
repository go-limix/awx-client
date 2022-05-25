package test

import (
	"context"
	"testing"

	awx "github.com/go-limix/awx-client"
)

func TestAdhoc(t *testing.T) {
	token := "7WwCmUaTK0PfmSfq6p0ETnQ5SOHJlF"
	var conf *awx.Configuration = awx.NewConfiguration()
	conf.BasePath = "http://localhost:8013"
	conf.AddDefaultHeader("Authorization", "Bearer "+token)
	var client *awx.APIClient = awx.NewAPIClient(conf)

	// 创建 ad-hoc
	var adHoc awx.AdHocCommandList
	adHoc.ModuleName = "command"
	adHoc.ModuleArgs = "pwd"
	adHoc.Credential = 5
	adHoc.Inventory = 1

	// 提交 ad-hoc
	adHoc, response, error := client.ApiApi.ApiAdHocCommandsCreate(context.Background(), adHoc, "v2")
	if error == nil {
		println(adHoc.Id)
	} else {
		if response != nil {
			println(response.Header)
		}
	}
}
