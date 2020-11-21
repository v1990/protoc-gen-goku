package helper

import (
	"fmt"
	"github.com/pkg/errors"
)

// 公共 管道 处理
func init() {
	RegisterFuncMap(map[string]interface{}{
		"Format": Format,
		"Set":    Set,
	})
}

type FormatPipeline interface {
	FormatPipeline(args Args) (string, error)
}

// 格式化为字符串
// 参数： (args ...interface{}, FormatPipeline )
//		  将 FormatPipeline 放在最后是为了适配 go template 的 pipeline 用法
// example: `$FormatPipeline|format $arg1 $arg2`  =>  `$FormatPipeline.Format($arg1,$arg2)`
func Format(_args ...interface{}) (string, error) {
	args := Args(_args)

	if args.Len() == 0 {
		return "", errors.New("no args")
	}

	switch obj := args.Last().(type) {
	case FormatPipeline:
		args.Pop()
		return obj.FormatPipeline(args)
	case string:
		args.Pop()
		return fmt.Sprintf(obj, args...), nil
	}

	return "", errors.New("can not find FormatPipeline")
}

type SetPipeline interface {
	SetPipeline(field string, value interface{}) interface{}
}

func Set(_args ...interface{}) interface{} {
	args := Args(_args)

	if args.Len() == 0 {
		Throws(errors.New("no args"))
		return nil
	}

	field := args.String(0, "")
	if field == "" {
		Throws(errors.New("function[set] first argument must be field name"))
		return nil
	}

	switch obj := args.Last().(type) {
	case SetPipeline:
		args.Pop()
		return obj.SetPipeline(field, args.Def(1, nil))
	}

	Throws(errors.New("unsupported SetPipeline"))
	return nil

}

func Throws(err error) {
	err = errors.WithStack(err)
	panic(err)
}
