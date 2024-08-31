package note

type NoteService struct {
	Store NoteStorage
}

func NewService(s NoteStorage) *NoteService {
	return &NoteService{Store: s}
}
