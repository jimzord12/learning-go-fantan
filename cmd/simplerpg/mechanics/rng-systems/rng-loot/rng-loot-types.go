package rngloot

import rpgitems "github.com/jimzord12/learning-go-fantan/cmd/simplerpg/models/rpg-items"

type ItemType = rpgitems.ItemType
type Material = rpgitems.Material
type PotionType = rpgitems.PotionType

type EnemyDrops struct {
	EquipmentType     ItemType   // 0-2
	EquipmentMaterial Material   // 0-100
	PotionDrop        PotionType // 0-100

}
