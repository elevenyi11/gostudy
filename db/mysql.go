package main

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/astaxie/beedb"

	_ "github.com/go-sql-driver/mysql"
)

// 跨服战冰火岛公会吃鸡排行榜
type GlobalMoonvstodGuildRanks struct {
	// 房间
	RoomId string `json:"RoomId"`

	// 赛季ID
	SeasonId int `json:"SeasonId"`

	// 公会Id
	GuildId string `json:"GuildId"`

	// 排名
	Rank int `json:"Rank"`

	// 占城积分
	NPCScore int `json:"NPCScore"`

	// 排名积分
	RankScore int `json:"RankScore"`

	// 总积分
	TotalScore int `json:"TotalScore"`

	// 淘汰时间
	WeedOutTime time.Time `json:"WeedOutTime"`

	// 公会名称
	GuildName string `json:"GuildName"`

	// 服务器ID
	ServerId int `json:"ServerId"`

	// 服务器名称
	ServerName string `json:"ServerName"`

	// 是否吃鸡
	IsChiJi bool `json:"IsChiJi"`

	// 旗帜信息
	BannerInfo string `json:"BannerInfo"`

	// 旗帜logo
	Logo string `json:"Logo"`
}

var orm beedb.Model

func init() {
	db, err := sql.Open("mysql", ".../fcl2_crossserver?charset=utf8&parseTime=true&loc=Local&timeout=30s")
	if err != nil {
		panic(err)
	}
	orm = beedb.New(db)
}

func query() ([]*GlobalMoonvstodGuildRanks, error) {
	db, err := sql.Open("mysql", ".../fcl2_crossserver?charset=utf8&parseTime=true&loc=Local&timeout=30s")
	if err != nil {
		panic(fmt.Errorf("初始化数据库失败，错误信息为：%s", err))
	}
	sql := "SELECT RoomId,SeasonId, GuildId, Rank, NPCScore, RankScore, TotalScore, WeedOutTime,GuildName,ServerId,ServerName,IsChiJi,BannerInfo,Logo FROM g_moonvstod_guild_rank_s "
	execResult, err := db.Query(sql)
	if err != nil {
		fmt.Println("GlobalMoonvstodGuildRanksDal.GetDataFromDb", err)
		return nil, err
	}
	defer execResult.Close()

	var result = make([]*GlobalMoonvstodGuildRanks, 0, 16)
	for execResult.Next() {
		var tempNewItem = &GlobalMoonvstodGuildRanks{}
		var isChiji []byte
		err = execResult.Scan(
			&tempNewItem.RoomId,
			&tempNewItem.SeasonId,
			&tempNewItem.GuildId,
			&tempNewItem.Rank,
			&tempNewItem.NPCScore,
			&tempNewItem.RankScore,
			&tempNewItem.TotalScore,
			&tempNewItem.WeedOutTime,
			&tempNewItem.GuildName,
			&tempNewItem.ServerId,
			&tempNewItem.ServerName,
			&isChiji,
			&tempNewItem.BannerInfo,
			&tempNewItem.Logo)

		tempNewItem.IsChiJi = IsTrue(isChiji)

		if err != nil {
			fmt.Println("GlobalMoonvstodGuildRanksDal.loadFromDb", err)
			return nil, err
		}

		result = append(result, tempNewItem)
	}
	return result, nil
}

func IsTrue(data []byte) bool {

	if data == nil {
		return false
	}

	if len(data) == 0 {
		return false
	}

	if data[0] == 0 {
		return false
	} else {
		return true
	}
}

func selectall() []GlobalMoonvstodGuildRanks {
	//get all data
	var alluser []GlobalMoonvstodGuildRanks
	orm.SetTable("g_moonvstod_guild_rank_s").Limit(10).FindAll(&alluser)
	return alluser
}
