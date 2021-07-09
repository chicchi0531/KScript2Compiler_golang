package ast

import(
	"ks2/compiler/vm"
	cm "ks2/compiler/common"
)

const (
	OP_EQUAL = iota
	OP_GT
	OP_GE
	OP_LT
	OP_LE
	OP_NEQ
	OP_NOT
	OP_ADD
	OP_SUB
	OP_MUL
	OP_DIV
	OP_MOD
	OP_AND
	OP_OR
	OP_INCR
	OP_DECR

	OP_INTEGER
	OP_FLOAT
	OP_STRING
)


type Node struct {
	Lineno int
	Left    vm.INode
	Right   vm.INode
	Op      int
	Driver  *vm.Driver

	Ival    int
	Fval    float32
	Sval    string
}

func (n *Node) Push() int {

	// 単項演算の場合
	switch n.Op {
	case OP_INCR:
		t := n.Left.Push()
		if t == cm.TYPE_STRING{
			n._err(cm.ERR_0005,"")
			return cm.TYPE_INTEGER
		}
		n.Driver.OpPushInteger(1)
		n.Driver.OpAdd()
		n.Left.Pop()//一旦代入してから再度PUSH
		n.Left.Push()
		return t

	case OP_DECR:
		t := n.Left.Push()
		if t == cm.TYPE_STRING{
			n._err(cm.ERR_0006,"")
			return cm.TYPE_INTEGER
		}
		n.Driver.OpPushInteger(1)
		n.Driver.OpSub()
		n.Left.Pop()//一旦代入してから再度PUSH
		n.Left.Push()
		return t
		
	case OP_NOT:
		t := n.Left.Push()
		if t == cm.TYPE_STRING{
			n._err(cm.ERR_0007, "")
			return cm.TYPE_INTEGER
		}
		n.Driver.OpNot()
		return t

	// const node
	case OP_INTEGER:
		n.Driver.OpPushInteger(n.Ival)
		return cm.TYPE_INTEGER
	case OP_FLOAT:
		n.Driver.OpPushFloat(n.Fval)
		return cm.TYPE_FLOAT
	case OP_STRING:
		n.Driver.OpPushString(n.Sval)
		return cm.TYPE_STRING
	}

	// 二項演算の場合
	leftType := n.Left.Push()
	rightType := n.Right.Push()

	// 型チェック
	if leftType == cm.TYPE_STRING && rightType != cm.TYPE_STRING ||
		leftType != cm.TYPE_STRING && rightType == cm.TYPE_STRING{
		n._err(cm.ERR_0012, "")
	}

	// 文字列演算
	if leftType == cm.TYPE_STRING || rightType == cm.TYPE_STRING{
		switch n.Op{
		case OP_ADD:
			n.Driver.OpAddString()
		default:
			n._err(cm.ERR_0013, "")
		}
		return cm.TYPE_STRING
	}

	// 数値演算
	switch n.Op {
	case OP_EQUAL:
		n.Driver.OpEqual()
	case OP_NEQ:
		n.Driver.OpNequ()
	case OP_ADD:
		n.Driver.OpAdd()
	case OP_SUB:
		n.Driver.OpSub()
	case OP_MUL:
		n.Driver.OpMul()
	case OP_DIV:
		n.Driver.OpDiv()
	case OP_MOD:
		n.Driver.OpMod()
	case OP_GT:
		n.Driver.OpGt()
	case OP_GE:
		n.Driver.OpGe()
	case OP_LT:
		n.Driver.OpLt()
	case OP_LE:
		n.Driver.OpLe()
	case OP_AND:
		n.Driver.OpAnd()
	case OP_OR:
		n.Driver.OpOr()
	default:
		n._err(cm.ERR_0014, "")
	}

	// どちらかの項がfloatの場合は、float項にキャストする
	if leftType == cm.TYPE_FLOAT || rightType == cm.TYPE_FLOAT{
		return cm.TYPE_FLOAT
	}
	return cm.TYPE_INTEGER
}

func (n *Node) Pop() int {
	n._err(cm.ERR_0008, "")
	return -1
}

// 内部エラー出力用
func (n *Node) _err(errorcode string, submsg string){
	n.Driver.Err.LogError(n.Driver.Filename, n.Lineno, errorcode, submsg)
}
func (n *Node) _warning(warningcode string, submsg string){
	n.Driver.Err.LogWarning(n.Driver.Filename, n.Lineno, warningcode, submsg)
}
