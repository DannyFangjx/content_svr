package logger

import (
	"content_svr/pub/requestid"
	"content_svr/pub/utils"
	"context"
	"fmt"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/sirupsen/logrus"
	"path/filepath"
	"strings"
	"time"
)

var logMap = map[logrus.Level]*logrus.Logger{}

const accessLogLev = 10001

func init() {
	wpath1, _ := utils.GetWorkPath()
	filepath.Abs(wpath1)
	s2 := filepath.Join(wpath1, "../")
	wpath, err := filepath.Abs(s2)
	if err != nil {
		panic(fmt.Sprintf("get workpath failed. err:%v", err.Error()))
	}
	path := wpath + "/log/"
	for _, lev := range logrus.AllLevels {
		fulPath := path + lev.String() + ".log"
		writer, _ := rotatelogs.New(
			fulPath+".%Y%m%d%H%M",
			rotatelogs.WithLinkName(fulPath),
			rotatelogs.WithMaxAge(time.Hour*24*7),     //最多保存30天日志。
			rotatelogs.WithRotationTime(time.Hour*24), //每1天分割一个日志。
		)

		logItem := &logrus.Logger{
			Out:       writer,
			Formatter: &logrus.TextFormatter{},
			Level:     lev,
		}
		//logItem.SetReportCaller(true)
		logMap[lev] = logItem
	}

	//增加access log
	fulPath := path + "access" + ".log"
	writer, _ := rotatelogs.New(
		fulPath+".%Y%m%d%H%M",
		rotatelogs.WithLinkName(fulPath),
		rotatelogs.WithMaxAge(time.Hour*24*7),
		rotatelogs.WithRotationTime(time.Hour*24),
	)
	logMap[accessLogLev] = &logrus.Logger{
		Out:       writer,
		Formatter: &logrus.TextFormatter{},
		Level:     logrus.TraceLevel,
	}
}

func getCtxInfo(ctx context.Context) string {
	return fmt.Sprintf("%v|%v|", requestid.GetRequestID(ctx), ctx.Value("ctx_user_id"))
}

func Warnf(ctx context.Context, format string, v ...interface{}) {
	logMap[logrus.WarnLevel].Warnf(fmt.Sprintf(getCtxInfo(ctx))+format, v...)
}

func Info(ctx context.Context, msg string) {
	logMap[logrus.InfoLevel].Infof(fmt.Sprintf(getCtxInfo(ctx))+"%v", msg)
}

func Infof(ctx context.Context, format string, v ...interface{}) {
	logMap[logrus.InfoLevel].Infof(fmt.Sprintf(getCtxInfo(ctx))+format, v...)
}

func Debug(ctx context.Context, msg string) {
	logMap[logrus.DebugLevel].Debugf(fmt.Sprintf(getCtxInfo(ctx))+"%v", msg)

}

func Debugf(ctx context.Context, format string, v ...interface{}) {
	logMap[logrus.DebugLevel].Debugf(fmt.Sprintf(getCtxInfo(ctx))+format, v...)
}

func Access(ctx context.Context, msg string) {
	msg = strings.Replace(msg, "\n", "", -1)
	logMap[accessLogLev].Tracef(fmt.Sprintf(getCtxInfo(ctx))+"%v", msg)
}

//
//func Data(ctx context.Context, msg string) {
//	msg = strings.Replace(msg, "\n", "", -1)
//	logger.WithMetaField(
//		splog.MetaRequestID, requestid.GetRequestID(ctx),
//	).Data(msg)
//}

func Error(ctx context.Context, msg string, err error) {
	errmsg := ""
	if err != nil {
		errmsg = err.Error()
	}
	logMap[logrus.ErrorLevel].Errorf(fmt.Sprintf(getCtxInfo(ctx))+"%v, err=%v", msg, errmsg)
}

func Errorf(ctx context.Context, format string, v ...interface{}) {
	logMap[logrus.ErrorLevel].Errorf(fmt.Sprintf(getCtxInfo(ctx))+format, v...)
}
