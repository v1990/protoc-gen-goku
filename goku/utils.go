package goku

import (
	"fmt"
	"log"
	"os"
	"runtime/debug"
	"strings"
)

func (g *Generator) Debug(format string, args ...interface{}) {
	if !g.debug {
		return
	}
	msg := fmt.Sprintf(format, args...)

	log.Output(2, msg+"\n")
	//log.Println(msg)
}

func (g *Generator) Warn(format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	log.Output(2, "[WARN]"+msg+"\n")
}

// FatalOnErr 当 err!=nil 时报错并退出进程
func (g *Generator) FatalOnErr(err error, format string, args ...interface{}) {
	if err == nil {
		return
	}

	msg := fmt.Sprintf(format, args...)
	msg += "  \n  Cause: " + err.Error()
	msg += "\n" + string(debug.Stack())

	log.Output(2, msg+"\n")
	os.Exit(1)
}
func (g *Generator) Fatal(format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	log.Output(2, "Fatal: "+msg+"\n")
	os.Exit(1)
}

func (g *Generator) Recover(errPtr *error, msgs ...string) {
	r := recover()
	if r == nil {
		return
	}

	msg := "panic:" + strings.Join(msgs, " ")

	var err error
	switch rr := r.(type) {
	case error:
		err = fmt.Errorf("%s %w", msg, rr)
	default:
		err = fmt.Errorf("%s %#v", msg, r)
	}

	if errPtr != nil {
		*errPtr = err
		return
	}

	g.Debug("%s \n %s", err, string(debug.Stack()))
	g.FatalOnErr(err, "")

}

func (g *Generator) ThrowsOnErr(err error) {
	if err == nil {
		return
	}

	if g.debug {
		err = fmt.Errorf("%s \n %+v", err, err)
	}
	panic(err)

}
