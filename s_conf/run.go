package s_conf

import "github.com/gin-gonic/gin"

func SimpleRun() {
	r := gin.Default()
	LoadRouter(r)
	r.Run(Conf.SERVING)
}
