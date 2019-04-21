package models

type BsSites struct {
	Siteid        int    `xorm:"not null pk autoincr INT(11)"`
	Url           string `xorm:"not null VARCHAR(255)"`
	Domains       string `xorm:"not null VARCHAR(255)"`
	Lastindex     int    `xorm:"not null default 0 comment('上次索引结束时间') index INT(11)"`
	Indexinterval int    `xorm:"not null default 3600 comment('索引间隔') index INT(11)"`
	Status        int    `xorm:"not null default -1 comment('-1未开始 0进行中 1已结束') index INT(1)"`
}
