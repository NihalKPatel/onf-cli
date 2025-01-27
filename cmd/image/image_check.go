package image

import (
	"fmt"
	"github.com/OnFinality-io/onf-cli/pkg/service"
	"github.com/OnFinality-io/onf-cli/pkg/watcher"
)

func ImageCheckProcess(imgRepo, version string, isPrintLog bool, checkSuccessFn func()) {
	watcherFlags := watcher.WatcherFlags{2, true}
	watcherFlags.ToWatch(func(done chan bool) {
		imageCheckPayload := &service.ImageCheckPayload{ImageRepository: imgRepo, Version: &version}
		imageCheckRet, err := service.CheckImage(imageCheckPayload)
		if err != nil {
			fmt.Println(err.Error())
			done <- true
			return
		}

		image := fmt.Sprintf("%s:%s", imgRepo, version)
		switch imageCheckRet.Status {
		case service.Pending:
			print(isPrintLog, "Image [%s] is checking.", image)
		case service.Fail:
			print(isPrintLog, "Image [%s] check status is %s. Reason is %s", image, imageCheckRet.Status, imageCheckRet.Config.CheckStatus.Reason)
			done <- true
		case service.Timeout:
			print(isPrintLog, "Image [%s] check status is %s. Reason is %s", image, imageCheckRet.Status, imageCheckRet.Config.CheckStatus.Reason)
			done <- true
		case service.Success:
			print(isPrintLog, "Image [%s] checked success.", image)
			checkSuccessFn()
			done <- true
		default:
			print(isPrintLog, "Check image [%s] status: %s", image, imageCheckRet.Status)
		}

	})
}

func print(isPrint bool, format string, a ...interface{}) {
	if isPrint {
		fmt.Println(fmt.Sprintf(format, a))
	}

}
