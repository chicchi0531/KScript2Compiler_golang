package ast

import(
	"ks2/compiler/vm"
)

type StateBlock struct{
	States []vm.IStatement
}

func (s *StateBlock) AddStates(n vm.IStatement) vm.IStateBlock{
	s.States = append(s.States, n)
	return s
}

func (s *StateBlock) Analyze(){
	for _, s := range s.States{
		if s!=nil{
			s.Analyze()
		}
	}
}
