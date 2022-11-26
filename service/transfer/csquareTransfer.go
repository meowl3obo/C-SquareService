package transfer

import (
	. "C-SquaredService/model"
	"encoding/json"
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

func MergeProductInventory(product Product, inventorys []Inventory) ProductInventory {
	var otherImgs []string
	err := json.Unmarshal([]byte(product.OtherImg), &otherImgs)
	productInventory := ProductInventory{
		Id:             product.Id,
		Name:           product.Name,
		MainImg:        product.MainImg,
		OtherImg:       product.OtherImg,
		Intro:          product.Intro,
		Illustrate:     product.Illustrate,
		ParentClassify: product.ParentClassify,
		ChildClassify:  product.ChildClassify,
		Status:         product.Status,
		Price:          product.Price,
		IsNew:          product.IsNew,
		IsSale:         product.IsSale,
		Inventorys:     inventorys,
		OtherImgs:      []string{},
	}
	if err == nil {
		productInventory.OtherImgs = otherImgs
	}
	return productInventory
}
