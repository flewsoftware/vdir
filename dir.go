package vdir

import (
	"path"
	"strings"
)

type VDir struct {
	currentLocation string
	locationExists  func(string) bool
}

func (vd *VDir) ChangeDir(newWD string) bool {
	if newWD == "." {
		newWD = vd.currentLocation
		return true
	}

	if newWD == "." {
		newWD = vd.currentLocation
		return true
	}
	if newWD == ".." {
		vd.currentLocation = GoBackADir(vd.currentLocation)
		return true
	} else if newWD == "/" {
		vd.currentLocation = "/"
		return true
	} else {
		if vd.locationExists(vd.currentLocation+"/"+newWD) == true {
			if vd.currentLocation == "/" {
				vd.currentLocation = vd.currentLocation + path.Clean(newWD)
				return true
			} else {
				vd.currentLocation = vd.currentLocation + "/" + path.Clean(newWD)
				return true
			}

		} else {
			return false
		}
	}
}

func GoBackADir(path string) string {
	a := PathToParts(path)
	var v = ""
	for i := 0; i < len(a)-1; i++ {
		var divider string
		if i == 0 {
			divider = ""
		} else {
			divider = "/"
		}
		v += divider + a[i]
	}
	return v
}

func (vd *VDir) ForceChangeDir(fullPath string) {
	vd.currentLocation = fullPath
}

func PathToParts(path string) []string {
	return strings.Split(path, "/")
}

func NewVPath(dirExistsHandler func(string) bool) VDir {
	return VDir{
		currentLocation: "/",
		locationExists:  dirExistsHandler,
	}
}

func (vd *VDir) GetDir() string {
	return vd.currentLocation
}

func Format(path string) string {
	return strings.Replace(path, "\\", "/", -1)
}
