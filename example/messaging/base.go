package messaging

import (
	"github.com/tas1999/tuya-connector-go/connector/constant"
	"github.com/tas1999/tuya-connector-go/connector/env/extension"
	"github.com/tas1999/tuya-connector-go/connector/logger"
	"github.com/tas1999/tuya-connector-go/connector/message/event"
)

func Listener() {
	extension.GetMessage(constant.TUYA_MESSAGE).InitMessageClient()

	extension.GetMessage(constant.TUYA_MESSAGE).SubEventMessage(func(m *event.NameUpdateMessage) {
		logger.Log.Info("=========== name update： ==========")
		logger.Log.Info(m)
	})

	extension.GetMessage(constant.TUYA_MESSAGE).SubEventMessage(func(m *event.StatusReportMessage) {
		logger.Log.Info("=========== report data： ==========")
		for _, v := range m.Status {
			logger.Log.Info(v.Code, v.Value)
		}
	})
}
