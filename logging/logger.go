package logging

import (
	"go.uber.org/zap"
	"os"
	"os/user"
	"path"
)

// TODO: zap allows to create Tree of Loggers
//  e.g. : for each workspace create a Namespace'd loggers with 1 main Logger on top

var (
	DEBUG     = os.Getenv("DEBUG")
	zapConfig zap.Config
)

// 			namespace string
func Logger() *zap.SugaredLogger {
	hd := path.Join(homeDir(), ".inspr")
	err := os.MkdirAll(hd, os.ModePerm)
	if err != nil {
		panic(err)
	}
	logPath := path.Join(hd, "log")
	zapConfig = zap.NewDevelopmentConfig()

	zapConfig.OutputPaths = []string{"stdout", logPath}
	zapConfig.ErrorOutputPaths = []string{"stderr", logPath}
	zapConfig.Development = false

	logger, err := zapConfig.Build()
	if err != nil {
		panic(err)
	}
	defer logger.Sync()

	if DEBUG == "" {
		zapConfig.Level.SetLevel(zap.InfoLevel)
	}

	//logger.With(zap.Namespace(namespace))
	return logger.Sugar()
}

func homeDir() string {
	usr, err := user.Current()
	if err != nil {
		panic(err)
	}
	return usr.HomeDir
}
