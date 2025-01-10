package weixin

import (
	"github.com/xxjwxc/public/mycache"
	wxpay "gopkg.in/go-with/wxpay.v1"
)

// UserInfo 用户信息
type UserInfo struct {
	OpenID     string   `json:"openid"`
	Nickname   string   `json:"nickname"`
	Sex        int32    `json:"sex"`
	Province   string   `json:"province"`
	City       string   `json:"city"`
	Country    string   `json:"country"`
	HeadImgURL string   `json:"headimgurl"`
	Privilege  []string `json:"privilege"`
	Unionid    string   `json:"unionid"`
}

// APITicket ...
type APITicket struct {
	Errcode   int    `json:"errcode"`
	Errmsg    string `json:"errmsg"`
	Ticket    string `json:"ticket"`
	ExpiresIn int    `json:"expires_in"`
}

// WxInfo 微信配置信息
type WxInfo struct {
	AppID     string // 微信公众平台应用ID
	AppSecret string // 微信支付商户平台商户号
	APIKey    string // 微信支付商户平台API密钥
	MchID     string // 商户号
	NotifyURL string // 通知地址
	ShearURL  string // 微信分享url
}

// TempMsg 订阅消息头
type TempMsg struct {
	Touser     string                       `json:"touser"`      //	是	接收者（用户）的 openid
	TemplateID string                       `json:"template_id"` //	是	所需下发的模板消息的id
	Page       string                       `json:"page"`        //	否	点击模板卡片后的跳转页面，仅限本小程序内的页面。支持带参数,（示例index?foo=bar）。该字段不填则模板无跳转。
	Data       map[string]map[string]string `json:"data"`        //是	模板内容，不填则下发空模板
}

// TempMsg 订阅消息头
type TempWebMsg struct {
	Touser     string                       `json:"touser"`      //	是	接收者（用户）的 openid
	TemplateID string                       `json:"template_id"` //	是	所需下发的模板消息的id
	Page       string                       `json:"url"`         //	否	点击模板卡片后的跳转页面，仅限本小程序内的页面。支持带参数,（示例index?foo=bar）。该字段不填则模板无跳转。
	Data       map[string]map[string]string `json:"data"`        //是	模板内容，不填则下发空模板
}

// ResTempMsg 模版消息返回值
type ResTempMsg struct {
	Errcode int    `json:"errcode"` //
	Errmsg  string `json:"errmsg"`
}

type wxTools struct {
	client     *wxpay.Client
	wxInfo     WxInfo
	certFile   string // 微信支付商户平台证书路径
	keyFile    string
	rootcaFile string
	cache      mycache.CacheIFS
}

// QrcodeRet ...
type QrcodeRet struct {
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
}

type wxPostdata struct {
	Scene string `json:"scene"`
	Page  string `json:"page"`
}

type wxQrcodedata struct {
	Path  string `json:"path"`  //路径
	Width int    `json:"width"` //二维码宽度
}

type AccessToken struct {
	AccessToken  string `json:"access_token"`  //	网页授权接口调用凭证,注意：此access_token与基础支持的access_token不同
	ExpiresIn    int    `json:"expires_in"`    // access_token接口调用凭证超时时间，单位（秒）
	RefreshToken string `json:"refresh_token"` //	用户刷新access_token
	Openid       string `json:"openid"`        // 用户唯一标识，请注意，在未关注公众号时，用户访问公众号的网页，也会产生一个用户和公众号唯一的OpenID
	Scope        string `json:"scope"`         // 用户授权的作用域，使用逗号（,）分隔
}

// WxUserinfo 微信用户信息
type WxUserinfo struct {
	Openid     string `json:"openid"`     // 微信用户唯一标识符,,微信用户唯一标识符
	NickName   string `json:"nickname"`   // 用户昵称
	Sex        int    `json:"sex"`        // 用户的性别
	City       string `json:"city"`       // 用户所在城市
	Province   string `json:"province"`   // 用户所在省份
	Country    string `json:"country"`    // 用户所在国家
	Headimgurl string `json:"headimgurl"` // 头像地址
	Unionid    string `json:"unionid"`    // 用户全局唯一标识
	// Privilege  []string `json:"privilege"`  // 户特权信息，json 数组，如微信沃卡用户为（chinaunicom）
}

