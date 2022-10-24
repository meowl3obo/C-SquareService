package model

type Classify struct {
	Id    int
	Name  string
	Path  string
	Child []ChildClassify
}
