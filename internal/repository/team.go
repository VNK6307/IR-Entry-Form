package repository

type TeamRepository interface {
	SaveTeam()
	//AddTeamMember()
	//RemoveTeamMember()
	//GetTeamMembers()
}

type Team struct {
	//teamName string
	members map[string][]string
}

func NewTeamRepository() *map[string][]string {
	members := make(map[string][]string)
	return &members
}

func (t Team) SaveTeam() {
	//TODO implement me
	panic("implement me")
}
