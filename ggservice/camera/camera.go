package camera

import (
	"log"
	"os"
	"os/exec"
	"strconv"
	"time"
)

var (
	cameraLoopHours int
	platform        string
)

func init() {
	platform = os.Getenv("GG_PLATFORM")

	var err error
	cameraLoopHours, err = strconv.Atoi(os.Getenv("GG_CAMERA_INTERVAL_HOURS"))
	if err != nil {
		log.Println(err)
	}
}

// Start : start the camera loop
func Start() {
	takePictureLoop()
}

func takePictureLoop() {
	for {
		// fswebcam -r 1280x720 --jpeg 100 -D 1 /opt/greenguard/static/images/camcapture/test.jpg
		cmd := exec.Command("fswebcam", "-r", "1280x720", "--jpeg", "100", "-D", "1", "/opt/greenguard/static/images/camcapture/gg_"+strconv.FormatInt(time.Now().Unix(), 10)+".jpeg")
		// cmd := exec.Command("streamer", "-f", "jpeg", "-o", "static/images/camcapture/gg_"+strconv.FormatInt(time.Now().Unix(), 10)+".jpeg")
		if platform == "pi3" {
			err := cmd.Run()
			if err != nil {
				log.Println(err)
			}
		}

		time.Sleep(time.Duration(cameraLoopHours) * time.Hour)
	}
}

func TakeNewPicture() {
	cmd := exec.Command("fswebcam", "-r", "1280x720", "--jpeg", "100", "-D", "1", "static/images/camcapture/gg_latest.jpeg")
	err := cmd.Run()

	if err != nil {
		log.Println(err)
	}

}
