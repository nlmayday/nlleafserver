package model

import "time"

// CREATE TABLE `plg_game_servers` (
// 	`id` int NOT NULL AUTO_INCREMENT PRIMARY KEY,
// 	`created_at` datetime NOT NULL,
// 	`updated_at` datetime NOT NULL,
// 	`deleted_at` datetime NULL,
// 	`game_id` int NOT NULL COMMENT '游戏id',
// 	`name` char NOT NULL COMMENT '游戏名字',
// 	`addr` varchar(64) NOT NULL COMMENT 'IP端口',
// 	`status` int NOT NULL DEFAULT '1' COMMENT '状态',
// 	`info_time` datetime NULL,
// 	`info_status` int NOT NULL DEFAULT '0'
//   );

type GameServers struct {
	ID         uint64 `gorm:"column:id"`
	CreatedAt  *time.Time
	UpdatedAt  *time.Time
	DeletedAt  *time.Time
	GameId     uint64 `gorm:"column:game_id"`
	Name       string `gorm:"column:name"`
	Addr       string `gorm:"column:addr"`
	Status     int    `gorm:"column:status"`
	InfoTime   *time.Time
	InfoStatus int `gorm:"column:info_status"`
}

func (GameServers) TableName() string {
	return "plg_game_servers"
}
