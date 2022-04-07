package dao

import (
	"Memo/dao/file"
	"Memo/dao/memory"
	"Memo/dao/user"
)

func InitDAOInst() {
	user.InitUserDBHandler()
	memory.InitDBHandler()
	file.InitStorageHandler()
}
