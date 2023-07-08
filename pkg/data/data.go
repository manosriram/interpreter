package data

type I_Context struct {
	Start   int32
	Current int32
	Line    int32
}

func New_IContext(start int32, current int32, line int32) *I_Context {
	return &I_Context{
		Start:   start,
		Current: current,
		Line:    line,
	}
}
