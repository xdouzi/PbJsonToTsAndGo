package Pb
/**
由 Net_Protocol.proto 文件生成 ...
author:yh  2024.11.27
*/
 //[消息模块]
 type CmdModule int32
 const (
  IdleModule CmdModule = 0
  Login CmdModule = 1001
  Lobby CmdModule = 1002
  World CmdModule = 1003
  Chat CmdModule = 1004
  Activity CmdModule = 1006
  Task CmdModule = 1007
  Theme CmdModule = 1008
  )

 //[ //消息头id added by yh @ 2023/08/25 14:09 注意:协议规则 协议号/1000 为模块id 取余为消息号  纠结了很久 最终还是坚持 Cmd.Login_Login //本是Cmd.Login_Login 这样的 代表登录模块的-登录消息 无赖C# PB不支持枚举下划线 Net_Login_LoginReq //客户端请求格式  通信_模块_消息名字Req Net_Login_LoginRet //服务器返回格式  通信_模块_消息名字Ret  //C#客户端 把To看做下划线使用 Cmd.Login_Login --> Cmd.Login_Login GameNet.Instance.RegisterProtocolHandler(Cmd.Login_Login,ProtocolLoginHandler); private void ProtocolLoginHandler(MsgPendingData msg_data){     Net_Login_LoginRet data = Net_Login_LoginRet.Parser.ParseFrom(msg_data.data_bytes); } ProtocolVersion =1;//协议版本号 ]
 type Cmd int32
 const (
  CmdIdle Cmd = 0
  Error Cmd = 1001001
  Login_Login Cmd = 1001002
  Login_Logout Cmd = 1001022
  Login_CloseServer Cmd = 1001023
  Login_KeepAlive Cmd = 1001003
  Login_Reconnect Cmd = 1001004
  Login_FacebookConnect Cmd = 1001005
  Login_FacebookLogout Cmd = 1001006
  Login_FacebookSetUserInfo Cmd = 1001007
  Login_CollectFacebookReward Cmd = 1001008
  Login_AppleLogin Cmd = 1001009
  Login_AppleLogout Cmd = 1001010
  Login_MultiUserLogin Cmd = 1001011
  Player_PlayerInfo Cmd = 1002001
  Player_UpdateMoneyInfo Cmd = 1004002
  Player_PlayerUpdatesInfo Cmd = 1002003
  Player_ChangeName Cmd = 1002004
  Lobby_AuthorizationLogin Cmd = 1003001
  Lobby_AuthorizationEnterGame Cmd = 1003002
  Lobby_LuckyWheel Cmd = 1003003
  Lobby_ThemeStore Cmd = 1003005
  Lobby_MoneyBank Cmd = 1003006
  Lobby_VipNewRule Cmd = 1003007
  Lobby_RateUs Cmd = 1003008
  Lobby_GetRanking Cmd = 1003009
  Lobby_EnterStampGame Cmd = 1003010
  Lobby_CollectLoginBonus Cmd = 1003011
  Lobby_ReissueLoginBonus Cmd = 1003012
  Lobby_GetDailyBonus Cmd = 1003015
  Lobby_CollectDailyBonus Cmd = 1003016
  Backpack_GetBackpackInfo Cmd = 1004001
  Backpack_BackpackUpdate Cmd = 1004003
  Backpack_UseItems Cmd = 1004004
  Backpack_OpenCell Cmd = 1004005
  Activity_UpdateActivity Cmd = 1006001
  Activity_ReceiveFireDragonTreasureRewards Cmd = 1006002
  Activity_UpdateAdsUserType Cmd = 1006003
  Task_UpdateTask Cmd = 1007001
  Task_UpdateCupTask Cmd = 1007002
  Task_GetCupTask Cmd = 1007003
  Task_ReceiveCupTaskAward Cmd = 1007004
  Task_ReceiveTaskReward Cmd = 1007005
  Shop_GetShopInfo Cmd = 1008001
  Shop_BuyItem Cmd = 1008002
  Shop_Pay Cmd = 1008003
  Inbox_UpdateItemsInfo Cmd = 1009001
  Inbox_SendEmail Cmd = 1009002
  Inbox_DeleteEmail Cmd = 1009003
  Inbox_ClaimRewardGifts Cmd = 1009004
  Lobby_GetFcmRewards Cmd = 1009005
  Lobby_CollectFcmRewards Cmd = 1009006
  Collect_WelcomeBack Cmd = 1009007
  Guide_MarkComplete Cmd = 1010001
  Ads_GetAdsTasks Cmd = 1011001
  Ads_CollectAdsReward Cmd = 1011002
  Match3_GetHeroList Cmd = 1005004
  Match3_UpdateHeroCard Cmd = 1005091
  Match3_HeroCardUnlock Cmd = 1005092
  Match3_HeroCardBuyUseNum Cmd = 1005093
  Match3_BuyTalentSkill Cmd = 1005094
  Match3_StartMatchingRoom Cmd = 1005001
  Match3_CancelMatching Cmd = 1005002
  Match3_StartMatchingRoomSuccess Cmd = 1005003
  Match3_EnterRoom Cmd = 1005005
  Match3_GetSceneInformation Cmd = 1005006
  Match3_OtherRoleEnterRoom Cmd = 1005007
  Match3_ExitRoom Cmd = 1005008
  Match3_OtherRoleExitRoom Cmd = 1005009
  Match3_NextMove Cmd = 1005011
  Match3_WaitOtherPlayerOfflineCountdown Cmd = 1005012
  Match3_RoleDeathJieSuan Cmd = 1005017
  Match3_GameOver Cmd = 1005013
  Match3_HeroAttack Cmd = 1005014
  Match3_ReleaseSkill Cmd = 1005015
  Match3_Interfere Cmd = 1005016
  Match3_NextRoundOperation Cmd = 1005018
  PassCheck_UpdateItemsInfo Cmd = 1012001
  PassCheck_ClaimReward Cmd = 1012002
  PassCheck_Buy Cmd = 1012003
  Slots_Spin Cmd = 1021001
  )

