package logger

import "go.uber.org/zap"

var Default *zap.SugaredLogger

func init() {
	b, _ := zap.NewProduction()
	Default = b.Sugar()
	Default.Named("Default")
}
