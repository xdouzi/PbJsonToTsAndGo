package pb
/**
由 Net_Login.proto 文件生成 ...
author:yh  2024.11.27
*/
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

