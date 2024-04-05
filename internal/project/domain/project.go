package project

type Project struct {
	title, description string
	members            []Member
	tasks              []Task
}

func NewProject(title string, description string) (*Project, error) {

	return &Project{
		title:       title,
		description: description,
	}, nil
}

func AddMembers(project Project, members []Member) (*Project, error) {
	if len(project.members) == 0 {
		return &Project{
			members: members,
		}, nil
	} else {
		newMembersList := append(project.members, members...)
		return &Project{
			members: newMembersList,
		}, nil
	}
}
