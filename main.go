package mediadevicesstate

import (
	"github.com/cakeholeDC/go-media-devices-state/pkg/camera"
	"github.com/cakeholeDC/go-media-devices-state/pkg/debug"
	"github.com/cakeholeDC/go-media-devices-state/pkg/microphone"
)

// IsCameraOn returns true is any camera in the system is ON
func IsCameraOn(debugEnabled bool) (bool, error) {
	return camera.IsCameraOn(debugEnabled)
}

// IsMicrophoneOn returns true is any microphone in the system is ON
func IsMicrophoneOn(debugEnabled bool) (bool, error) {
	return microphone.IsMicrophoneOn(debugEnabled)
}

// Debug calls all available device functions and prints the results
func Debug() {
	debug.Debug()
}
