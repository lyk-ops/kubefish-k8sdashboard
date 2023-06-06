import { request } from "@/utils/service"
import { commonListReq, commonReq, commonRelatedReq } from "./entry"
import { type ReplicaSetBriefList } from "./replicaSet"

export interface DeploymentBrief {
  name: string
  namespace: string
  AvailableReplicas: number
  Replicas: number
  CreationTimestamp: string
}

interface DeploymentBriefList extends PageInfo {
  list: DeploymentBrief[]
  total: number
}

// 获取deployment list
export const getDeploymentsApi = (data: commonListReq) => {
  return request<ApiResponseData<DeploymentBriefList>>({
    url: "/k8s/deployment/getDeployments",
    method: "post",
    data
  })
}

export interface DeploymentDetail {
  metadata: object
  spec: {
    replicas: number
    selector: object
    strategy: {
      type: string
      rollingUpdate: {
        maxUnavailable: number
        maxSurge: string
      }
    }
  }
  status: {
    replicas: number
    updatedReplicas: number
    readyReplicas: number
    availableReplicas: number
    conditions: []
  }
}

// 获取deployment详情
export const getDeploymentDetailApi = (data: commonReq) => {
  return request<ApiResponseData<DeploymentDetail>>({
    url: "/k8s/deployment/getDeploymentDetail",
    method: "post",
    data
  })
}

// 删除deployment
export const deleteDeploymentApi = (data: commonReq) => {
  return request<ApiResponseData<null>>({
    url: "/k8s/deployment/deleteDeployment",
    method: "post",
    data
  })
}

// 获取deployment关联的rs
export const getDeploymentRsApi = (data: commonRelatedReq) => {
  return request<ApiResponseData<ReplicaSetBriefList>>({
    url: '/k8s/deployment/getDeploymentRs',
    method: 'post',
    data
  })
}
