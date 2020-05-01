package httpclient

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"testing"
)

func TestClientFunction(b *testing.T){
	client := NewClient(SetTimeout(3))
	MonitorURL := "http://monitor-api.server.com/get_day_data"

	var reqDat = url.Values{}
	reqDat.Set("idtype", "2")
	reqDat.Set("begtime", "1578499200")
	reqDat.Set("endtime", "1578621125")
	reqDat.Set("query", `[{"id":6354, "attrid":399536},{"id":6354, "attrid":2773697}]`)
	headers := http.Header{}
	headers.Set("Content-Type", "application/x-www-form-urlencoded")
	res ,err := client.POST(MonitorURL, headers, strings.NewReader(reqDat.Encode()))
	if err != nil{
		fmt.Println(err.Error())
		return
	}
	fmt.Println(res ,err)
	defer res.Body.Close()
}
