package main

import (
	"time"
)

type UserAlbum struct {
	Id        int    `xorm:"not null pk autoincr INT(11)"`
	Uid       int64  `xorm:"not null default 0 comment('用户id') index BIGINT(20)"`
	PhotoUrl  string `xorm:"not null default '' comment('照片url') VARCHAR(256)"`
	PhotoFile string `xorm:"not null default '' comment('本地文件路径') VARCHAR(128)"`
	Visible   int    `xorm:"not null default b'1' comment('是否可见（如果用户删除，那么标识为不可见）') BIT(1)"`
}

type ModelCfg struct {
	Id                int    `xorm:"not null pk autoincr comment('自增id') INT(16)"`
	ModelId           int64  `xorm:"not null comment('id') BIGINT(32)"`
	ModelType         int    `xorm:"not null comment('类型') INT(16)"`
	ModelCnName       string `xorm:"not null default '' comment('中文名') VARCHAR(32)"`
	ModelEnName       string `xorm:"not null default '' comment('英文名') VARCHAR(32)"`
	ModelPrice        int64  `xorm:"not null comment('价格') BIGINT(16)"`
	ModelSex          int    `xorm:"not null comment('性别') TINYINT(4)"`
	ModelIcon         string `xorm:"not null default '' comment('icon') VARCHAR(64)"`
	ModelModel        string `xorm:"not null default '' comment('model') VARCHAR(64)"`
	ModelCurrencyType int    `xorm:"not null comment('货币类型') INT(8)"`
	LimitCount        int64  `xorm:"not null comment('限购数量') BIGINT(16)"`
	DisplayShop       int    `xorm:"not null comment('是否在商城中显示') TINYINT(4)"`
	ModelTag          int    `xorm:"not null comment('热卖等标签') TINYINT(4)"`
}

type UserModel struct {
	Id      int64 `xorm:"pk autoincr comment('自增id') BIGINT(22)"`
	Uid     int64 `xorm:"not null comment('uid') index BIGINT(22)"`
	ModelId int64 `xorm:"not null comment('model id') BIGINT(21)"`
	Status  int   `xorm:"not null comment('是否使用') TINYINT(2)"`
	BuyTime int64 `xorm:"not null comment('购买时间') BIGINT(32)"`
	Count   int64 `xorm:"not null comment('数量') BIGINT(20)"`
}

type UserShoppingCart struct {
	Id      int64 `xorm:"pk autoincr comment('自增id') BIGINT(22)"`
	Uid     int64 `xorm:"not null comment('uid') index BIGINT(22)"`
	ModelId int64 `xorm:"not null comment('model id') BIGINT(21)"`
	AddTime int64 `xorm:"not null comment('加入购物车时间') BIGINT(32)"`
	Count   int64 `xorm:"not null comment('数量') BIGINT(20)"`
}

type DcLog struct {
	Id            int64  `xorm:"pk autoincr comment('自增id') BIGINT(20)"`
	DcType        int32  `xorm:"not null default 0 comment('dc类型') INT(11)"`
	Oid           string `xorm:"not null default '' comment('订单id')  VARCHAR(100)"`
	Suid          int64  `xorm:"not null default 0 comment('赠送用户id')  BIGINT(20)"`
	Ruid          int64  `xorm:"not null default 0 comment('接收用户id')  BIGINT(20)"`
	ChangeType    int32  `xorm:"not null default 0 comment('改变类型')  INT(11)"`
	Gid           int32  `xorm:"not null default 0 comment('礼物id')  INT(11)"`
	Cnt           int32  `xorm:"not null default 0 comment('礼物数量') INT(11)"`
	Giftprice     int32  `xorm:"not null default 0 comment('礼物单价') INT(11)"`
	DuserType     int32  `xorm:"not null default 1 comment('目标用户身份类型：0普通用户1主播') TINYINT(4)"`
	DcChanges     int64  `xorm:"not null default 0 comment('目标用户金币变化量') BIGINT(20)"`
	DcNew         int64  `xorm:"not null default 0 comment('接收者接收前金币余额') BIGINT(20)"`
	RoomId        int64  `xorm:"not null default 0 comment('房间id') BIGINT(20)"`
	RoomType      int32  `xorm:"not null default 0 comment('房间类型') INT(11)"`
	DcDesc        string `xorm:"not null default '' comment('交易描述') VARCHAR(50)"`
	DcHash        string `xorm:"not null default '' comment('交易hash') VARCHAR(50)"`
	DcService     int64  `xorm:"not null default 0 comment('交易手续费') BIGINT(20)"`
	DcAddress     string `xorm:"not null default '' comment('交易地址') VARCHAR(100)"`
	DcStatus      int32  `xorm:"not null default 0 comment('交易状态 0:审核中 1:成功 2:审核通过 3:审核不通过') TINYINT(4)"`
	DcServiceType int32  `xorm:"not null default 0 comment('交易手续费类型 0 ：默认跟交易类型一样') INT(11)"`
	DcNum         string `xorm:"not null default 0 comment('交易金额') VARCHAR(100)"`
}

