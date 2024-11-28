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
  LoginListInfo   []*Base_Login_LoginInfo
  LoginMapInfo map[int32]*Base_Login_LoginInfo 
  LoginInfo   *Base_Login_LoginInfo
 }

 type Net_Login_LoginRet struct {
  Ret   int32
 }

