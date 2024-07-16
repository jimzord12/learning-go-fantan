package dungpkg

var ActiveDungeon *Dungeon

func DungeonInit(diff Difficulty) {
	ActiveDungeon = &Dungeon{
		Difficulty: diff,
	}
}
