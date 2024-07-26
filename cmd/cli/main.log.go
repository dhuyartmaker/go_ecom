package main

import "go.uber.org/zap"

func main() {
	sugar := zap.NewExample().Sugar()
	sugar.Infof("Hello name :%s", "Huygo")

	logger := zap.NewExample()
	logger.Info("Hello", zap.String("name", "TipsGo"))
}