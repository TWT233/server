// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package dao

import (
	"time"
)

const TableNameAccessToken = "chii_oauth_access_tokens"

// AccessToken mapped from table <chii_oauth_access_tokens>
type AccessToken struct {
	ID          uint32    `gorm:"column:id;type:mediumint(8);primaryKey;autoIncrement:true"`
	Type        uint8     `gorm:"column:type;type:tinyint(1) unsigned;not null"`
	AccessToken string    `gorm:"column:access_token;type:varchar(40);not null"`
	ClientID    string    `gorm:"column:client_id;type:varchar(80);not null"`
	UserID      string    `gorm:"column:user_id;type:varchar(80)"`
	ExpiredAt   time.Time `gorm:"column:expires;type:timestamp;not null;default:CURRENT_TIMESTAMP"`
	Scope       *string   `gorm:"column:scope;type:varchar(4000)"`
	Info        []byte    `gorm:"column:info;type:varchar(255);not null"`
}

// TableName AccessToken's table name
func (*AccessToken) TableName() string {
	return TableNameAccessToken
}
