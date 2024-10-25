package handler

import (
	"monutil/common/domain"
	"monutil/server/logger"
	"reflect"
	"time"

	"github.com/mitchellh/mapstructure"

	"github.com/go-netty/go-netty"
)

type DoHandler struct{}

// 业务处理函数
func (DoHandler) HandleRead(ctx netty.InboundContext, message netty.Message) {
	// TODO [msg_type]: bussniss data
	msg := message.(map[string]interface{})
	if msg["msgType"] == "metrics" {
		// monstat := msg["msg"]
		var monstat domain.MonStat
		err := msdecode(msg, &monstat)
		if err != nil {
			logger.Errorf("监控字典转码失败:%s", err.Error())
		}
		logger.Infof("收到监测数据：%v", monstat)
		// TODO: 将monstat发送至需要的地方。
		// dofunction(monstat){}
	} else {
		ctx.HandleRead(message)
	}
}

func toTimeHookFunc() mapstructure.DecodeHookFunc {
	return func(
		f reflect.Type,
		t reflect.Type,
		data interface{}) (interface{}, error) {
		if t != reflect.TypeOf(time.Time{}) {
			return data, nil
		}

		switch f.Kind() {
		case reflect.String:
			return time.Parse(time.RFC3339, data.(string))
		case reflect.Float64:
			return time.Unix(0, int64(data.(float64))*int64(time.Millisecond)), nil
		case reflect.Int64:
			return time.Unix(0, data.(int64)*int64(time.Millisecond)), nil
		default:
			return data, nil
		}
		// Convert it by parsing
	}
}

func msdecode(input map[string]interface{}, result interface{}) error {
	decoder, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
		Metadata: nil,
		DecodeHook: mapstructure.ComposeDecodeHookFunc(
			toTimeHookFunc()),
		Result: result,
	})
	if err != nil {
		return err
	}

	if err := decoder.Decode(input); err != nil {
		return err
	}
	return err
}
