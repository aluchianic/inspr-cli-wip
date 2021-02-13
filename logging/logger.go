package logging

import (
	"go.uber.org/zap"
	"os"
	"os/user"
	"path"
)

// 			namespace string
func Logger() *zap.Logger {
	hd := path.Join(homeDir(), ".inspr")
	err := os.MkdirAll(hd, os.ModePerm)
	if err != nil {
		panic(err)
	}
	logPath := path.Join(hd, "log")

	zapConfig := zap.NewDevelopmentConfig()

	zapConfig.OutputPaths = []string{"stdout", logPath}
	zapConfig.ErrorOutputPaths = []string{"stderr", logPath}

	logger, err := zapConfig.Build()
	if err != nil {
		panic(err)
	}
	defer logger.Sync()

	//logger.With(zap.Namespace(namespace))

	return logger
}

func homeDir() string {
	usr, err := user.Current()
	if err != nil {
		panic(err)
	}
	return usr.HomeDir
}