type WithdrawLog struct {
	Id          int64     `xorm:"pk autoincr comment('自增id') BIGINT(20)"`
	TradeNo     string    `xorm:"not null default '' comment('订单id') VARCHAR(80)"`
	Uid         int64     `xorm:"not null default 0 comment('uid') BIGINT(20)"`
	Balance     string    `xorm:"not null default '' comment('价格') VARCHAR(80)"`
	BalanceType string    `xorm:"not null default '' comment('货币类型') VARCHAR(10)"`
	CreateTime  time.Time `xorm:"not null default '2020-02-02 02:02:02' comment('时间') DATETIME"`
	TradeType   int       `xorm:"not null default 0 comment('交易类型') INT(11)"`
	TradeStatus int       `xorm:"not null default 0 comment('交易类型') INT(4)"`
}

type Microphone struct {
	Id            int64   `xorm:"pk autoincr BIGINT(64)"`
	TokenId       int64   `xorm:"not null comment('tokenId') unique BIGINT(22)"`
	Uid           int64   `xorm:"not null comment('uid') index BIGINT(22)"`
	Quality       int64   `xorm:"not null comment('quality') BIGINT(22)"`
	InitLevel     int64   `xorm:"not null comment('初始等级') BIGINT(22)"`
	CurLevel      int64   `xorm:"not null comment('current level') BIGINT(22)"`
	Status        int     `xorm:"not null comment('status') BIGINT(22)"` // 麦克风当前的状态
	Intact        int64   `xorm:"not null comment('当前完整度') BIGINT(22)"`
	Ability       float64 `xorm:"not null comment('能力值') FLOAT(22)"` // 麦克风能力值
	Luck          float64 `xorm:"not null comment('幸运值') FLOAT(22)"` // 麦克风幸运值
	Endurance     float64 `xorm:"not null comment('耐力值') FLOAT(22)"`
	Power         int64   `xorm:"not null comment('基础电量') BIGINT(22)"`
	CapPower      float64 `xorm:"not null comment('蓄电量') FLOAT(22)"`
	MintCount     int     `xorm:"not null comment('合成次数') BIGINT(22)"` // 麦克风合成次数
	GainTime      int64   `xorm:"not null comment('获取时间') BIGINT(22)"`
	MintTime      int64   `xorm:"not null comment('合成时间') BIGINT(22)"`
	NextLevelTime int64   `xorm:"not null comment('下次升级时间') BIGINT(22)"`
	BuyPrice      int64   `xorm:"not null comment('购买价格') BIGINT(16)"`
	ImageUrl      string  `xorm:"not null comment('图片地址') VARCHAR(200)"`
	TokenName     string  `xorm:"not null comment('token编号') VARCHAR(50)"`
}

type MicrophoneGlobal struct {
	Id              int64   `xorm:"pk autoincr bigint(64)"`
	Uid             int64   `xorm:"not null comment('uid') unique BIGINT(22)"`
	UsedPower       float64 `xorm:"not null comment('当前已经消耗的电量') FLOAT(22)"`
	CurIncomeDay    int64   `xorm:"not null comment('当日当前的收益') BIGINT(22)"`
	RealIncomeDay   int64   `xorm:"not null comment('当日实际获得的收益') BIGINT(22)"`
	OnmicroTimeDay  int64   `xorm:"not null comment('当日上麦时长') BIGINT(22)"`
	ConSumePowerDay float64 `xorm:"not null comment('当日累计消耗的电量') FLOAT(22)"`
}

// 收益锁进度
type IncomeLockProgress struct {
	Id                  int64 `xorm:"pk autoincr bigint(64)"`
	Uid                 int64 `xorm:"not null comment('uid') index BIGINT(22)"`
	NeedIncomeProgresss int64 `xorm:"not null comment('收益进度锁') BIGINT(22)"`
	LockStatus          int64 `xorm:"not null comment('锁状态') BIGINT(22)"` // 0:未解锁 1:已解锁
	UnlockUid           int64 `xorm:"not null comment('解锁用户id') BIGINT(22)"`
	UnlockTime          int64 `xorm:"not null comment('解锁时间') BIGINT(22)"`
}

