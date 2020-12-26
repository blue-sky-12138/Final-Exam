package database

import "fmt"

type History struct {
	Sender string
	Time string
	Receiver string
	Value float64
	Remark string
}
type Histories []History

//查找历史记录
func FindAllHistory(histories *Histories,sender string) {
	var temHistory History
	prepare:=fmt.Sprintf("select sender,time,receiver,value,remark from history where sender='%s' or receiver='%s' order by id desc",sender,sender)
	stmt,err:= DataBase.Query(prepare)
	CheckError("FindHistory error",err)
	defer stmt.Close()
	for stmt.Next(){
		stmt.Scan(&temHistory.Sender,&temHistory.Time,&temHistory.Receiver,&temHistory.Value)
		*histories=append(*histories,temHistory)
	}
	fmt.Println(*histories)
}

//添加记录
func InsertHistory(sender string,time string,recevier string,value int,remark string){
	prepare:=fmt.Sprintf("insert history (sender,time,receiver,value,remark)value('%s','%s','%s',%d,'%s')",
		sender,time,recevier,value,remark)
	stmt,err:= DataBase.Prepare(prepare)
	defer stmt.Close()
	CheckError("InsertHistory error",err)
	stmt.Exec()
}