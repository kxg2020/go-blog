package logs

import (
	"encoding/json"
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"log"
	"os"
	"time"
	"wechat/pkg/tool"
)

const FormatFile  = `{"level": "%s","encoding": "json","outputPaths": ["%s"],"errorOutputPaths": ["%s"]}`
const FormatStd   = `{"level": "%s","encoding": "json","outputPaths": ["stdout"],"errorOutputPaths": ["stdout"]}`
const LogPath     = "pkg/logs/"
var Zap *zap.Logger

func ZapLogInitialize(){
	zapInit(LogPath + time.Now().Format("2006/01"),"error",false)
}
func zapInit(path string,level string,debug bool) {
	var format,logFile string
	var config zap.Config
	var err    error
	if ret,_ := tool.DirOrFileExist(path);ret == false{
		err := os.MkdirAll(path,os.ModePerm)
		if err != nil {
			log.Fatal(err.Error())
		}
	}
	logFile = path + "/" + time.Now().Format("01") + ".log"
	if res,_ := tool.DirOrFileExist(logFile);res == false{
		_,err = os.Create(logFile)
		if err != nil {
			log.Fatal(err.Error())
		}
	}
	if debug {
		format = fmt.Sprintf(FormatStd, level)
	}else {
		format = fmt.Sprintf(FormatFile, level, logFile, logFile)
	}
	if err := json.Unmarshal([]byte(format), &config); err != nil {
		log.Fatal(err.Error())
	}
	config.EncoderConfig = zap.NewProductionEncoderConfig()
	config.EncoderConfig.EncodeTime = func(i time.Time, encoder zapcore.PrimitiveArrayEncoder) {
		encoder.AppendString(i.Format("2006-01-02 15:04:05"))
	}
	Zap, err = config.Build()
	if err != nil {
		log.Fatal(err.Error())
	}
}

