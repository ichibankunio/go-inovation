// +build darwin,!arm,!arm64 linux windows
// +build !android
// +build !ios

package input

func isTouchPrimaryInput() bool {
	return false
}
