package repo

//SingleCurdRepo Default imp
type SingleCurdRepo struct {
	DefaultTableManager
	BaseRepo
}

//CreateNew create new
func CreateNew(config *DatabaseConfig, entity ...interface{}) (*SingleCurdRepo, error) {
	xp := &SingleCurdRepo{BaseRepo: BaseRepo{TableManager: &SingleCurdRepo{}}}
	return xp, xp.Initializate(config, entity...)
}
