// Copyright 2014 mqant Author. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package log 日志初始化
package log

import (
	beegolog "github.com/tofindme/mqant/log/beego"
)

var beego *beegolog.BeeLogger
var bi *beegolog.BeeLogger

// InitLog 初始化日志
func InitLog(debug bool, ProcessID string, Logdir string, settings map[string]interface{}) {
	beego = NewBeegoLogger(debug, ProcessID, Logdir, settings)
}

// InitBI 初始化BI日志
func InitBI(debug bool, ProcessID string, Logdir string, settings map[string]interface{}) {
	bi = NewBeegoLogger(debug, ProcessID, Logdir, settings)
}

// LogBeego LogBeego
func LogBeego() *beegolog.BeeLogger {
	if beego == nil {
		beego = beegolog.NewLogger()
	}
	return beego
}

// BiBeego BiBeego
func BiBeego() *beegolog.BeeLogger {
	return bi
}

// Debug Debug
func Debug(format string, a ...interface{}) {
	//gLogger.doPrintf(debugLevel, printDebugLevel, format, a...)
	LogBeego().Debug( format, a...)
}

// Info Info
func Info(format string, a ...interface{}) {
	//gLogger.doPrintf(releaseLevel, printReleaseLevel, format, a...)
	LogBeego().Info( format, a...)
}

// Error Error
func Error(format string, a ...interface{}) {
	//gLogger.doPrintf(errorLevel, printErrorLevel, format, a...)
	LogBeego().Error( format, a...)
}

// Warning Warning
func Warning(format string, a ...interface{}) {
	//gLogger.doPrintf(fatalLevel, printFatalLevel, format, a...)
	LogBeego().Warning( format, a...)
}


// Close Close
func Close() {
	LogBeego().Close()
}
