package models

type WpTermRelationships struct {
	ObjectId       int64 `xorm:"not null pk default 0 BIGINT(20)"`
	TermTaxonomyId int64 `xorm:"not null pk default 0 index BIGINT(20)"`
	TermOrder      int   `xorm:"not null default 0 INT(11)"`
}
