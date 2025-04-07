package Models

import (
	"strings"
)

/* Goes through the entire Item Tree Set and masks all container and item information */
func MaskItemsTreeResult(ItemsTree *ItemsTreeResult) {

	RecurseContainers(&ItemsTree.Containers)

	/* Masks the Items Array at the root-level */
	if len(ItemsTree.Items) > 0 {
		for _, SelectedItem := range ItemsTree.Items {
			MaskItem(&SelectedItem)
		}
	}
}

/* Goes through every container recursively */
func RecurseContainers(SelectedContainer *[]Containers) {
	if len(*SelectedContainer) > 0 {
		for j, ChildContainer := range *SelectedContainer {
			(*SelectedContainer)[j] = MaskContainer(&ChildContainer)
			RecurseContainers(&ChildContainer.Containers)
		}
	}
}

/* Masks all data in a Container */
func MaskContainer(ChildContainer *Containers) Containers {
	ChildContainer.SerialNumber = MaskText(ChildContainer.SerialNumber)
	ChildContainer.Number = MaskText(ChildContainer.Number)
	ChildContainer.LotNum = MaskText(ChildContainer.LotNum)
	ChildContainer.Expiry = MaskText(ChildContainer.Expiry)

	/* Goes through the list of items and masks thier information */
	if len(ChildContainer.Items) > 0 {
		for i, ChildItem := range ChildContainer.Items {
			ChildContainer.Items[i] = MaskItem(&ChildItem)
		}
	}

	return *ChildContainer
}

/* Goes through the ItemSet Array and Masks all info */
func MaskItemList(ItemList []Items) []Items {
	for i, SelectedItem := range ItemList {
		ItemList[i] = MaskItem(&SelectedItem)
	}
	return ItemList
}

/* Masks all info in the Items */
func MaskItem(ChildItem *Items) Items {
	ChildItem.Gtin = MaskText(ChildItem.Gtin)
	ChildItem.LotNum = MaskText(ChildItem.LotNum)
	ChildItem.Expiry = MaskText(ChildItem.Expiry)
	ChildItem.Ndc = MaskText(ChildItem.Ndc)
	ChildItem.Upc = MaskText(ChildItem.Upc)
	ChildItem.SerialNumber = MaskText((ChildItem.SerialNumber))
	return *ChildItem
}

/* Masks all information presented with the string */
func MaskText(StringSet string) string {
	return strings.Repeat("*", len(StringSet))
}
