package model

import (
	"github.com/KubeOperator/KubeOperator/pkg/model/common"
	uuid "github.com/satori/go.uuid"
	"time"
)

type ClusterStatusCondition struct {
	common.BaseModel
	ID              string `json:"-"`
	Name            string `json:"name"`
	ClusterStatusID string `json:"-"`
	Status          string `json:"status"`
	Message         string `json:"message" gorm:"type:text"`
	OrderNum        int    `json:"-"`
	LastProbeTime   time.Time
}

func (c *ClusterStatusCondition) BeforeCreate() (err error) {
	c.ID = uuid.NewV4().String()
	return nil
}

func (c ClusterStatusCondition) TableName() string {
	return "ko_cluster_status_condition"
}
