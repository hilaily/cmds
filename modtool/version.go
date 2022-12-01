package main

import (
	"fmt"
	"strings"

	"github.com/Masterminds/semver"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cast"
)

const (
	verMajor = "major"
	verMinor = "minor"
	verPatch = "patch"
	verPre   = "pre"
)

type verType string

func newVersion(ver string) (*version, error) {
	v, err := semver.NewVersion(ver)
	if err != nil {
		return nil, fmt.Errorf("parse version fail %w", err)
	}
	return &version{Version: v}, nil
}

// Return 2: if it has max tags
func max(ver []string, preReleasePrefix string) (*version, bool, error) {
	if len(ver) == 0 {
		return nil, false, nil
	}
	arr, err := parseVer(ver)
	if err != nil {
		return nil, false, err
	}
	maxOfAll := arr[0]
	var maxOfPre *version
	for _, v := range arr {
		if maxOfAll.Compare(v.Version) < 0 {
			maxOfAll = v
		}
		if preReleasePrefix != "" {
			pre := v.Version.Prerelease()
			if strings.HasPrefix(pre, preReleasePrefix) {
				if maxOfPre == nil || maxOfPre.Compare(v.Version) < 0 {
					maxOfPre = v
				}
			}
		}
	}

	if preReleasePrefix == "" {
		return maxOfAll, true, nil
	}
	// pre release prefix is not emtpy
	// 1. we don't find version with this prefix,
	// like v1.2.3, v1.2.4-otherpre01, return v1.2.4-otherpre01
	if maxOfPre == nil {
		return maxOfAll, true, nil
	}
	// 2. if there is a max pre, and max normal > max pre,
	// like v1.2.1, v1.2.2-z01, v1.2.0-pre10.
	mainPre, err := maxOfPre.SetPrerelease("")
	if err != nil {
		return nil, false, fmt.Errorf("set pre release fail, %w", err)
	}
	if maxOfAll.Compare(&mainPre) > 0 {
		return maxOfAll, true, nil
	}
	// 3. if there is a max pre, and max normal < max pre, like v1.2.1,v1.2.2-pre10.
	return maxOfPre, true, nil
}

func parseVer(ver []string) ([]*version, error) {
	arr := make([]*version, 0, len(ver))
	for _, v := range ver {
		vv, err := newVersion(v)
		if err != nil {
			return nil, err
		}
		arr = append(arr, vv)
	}
	return arr, nil
}

func firstVersion(typ verType, preReleasePrefix string) *version {
	switch typ {
	case verMajor:
		v, _ := newVersion("v1.0.0")
		return v
	case verMinor:
		v, _ := newVersion("v0.1.0")
		return v
	case verPatch:
		v, _ := newVersion("v0.0.1")
		return v
	default:
		v, _ := newVersion("v0.0.1-" + preReleasePrefix + "01")
		return v
	}
}

type version struct {
	*semver.Version
}

func (v *version) inc(typ verType, preReleasePrefix string) *version {
	var vv semver.Version
	switch typ {
	case verMajor:
		vv = v.IncMajor()
		return &version{Version: &vv}
	case verMinor:
		vv = v.IncMinor()
		return &version{Version: &vv}
	case verPatch:
		vv = v.IncPatch()
		return &version{Version: &vv}
	}
	pre := v.Prerelease()
	logrus.Debugf("inc: %s,%s", preReleasePrefix, pre)
	if pre != "" || strings.HasPrefix(pre, preReleasePrefix) {
		preLast := strings.ReplaceAll(pre, preReleasePrefix, "")
		i, err := cast.ToIntE(preLast)
		if err != nil {
			vv, _ = v.SetPrerelease(pre + "01")
			return &version{Version: &vv}
		}
		vv, _ = v.SetPrerelease(preReleasePrefix + fmt.Sprintf("%02d", i+1))
		return &version{Version: &vv}
	}
	vv, _ = v.SetPrerelease(preReleasePrefix + "01")
	return &version{Version: &vv}
}
