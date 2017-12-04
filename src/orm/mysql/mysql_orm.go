package mysql

import (
	"sync"
	"database/sql"
	"errors"
	"orm"
)

type MySqlOrm struct{
	sync.RWMutex
	db *sql.DB
}

func (orm *MySqlOrm)Open(dataSrc string) error{
	db, err := sql.Open("mysql", dataSrc)
	if err != nil{
		panic(err)
		return errors.New("数据库连接错误.")
	}
	orm.db = db
	return nil
}

func (orm *MySqlOrm)Close(){
	if orm.db != nil{
		orm.db.Close()
		orm.db = nil
	}
}

func (orm *MySqlOrm)Add(object orm.Object) error{
	orm.Lock()
	defer orm.Unlock()

	stmt, err := orm.db.Prepare(object.GetAddSql())
	defer stmt.Close()

	if err != nil{
		return errors.New("sql prepare error.")
	}
	stmt.Exec(object.GetAddValues()...)

	return nil
}

func (orm *MySqlOrm)Del(object orm.Object) error{
	orm.Lock()
	defer orm.Unlock()

	stmt, err := orm.db.Prepare(object.GetDelSql())
	defer stmt.Close()

	if err != nil{
		return errors.New("sql prepare error.")
	}
	stmt.Exec(object.GetDelValues()...)

	return nil
}

func (orm *MySqlOrm)Save(object orm.Object) error{
	orm.Lock()
	defer orm.Unlock()

	stmt, err := orm.db.Prepare(object.GetSaveSql())
	defer stmt.Close()

	if err != nil{
		return errors.New("sql prepare error.")
	}
	stmt.Exec(object.GetSaveValues()...)

	return nil
}

func (orm *MySqlOrm)Load(object orm.Object) error{
	orm.Lock()
	defer orm.Unlock()

	row := orm.db.QueryRow(object.GetLoadSql(), object.GetLoadValues()...)
	row.Scan(object.GetLoadAddrs()...)

	return nil
}