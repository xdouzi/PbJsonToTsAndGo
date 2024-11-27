/**
由 Net_Login.xlsx %!s(MISSING) excel文件生成 ...
author:yh 
*/
export namespace pb {
 export enum  ANTIWALLOW_RET {
  ANTIWALLOW_SUC = 0,
  ANTIWALLOW_ID_INVALID = -1,
  ANTIWALLOW_ID_REPEATE = -2,
  ANTIWALLOW_FAIL = -3,
  ANTIWALLOW_NOT_ENOUGH_18 = -4,
 }

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
