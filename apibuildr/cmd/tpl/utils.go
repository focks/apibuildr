package tpl

func GetTestUtilTemplate() []byte {

	return []byte(`package cmd
import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func getTestingLogger() *zap.Logger {
	atom := zap.NewAtomicLevel()
	encoderCfg := zap.NewProductionEncoderConfig()
	encoderCfg.TimeKey = ""
	logger := zap.New(zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderCfg),
		zapcore.Lock(os.Stdout),
		atom,
	))
	defer logger.Sync()
	atom.SetLevel(zap.ErrorLevel)

	return logger
}
	`)

}
