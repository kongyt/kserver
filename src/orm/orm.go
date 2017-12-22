package orm

import (
	"reflect"
	"sync"
	"database/sql"
	_"github.com/go-sql-driver/mysql"
	"errors"
	"bytes"
)

type DBEntityInfo struct{
	TableName               string      // 该结构体对应的数据库表名称

	PrimaryKeyFieldIdx      int         // 数据库主键字段对应的结构体域索引
	PrimaryKeyName			string      // 数据库主键字段名称
	PrimaryKeyAutoIncrement	bool        // 主键是否自增

	FieldIndex              []int       // 数据库非主键字段对应的结构体域索引
	FieldNames              []string    // 数据库非主键字段名称

	AddSql                  string      // 增sql语句，注册时生成
	DelSql                  string      // 删sql语句，注册时生成
	SaveSql                 string      // 改sql语句，注册时生成
	LoadSql                 string      // 查sql语句，注册时生成
}

// 返回非主键域变量地址
func (info *DBEntityInfo)GetFieldAddress(entity interface{}) []interface{}{
	values := make([]interface{}, len(info.FieldNames))
	s := reflect.ValueOf(entity).Elem()
	for i := 0; i < len(info.FieldIndex); i++{
		values[i] = s.Field(info.FieldIndex[i]).Addr().Interface()
	}
	return values
}

// 返回主键域变量地址
func (info *DBEntityInfo)GetPrimaryKeyAddress(entity interface{}) interface{}{
	s := reflect.ValueOf(entity).Elem()
	return s.Field(info.PrimaryKeyFieldIdx).Addr().Interface()
}

func (info *DBEntityInfo)GenerateAddSql(){
	sqlBuf := bytes.Buffer{}
	sqlBuf.WriteString("INSERT INTO ")
	sqlBuf.WriteString(info.TableName)
	sqlBuf.WriteString(" (")
	if info.PrimaryKeyAutoIncrement == false{
		sqlBuf.WriteString(info.PrimaryKeyName)
		sqlBuf.WriteString(",")
	}

	for i := 0; i < len(info.FieldNames) -1; i++{
		sqlBuf.WriteString(info.FieldNames[i])
		sqlBuf.WriteString(",")
	}
	sqlBuf.WriteString(info.FieldNames[len(info.FieldNames)-1])
	sqlBuf.WriteString(") VALUES (")
	if info.PrimaryKeyAutoIncrement == false{
		sqlBuf.WriteString("?,")
	}
	for i := 0; i < len(info.FieldNames) -1; i++{
		sqlBuf.WriteString("?,")
	}
	sqlBuf.WriteString("?)")

	info.AddSql = sqlBuf.String()
}

func (info *DBEntityInfo)GenerateDelSql(){
	sqlBuf := bytes.Buffer{}
	sqlBuf.WriteString("DELETE FROM ")
	sqlBuf.WriteString(info.TableName)
	sqlBuf.WriteString(" WHERE ")
	sqlBuf.WriteString(info.PrimaryKeyName)
	sqlBuf.WriteString(" = ?")

	info.DelSql = sqlBuf.String()
}

func (info *DBEntityInfo)GenerateSaveSql(){
	sqlBuf := bytes.Buffer{}
	sqlBuf.WriteString("UPDATE ")
	sqlBuf.WriteString(info.TableName)
	sqlBuf.WriteString(" SET ")
	for i := 0; i < len(info.FieldNames) -1; i++{
		sqlBuf.WriteString(info.FieldNames[i])
		sqlBuf.WriteString(" = ?, ")
	}
	sqlBuf.WriteString(info.FieldNames[len(info.FieldNames)-1])
	sqlBuf.WriteString(" = ? WHERE ")
	sqlBuf.WriteString(info.PrimaryKeyName)
	sqlBuf.WriteString(" = ?")

	info.SaveSql = sqlBuf.String()
}

func (info *DBEntityInfo)GenerateLoadSql(){
	sqlBuf := bytes.Buffer{}
	sqlBuf.WriteString("SELECT ")
	for i := 0; i < len(info.FieldNames) -1; i++{
		sqlBuf.WriteString(info.FieldNames[i])
		sqlBuf.WriteString(", ")
	}
	sqlBuf.WriteString(info.FieldNames[len(info.FieldNames)-1])
	sqlBuf.WriteString(" FROM ")
	sqlBuf.WriteString(info.TableName)
	sqlBuf.WriteString(" WHERE ")
	sqlBuf.WriteString(info.PrimaryKeyName)
	sqlBuf.WriteString(" = ?")

	info.LoadSql = sqlBuf.String()
}

type Orm struct{
	sync.RWMutex
	db *sql.DB
	dbEntitiesInfo	map[reflect.Type] *DBEntityInfo
}


func (orm *Orm)Open(driverName string, dataSrc string) error{
	db, err := sql.Open(driverName, dataSrc)
	if err != nil{
		panic(err)
		return errors.New("database connect failed")
	}
	orm.db = db
	return nil
}

func (orm *Orm)Close(){
	if orm.db != nil{
		orm.db.Close()
		orm.db = nil
	}
}

