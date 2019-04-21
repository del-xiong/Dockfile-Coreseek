package models

import (
	"time"
)

type WpPosts struct {
	Id                  int64     `xorm:"pk autoincr index(type_status_date) BIGINT(20)"`
	PostAuthor          int64     `xorm:"not null default 0 index BIGINT(20)"`
	PostDate            time.Time `xorm:"not null default '0000-00-00 00:00:00' index(type_status_date) DATETIME"`
	PostDateGmt         time.Time `xorm:"not null default '0000-00-00 00:00:00' DATETIME"`
	PostContent         string    `xorm:"not null LONGTEXT"`
	PostTitle           string    `xorm:"not null TEXT"`
	PostExcerpt         string    `xorm:"not null TEXT"`
	PostStatus          string    `xorm:"not null default 'publish' index index(type_status_date) VARCHAR(20)"`
	CommentStatus       string    `xorm:"not null default 'open' VARCHAR(20)"`
	PingStatus          string    `xorm:"not null default 'open' VARCHAR(20)"`
	PostPassword        string    `xorm:"not null default '' VARCHAR(255)"`
	PostName            string    `xorm:"not null default '' index VARCHAR(200)"`
	ToPing              string    `xorm:"not null TEXT"`
	Pinged              string    `xorm:"not null TEXT"`
	PostModified        time.Time `xorm:"not null default '0000-00-00 00:00:00' DATETIME"`
	PostModifiedGmt     time.Time `xorm:"not null default '0000-00-00 00:00:00' DATETIME"`
	PostContentFiltered string    `xorm:"not null LONGTEXT"`
	PostParent          int64     `xorm:"not null default 0 index BIGINT(20)"`
	Guid                string    `xorm:"not null default '' VARCHAR(255)"`
	MenuOrder           int       `xorm:"not null default 0 INT(11)"`
	PostType            string    `xorm:"not null default 'post' index index(type_status_date) VARCHAR(20)"`
	PostMimeType        string    `xorm:"not null default '' VARCHAR(100)"`
	CommentCount        int64     `xorm:"not null default 0 BIGINT(20)"`
	Indexed             int       `xorm:"not null default 0 index INT(1)"`
}
