package sms

type TencentConfig struct {
	SecretID   string `mapstructure:"secretID" json:"secretID" yaml:"secretID"`
	SecretKey  string `mapstructure:"secretKey" json:"secretKey" yaml:"secretKey"`
	SdkAppID   string `mapstructure:"sdkAppID" json:"sdkAppID" yaml:"sdkAppID"`
	Sign       string `mapstructure:"sign" json:"sign" yaml:"sign"`
	TemplateID string `mapstructure:"templateID" json:"templateID" yaml:"templateID"`
}

type VolcConfig struct {
	// 登录 https://console.volcengine.com/iam/keymanage 获取ak、sk
	AccessKey string `mapstructure:"accessKey" json:"accessKey" yaml:"accessKey"` // Access Key ID
	SecretKey string `mapstructure:"secretKey" json:"secretKey" yaml:"secretKey"` // Secret Access Key
	// 您可在该页面 https://console.volcengine.com/sms/subaccount/list?subAccountId=7331c4b7 查看消息组ID。
	SmsAccount string `mapstructure:"smsAccount" json:"smsAccount" yaml:"smsAccount"` // 消息组 ID
	Sign       string `mapstructure:"sign" json:"sign" yaml:"sign"`                   // 短信签名内容
	TemplateID string `mapstructure:"templateID" json:"templateID" yaml:"templateID"` // 短信模板 ID
}