// 注册
func (orm *Orm)Register(entity interface{}) error{
	s := reflect.TypeOf(entity).Elem() //通过反射获取type定义

	_, ok := orm.dbEntitiesInfo[reflect.TypeOf(entity)]
	if ok{
		return errors.New("entity type already register")
	}

	dbEntityInfo := new(DBEntityInfo)
	alreadyPk := false

	for i := 0; i < s.NumField(); i++ {
		tag := s.Field(i).Tag
		fieldName, ok := tag.Lookup("field")	// 可持久化字段
		if ok{
			index, ok := tag.Lookup("index")	// 可索引字段
			if ok{
				if index == "pk"{	// 索引类型为主键
					if alreadyPk == false{ // 主键重定义检查
						alreadyPk = true
						dbEntityInfo.PrimaryKeyFieldIdx = i
						dbEntityInfo.PrimaryKeyName = fieldName

						// 判断主键是否自增
						ai, ok := tag.Lookup("auto")
						if ok {
							if ai == "true"{
								dbEntityInfo.PrimaryKeyAutoIncrement = true
							}else{
								dbEntityInfo.PrimaryKeyAutoIncrement = false
							}
						}else{
							dbEntityInfo.PrimaryKeyAutoIncrement = false
						}

						// 判断是否设置表名
						tableName, ok := tag.Lookup("table")
						if ok{
							dbEntityInfo.TableName = tableName
						}else{
							return errors.New("pk field don't have table name")
						}
					}else{
						return errors.New("multiple definitions of primary key")
					}
				}
			}else{	// 普通字段
				dbEntityInfo.FieldIndex = append(dbEntityInfo.FieldIndex, i)
				dbEntityInfo.FieldNames = append(dbEntityInfo.FieldNames, fieldName)
			}
		}
	}

	if !alreadyPk {
		return errors.New("not set primary key")
	}

	dbEntityInfo.GenerateAddSql()
	dbEntityInfo.GenerateDelSql()
	dbEntityInfo.GenerateSaveSql()
	dbEntityInfo.GenerateLoadSql()

	orm.dbEntitiesInfo[reflect.TypeOf(entity)] = dbEntityInfo

	return nil
}

//增加对象
func (orm *Orm)Add(entity interface{}) error{
	orm.Lock()
	defer orm.Unlock()
	dbEntityInfo, ok := orm.dbEntitiesInfo[reflect.TypeOf(entity)]
	if !ok{
		return errors.New("entity type not register")
	}

	stmt, err := orm.db.Prepare(dbEntityInfo.AddSql)
	defer stmt.Close()

	if err != nil{
		return err
	}

	if dbEntityInfo.PrimaryKeyAutoIncrement == false{
		_, err := stmt.Exec(dbEntityInfo.GetPrimaryKeyAddress(entity), dbEntityInfo.GetFieldAddress(entity))
		if err != nil{
			panic(err)
		}
	}else {
		res, err := stmt.Exec(dbEntityInfo.GetFieldAddress(entity)...)
		if err != nil{
			panic(err)
		}

		// 返回增长ID
		lastId, err := res.LastInsertId()
		if err != nil{
			panic(err)
		}
		pkAddr := dbEntityInfo.GetPrimaryKeyAddress(entity).(*int)
		*pkAddr = (int)(lastId)

	}

	return nil
}

// 删除对象
func (orm *Orm)Del(entity interface{}) error{
	orm.Lock()
	defer orm.Unlock()

	dbEntityInfo, ok := orm.dbEntitiesInfo[reflect.TypeOf(entity)]
	if !ok{
		return errors.New("entity type not register")
	}

	stmt, err := orm.db.Prepare(dbEntityInfo.DelSql)
	defer stmt.Close()

	if err != nil{
		return errors.New("sql prepare error")
	}

	stmt.Exec(dbEntityInfo.GetPrimaryKeyAddress(entity))

	return nil
}

// 修改对象
func (orm *Orm)Save(entity interface{}) error{
	orm.Lock()
	defer orm.Unlock()

	dbEntityInfo, ok := orm.dbEntitiesInfo[reflect.TypeOf(entity)]
	if !ok{
		return errors.New("entity type not register")
	}

	stmt, err := orm.db.Prepare(dbEntityInfo.SaveSql)
	defer stmt.Close()

	if err != nil{
		return errors.New("sql prepare error")
	}

	args := dbEntityInfo.GetFieldAddress(entity)
	args = append(args, dbEntityInfo.GetPrimaryKeyAddress(entity))
	stmt.Exec(args...)

	return nil
}

// 查找对象
func (orm *Orm)Load(entity interface{}) error{
	orm.Lock()
	defer orm.Unlock()

	dbEntityInfo, ok := orm.dbEntitiesInfo[reflect.TypeOf(entity)]
	if !ok{
		return errors.New("entity type not register")
	}

	row := orm.db.QueryRow(dbEntityInfo.LoadSql, dbEntityInfo.GetPrimaryKeyAddress(entity))
	err := row.Scan(dbEntityInfo.GetFieldAddress(entity)...)

	return err
}

func NewOrm() *Orm{
	orm := new(Orm)
	orm.dbEntitiesInfo = make(map[reflect.Type] *DBEntityInfo)
	return orm
} 