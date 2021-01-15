# Dockfile-Coreseek

Coreseek 是一款中文全文检索/搜索软件，以 GPLv2 许可协议开源发布，基于 Sphinx 研发并独立发布，专攻中文搜索和信息处理领域，适用于行业/垂直搜索、论坛/站内搜索、数据库搜索、文档/文献检索、信息检索、数据挖掘等应用场景，用户可以免费下载使用

演示启动请看**example目录**下的demo示例

-
#需要创建个配置文件： `/path/sphinx/sphinx.conf`

```
################# sphinx config #######################
source search
{
	type = mysql
	sql_host = dbhost
	sql_user = dbname
	sql_pass = dbpass
	sql_db = db
	sql_port = 3306
	sql_query_pre = SET NAMES utf8
	sql_query_pre = SET SESSION query_cache_type=OFF
	sql_query = SELECT * FROM docs
	sql_attr_uint = item_id
	sql_attr_timestamp = updated_at
	sql_attr_timestamp = created_at
	sql_ranged_throttle = 0
}


# 基于mmseg3的正常中英分词配置
index search
{
	source = search
	path = /var/sphinx/data/search
	docinfo = extern
	mlock = 0
 	preopen = 1
	min_word_len = 1
	charset_type = zh_cn.utf-8
	charset_dictpath = /usr/local/mmseg3/etc/
	min_prefix_len = 0
	min_infix_len = 1
	ngram_len = 0
}
# 中英一元分词配置
index search
{
	source = search
	path = /var/sphinx/data/search
    docinfo                 = extern
    mlock                   = 0
    morphology              = none
    min_word_len            = 1
    charset_type            = utf-8
    min_prefix_len          = 0
    html_strip              = 1
    charset_table           = 0..9, A..Z->a..z, _, a..z, U+410..U+42F->U+430..U+44F, U+430..U+44F
    ngram_len               = 1
    ngram_chars             = 0..9, A..Z->a..z, _, a..z, U+3000..U+2FA1F
}

indexer
{
	mem_limit = 1024M
	write_buffer = 4M
}

searchd
{
	listen = 0.0.0.0:9312
	log = /var/sphinx/log/sphinx.log
	query_log = /var/sphinx/log/query.log
	read_timeout = 2
	max_children = 0
	pid_file = /var/run/sphinx.pid
	max_matches = 100000
	seamless_rotate = 1
	preopen_indexes = 0
	unlink_old = 1
	read_buffer = 8M
	compat_sphinxql_magics = 0
}
# EOF


```

# 启动 `docker`
```

docker run --name coreseek_sphinx -v /root/sphinx/sphinx:/usr/local/etc/sphinx -p 9312:9312 -i pastyouth/coreseek

```
# 重建mmseg字库
```
/usr/local/bin/mmseg -u /path/to/unigram.txt
```
# coreseek操作
```
# 进入容器
docker exec -it [containerid] /bin/sh

#全索引
/usr/local/bin/indexer --all --rotate

# 索引index search2
/usr/local/bin/indexer search2 --rotate

# 查看帮助
/usr/local/bin/searchd -h

# 停止
/usr/local/bin/searchd --stop
```