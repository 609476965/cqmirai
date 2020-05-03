package main

import (
	"gitee.com/LXY1226/logging"
	"github.com/json-iterator/go"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

func (c *CMiraiWSRConn) TransMsgToMirai(msg []byte) []byte {
	//iter := IteratorPool.BorrowIterator(msg)
	//iter.
	req := new(cqRequest)
	err := json.Unmarshal(msg, &req)
	if err != nil {
		logging.WARN("解析CQ消息失败: ", err.Error())
		return nil
	}
	switch req.Action {
	case "send_msg":
		cqResp := c.sendMsg(req.Params.ToString())
		cqResp.Echo = req.Echo
		o, err := json.Marshal(cqResp)
		if err != nil {
			logging.WARN("生成CQ回执失败: ", err.Error())
			return nil
		}
		return o
	case "get_group_member_info":
		//return c.getGroupMemberInfo(j)
		return append([]byte("{"), append(req.Echo, '}')...)
	default:
		return append([]byte("{"), append(req.Echo, '}')...)
	}
}

func (c *CMiraiWSRConn) TransMsgToCQ(msg []byte) []byte {
	miraiMsg := new(Message)
	err := json.Unmarshal(msg, miraiMsg)
	if err != nil {
		logging.WARN("解析Mirai消息失败: ", err.Error())
		return nil
	}
	switch miraiMsg.Type {
	case "GroupMessage":
		return c.MiraiGroupMessage(miraiMsg)
	case "FriendMessage":
		return c.MiraiFriendMessage(miraiMsg)
	default:
		return nil

	}
}
