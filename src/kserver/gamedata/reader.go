package gamedata

import (
	"github.com/kongyt/leaf/recordfile"
	"github.com/kongyt/leaf/log"
	"reflect"
)

func readRf(st interface{}) *recordfile.RecordFile{
	rf, err := recordfile.New(st)
	if err != nil{
		log.Fatal("%v", err)
	}
	fn := reflect.TypeOf(st).Name()+".txt"
	err = rf.Read("gamedata/"+fn)
	if err != nil{
		log.Fatal("%v: %v", fn, err)
	}
	return rf
}
