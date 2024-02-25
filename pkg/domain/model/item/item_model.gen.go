// Package item アイテム
package item

type Items []*Item

type Item struct {
	MasterItemId int64
	Count        int32
}

func NewItem() *Item {
	return &Item{}
}

func NewItems() Items {
	return Items{}
}

func SetItem(masterItemId int64, count int32) *Item {
	return &Item{
		MasterItemId: masterItemId,
		Count:        count,
	}
}
