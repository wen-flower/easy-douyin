namespace go chat

include "common.thrift"

struct MessageListParam {
    1: i64 logged_user_id // 登录用户的 ID
    2: i64 to_user_id // 对方的ID
}

struct MessageListResp {
    1: common.BaseResp base_resp
    2: list<common.MessageInfo> message_list
}

struct MessageActionParam {
    1: i64 logged_user_id // 登录用户的 ID
    2: i64 to_user_id // 对方的ID
    3: i16 action_type // 1-发送消息
    4: string content // 消息内容
}

struct MessageActionResp {
    1: common.BaseResp base_resp
}

service ChatService {
    MessageListResp MessageList(1: MessageListParam param)
    MessageActionResp MessageAction(1: MessageActionParam param)
}