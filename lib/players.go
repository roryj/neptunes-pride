package lib

type PlayerIntel struct {
	TotalStars int `json:"ts"`
	TotalEconomy int `json:"e"`
	TotalIndustry int `json:"i"`
	TotalScience int `json:"s"`
	TotalShips int `json:"sh"`
	WeaponsTech int `json:"wt"`
	BankingTech int `json:"bt"`
	ManufacturingTech int `json:"mt"`
	HyperspaceTech int `json:"ht"`
	ScanningTech int `json:"st"`
	GeneralResearchTech int `json:"gt"`
	TerraformingTech int `json:"tt"`
	UID int `json:"uid"`
}

type Player struct {
	Uid int
	Alias string
	PlayerName string
	PlayerIntel
}

func GetAllPlayers() []*Player {
	var allPlayers []*Player

	allPlayers = append(allPlayers, &Player{
		Uid: 0,
		Alias: "Kodos",
		PlayerName: "Jordy",
	})
	allPlayers = append(allPlayers, &Player{
		Uid: 1,
		Alias: "ManOfSax",
		PlayerName: "Zack",
	})
	allPlayers = append(allPlayers, &Player{
		Uid: 2,
		Alias: "Lords of Kobol",
		PlayerName: "Matt",
	})
	allPlayers = append(allPlayers, &Player{
		Uid: 3,
		Alias: "Soggy Sea Dongs",
		PlayerName: "Elliot",
	})
	allPlayers = append(allPlayers, &Player{
		Uid: 4,
		Alias: "Slaybraham",
		PlayerName: "Edd",
	})
	allPlayers = append(allPlayers, &Player{
		Uid: 5,
		Alias: "mr worldwide",
		PlayerName: "Rory",
	})
	allPlayers = append(allPlayers, &Player{
		Uid: 6,
		Alias: "Taint",
		PlayerName: "Nick",
	})
	allPlayers = append(allPlayers, &Player{
		Uid: 7,
		Alias: "Sir Hotdogsa Sandwich",
		PlayerName: "Casey",
	})

	return allPlayers
}
