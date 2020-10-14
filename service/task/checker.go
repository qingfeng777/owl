package task

type sqlChecker interface {
	SqlCheck(sql, charset, collation string, info *DBInfo) (pass bool, suggestion string, affectRow int)
	ListRules() interface{}
}

var checker sqlChecker

func setChecker(impl sqlChecker) {
	checker = impl
}
