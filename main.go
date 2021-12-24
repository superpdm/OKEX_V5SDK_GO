package main

import (
	"context"
	. "github.com/superpdm/OKEX_V5SDK_GO/rest"
	. "github.com/superpdm/OKEX_V5SDK_GO/ws"
	"log"
	"time"
)

/*
	rest API请求
	更多示例请查看 rest/rest_test.go
*/
func REST() {
	// 设置您的APIKey
	apikey := APIKeyInfo{
		ApiKey:     "xxxx",
		SecKey:     "xxxx",
		PassPhrase: "xxxx",
	}

	// 第三个参数代表是否为模拟环境，更多信息查看接口说明
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

// 订阅私有频道
func wsPriv() {
	ep := "wss://ws.okex.com:8443/ws/v5/private?brokerId=9999"

	// 填写您自己的APIKey信息
	apikey := "xxxx"
	secretKey := "xxxxx"
	passphrase := "xxxxx"

	// 创建ws客户端
	r, err := NewWsClient(ep)
	if err != nil {
		log.Println(err)
		return
	}

	// 设置连接超时
	r.SetDailTimeout(time.Second * 2)
	err = r.Start()
	if err != nil {
		log.Println(err)
		return
	}
	defer r.Stop()
	var res bool

	res, _, err = r.Login(apikey, secretKey, passphrase)
	if res {
		log.Println("login success")
	} else {
		log.Println("login failed", err)
		return
	}

	// 订阅账户频道
	var args []map[string]string
	arg := make(map[string]string)
	arg["ccy"] = "BTC"
	args = append(args, arg)

	start := time.Now()
	res, _, err = r.PrivAccout(OP_SUBSCRIBE, args)
	if res {
		usedTime := time.Since(start)
		log.Println("subscribe success, cost: ", usedTime.String())
	} else {
		log.Println("subscribe failed, ", err)
	}

	time.Sleep(100 * time.Second)
	start = time.Now()
	res, _, err = r.PrivAccout(OP_UNSUBSCRIBE, args)
	if res {
		usedTime := time.Since(start)
		log.Println("unsubscribe success, ", usedTime.String())
	} else {
		log.Println("unsubscribe failed, ", err)
	}

}

// 订阅公共频道
func wsPub() {
	ep := "wss://ws.okex.com:8443/ws/v5/public?brokerId=9999"

	// 创建ws客户端
	r, err := NewWsClient(ep)
	if err != nil {
		log.Println(err)
		return
	}

	// 设置连接超时
	r.SetDailTimeout(time.Second * 2)
	err = r.Start()
	if err != nil {
		log.Println(err)
		return
	}
	defer r.Stop()
	// 订阅产品频道
	var args []map[string]string
	arg := make(map[string]string)
	arg["instType"] = FUTURES
	//arg["instType"] = OPTION
	args = append(args, arg)

	start := time.Now()
	res, _, err := r.PubInstruemnts(OP_SUBSCRIBE, args)
	if res {
		usedTime := time.Since(start)
		log.Println("subscribe success", usedTime.String())
	} else {
		log.Println("subscribe failed", err)
	}

	time.Sleep(30 * time.Second)

	start = time.Now()
	res, _, err = r.PubInstruemnts(OP_UNSUBSCRIBE, args)
	if res {
		usedTime := time.Since(start)
		log.Println("unsubscribe success", usedTime.String())
	} else {
		log.Println("unsubscribe failed", err)
	}
}

// websocket交易
func wsJrpc() {
	ep := "wss://ws.okex.com:8443/ws/v5/private?brokerId=9999"

	// 填写您自己的APIKey信息
	apikey := "xxxx"
	secretKey := "xxxxx"
	passphrase := "xxxxx"

	var res bool
	var req_id string

	// 创建ws客户端
	r, err := NewWsClient(ep)
	if err != nil {
		log.Println(err)
		return
	}

	// 设置连接超时
	r.SetDailTimeout(time.Second * 2)
	err = r.Start()
	if err != nil {
		log.Println(err)
		return
	}

	defer r.Stop()

	res, _, err = r.Login(apikey, secretKey, passphrase)
	if res {
		log.Println("login success")
	} else {
		log.Println("login failed", err)
		return
	}

	start := time.Now()
	param := map[string]interface{}{}
	param["instId"] = "BTC-USDT"
	param["tdMode"] = "cash"
	param["side"] = "buy"
	param["ordType"] = "market"
	param["sz"] = "200"
	req_id = "00001"

	res, _, err = r.PlaceOrder(req_id, param)
	if res {
		usedTime := time.Since(start)
		log.Println("place order success", usedTime.String())
	} else {
		usedTime := time.Since(start)
		log.Println("place order failed", usedTime.String(), err)
	}
}

func main() {
	// 公共订阅
	wsPub()

	// 私有订阅
	wsPriv()

	// websocket交易
	wsJrpc()

	// rest请求
	REST()
}