// 松果钥匙锁进度
type ConekeyLockProgress struct {
	Id                int64 `xorm:"pk autoincr bigint(64)"` // keyID
	Uid               int64 `xorm:"not null comment('uid') index BIGINT(22)"`
	NeedProgroessTime int64 `xorm:"not null comment('获取所需的进度') BIGINT(22)"`
	ConekeyStatus     int64 `xorm:"not null comment('状态') BIGINT(22)"` // 0: 未获得 1:已获得
}

// 背包中的松果钥匙
type BagConeKey struct {
	Id            int64 `xorm:"pk autoincr bigint(64)"` // keyID
	Uid           int64 `xorm:"not null comment('uid') index BIGINT(22)"`
	KeyId         int64 `xorm:"not null comment('keyId') BIGINT(22)"`
	GainTime      int64 `xorm:"not null comment('获取时间') BIGINT(22)"`
	ConekeyStatus int64 `xorm:"not null comment('背包中钥匙的状态') BIGINT(22)"` // 2:背包中；3：已过期；4：已送出
}

// 装备上麦卸载下麦最新记录
type OnequipoOrOnMicroRecord struct {
	Id          int64   `xorm:"pk autoincr bigint(64)"` // keyID
	Uid         int64   `xorm:"not null comment('uid') unique BIGINT(22)"`
	TokenId     int64   `xorm:"not null comment('麦克风ID') BIGINT(22)"` // 装备时赋值tokenID
	RoomId      int64   `xorm:"not null comment('房间ID') BIGINT(22)"`  // 上麦时赋值rommID
	EquipTime   int64   `xorm:"not null comment('装备时间') BIGINT(22)"`
	OnMicroTime int64   `xorm:"not null comment('上麦时间') BIGINT(22)"`
	UsedPower   float64 `xorm:"not null comment('实际产生收益那一刻已经使用的电量') FLOAT(22)"`
	MaxPower    int64   `xorm:"not null comment('实际产生收益那一刻的最大电量') BIGINT(22)"`
}

// 每次上麦产生收益记录表
type IncomeRecord struct {
	Id       int64     `xorm:"pk autoincr bigint(64)"`
	Uid      int64     `xorm:"not null comment('uid') index BIGINT(22)"`
	GainTime time.Time `xorm:"not null default '2020-02-02 02:02:02' comment('时间') DATETIME"` // 获得收益的时间
	InCome   int64     `xorm:"not null comment('收益FCR') BIGINT(22)"`                          // 收益FCR
}

// MintRecord 合成记录
type MintRecord struct {
	Id      int64 `xorm:"pk autoincr bigint(64)"` // keyID
	TokenId int64 `xorm:"not null comment('麦克风ID') unique BIGINT(22)"`
	Uid     int64 `xorm:"not null comment('uid') BIGINT(22)"`
	Mother  int64 `xorm:"not null comment('父麦克风ID') BIGINT(22)"`
	Father  int64 `xorm:"not null comment('母麦克风ID') BIGINT(22)"`
}

// NFT市场
type NftMarket struct {
	Id          int64  `xorm:"pk autoincr BIGINT(64)"`
	Uid         int64  `xorm:"not null comment('上架用户ID') BIGINT(22)"`
	Price       int64  `xorm:"not null comment('售卖价格') BIGINT(16)"`
	PutawayTime int64  `xorm:"not null comment('上架时间') BIGINT(22)"`
	NftId       int64  `xorm:"not null comment('nft ID') BIGINT(22)"`
	NftType     string `xorm:"not null default '' comment('nft 类型') VARCHAR(256)"`
}

// 盲盒表
type BlindBox struct {
	Id      int64 `xorm:"pk autoincr bigint(64)"`
	Quality int64 `xorm:"not null comment('盲盒品质') BIGINT(22)"`
	Uid     int64 `xorm:"not null comment('uid') index BIGINT(22)"`
	Mother  int64 `xorm:"not null comment('该盲盒的父MicroId') BIGINT(22)"`
	Father  int64 `xorm:"not null comment('该盲盒的母MicroId') BIGINT(22)"`
	Status  int64 `xorm:"not null comment('盲盒的状态') BIGINT(22)"`
}

