package levelsystem

func CalcRequiredExp(currentLevel int) float64 {
	if currentLevel < 1 {
		return 0
	}

	return float64(int(BaseExp) * (currentLevel * (currentLevel - 1) / 2))
}
