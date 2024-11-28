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
   public String FacebookId;
   public String AppleId;
   public String HmsId;
   public String Credential;
   public String Package;
   public String DeviceModel;
   public String SysVersion;
   public Base_Login_LoginInfo LoginInfo;
 }

 class Net_Login_LoginRet {
   public int Ret;
   public long Uid;
   public int AntiWallow;
   public String Credential;
   public long ServerTime;
   public int DifferentDevice;
   public Boolean IsFirstLogin;
 }