type MicroRechargeError struct {
	Id        int64  `xorm:"pk autoincr bigint(64)"`
	Uid       int64  `xorm:"not null comment('uid') BIGINT(22)"`
	TokenId   int64  `xorm:"not null comment('麦克风ID') unique BIGINT(22)"`
	TokenUrl  string `xorm:"not null comment('tokenUrl') VARCHAR(200)"`
	EventType int64  `xorm:"not null comment('uid') BIGINT(22)"`
}

// 使用麦克风产生收益详细记录表
type UseMicroincomeDescrecord struct {
	Id            int64   `xorm:"pk autoincr bigint(64)"`
	Uid           int64   `xorm:"not null comment('uid') BIGINT(22)"`
	RoomId        int64   `xorm:"not null comment('房间ID') BIGINT(22)"`
	TokenId       int64   `xorm:"not null comment('麦克风ID') BIGINT(22)"`
	Level         int64   `xorm:"not null comment('等级') BIGINT(22)"`
	TokenName     string  `xorm:"not null comment('编号') VARCHAR(200)"`
	ImageUrl      string  `xorm:"not null comment('图片地址') VARCHAR(200)"`
	Quality       int64   `xorm:"not null comment('品质') BIGINT(22)"`
	RoomUserNums  int64   `xorm:"not null comment('房间人数') BIGINT(22)"`
	EquipTime     int64   `xorm:"not null comment('装备时间') BIGINT(22)"`
	UninstallTime int64   `xorm:"not null comment('卸载装备时间') BIGINT(22)"`
	OnMicroTime   int64   `xorm:"not null comment('上麦时间') BIGINT(22)"`
	OutMicroTime  int64   `xorm:"not null comment('下麦时间') BIGINT(22)"`
	UseTime       int64   `xorm:"not null comment('使用时长') BIGINT(22)"`
	ConsumePower  float64 `xorm:"not null comment('消耗的电量') FLOAT(22)"`
	OldIntact     int64   `xorm:"not null comment('原磨损度') BIGINT(22)"`
	NewIntact     int64   `xorm:"not null comment('新磨损度') BIGINT(22)"`
	UnlockInCome  int64   `xorm:"not null comment('未锁住收益FCR') BIGINT(22)"`
	LockInCome    int64   `xorm:"not null comment('锁住收益FCR') BIGINT(22)"`
}

// 联合查表需要把多个表结构合在一起
type NftMarketAllData struct {
	NftMarket  `xorm:"extends"`
	Microphone `xorm:"extends"`
}

type GiftInfo struct {
	Gid            int32     `xorm:"not null pk autoincr comment('礼物id') INT(11)"`
	GiftType       int32     `xorm:"not null default 1 comment('礼物类型：1现金礼物2幸运礼物3热门礼物4奢侈礼物5彩蛋6彩蛋爆出的礼物') INT(11)"`
	GnameEn        string    `xorm:"not null default '' comment('礼物名称') VARCHAR(64)"`
	GnameZh        string    `xorm:"not null default '' comment('礼物名称中文') VARCHAR(64)"`
	Icon           string    `xorm:"not null default '' comment('礼物图标') VARCHAR(200)"`
	Price          int32     `xorm:"not null default 0 comment('礼物价格单位( 分)') INT(11)"`
	Exp            int64     `xorm:"not null default 0 comment('赠送获取经验值') BIGINT(20)"`
	AnchorExp      int64     `xorm:"not null default 0 comment('如果是送给主播，主播获取的经验值') BIGINT(20)"`
	AnchorGetgold  int64     `xorm:"not null default 0 comment('如果是送给主播，主播获取的金币值') BIGINT(20)"`
	UserGetgold    int64     `xorm:"not null default 0 comment('如果是送非主播用户，对方获取的金币值') BIGINT(20)"`
	SysGetgold     int64     `xorm:"not null default 0 comment('系统收益') BIGINT(20)"`
	Pos            int32     `xorm:"not null default 0 comment('位置顺序') INT(11)"`
	Visible        int32     `xorm:"not null default 1 comment('是否可见') INT(11)"`
	Remark         string    `xorm:"not null default '' comment('备注') VARCHAR(200)"`
	Updatetime     time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' comment('修改时间') DATETIME"`
	Createtime     time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' comment('创建时间') DATETIME"`
	LabelPic       string    `xorm:"not null default '' comment('角标图片') VARCHAR(256)"`
	EnableLabel    int32     `xorm:"not null default b'0' comment('是否使用角标') BIT(1)"`
	IsBunch        int32     `xorm:"not null default b'0' comment('能否多个一起送') BIT(1)"`
	Bunch          string    `xorm:"not null default '' comment('连击配置') VARCHAR(128)"`
	Vip            int32     `xorm:"not null default 0 comment('VIP等级，大于等于此等级的用户才能赠送此礼物') INT(11)"`
	IsShowChatInfo int32     `xorm:"not null default b'0' comment('送礼物是否显示聊天信息') BIT(1)"`
	VisibleLv      int32     `xorm:"not null default 0 comment('财富等级达到这个字段指定的等级，才能显示') INT(11)"`
	CanGraffiti    int32     `xorm:"not null default b'0' comment('是否可以涂鸦') BIT(1)"`
	Spine          string    `xorm:"not null default '' comment('特效名字') VARCHAR(128)"`
	Hot            int32     `xorm:"not null default 0 comment('热度') INT(11)"`
}

