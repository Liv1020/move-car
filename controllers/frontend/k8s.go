package frontend

import "github.com/gin-gonic/gin"

type k8s struct{}

// K8S K8S
var K8S = k8s{}

// V1 V1
func (t *k8s) V1(c *gin.Context) {
	c.String(200, "Hello k8s v1")
}

// V2 V2
func (t *k8s) V2(c *gin.Context) {
	c.String(200, "Hello k8s v2")
}
