package rest

import (
	"context"
	"log"
	"testing"
)

/*
	GET请求
*/
func TestRESTAPIGet(t *testing.T) {

	rest := NewRESTAPI("https://www.okex.win", GET, "/api/v5/account/balance", nil)
	rest.SetSimulate(true).SetAPIKey("xxxx", "xxxx", "xxxx")
	rest.SetUserId("xxxxx")
	response, err := rest.Run(context.Background())
	if err != nil {
		log.Println(err)
		return
	}

	log.Println("Response:")
	log.Println("\thttp code: ", response.Code)
	log.Println("\ttotal time: ", response.TotalUsedTime)
	log.Println("\trequest time: ", response.ReqUsedTime)
	log.Println("\tresponse: ", response.Body)
	log.Println("\terrCode: ", response.V5Response.Code)
	log.Println("\terrMsg: ", response.V5Response.Msg)
	log.Println("\tdata: ", response.V5Response.Data)

	// 请求的另一种方式
	apikey := APIKeyInfo{
		ApiKey:     "xxxxx",
		SecKey:     "xxxxx",
		PassPhrase: "xxx",
	}

	cli := NewRESTClient("https://www.okex.win", &apikey, true)
	rsp, err := cli.Get(context.Background(), "/api/v5/account/balance", nil)
	if err != nil {
		return
	}

	log.Println("Response:")
	log.Println("\thttp code: ", rsp.Code)
	log.Println("\ttotal time: ", rsp.TotalUsedTime)
	log.Println("\trequest time: ", rsp.ReqUsedTime)
	log.Println("\tresponse: ", rsp.Body)
	log.Println("\terrCode: ", rsp.V5Response.Code)
	log.Println("\terrMsg: ", rsp.V5Response.Msg)
	log.Println("\tdata: ", rsp.V5Response.Data)
}

/*
	POST请求
*/
func TestRESTAPIPost(t *testing.T) {
	param := make(map[string]interface{})
	param["greeksType"] = "PA"

	rest := NewRESTAPI("https://www.okex.win", POST, "/api/v5/account/set-greeks", &param)
	rest.SetSimulate(true).SetAPIKey("xxxx", "xxxx", "xxxx")
	response, err := rest.Run(context.Background())
	if err != nil {
		log.Println(err)
		return
	}

	log.Println("Response:")
	log.Println("\thttp code: ", response.Code)
	log.Println("\ttotal time: ", response.TotalUsedTime)
	log.Println("\trequest time: ", response.ReqUsedTime)
	log.Println("\tresponse: ", response.Body)
	log.Println("\terrCode: ", response.V5Response.Code)
	log.Println("\terrMsg: ", response.V5Response.Msg)
	log.Println("\tdata: ", response.V5Response.Data)

	// 请求的另一种方式
	apikey := APIKeyInfo{
		ApiKey:     "xxxx",
		SecKey:     "xxxxx",
		PassPhrase: "xxxx",
	}

	cli := NewRESTClient("https://www.okex.win", &apikey, true)
	rsp, err := cli.Post(context.Background(), "/api/v5/account/set-greeks", &param)
	if err != nil {
		return
	}

	log.Println("Response:")
	log.Println("\thttp code: ", rsp.Code)
	log.Println("\ttotal time: ", rsp.TotalUsedTime)
	log.Println("\trequest time: ", rsp.ReqUsedTime)
	log.Println("\tresponse: ", rsp.Body)
	log.Println("\terrCode: ", rsp.V5Response.Code)
	log.Println("\terrMsg: ", rsp.V5Response.Msg)
	log.Println("\tdata: ", rsp.V5Response.Data)
}
