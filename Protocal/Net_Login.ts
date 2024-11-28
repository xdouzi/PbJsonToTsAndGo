/**
由 Net_Login.xlsx %!s(MISSING) excel文件生成 ...
author:yh 
*/
export namespace pb {
 export class Base_Login_LoginInfo{
  public DeviceId:string;
 }

 export class Net_Login_LoginReq{
  public DeviceId:string;
  public LoginListInfo:Base_Login_LoginInfo[];
  public LoginMapInfo: Map<number, Base_Login_LoginInfo>;
  public LoginInfo:Base_Login_LoginInfo;
 }

 export class Net_Login_LoginRet{
  public Ret:number;
 }

}
