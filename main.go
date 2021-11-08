package main

import (
	obsClient "obsTest/obsclient"

	"github.com/Lincyaw/huaweicloud-sdk-go-obs/obs"
)

func main() {
	//---- init log ----
	defer obs.CloseLog()
	obs.InitLog("/temp/OBS-SDK.log", 1024*1024*100, 5, obs.LEVEL_WARN, false)
	obsClient.ListBuckets()
}
