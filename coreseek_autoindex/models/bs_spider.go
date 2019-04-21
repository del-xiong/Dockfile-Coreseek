package models

type BsSpider struct {
	Id      int    `xorm:"not null pk autoincr INT(11)"`
	Url     string `xorm:"not null VARCHAR(255)"`
	Title   string `xorm:"not null VARCHAR(255)"`
	Content string `xorm:"not null TEXT"`
	Utctime string `xorm:"not null VARCHAR(30)"`
}
