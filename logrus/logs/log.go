package logs

import "github.com/sirupsen/logrus"

type Log struct {
	*logrus.Entry
	//*logrus.Logger
	LogWriter
}

func (l *Log) Flush() {
	l.LogWriter.Flush()
}

type LogConf struct {
	Level       logrus.Level
	AdapterName string
}

func InitLog(conf LogConf) *Log {
	adapterName := "std"
	if conf.AdapterName != "" {
		adapterName = conf.AdapterName
	}
	writer, ok := writeAdapter[adapterName]
	if !ok {
		adapterName = "std"
		writer, _ = writeAdapter[adapterName]
	}
	log := &Log{
		logrus.NewEntry(logrus.New()),
		//logrus.New(),
		writer(),
	}
	log.Logger.SetOutput(log.LogWriter)
	if conf.Level != 0 {
		log.Logger.SetLevel(conf.Level)
	}
	log.Logger.SetFormatter(&logrus.JSONFormatter{})
	log.Logger.SetReportCaller(true)
	return log
}
