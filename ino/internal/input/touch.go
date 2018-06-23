// +build darwin,!arm,!arm64 freebsd linux windows js
// +build !android
// +build !ios

package input

func isTouchPrimaryInput() bool {
	return false
}