type InviteLog struct {
	Id           int64 `xorm:"pk autoincr comment('自增') BIGINT(22)"`
	Fid          int64 `xorm:"not null default 0 comment('邀请id') BIGINT(22)"`
	Uid          int64 `xorm:"not null default 0 comment('被邀请id') BIGINT(20)"`
	InviteStatus int   `xorm:"not null default 1 comment('绑定状态：1：绑定，2：解绑') TINYINT(3)"`
}

type NftLog struct {
	Id              int64     `xorm:"pk autoincr bigint(64)"`
	CreateTime      time.Time `xorm:"not null default '2016-08-31 04:25:00' comment('时间') DATETIME"`
	NftType         string    `xorm:"not null default '' comment('nft类型(麦克风:fcm)') VARCHAR(22)"`
	WithDrawOrderId string    `xorm:"default '' comment('提取订单ID') index VARCHAR(100)"`
	SrcUid          int64     `xorm:"not null default 0 comment('赠送用户id')  BIGINT(20)"`
	ToUid           int64     `xorm:"not null default 0 comment('接收用户id')  BIGINT(20)"`
	ToAddress       string    `xorm:"not null default '' comment('接受地址') VARCHAR(100)"`
	FromAddress     string    `xorm:"not null default '' comment('源地址') VARCHAR(100)"`
	ItemId          int64     `xorm:"not null comment('tokenId/blankBoxId') BIGINT(22)"`
	TxStatus        int32     `xorm:"not null default 0 comment('交易状态 1:发起 2:成功 3:失败') TINYINT(4)"`
	TxHash          string    `xorm:"default '' comment('交易hash') index VARCHAR(200)"`
	TxType          int32     `xorm:"not null default 0 comment('交易类型(1:提取 2：充值 3:mint)') INT(11)"`
	Cost            int64     `xorm:"not null default 0 comment('手续费') INT(22)"`
}

type Chat struct {
	Id           int    `xorm:"not null pk autoincr INT(11)"`
	RoomId       int64  `xorm:"not null default 0 comment('房间id') index BIGINT(20)"`
	UserId       int64  `xorm:"not null default 0 comment('发送者用户id') index BIGINT(20)"`
	Nick         string `xorm:"not null default '' comment('发送者用户昵称') VARCHAR(64)"`
	Lv           int    `xorm:"not null default 0 comment('发送者用户财富等级') INT(11)"`
	Vip          int    `xorm:"not null default 0 comment('发送者用户vip') INT(11)"`
	VipExpired   int    `xorm:"not null default b'0' comment('发送者的VIP是否已经过期') BIT(1)"`
	AnchorType   int    `xorm:"not null default 0 comment('主播类型') INT(11)"`
	AnchorLv     int    `xorm:"not null default 0 comment('发送者用户主播等级') INT(11)"`
	RecvUserId   int64  `xorm:"not null default 0 comment('接手者用户id') index BIGINT(20)"`
	RecvNick     string `xorm:"not null default '' comment('@ 用户昵称;没有填“”;') VARCHAR(64)"`
	ChatCategory int    `xorm:"not null default 0 comment('聊天内容分类：0:文本; 1:礼物; 2:语音;') index INT(11)"`
	ChatType     int    `xorm:"not null default 0 comment('聊天类型 0:公聊 ; 1:私聊; 2:悄悄话;') index INT(11)"`
	Content      string `xorm:"not null default '' comment('聊天内容') VARCHAR(256)"`
	Duration     int    `xorm:"not null default 0 comment('如果是语音，表示语音时长（单位秒）') INT(11)"`
}

