package main

import (
	"./models"
	"bytes"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"github.com/marcsauter/single"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)
var engine *xorm.Engine
const maxlimit int = 1000
var linebuffer map[string]string = map[string]string{}
func main() {
	log.Println("start...")
	s := single.New("coreseek.autoindex")
	if err := s.CheckLock(); err != nil && err == single.ErrAlreadyRunning {
		log.Fatal("another instance of the app is already running, exiting")
	} else if err != nil {
		// Another error occurred, might be worth handling it as well
		log.Fatalf("failed to acquire exclusive app lock: %v", err)
	}
	defer s.TryUnlock()

	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	b, err := ioutil.ReadFile(dir+"/sphinx.conf")
	if err != nil {
		log.Println(err)
	}
	body := strings.Replace(string(b), "\r\n", "\n", -1)
	data := strings.Split(body, "\n")
	for _,line := range data {
		tmp := strings.Split(strings.TrimSpace(line), "=")
		if len(tmp) == 2 {
			linebuffer[strings.TrimSpace(tmp[0])] = strings.TrimSpace(tmp[1])
		}
	}

	engine, err = xorm.NewEngine("mysql", linebuffer["sql_user"]+":"+linebuffer["sql_pass"]+"@tcp("+linebuffer["sql_host"]+")/"+linebuffer["sql_db"]+"?charset=utf8")
	if err != nil {
		panic(err)
	}

	for {
		go rotate_index()
		time.Sleep(time.Second * 120)
	}
}
func rotate_index() {
	var iddict bytes.Buffer
	var wpposts = make([]models.WpPosts, 0)
	// var err error
	// 先取出一批需要建立索引的数据(-1未索引)
	engine.Where("post_type=? and post_status=? and indexed=?", "sucai","publish",-1).Cols("ID").Limit(maxlimit).Find(&wpposts)
	log.Println(len(wpposts))
	if len(wpposts) ==0 {
		return
	}
	for k,d := range wpposts {
		iddict.WriteString(strconv.Itoa(int(d.Id)))
		if k<len(wpposts)-1 {
			iddict.WriteString(",")
		}
	}
	cmd := exec.Command("/usr/local/bin/indexer", "--all","--rotate")
	cmd.Output()
	// 设置待索引数据状态为 0 待索引
	// engine.Exec("update wp_posts set indexed=? where ID in ("+iddict.String()+")", 0)
	// cmd := exec.Command("")
	// if _, err = os.Stat("/var/sphinx/data/search_delta.sph"); err != nil {
	// 	// delta不存在 创建delta全量索引
	// 	cmd = exec.Command("/usr/local/bin/indexer", "delta")
	// 	cmd.Output()
	// } else {
	// 		// 创建delta增量索引
	// 	cmd = exec.Command("/usr/local/bin/indexer", "delta","--rotate")
	// 	cmd.Output()
	// }
	// // 合并入主索引
	// cmd = exec.Command("/usr/local/bin/indexer","--merge","search" ,"delta","--rotate")
	// cmd.Output()
	// // 设置索引状态为1 索引完成
	// engine.Exec("update wp_posts set indexed=? where indexed=?", 1, 0)
}