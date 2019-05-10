package device

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"os/exec"
	"strconv"
	"strings"

	"github.com/davecgh/go-spew/spew"
)

var (
	ambientDHTPin string
	tentDHTPin    string
	platform      string
)

func init() {
	var err error
	platform = os.Getenv("GG_PLATFORM")
	ambientDHTPin = os.Getenv("GG_AMBIENT_DHT_PIN")
	tentDHTPin = os.Getenv("GG_TENT_DHT_PIN")

	if err != nil {
		log.Println(err)
	}
}

/*

Might need to run...
GetThermalInfo(dhtPin int) (float, int, int, err) {

}
```
go get -d -u github.com/vacoj/dht
go generate github.com/vacoj/dht
```

...on the device to get it working, or at least compile on a pi
*/

// GetDeviceID returns the serial number of the board on a Pi3 running raspbian
func GetDeviceID() string {
	// q := `grep -Po '^Serial\s*:\s*\K[[:xdigit:]]{16}' /proc/cpuinfo`
	// out, err := exec.Command(q).Output()

	data, err := ioutil.ReadFile(`/proc/cpuinfo`)
	if err != nil {
		fmt.Println(err)
	}
	sData := string(data)
	serialSplitter := "Serial\t\t: "
	serial := strings.Trim(strings.Split(sData, serialSplitter)[1], "\n")

	if err != nil {
		fmt.Println(err)
	}

	return serial
}

// GetThermalInfo returns CPU Temp, Ambient Temp, and Humidity if no error
func GetThermalInfo() (float64, float64, float64, float64, float64) {
	var cpuT float64
	if platform == "pi3" {
		cpuT = getCPUTemperature()
		ah, at, th, tt := GetAmbientInfo()
		spew.Dump(cpuT, ah, at, th, tt)
		return cpuT, ah, at, th, tt

	}

	if os.Getenv("GG_DEBUG_MODE") == "True" {
		return randFloat64(), randFloat64(), randFloat64(), randFloat64(), randFloat64()
	}
	return 0, 0, 0.0, 0.0, 0.0
}

// GetAmbientInfo returns ambient humidity and temperature if no errors occur
func GetAmbientInfo() (float64, float64, float64, float64) {

	if os.Getenv("GG_DEBUG_MODE") == "True" {
		return randFloat64(), randFloat64(), randFloat64(), randFloat64()
	}
	// AMBIENT

	cmdA := exec.Command("bin/dht11", ambientDHTPin)
	var outA bytes.Buffer
	cmdA.Stdout = &outA
	err := cmdA.Run()

	if err != nil {
		log.Println(err)
	}

	outStrA := strings.Replace(outA.String(), "\n", "", -1)
	tempsA := strings.Split(outStrA, " ")

	// TENT

	cmdT := exec.Command("bin/dht11", tentDHTPin)
	var outT bytes.Buffer
	cmdT.Stdout = &outT
	err = cmdT.Run()

	if err != nil {
		log.Println(err)
	}

	outStrT := strings.Replace(outT.String(), "\n", "", -1)
	tempsT := strings.Split(outStrT, " ")

	ah, err := strconv.ParseFloat(tempsA[0], 64)
	at, err := strconv.ParseFloat(tempsA[1], 64)

	th, err := strconv.ParseFloat(tempsT[0], 64)
	tt, err := strconv.ParseFloat(tempsT[1], 64)

	if err != nil {
		log.Println(err)
	}

	return ah, at, th, tt

}

// GetCPUTemp returns the CPU temperature in metric Celsius
func getCPUTemperature() float64 {
	temp := 0.0
	var err error
	var out bytes.Buffer

	if os.Getenv("GG_DEBUG_MODE") == "True" {
		return randFloat64()
	}

	if platform == "pi3" {
		cmd := exec.Command("/opt/vc/bin/vcgencmd", "measure_temp")
		cmd.Stdout = &out
		err = cmd.Run()
	}

	if err != nil {
		log.Println(err)
	}

	minusText := strings.Replace(out.String(), "temp=", "", -1)
	minusUnit := strings.Replace(string(minusText), "'C\n", "", -1)

	temp, err = strconv.ParseFloat(minusUnit, 64)

	if err != nil {
		log.Println(err)
	}
	return temp
}

func randFloat64() float64 {
	return float64(rand.Intn(50))
}
