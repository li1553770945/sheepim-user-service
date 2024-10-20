namespace go auth

include "base.thrift"

struct LoginReq {
    1: required string username
    2: required string password
}

struct LoginResp {
    1: required base.BaseResp baseResp
    2: optional string token
}


struct LogoutReq {
    1: required string token
}

struct LogoutResp {
    1: required base.BaseResp baseResp
}


struct GenerateActiveCodeReq {
    1: required string username
}

struct GenerateActiveCodeResp {
    1: required base.BaseResp baseResp
    2: optional string token
}

service AuthService {
    LoginResp Login(LoginReq req)
    LogoutResp Logout(LogoutReq req)
    GenerateActiveCodeResp GenerateActiveCode(GenerateActiveCodeReq req)

}
