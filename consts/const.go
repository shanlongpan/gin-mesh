/**
* @Author:Tristan
* @Date: 2021/12/1 2:49 下午
 */

package consts

import "github.com/shanlongpan/gin-mesh/config"

var LogFileDir = config.Conf.LogDir

const DefaultTraceIdHeader = "trace_id"
const (
	Suffix             = ".log"
)
