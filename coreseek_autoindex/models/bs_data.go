package models

type BsData struct {
	Id          int    `xorm:"not null pk autoincr INT(11)"`
	Url         string `xorm:"not null VARCHAR(255)"`
	Urlhash     string `xorm:"not null unique VARCHAR(12)"`
	Title       string `xorm:"not null VARCHAR(255)"`
	Content     string `xorm:"not null TEXT"`
	Contenthash string `xorm:"not null VARCHAR(12)"`
	Lastupdate  string `xorm:"not null VARCHAR(30)"`
	Indexed     int    `xorm:"not null default -1 comment('-1未索引 0 索引中 1已索引') index INT(1)"`
}
