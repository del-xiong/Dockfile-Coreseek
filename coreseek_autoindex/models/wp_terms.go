package models

type WpTerms struct {
	TermId    int64  `xorm:"not null pk autoincr BIGINT(20)"`
	Name      string `xorm:"not null default '' index VARCHAR(200)"`
	Slug      string `xorm:"not null default '' index VARCHAR(200)"`
	TermGroup int64  `xorm:"not null default 0 BIGINT(10)"`
}
