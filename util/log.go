package util

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/lupc/go-myzap"
	"go.uber.org/zap"
)

var logger *zap.Logger

func GetLogger() *zap.Logger {
	if logger == nil {
		var myZapConfig = myzap.NewConfigByName("acs")
		myZapConfig.TimeFormat = "2006-01-02 15:04:05.000"
		logger = myZapConfig.BuildLogger().
			WithOptions(zap.WithCaller(false))
	}
	return logger
}

func LogRecoverToError(reason string) {
	if err := recover(); err != nil {
		var msg = fmt.Sprintf("panic:%v", reason)
		GetLogger().Error(msg, zap.Error(err.(error)))
	}
}

func RunWithRecover(panicReason string, action func()) {
	defer LogRecoverToError(panicReason)
	action()
}

func GoWithRecover(panicReason string, action func()) {
	go func() {
		RunWithRecover(panicReason, action)
	}()
}

func HighlightSuccess(msg string) string {
	colorFormat := color.New(color.FgWhite, color.BgGreen).SprintFunc()
	return colorFormat(" " + msg + " ")
}

func HighlightFail(msg string) string {
	colorFormat := color.New(color.FgWhite, color.BgRed).SprintFunc()
	return colorFormat(" " + msg + " ")
}

func HighlightConnected(msg string) string {
	colorFormat := color.New(color.FgGreen, color.BgWhite).SprintFunc()
	return colorFormat(" " + msg + " ")
}

func HighlightDisconnected(msg string) string {
	colorFormat := color.New(color.FgHiBlack, color.BgWhite).SprintFunc()
	return colorFormat(" " + msg + " ")
}

func HighlightError(msg string) string {
	colorFormat := color.New(color.FgRed).SprintFunc()
	return colorFormat(" " + msg + " ")
}
