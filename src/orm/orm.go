package orm

type Object interface{
	GetAddSql() string
	GetAddValues()[]interface{}

	GetDelSql() string
	GetDelValues() []interface{}

	GetSaveSql() string
	GetSaveValues()[]interface{}

	GetLoadSql() string
	GetLoadValues()[]interface{}
	GetLoadAddrs() []interface{}
}

type Orm interface{
	Open(dataSrc string) error
	Close()
	Add(object *Object) error
	Del(object *Object) error
	Save(object *Object) error
	Load(object *Object) error
}
