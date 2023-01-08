package core

var SetType = &Type{
	Name:      "Set",
	Parent:    nil,
	Functions: nil,
}

func NewSetSubType(children []*Type) *Type {
	return &Type{
		Name:      "Set",
		Parent:    SetType,
		Functions: nil,
		Generic:   true,
		Generics:  children,
	}
}

type SetVariable struct {
	Children []*Pointer
}

func (*SetVariable) GetType() *Type {
	return SetType
}

func NewSetPointer(children []*Pointer) *Pointer {
	return &Pointer{
		Typ:      SetType,
		Variable: &SetVariable{Children: children},
	}
}
