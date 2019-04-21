package models

type TestData struct {
	Id      int    `xorm:"not null pk autoincr INT(11)"`
	Title   string `xorm:"not null VARCHAR(255)"`
	Content string `xorm:"not null VARCHAR(255)"`
	Indexed int    `xorm:"not null default 0 index INT(1)"`
}
