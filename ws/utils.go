package ws

import (
	"errors"
	. "github.com/superpdm/OKEX_V5SDK_GO/ws/wImpl"
	. "github.com/superpdm/OKEX_V5SDK_GO/ws/wInterface"
	"log"
	"runtime/debug"
)

// 判断返回结果成功失败
func checkResult(wsReq WSReqData, wsRsps []*Msg) (res bool, err error) {
	defer func() {
		a := recover()
		if a != nil {
			log.Printf("Receive End. Recover msg: %+v", a)
			debug.PrintStack()
		}
		return
	}()

	res = false
	if len(wsRsps) == 0 {
		return
	}

	for _, v := range wsRsps {
		switch v.Info.(type) {
		case ErrData:
			return
		}
		if wsReq.GetType() != v.Info.(WSRspData).MsgType() {
			err = errors.New("message type inconsistent")
			return
		}
	}

	//检查所有频道是否都更新成功
	if wsReq.GetType() == MSG_NORMAL {
		req, ok := wsReq.(ReqData)
		if !ok {
			log.Println("type cast failed", req)
			err = errors.New("type cast failed")
			return
		}

		for idx, _ := range req.Args {
			ok := false
			i_req := req.Args[idx]
			for i, _ := range wsRsps {
				info, _ := wsRsps[i].Info.(RspData)
				//log.Println("<<",info)
				if info.Event == req.Op && info.Arg["channel"] == i_req["channel"] && info.Arg["instType"] == i_req["instType"] {
					ok = true
					continue
				}
			}
			if !ok {
				err = errors.New("not all expected response")
				return
			}
		}
	} else {
		for i, _ := range wsRsps {
			info, _ := wsRsps[i].Info.(JRPCRsp)
			if info.Code != "0" {
				return
			}
		}
	}

	res = true
	return
}
