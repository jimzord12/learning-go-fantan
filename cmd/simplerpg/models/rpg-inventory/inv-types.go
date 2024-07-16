package invpkg

import rpgitems "github.com/jimzord12/learning-go-fantan/cmd/simplerpg/models/rpg-items"

type Item = rpgitems.Item

type Inventory struct {
	Size  int
	Items []*Item
}

type Equipment struct {
	WeaponSlot    *Item
	ArmorSlot     *Item
	AccessorySlot *Item
}