type WxMenu struct {
	Button []WxMenuButton `json:"button"`
}

type WxMenuButton struct {
	Type      string      `json:"type,omitempty"`
	Name      string      `json:"name,omitempty"`
	Key       string      `json:"key,omitempty"`
	Url       string      `json:"url,omitempty"`
	SubButton []SubButton `json:"sub_button"`
}

type SubButton struct {
	Type string `json:"type,omitempty"`
	Name string `json:"name,omitempty"`
	Key  string `json:"key,omitempty"`
	Url  string `json:"url,omitempty"`
}

// CustomMsg 客服消息头
type CustomMsg struct {
	Touser  string       `json:"touser"`          //	是	接收者（用户）的 openid
	Msgtype string       `json:"msgtype"`         //	是	所需下发的模板消息的id
	Text    *CustomText  `json:"text,omitempty"`  //	文本类容
	Voice   *CustomVoice `json:"voice,omitempty"` //	语音
	Music   *CustomMusic `json:"music,omitempty"` //	音乐消息
	Image   *CustomImage `json:"image,omitempty"` //	图片消息
}
type CustomImage struct {
	MediaId string `json:"media_id"` //	图片
}

type CustomText struct {
	Content string `json:"content"` //	文本类容
}

type CustomVoice struct {
	MediaId string `json:"media_id"` //	语音
}

type CustomMusic struct {
	Title        string `json:"title"`          // 标题
	Description  string `json:"description"`    // 描述
	MusicUrl     string `json:"musicurl"`       // 链接
	HQMusicUrl   string `json:"hqmusicurl"`     // 链接
	ThumbMediaId string `json:"thumb_media_id"` // 缩略图
}

type GuideConfig struct {
	GuideAccount       string               `json:"guide_account"`         // 顾问号
	IsDelete           bool                 `json:"is_delete"`             // 操作类型，false表示添加 true表示删除
	GuideFastReplyList []GuideFastReplyList `json:"guide_fast_reply_list"` // 	快捷回复列表
	GuideAutoReply     GuideAutoReply       `json:"guide_auto_reply"`      //	第一条新客户关注自动回复
	GuideAutoReplyPlus GuideAutoReply       `json:"guide_auto_reply_plus"` //	第二条新客户关注自动回复
}

// GuideFastReplyList 快捷回复列表
type GuideFastReplyList struct {
	Content string `json:"content"` // 快捷回复
}

type GuideAutoReply struct {
	Content string `json:"content"` // 新客户关注自动回复内容,图片填mediaid,获取方式同图片素材,小程序卡片填下面请求demo中字段的json格式
	Msgtype int    `json:"msgtype"` // 1表示文字，2表示图片，3表示小程序卡片
}

// WxJsSign
type WxJsSign struct {
	Appid       string `json:"appid"`
	Noncestr    string `json:"noncestr"`
	Timestamp   string `json:"timestamp"`
	Url         string `json:"url"`
	Signature   string `json:"signature"`
	JsapiTicket string `json:"jsapi_ticket"`
}

type WxGetUser struct {
	Total int64         `json:"total"`
	Count int64         `json:"count"`
	Data  WxGetUserData `json:"data"`
}

type WxGetUserData struct {
	Openid     []string `json:"openid"`
	NextOpenid string   `json:"next_openid"`
}

type FreepublishiInfo struct { // 公众号文章
	ArticleId        string `json:"article_id"`         // 成功发布的图文消息id
	Title            string `json:"title"`              // 文章标题
	Author           string `json:"author"`             // 作者
	Digest           string `json:"digest"`             // 摘要
	ContentSourceUrl string `json:"content_source_url"` // 图文消息的原文地址，即点击“阅读原文”后的URL
	Url              string `json:"url"`                // 图文消息的URL
	ThumbMediaId     string `json:"thumb_media_id"`     //  图文消息的封面图片素材id（一定是永久MediaID）
	IsDeleted        bool   `json:"is_deleted"`         // 该图文是否被删除
	UpdateTime       int64  `json:"update_time"`        // 更新时间
}

type FreepublishiInfoReq struct {
	Offset    int64 `json:"offset"`
	Count     int64 `json:"count"`
	NoContent int64 `json:"no_content"`
}

