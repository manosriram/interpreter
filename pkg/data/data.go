package data

import "interpreter/pkg/file"

type I_Context struct {
	Start   int32
	Current int32
	Line    int32
	Tokens  []*Token
}

func New_IContext(start int32, current int32, line int32) *I_Context {
	return &I_Context{
		Start:   start,
		Current: current,
		Line:    line,
		Tokens:  make([]*Token, 0),
	}
}

type GlobalCtx struct {
	F   *file.File
	Ctx *I_Context
}

const (
	INT32 = 32
	INT64 = 64
)
