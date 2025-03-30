package registries

type Identifiable struct {
	ID string
}


type InMemoryRegistry map[string]Identifiable


func (reg InMemoryRegistry) Get(id string) (Identifiable, bool) {
	note, err := reg[id]
	return note, err
}
func (reg InMemoryRegistry) Update(note Identifiable) Identifiable {
	reg[note.ID] = note
	return reg[note.ID]
}

func (req InMemoryRegistry) Delete(note Identifiable) Identifiable {
	deletedT := req[note.ID]
	delete(req, note.ID)
	return deletedT
}

