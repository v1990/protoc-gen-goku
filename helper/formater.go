package helper

import (
	"errors"
	"fmt"
)

type Formatter interface {
	Format(_args ...interface{}) (string, error)
}

// 格式化为字符串
// 参数： (args ...interface{}, Formatter )
//		  将 Formatter 放在最后是为了适配 go template 的 pipeline 用法
// example: `$Formatter|format $arg1 $arg2`  =>  `$Formatter.Format($arg1,$arg2)`
func Format(_args ...interface{}) (string, error) {
	args := Args(_args)

	if args.Len() == 0 {
		return "", errors.New("no args")
	}

	switch obj := args.Last().(type) {
	case Formatter:
		args.Pop()
		return obj.Format(args...)
	case string:
		args.Pop()
		return fmt.Sprintf(obj, args...), nil
	}

	return "", errors.New("can not find Formatter")
}
