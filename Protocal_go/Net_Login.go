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

 export class Base_Login_LoginInfo{
  public DeviceId:string;
 }

 export class Net_Login_LoginReq{
  public DeviceId:string;
  public FacebookId:string;
  public AppleId:string;
  public HmsId:string;
  public Credential:string;
  public Package:string;
  public DeviceModel:string;
  public SysVersion:string;
  public LoginInfo:Base_Login_LoginInfo;
 }

 export class Net_Login_LoginRet{
  public Ret:number;
  public Uid:number;
  public AntiWallow:number;
  public Credential:string;
  public ServerTime:number;
  public DifferentDevice:number;
  public IsFirstLogin:boolean;
 }

}
