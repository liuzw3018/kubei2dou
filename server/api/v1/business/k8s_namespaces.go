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

type k8sNameSpaces struct{}

func NewK8SNameSpaces() *k8sNameSpaces {
	return &k8sNameSpaces{}
}

func (k *k8sNameSpaces) GetAll(c *gin.Context) {
	var (
		code public.ResponseCode
	)
	namespaceList, err := k8s.ClientSet.CoreV1().Namespaces().List(c, metav1.ListOptions{})
	if err != nil {
		code = 2002
		goto ERR
	}

	response.Success(c, public.SuccessCode, "ok", namespaceList)
	return

ERR:
	response.Error(c, code, err)
	return
}

func (k *k8sNameSpaces) Get(c *gin.Context) {

}

func (k *k8sNameSpaces) Update(c *gin.Context) {

}

func (k *k8sNameSpaces) UpdateAll(c *gin.Context) {

}

func (k *k8sNameSpaces) Add(c *gin.Context) {

}

func (k *k8sNameSpaces) DeleteOne(c *gin.Context) {

}

func (k *k8sNameSpaces) DeleteMore(c *gin.Context) {

}
