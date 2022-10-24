package transfer

import (
	. "C-SquaredService/model"
	"fmt"
)

func MergeClassify(parentClassifyList []ParentClassify, childClassifyList []ChildClassify) []Classify {
	var classifyList = []Classify{}
	for _, parentClassify := range parentClassifyList {
		classify := Classify{
			Id:    parentClassify.Id,
			Name:  parentClassify.Name,
			Path:  parentClassify.Path,
			Child: []ChildClassify{},
		}
		for _, childClassify := range childClassifyList {
			if childClassify.ParentId == parentClassify.Id {
				newChildClassify := childClassify
				newChildClassify.Path = fmt.Sprintf("%v%v", parentClassify.Path, childClassify.Path)
				classify.Child = append(classify.Child, newChildClassify)
			}
		}
		classifyList = append(classifyList, classify)
	}
	return classifyList
}