type FreepublishiInfoResp struct {
	TotalCount int                    `json:"total_count"`
	ItemCount  int                    `json:"item_count"`
	Item       []FreepublishiInfoItem `json:"item"`
}

type FreepublishiInfoItem struct {
	ArticleId  string                  `json:"article_id"`  // 成功发布的图文消息id
	UpdateTime int64                   `json:"update_time"` // 更新时间
	Content    FreepublishiInfoContent `json:"content"`
}

type FreepublishiInfoContent struct {
	NewsItem []FreepublishiInfoNewItem `json:"news_item"` // 成功发布的图文消息id
}

type FreepublishiInfoNewItem struct {
	Title            string `json:"title"`              // 文章标题
	Author           string `json:"author"`             // 作者
	Digest           string `json:"digest"`             // 再要
	ContentSourceUrl string `json:"content_source_url"` // 图文消息的原文地址，即点击“阅读原文”后的URL
	Url              string `json:"url"`                // 图文消息的URL
	ThumbMediaId     string `json:"thumb_media_id"`     //  图文消息的封面图片素材id（一定是永久MediaID）
	IsDeleted        bool   `json:"is_deleted"`         // 该图文是否被删除
}

type MediaIdReq struct {
	MediaId string `json:"media_id"` // 永久素材id
}

type MediaResp struct {
	Title     string     `json:"title"`     // 素材标题
	DownUrl   string     `json:"down_url"`  // 素材下载地址
	NewsItemS []NewsItem `json:"news_item"` //
}

type NewsItem struct {
	Title            string `json:"title"`              // 素材标题
	Url              string `json:"url"`                // 图文页的URL
	ContentSourceUrl string `json:"content_source_url"` // 图文消息的原文地址，即点击“阅读原文”后的URL
}

type GetblacklistReq struct {
	BeginOpenid string `json:"begin_openid"` // 开始的openid
}

type GetblacklistResp struct {
	Total      int64      `json:"total"` // 总量
	Count      int64      `json:"count"` // 素材下载地址
	Data       OpenIDData `json:"data"`  //
	NextOpenid string     `json:"next_openid"`
}

type OpenIDData struct {
	Openid []string `json:"openid"`
}

type UploadTmpFileResp struct {
	Tp        string `json:"type"`
	MediaId   string `json:"media_id"`
	CreatedAt int64  `json:"created_at"`
}

// WxUserinfo 微信用户信息
type WxPhoneResp struct {
	Errcode     int         `json:"errcode"`
	Errmsg      string      `json:"errmsg"`
	WxPhoneinfo WxPhoneinfo `json:"phone_info"`
}

// WxUserinfo 微信用户信息
type WxPhoneinfo struct {
	PhoneNumber     string `json:"phoneNumber"`
	PurePhoneNumber string `json:"purePhoneNumber"`
	CountryCode     int    `json:"countryCode"`
}

type H5QrcodeReq struct {
	ExpireSeconds int        `json:"expire_seconds"` // 该二维码有效时间，以秒为单位。 最大不超过2592000（即30天），此字段如果不填，则默认有效期为60秒。
	ActionName    string     `json:"action_name"`    // 二维码类型，QR_SCENE为临时的整型参数值，QR_STR_SCENE为临时的字符串参数值，QR_LIMIT_SCENE为永久的整型参数值，QR_LIMIT_STR_SCENE为永久的字符串参数值
	SceneId       int        `json:"scene_id"`       // 场景值ID，临时二维码时为32位非0整型，永久二维码时最大值为100000（目前参数只支持1--100000
	ActionInfo    ActionInfo `json:"action_info"`    // 二维码详细信息
}

type ActionInfo struct {
	Scene ActionScene `json:"scene"`
}

type ActionScene struct {
	SceneStr string `json:"scene_str"` // 场景值ID（字符串形式的ID），字符串类型，长度限制为1到64
}

type H5QrcodeResp struct {
	Ticket        string `json:"ticket"`         // 获取的二维码ticket，凭借此ticket可以在有效时间内换取二维码。
	ExpireSeconds int    `json:"expire_seconds"` // 该二维码有效时间，以秒为单位。 最大不超过2592000（即30天），此字段如果不填，则默认有效期为60秒。
	URL           string `json:"url"`            // 二维码图片解析后的地址，开发者可根据该地址自行生成需要的二维码图片
}
