package pb
/**
由 Net_Login.xlsx %!s(MISSING) excel文件生成 ...
author:yh 
*/
 //[登录]
 type ANTIWALLOW_RET int32
 const (
  ANTIWALLOW_SUC ANTIWALLOW_RET = 0
  ANTIWALLOW_ID_INVALID ANTIWALLOW_RET = -1
  ANTIWALLOW_ID_REPEATE ANTIWALLOW_RET = -2
  ANTIWALLOW_FAIL ANTIWALLOW_RET = -3
  ANTIWALLOW_NOT_ENOUGH_18 ANTIWALLOW_RET = -4
  )

 type Base_Login_LoginInfo struct {
  DeviceId   string
 }

 type Net_Login_LoginReq struct {
  DeviceId   string
  FacebookId   string
  AppleId   string
  HmsId   string
  Credential   string
  Package   string
  DeviceModel   string
  SysVersion   string
  LoginInfo   *Base_Login_LoginInfo
 }

 type Net_Login_LoginRet struct {
  Ret   int32
  Uid   int64
  AntiWallow   int32
  Credential   string
  ServerTime   int64
  DifferentDevice   int32
  IsFirstLogin   bool
 }