type WithdrawCfg struct {
	Id          int    `xorm:"not null pk autoincr comment('自增') INT(10)"`
	ServiceType int    `xorm:"not null default 0 comment('货币类型手续费(1:fcc2:fcr3:usdt4:matic)') INT(10)"`
	ServiceNum  int64  `xorm:"not null default 0 comment(' 手续费') BIGINT(20)"`
	RateUsd     string `xorm:"not null default '' comment('汇率') VARCHAR(50)"`
	RateEur     string `xorm:"not null default '' comment('汇率') VARCHAR(50)"`
	RateVnd     string `xorm:"not null default '' comment('汇率') VARCHAR(50)"`
	RateInr     string `xorm:"not null default '' comment('汇率') VARCHAR(50)"`
	ServiceName string `xorm:"not null default '' comment('货币名称') VARCHAR(50)"`
}

type PayLog struct {
	Id              int    `xorm:"not null pk autoincr INT(16)"`
	Uid             int64  `xorm:"not null comment('玩家id') BIGINT(20)"`
	TransactionId   string `xorm:"not null comment('交易id') VARCHAR(64)"`
	PayId           string `xorm:"comment('第三方订单id') VARCHAR(64)"`
	PayType         int    `xorm:"not null default 0 comment('第三方支付公司 0,mercuryo;1,alchemy') TINYINT(2)"`
	Address         string `xorm:"not null comment('钱包地址') VARCHAR(128)"`
	TransactionHash string `xorm:"comment('交易hash') VARCHAR(128)"`
	CardNumber      string `xorm:"comment('支付卡号') VARCHAR(64)"`
	Email           string `xorm:"comment('邮箱') VARCHAR(32)"`
	Currency        string `xorm:"comment('虚拟币类型') VARCHAR(16)"`
	Amount          string `xorm:"comment('虚拟币数量') VARCHAR(64)"`
	FiatCurrency    string `xorm:"comment('法币类型') VARCHAR(16)"`
	FiatAmount      string `xorm:"comment('法币数量') VARCHAR(64)"`
	Status          string `xorm:"comment('状态') VARCHAR(16)"`
	CreatedAt       int64  `xorm:"comment('创建时间') BIGINT(22)"`
	UpdatedAt       int64  `xorm:"comment('更新时间') BIGINT(22)"`
}

type OfficialInviteCode struct {
	Id   int    `xorm:"not null pk autoincr comment('自增id') INT(11)"`
	Code string `xorm:"not null default '' comment('邀请码') VARCHAR(32)"`
}

type UserStat struct {
	Uid            int64   `xorm:"not null pk comment('用户id') BIGINT(20)"`
	Lv             int32   `xorm:"not null default 0 comment('用户等级') INT(11)"`
	Exp            int64   `xorm:"not null default 0 comment('经验值') BIGINT(20)"`
	Vip            int32   `xorm:"not null default 0 comment('vip等级') INT(11)"`
	FriendCnt      int32   `xorm:"not null default 0 comment('关注数') INT(11)"`
	FansCnt        int32   `xorm:"not null default 0 comment('粉丝数') INT(11)"`
	FollowCnt      int32   `xorm:"not null default 0 comment('关注数') INT(11)"`
	GiftSendCnt    int64   `xorm:"not null default 0 comment('送礼数') BIGINT(20)"`
	GiftRecvCnt    int64   `xorm:"not null default 0 comment('收礼数') BIGINT(20)"`
	Isnew          bool    `xorm:"not null default b'1' comment('是否是新人') BIT(1)"`
	Tags           string  `xorm:"not null default '' comment('用户标签') VARCHAR(200)"`
	Latitude       float64 `xorm:"not null default 0 comment('纬度') DOUBLE"`
	Longitude      float64 `xorm:"not null default 0 comment('经度') DOUBLE"`
	Location       string  `xorm:"not null default '' comment('城市信息') VARCHAR(128)"`
	Status         int32   `xorm:"not null default 0 comment('账号状态：0正常，1禁言，2踢出，3冻结，4封号') TINYINT(4)"`
	Statustime     int32   `xorm:"not null default 0 comment('账号非正常状态时长，单位：分钟') INT(11)"`
	Isonline       int32   `xorm:"not null default 0 comment('是否在线') TINYINT(4)"`
	Isidle         int32   `xorm:"not null default 1 comment('是否空闲') TINYINT(4)"`
	PersonalStatus string  `xorm:"not null default 'Hey,I am a HALA LIVE member.' comment('个性签名') VARCHAR(512)"`
	RegistrationId string  `xorm:"not null default '' comment('极光推送使用的设备标识') VARCHAR(256)"`
	Language       string  `xorm:"not null default 'en' comment('用户使用的语言') VARCHAR(32)"`
	UniqueId       int64   `xorm:"not null default 0 comment('靓号id') BIGINT(20)"`
	Fcr            int64   `xorm:"not null default 0 comment('fcr') BIGINT(20)"`
	Fcc            int64   `xorm:"not null default 0 comment('fcc') BIGINT(20)"`
	Usdt           int64   `xorm:"not null default 0 comment('usdt') BIGINT(20)"`
	Eth            int64   `xorm:"not null default 0 comment('Eth') BIGINT(20)"`
	PinCode        string  `xorm:"not null default '' comment('提现密码') VARCHAR(10)"`
	MoneyAddress   string  `xorm:"not null default '' comment('钱包地址') VARCHAR(255)"`
}

