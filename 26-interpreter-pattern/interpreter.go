package interpreter

import "strings"

/*
	设计思想：
		1. 表达式接口 Expression interface
		2. 具体的表达式类 ConcreteExpression struct
		3. 上下文类 Context struct
*/

type Exp int

const (
	Equ = iota
	Cont
)
/*创建Expression*/
type Expression interface {
	Interpret() bool
}

/*创建context结构， 用作需要解释的上下文信息*/
type Context struct {
	val string
}

func (con *Context) GetVal() string {
	return con.val
}
//Equal表达式类 implements expression
type Equal struct {
	left Context
	right Context
}

func (e *Equal) Interpret() bool {
	return e.left.GetVal() == e.right.GetVal()
}

//Contain表达式类, implements expression
type Contain struct {
	left Context
	right Context
}

func (con *Contain) Interpret() bool {
	return strings.Contains(con.left.GetVal(), con.right.GetVal())
}

func CreateExpression(exp Exp, left, right Context) Expression {
	switch exp {
	case Equ:
		return &Equal{
			left: left,
			right: right,
		}
	case Cont:
		return &Contain{
			left: left,
			right: right,
		}
	default:
		return nil
	}
}
