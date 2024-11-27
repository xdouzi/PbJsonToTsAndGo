package pb;
/**
由 Net_Login.proto 文件生成 ...
author:yh  2024.11.27
*/
 //[登录]
 enum Cmd {
  Login_Login(0),
  Login_Ok(-1),
  ANTIWALLOW_ID_REPEATE(-2),
  ANTIWALLOW_FAIL(-3),
  ANTIWALLOW_NOT_ENOUGH_18(-4),
 	;
 
 private final int code;
 // 构造器
 Cmd(int code) {
 	this.code = code;
 }
 // 获取枚举值的代码
 public int getCode() {
 	return code;
 }
  }

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

