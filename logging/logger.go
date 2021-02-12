package logging

import (
	"go.uber.org/zap"
	"log"
	"os/user"
	"path"
)

// 			namespace string
func Logger() *zap.Logger {
	hd := path.Join(homeDir(), ".inspr")
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
		log.Fatal(err)
	}
	return usr.HomeDir
}
