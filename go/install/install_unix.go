// Copyright 2015 Keybase, Inc. All rights reserved. Use of
// this source code is governed by the included BSD license.

// +build linux freebsd openbsd

package install

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"
	"time"

	"github.com/keybase/client/go/libkb"
)

// Below paragraph is kept for historical purposes -------------------------------
// This is no longer what happens.
// | Similar to the Brew install on OSX, the Unix install happens in two steps.
// | First, the system package manager installs all the binaries as root. Second,
// | an autostart file needs to be written to the user's home dir, so that
// | Keybase launches when that user logs in. The second step is done the first
// | time the user starts Keybase.
// | ".desktop" files and the ~/.config/autostart directory are part of the
// | freedesktop.org set of standards, which the popular desktop environments
// | like Gnome and KDE all support. See
// | http://standards.freedesktop.org/desktop-entry-spec/latest/.
// -------------------------------------------------------------------------------
// New way: we package /etc/xdg/autostart/keybase.desktop. If a user wishes
// to toggle off, they put a dummy file in $XDG_CONFIG_HOME/autostart/keybase.desktop.
// To toggle on, simply delete that file (this will allow them to receive updates too).
// User can configure this using the command line client.

const autostartFileText = `# This file allows users with desktop environments
# that respect the XDG autostart standard to autostart Keybase on boot.
# Users can enable autostart by running
# $ keybase install autostart on
# or disable autostart by running
# $ keybase install autostart off
# which will toggle this behavior for the current user.
# If the user wishes, they may explicitly create
# a custom $XDG_CONFIG_HOME/autostart/keybase.desktop
# file. Trying to toggle autostart on and off after that
# will require --force to override the custom setting.
# These settings will be preserved across updates.

[Desktop Entry]
Name=Keybase
Comment=Keybase Service, Filesystem, and GUI
Type=Application
Exec=env KEYBASE_AUTOSTART=1 run_keybase
`

const historicalAutostartFileText = `#
`

func autostartDir(context Context) string {
	// strip off the "keybase" folder on the end of the config dir
	return path.Join(context.GetConfigDir(), "..", "autostart")
}

func autostartFilePath(context Context) string {
	return path.Join(autostartDir(context), "keybase_autostart.desktop")
}

// AutoInstall installs auto start on unix
func AutoInstall(context Context, _ string, _ bool, timeout time.Duration, log Log) ( /* newProc */ bool, error) {
	autostartFilename := autostartFilePath(context)
	autostartText, err := ioutil.ReadFile(autostartFilename)
	if err != nil {
		// There is no autostart file or there was an unexpected error; do
		// nothing.
		log.Debug("Unable to read autostart file; err=%s", err)
		return false, nil
	}
	isHistoricalFile := autostartText == historicalAutostartFileText
	if isHistoricalFile {
		log.Debug("Autostart file same as historical; deleting user local version.")
		// Delete the file; it is now provided in /etc/xdg/autostart
		err := os.Remove(autostartFilename)
		if err != nil {
			log.Debug("Unable to delete autostart file: err=%s", err)
			return nil, err
		}
	}

	return false, nil
}

// CheckIfValidLocation is not used on unix
func CheckIfValidLocation() error {
	return nil
}

// KBFSBinPath returns the path to the KBFS executable
func KBFSBinPath(runMode libkb.RunMode, binPath string) (string, error) {
	return kbfsBinPathDefault(runMode, binPath)
}

// kbfsBinName returns the name for the KBFS executable
func kbfsBinName() string {
	return "kbfsfuse"
}

func updaterBinName() (string, error) {
	return "", fmt.Errorf("Updater isn't supported on unix")
}

// RunApp starts the app
func RunApp(context Context, log Log) error {
	// TODO: Start app, see run_keybase: /opt/keybase/Keybase
	return nil
}

func InstallLogPath() (string, error) {
	return "", nil
}

// WatchdogLogPath doesn't exist on linux as an independent log file
func WatchdogLogPath(string) (string, error) {
	return "", nil
}

func SystemLogPath() string {
	return ""
}
