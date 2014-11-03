// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package pay3

// 统一支付接口 请求参数
type UnifiedOrderRequest map[string]string

// 设置签名字段.
//  appKey: 商户支付密钥Key
//
//  NOTE: 要求在 UnifiedOrderRequest 其他字段设置完毕后才能调用这个函数, 否则签名就不正确.
func (req UnifiedOrderRequest) SetSignature(appKey string) (err error) {
	return SetSignature(req, appKey)
}