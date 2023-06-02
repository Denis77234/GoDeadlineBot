package dbinterface

import(
	"time"
	"katun/internal/database/member"
)

type Database interface {
	AddVkat(userId int64, userTag string, time time.Time) (int, error)
	DeleteVkat(userid int64) (error)
	UpdateVkat(userid int64, time time.Time) (error)
	Time(user int64) (string, string, error)
	VkatMembers() ([]member.Member, error)
}