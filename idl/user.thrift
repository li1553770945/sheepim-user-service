namespace go user

include "base.thrift"

struct CheckUsernameAndPasswdReq {
    1: required string username
    2: required string password
}

struct CheckUsernameAndPasswdResp {
    1: required base.BaseResp baseResp
    2: optional i64 userId
}

struct AddUserReq {
    1: required string username
    2: required string password
    3: required string nickname
    4: required string active_code
}

struct AddUserResp {
    1: required base.BaseResp baseResp
    2: optional i64 userId
}


struct UserInfoReq {
    1: required i64 userId
}

struct UserInfoResp {
    1: required base.BaseResp baseResp
    2: optional i64 userId
    3: optional string username
    4: optional string nickname
    5: optional string role
}

service UserService {
    CheckUsernameAndPasswdResp  CheckUsernameAndPasswd(1: CheckUsernameAndPasswdReq req)
    AddUserResp AddUser(1:  AddUserReq req)
    UserInfoResp GetUserInfo(1: UserInfoReq req)
}
