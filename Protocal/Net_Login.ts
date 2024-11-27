/**
由 Net_Login.xlsx %!s(MISSING) excel文件生成 ...
author:yh 
*/
export class Net_Login_LoginReq{
 public DeviceId:string;
 public FacebookId:string;
 public AppleId:string;
 public HmsId:string;
 public Credential:string;
 public Package:string;
 public DeviceModel:string;
 public SysVersion:string;
}

export class Net_Login_LoginRet{
 public Ret:int32;
 public Uid:int64;
 public AntiWallow:int32;
 public Credential:string;
 public ServerTime:int64;
 public DifferentDevice:int32;
 public IsFirstLogin:bool;
}

}
