// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameWithdrawalRecord = "withdrawal_record"

// WithdrawalRecord mapped from table <withdrawal_record>
type WithdrawalRecord struct {
	ID         int64     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	UserID     int64     `gorm:"column:user_id;not null;comment:申请的用户id" json:"user_id"`                    // 申请的用户id
	Amount     string    `gorm:"column:amount;not null;comment:提现金额 分" json:"amount"`                       // 提现金额 分
	Type       string    `gorm:"column:type;not null;comment:提现方式" json:"type"`                             // 提现方式
	Name       string    `gorm:"column:name;not null;comment:收款人姓名" json:"name"`                            // 收款人姓名
	Contact    string    `gorm:"column:contact;not null;comment:联系方式" json:"contact"`                       // 联系方式
	Account    string    `gorm:"column:account;not null;comment:收款账号" json:"account"`                       // 收款账号
	Remarks    string    `gorm:"column:remarks;not null;comment:评论" json:"remarks"`                         // 评论
	Message    string    `gorm:"column:message;not null;comment:消息" json:"message"`                         // 消息
	Status     int32     `gorm:"column:status;not null;default:3;comment:0异常 1正常 3审核中 6等待下发" json:"status"` // 0异常 1正常 3审核中 6等待下发
	IP         string    `gorm:"column:ip;not null" json:"ip"`
	UserAgent  string    `gorm:"column:user_agent;not null;comment:ua" json:"user_agent"` // ua
	CreateTime time.Time `gorm:"column:create_time;not null;default:CURRENT_TIMESTAMP" json:"create_time"`
	UpdateTime time.Time `gorm:"column:update_time;not null;default:CURRENT_TIMESTAMP" json:"update_time"`
}

// TableName WithdrawalRecord's table name
func (*WithdrawalRecord) TableName() string {
	return TableNameWithdrawalRecord
}