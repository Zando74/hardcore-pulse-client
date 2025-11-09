package port

type PlayerHashRepository interface {
	SaveAll(playerhash []string) error
	Exist(playerhash string) bool
}