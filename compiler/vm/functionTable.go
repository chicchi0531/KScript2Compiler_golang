package vm

import (
	cm "ks2/compiler/common"
)

type Argument struct {
	Name      string
	VarType   *VariableTypeTag
	IsPointer bool
	ArraySize int
}

func MakeArgument(name string, vartype *VariableTypeTag, ispointer bool, arraysize int) *Argument {
	a := new(Argument)
	a.Name = name
	a.VarType = vartype
	a.IsPointer = ispointer
	a.ArraySize = arraysize
	return a
}

type FunctionTag struct {
	Name       string
	Args       []*Argument
	Address    int
	ReturnType *VariableTag
	Defined    bool //定義済みかどうか
}

func MakeFunctionTag(name string, args []*Argument, rettype *VariableTag, defined bool) *FunctionTag {
	f := new(FunctionTag)
	f.Name = name
	f.Args = args
	f.Address = 0
	f.ReturnType = rettype
	f.Defined = defined
	return f
}

// function table
type FunctionTable struct {
	Functions []*FunctionTag
	driver    *Driver
}

func MakeFunctionTable(d *Driver) *FunctionTable {
	f := new(FunctionTable)
	f.Functions = make([]*FunctionTag, 0)
	f.driver = d
	return f
}

func (t *FunctionTable) Add(tag *FunctionTag, lineno int) *FunctionTag {
	//宣言済みかチェック
	f := t.Find(tag.Name)
	if f != nil {
		t.driver.Err.LogError(t.driver.Filename, lineno, cm.ERR_0023, "関数："+tag.Name)
		return f
	}

	//call用のアドレスを設定
	if tag.Name == "main" {
		tag.Address = 0
	} else {
		tag.Address = t.driver.MakeLabel()
	}
	t.Functions = append(t.Functions, tag)

	return tag
}

// 同名の関数を探索
func (t *FunctionTable) Find(name string) *FunctionTag {
	for _, f := range t.Functions {
		if f.Name == name {
			return f
		}
	}
	return nil
}
