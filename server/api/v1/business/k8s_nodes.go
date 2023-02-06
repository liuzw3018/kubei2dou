package business

import (
	"github.com/gin-gonic/gin"
	"github.com/liuzw3018/kubei2dou/k8s"
	"github.com/liuzw3018/kubei2dou/server/pkg/public"
	"github.com/liuzw3018/kubei2dou/server/pkg/response"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

/**
 * @Author: liuzw
 * @Date: 2023/2/3 09:42
 * @File: k8s_namespaces.go
 * @Description: //TODO $
 * @Version:
 */

type k8sNodes struct{}

func NewK8SNodes() *k8sNodes {
	return &k8sNodes{}
}

func (k *k8sNodes) GetAll(c *gin.Context) {
	var (
		code public.ResponseCode
	)
	nodeList, err := k8s.ClientSet.CoreV1().Nodes().List(c, metav1.ListOptions{})
	if err != nil {
		code = 2002
		goto ERR
	}

	response.Success(c, public.SuccessCode, "ok", nodeList)
	return

ERR:
	response.Error(c, code, err)
	return
}

func (k *k8sNodes) Get(c *gin.Context) {

}

func (k *k8sNodes) Update(c *gin.Context) {

}

func (k *k8sNodes) UpdateAll(c *gin.Context) {

}

func (k *k8sNodes) Add(c *gin.Context) {

}

func (k *k8sNodes) DeleteOne(c *gin.Context) {

}

func (k *k8sNodes) DeleteMore(c *gin.Context) {

}