type User struct {
	Uid         int64  `xorm:"not null pk comment('用户id') BIGINT(20)"`
	Uname       string `xorm:"not null default '' comment('用户账号') VARCHAR(64)"`
	Pwd         string `xorm:"not null default '' comment('密码') VARCHAR(32)"`
	Nick        string `xorm:"not null pk default '' comment('昵称') VARCHAR(64)"`
	Accounttype int32  `xorm:"not null default 0 comment('账号类型；0自有注册1FaceBook2Google3Twitter') INT(11)"`
	Openid      string `xorm:"not null default '' comment('openid') VARCHAR(64)"`
	Unionid     string `xorm:"not null default '' comment('unionid') VARCHAR(64)"`
	Pretty      int64  `xorm:"not null default 0 comment('靓号') BIGINT(20)"`
	Sex         int32  `xorm:"not null default 0 comment('性别：0未知1男2女3第三性别') TINYINT(4)"`
	Age         int32  `xorm:"not null default 0 comment('年龄') TINYINT(4)"`
	Head        string `xorm:"not null default '' comment('头像url') VARCHAR(500)"`
	Country     int32  `xorm:"not null default 0 comment('国家代码') INT(11)"`
	Countryname string `xorm:"not null default '' VARCHAR(64)"`
	City        string `xorm:"not null default '' comment('所在城市') VARCHAR(64)"`
	Zone        int32  `xorm:"not null default 0 comment('区号') INT(11)"`
	Phone       int64  `xorm:"not null default 0 comment('手机号') BIGINT(32)"`
	Regip       string `xorm:"not null default '' comment('注册IP') VARCHAR(32)"`
	Platform    string `xorm:"not null default '' comment('平台类型') VARCHAR(32)"`
	DevUuid     string `xorm:"not null default '' comment('设备UUID') VARCHAR(64)"`
	Ver         string `xorm:"not null default '' comment('客户端版本号') VARCHAR(16)"`
	Isrobot     bool   `xorm:"not null default b'0' BIT(1)"`
	Headtemp    string `xorm:"not null default '' VARCHAR(500)"`
	Email       string `xorm:"not null default '' comment('email') VARCHAR(64)"`
	StarId      int32  `xorm:"not null default 7 comment('星座id(默认为10，这个和birthday为1月1日对应)') TINYINT(4)"`
	Frozen      bool   `xorm:"not null default b'0' comment('是否封号状态') BIT(1)"`
	PkgId       int32  `xorm:"not null default 0 comment('包id，具体信息从package中找') INT(11)"`
	FlavorId    int32  `xorm:"not null default 0 comment('渠道商id，具体信息从flavors表中找') INT(11)"`
	CertName    string `xorm:"not null default '' comment('真实姓名') VARCHAR(64)"`
	CertNo      string `xorm:"not null default '' comment('身份证号') VARCHAR(32)"`
	First       bool   `xorm:"not null default b'0' comment('是否需要显示设置用户信息界面') BIT(1)"`
	Isyoungmod  bool   `xorm:"not null default b'0' comment('是否开启青少年模式') BIT(1)"`
	YoungModPwd string `xorm:"not null default '' comment('青少年模式密码') VARCHAR(255)"`
	RegPkgid    int32  `xorm:"not null default 0 comment('注册时包id，具体信息从package中找') INT(11)"`
	RegFlavorid int32  `xorm:"not null default 0 comment('注册时的渠道商id，具体信息从flavors表中找') INT(11)"`
	InviteCode  string `xorm:"not null default '' comment('邀请码') VARCHAR(32)"`
	UserType    int32  `xorm:"not null default 0 comment('0:普通 1:社区 2:城市 3:国家 ') INT(11)"`
}

