package fbook

type Friends []*User

func (f Friends) IDs() []string {
	var ids []string

	for _, user := range f {
		ids = append(ids, user.ID)
	}

	return ids
}
