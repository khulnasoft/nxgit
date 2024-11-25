package migrations

import (
	"github.com/go-xorm/xorm"
	"go.khulnasoft.com/nxgit/modules/util"
)

func addU2FReg(x *xorm.Engine) error {
	type U2FRegistration struct {
		ID          int64 `xorm:"pk autoincr"`
		Name        string
		UserID      int64 `xorm:"INDEX"`
		Raw         []byte
		Counter     uint32
		CreatedUnix util.TimeStamp `xorm:"INDEX created"`
		UpdatedUnix util.TimeStamp `xorm:"INDEX updated"`
	}
	return x.Sync2(&U2FRegistration{})
}
