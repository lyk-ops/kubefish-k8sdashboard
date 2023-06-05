package request

import "github.com/pddzl/kubefish/server/model/common/request"

// RsListReq 获取replicaSet列表
type RsListReq struct {
	request.PageInfo
	Namespace string `json:"namespace"`
}

type ReplicaSetCommon struct {
	Namespace  string `json:"namespace" validate:"required"`
	ReplicaSet string `json:"replicaSet" validate:"required"`
}

// RsPods 获取rs关联的pod
type RsPods struct {
	ReplicaSetCommon
	request.PageInfo
}
