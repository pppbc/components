package sms

import (
	"encoding/json"

	"components/errors"

	"github.com/sirupsen/logrus"
	"github.com/volcengine/volc-sdk-golang/service/sms"
)

type VolcSmsClient struct {
	Conf VolcConfig
}

// NewVolcSmsClient
func NewVolcSmsClient(conf VolcConfig) *VolcSmsClient {
	return &VolcSmsClient{Conf: conf}
}

// Send 发送验证码
func (s *VolcSmsClient) Send(mobile string, code string) error {
	// ak、sk
	sms.DefaultInstance.Client.SetAccessKey(s.Conf.AccessKey)
	sms.DefaultInstance.Client.SetSecretKey(s.Conf.SecretKey)

	// req body
	tp := TemplateParam{Code: code, Expiration: 10}
	tpj, _ := json.Marshal(tp)
	req := &sms.SmsRequest{
		SmsAccount:    s.Conf.SmsAccount,
		Sign:          s.Conf.Sign,
		TemplateID:    s.Conf.TemplateID,
		TemplateParam: string(tpj),
		PhoneNumbers:  mobile,
	}
	result, statusCode, err := sms.DefaultInstance.Send(req)
	if err != nil {
		logrus.Errorf("sendSms volcengineSDKError %s", err.Error())
		return err
	}
	if statusCode != 200 || result.ResponseMetadata.Error != nil {
		logrus.Errorf("sendSms json is :%+v, err: %+v, statusCode is:%d", result, result.ResponseMetadata.Error, statusCode)
		return errors.ErrThird
	}

	// 打印返回信息
	logrus.Debugf("sendSms json is :%+v", result)
	return nil
}

type TemplateParam struct {
	Code       string `json:"code"`       // 验证码
	Expiration int    `json:"expiration"` // 有效期，单位分钟
}