type Fans struct {
	Uid      int64  `xorm:"not null pk comment('用户id') BIGINT(20)"`
	Fuid     int64  `xorm:"not null pk default 0 comment('粉丝用户id') BIGINT(20)"`
	Eachfans bool   `xorm:"not null default b'0' comment('是否互粉') BIT(1)"`
	Country  int32  `xorm:"not null default 86 comment('国家') INT(11)"`
	City     string `xorm:"not null default '' comment('城市') VARCHAR(64)"`
}

type NftOrder struct {
	Id              int64     `xorm:"not null pk autoincr BIGINT(64)"`
	Suid            int64     `xorm:"not null default 0 comment('买家用户id')  BIGINT(20)"`
	Ruid            int64     `xorm:"not null default 0 comment('卖家用户id')  BIGINT(20)"`
	DcType          int32     `xorm:"not null default 0 comment('dc类型') INT(11)"`
	ChangeType      int32     `xorm:"not null default 0 comment('改变类型')  INT(11)"`
	Price           int64     `xorm:"not null default 0 comment('付款价格单位( 分)') INT(11)"`
	Collection      int64     `xorm:"not null default 0 comment('收款价格单位( 分)') INT(11)"`
	NftId           int64     `xorm:"not null default 0 comment('nft ID') INT(11)"`
	WithDrawOrderId string    `xorm:"default '' comment('提取订单ID') index VARCHAR(100)"`
	TxHash          string    `xorm:"default '' comment('交易hash') index VARCHAR(100)"`
	NftType         string    `xorm:"default '' comment('nft类型') index VARCHAR(10)"`
	FromAddress     string    `xorm:"not null default '' comment('源地址') VARCHAR(100)"`
	ToAddress       string    `xorm:"not null default '' comment('接受地址') VARCHAR(100)"`
	Status          int32     `xorm:"not null default 0 comment('状态 0:发起 1:成功 2:失败') TINYINT(4)"`
	CreateTime      time.Time `xorm:"not null default '2016-08-31 04:25:00' comment('时间') DATETIME"`
	UpdateTime      time.Time `xorm:"not null default '2016-08-31 04:25:00' comment('时间') DATETIME"`
}

type QuestionCfg struct {
	Id         int32     `xorm:"not null pk autoincr comment('自增') INT(11)"`
	Pid        int32     `xorm:"not null default 0 comment('上级id') INT(11)"`
	RemarkZh   string    `xorm:"not null default '' comment('描述 简体中文') VARCHAR(50)"`
	RemarkZhHk string    `xorm:"not null default '' comment('描述 繁体中文') VARCHAR(50)"`
	RemarkEn   string    `xorm:"not null default '' comment('描述 英文') VARCHAR(100)"`
	RemarkVn   string    `xorm:"not null default '' comment('描述 越南') VARCHAR(100)"`
	RemarkBi   string    `xorm:"not null default '' comment('描述 印尼') VARCHAR(100)"`
	Status     int32     `xorm:"not null default 0 comment('状态（0=启用；1=禁用）') TINYINT(1)"`
	CreateTime time.Time `xorm:"not null default '2021-08-31 04:25:00' comment('创建时间') DATETIME"`
	UpdateTime time.Time `xorm:"not null default '2021-08-31 04:25:00' comment('修改时间') DATETIME"`
}

type Feedback struct {
	Id         int64     `xorm:"not null pk autoincr comment('自增') BIGINT(20)"`
	Uid        int64     `xorm:"not null default 0 comment('用户id') BIGINT(20)"`
	Model      int32     `xorm:"default 0 comment('问题场景id') INT(11)"`
	Type       int32     `xorm:"default 0 comment('问题类型id') INT(11)"`
	Remark     string    `xorm:"not null default '' comment('反馈及建议') VARCHAR(200)"`
	Images     string    `xorm:"default '' comment('图片（最多四张;号分割）') VARCHAR(500)"`
	CreateTime time.Time `xorm:"not null default '2022-08-31 04:25:00' comment('创建时间') DATETIME"`
}

// 联合查表需要把多个表结构合在一起
type FeedbackAllData struct {
	Feedback Feedback    `xorm:"extends"`
	Model    QuestionCfg `xorm:"extends"`
	Ques     QuestionCfg `xorm:"extends"`
}
