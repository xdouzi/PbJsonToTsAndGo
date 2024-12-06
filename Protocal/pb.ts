export namespace pb {
/**
由 Net_Protocol.xlsx %!!(MISSING)s(MISSING) excel文件生成 ...
author:yh 
*/
 export enum  CmdModule {
  IdleModule = 0,
  Login = 1001,
  Lobby = 1002,
  World = 1003,
  Chat = 1004,
  Activity = 1006,
  Task = 1007,
  Theme = 1008,
 }

 export enum  Cmd {
  CmdIdle = 0,
  Error = 1001001,
  Login_Login = 1001002,
  Login_Logout = 1001022,
  Login_CloseServer = 1001023,
  Login_KeepAlive = 1001003,
  Login_Reconnect = 1001004,
  Login_FacebookConnect = 1001005,
  Login_FacebookLogout = 1001006,
  Login_FacebookSetUserInfo = 1001007,
  Login_CollectFacebookReward = 1001008,
  Login_AppleLogin = 1001009,
  Login_AppleLogout = 1001010,
  Login_MultiUserLogin = 1001011,
  Player_PlayerInfo = 1002001,
  Player_UpdateMoneyInfo = 1004002,
  Player_PlayerUpdatesInfo = 1002003,
  Player_ChangeName = 1002004,
  Lobby_AuthorizationLogin = 1003001,
  Lobby_AuthorizationEnterGame = 1003002,
  Lobby_LuckyWheel = 1003003,
  Lobby_ThemeStore = 1003005,
  Lobby_MoneyBank = 1003006,
  Lobby_VipNewRule = 1003007,
  Lobby_RateUs = 1003008,
  Lobby_GetRanking = 1003009,
  Lobby_EnterStampGame = 1003010,
  Lobby_CollectLoginBonus = 1003011,
  Lobby_ReissueLoginBonus = 1003012,
  Lobby_GetDailyBonus = 1003015,
  Lobby_CollectDailyBonus = 1003016,
  Backpack_GetBackpackInfo = 1004001,
  Backpack_BackpackUpdate = 1004003,
  Backpack_UseItems = 1004004,
  Backpack_OpenCell = 1004005,
  Activity_UpdateActivity = 1006001,
  Activity_ReceiveFireDragonTreasureRewards = 1006002,
  Activity_UpdateAdsUserType = 1006003,
  Task_UpdateTask = 1007001,
  Task_UpdateCupTask = 1007002,
  Task_GetCupTask = 1007003,
  Task_ReceiveCupTaskAward = 1007004,
  Task_ReceiveTaskReward = 1007005,
  Shop_GetShopInfo = 1008001,
  Shop_BuyItem = 1008002,
  Shop_Pay = 1008003,
  Inbox_UpdateItemsInfo = 1009001,
  Inbox_SendEmail = 1009002,
  Inbox_DeleteEmail = 1009003,
  Inbox_ClaimRewardGifts = 1009004,
  Lobby_GetFcmRewards = 1009005,
  Lobby_CollectFcmRewards = 1009006,
  Collect_WelcomeBack = 1009007,
  Guide_MarkComplete = 1010001,
  Ads_GetAdsTasks = 1011001,
  Ads_CollectAdsReward = 1011002,
  Match3_GetHeroList = 1005004,
  Match3_UpdateHeroCard = 1005091,
  Match3_HeroCardUnlock = 1005092,
  Match3_HeroCardBuyUseNum = 1005093,
  Match3_BuyTalentSkill = 1005094,
  Match3_StartMatchingRoom = 1005001,
  Match3_CancelMatching = 1005002,
  Match3_StartMatchingRoomSuccess = 1005003,
  Match3_EnterRoom = 1005005,
  Match3_GetSceneInformation = 1005006,
  Match3_OtherRoleEnterRoom = 1005007,
  Match3_ExitRoom = 1005008,
  Match3_OtherRoleExitRoom = 1005009,
  Match3_NextMove = 1005011,
  Match3_WaitOtherPlayerOfflineCountdown = 1005012,
  Match3_RoleDeathJieSuan = 1005017,
  Match3_GameOver = 1005013,
  Match3_HeroAttack = 1005014,
  Match3_ReleaseSkill = 1005015,
  Match3_Interfere = 1005016,
  Match3_NextRoundOperation = 1005018,
  PassCheck_UpdateItemsInfo = 1012001,
  PassCheck_ClaimReward = 1012002,
  PassCheck_Buy = 1012003,
  Slots_Spin = 1021001,
 }


/**
由 Net_Login.xlsx %!!(MISSING)s(MISSING) excel文件生成 ...
author:yh 
*/
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
