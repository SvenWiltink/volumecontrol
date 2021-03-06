package main

import (
	"flag"
	"github.com/svenwiltink/volumecontrol"
)

func main() {
	var volume int
	flag.IntVar(&volume, "v", 0, "Volume")
	flag.Parse()
	err := volumecontrol.DecreaseVolume(volume)
	if err != nil {
		panic(err)
	}
}
