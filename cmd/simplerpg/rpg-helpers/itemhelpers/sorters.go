package itemhelpers

import (
	"cmp"
	"slices"
	"strings"

	rpgitems "github.com/jimzord12/learning-go-fantan/cmd/simplerpg/models/rpg-items"
)

type Item = rpgitems.Item

// Sort weapons by their type (name)
func SortByType(items []*Item) {
	slices.SortStableFunc(items, func(a, b *Item) int {
		a_name := strings.Split(a.Name, " ")[1]
		b_name := strings.Split(b.Name, " ")[1]
		return cmp.Compare(a_name, b_name)
	})
}

func SortByValue(weapons []*Item) {
	slices.SortStableFunc(weapons, func(a, b *Item) int {
		return cmp.Compare(a.Value, b.Value)
	})
}

// Sort weapons by their material
func SortByMaterial(weapons []*Item) {
	slices.SortStableFunc(weapons, func(i, j *Item) int {
		a := strings.Split(i.Name, " ")[0]
		b := strings.Split(j.Name, " ")[0]
		return cmp.Compare(a, b)
	})
}
