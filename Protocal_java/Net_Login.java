package pb;
/**
由 Net_Login.proto 文件生成 ...
author:yh  2024.11.27
*/
 class Base_Login_LoginInfo {
   public String DeviceId;
 }

 class Net_Login_LoginReq {
   public String DeviceId;
   public Base_Login_LoginInfo []LoginListInfo;
   public map[int]Base_Login_LoginInfo LoginMapInfo; 
   public Base_Login_LoginInfo LoginInfo;
 }

 class Net_Login_LoginRet {
   public int Ret;
 }

