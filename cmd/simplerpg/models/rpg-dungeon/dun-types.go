package dungpkg

type Difficulty int

const (
	WOOD_DIF Difficulty = iota + 1
	BRONZE_DIF
	IRON_DIF
	STEEL_DIF
	TITANIUM_DIF
	MYTHRIL_DIF
)

type Dungeon struct {
	Difficulty
}
