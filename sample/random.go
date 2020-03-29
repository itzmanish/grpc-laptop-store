package sample

import (
	"math/rand"
	"time"

	"github.com/google/uuid"
	"github.com/itzmanish/laptopstore/pb"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func randomScreenPanel() pb.Screen_Panel {
	if rand.Intn(2) == 1 {
		return pb.Screen_OLED
	}
	return pb.Screen_IPS
}

func randomScreenResolution() *pb.Screen_Resolution {
	height := randomInt(1000, 4330)
	width := height * 16 / 9
	return &pb.Screen_Resolution{
		Height: uint32(height),
		Width:  uint32(width),
	}
}

func randomKeyboardLayout() pb.Keyboard_Layout {
	switch rand.Intn(3) {
	case 0:
		return pb.Keyboard_QUERTY
	case 1:
		return pb.Keyboard_QUERTZ
	case 2:
		return pb.Keyboard_AZERTY
	default:
		return pb.Keyboard_DEFAULT
	}
}

func randomID() string {
	return uuid.New().String()
}

func randomBool() bool {
	return rand.Intn(2) == 1
}

func randomCPUBrand() string {
	return randStringFromSet("Intel", "Amd")
}

func randomCPUName(brand string) string {
	switch brand {
	case "Intel":
		return randStringFromSet(
			"i9-7980XE",
			"i7-7820X",
			"Xeon 5160",
			"i5-7640X",
			"i3-7350K",
		)
	case "Amd":
		return randStringFromSet(
			"Ryzen 5 1400	",
			"Ryzen 7 1700",
			"Ryzen Threadripper 1900X",
			"Ryzen 5 2400GE",
			"Ryzen 3 2300X",
		)
	default:
		return ""
	}

}

func randomGPUBrand() string {
	return randStringFromSet("AMD", "NVIDIA")
}

func randomGPUName(brand string) string {
	switch brand {
	case "AMD":
		return randStringFromSet(
			"Radeon RX 5700",
			"Radeon RX 5600 XT",
			"Radeon RX 5700 XT",
			"Radeon VII",
			"Radeon RX 5500 XT",
		)
	case "NVIDIA":
		return randStringFromSet(
			"GeForce 1060",
			"GeForce GTX 260",
			"GeForce GTX 275",
			"GeForce GTX 295",
			"GeForce GTX 260 Core 216",
		)
	default:
		return ""
	}

}
func randomLaptopBrand() string {
	return randStringFromSet("APPLE", "DELL", "LENOVO")
}

func randomLaptopName(brand string) string {
	switch brand {
	case "APPLE":
		return randStringFromSet(
			"Macbook Air",
			"Macbook Pro",
		)
	case "DELL":
		return randStringFromSet(
			"Inspirion",
			"Lattitude",
			"Vostro",
			"XPS",
			"Alienware",
		)
	case "LENOVO":
		return randStringFromSet(
			"Thinkpad",
			"Ideapad",
			"Helix",
			"X1",
			"Yoga",
		)
	default:
		return ""
	}

}

func randStringFromSet(a ...string) string {
	n := len(a)
	if n == 0 {
		return ""
	}
	return a[rand.Intn(n)]
}

func randomInt(min, max int) int {
	return min + rand.Intn(max-min+1)
}

func randomFloat64(min, max float64) float64 {
	return min + rand.Float64()*(max-min)
}
func randomFloat32(min, max float32) float32 {
	return min + rand.Float32()*(max-min)
}
