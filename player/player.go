package player

type Player struct {
	Username	string
	Skill_lvl	int
}

func CreatePlayer(username string) *Player {
	return &Player{
		Username: username,
		Skill_lvl: generateRandomSkillLevel(),
	} // Pick a random number for skill level
}

func generateRandomSkillLevel() int {
	return 1500
}	
