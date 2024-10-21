# Rest Reminder for macOS

This is a Go-based reminder application for macOS that sends a notification every hour to remind you to take a break. The timer resets whenever the Mac wakes from sleep, ensuring that notifications are relevant only while you are actively using your computer.

## Prerequisites

- **macOS** (This script uses macOS-specific commands like `pmset` and `osascript`).
- **Go (Golang)** installed on your system. If not installed, download and install it from [https://go.dev/dl/](https://go.dev/dl/).
- **Git** installed on your system.

## Installation

### Step 1: Clone the Repository

1. Clone the repository:

   ```bash
   git clone https://github.com/yourusername/rest-reminder.git
   cd rest-reminder
Replace https://github.com/yourusername/rest-reminder.git with the URL of your repository.
### Step 2: Build the Program
2.Build the Go program:
```bash
go build -o rest_reminder main.go
```
This will generate an executable file named rest_reminder in the current directory.
3.(Optional) Move the executable to /usr/local/bin to make it globally accessible:
```bash
sudo mv rest_reminder /usr/local/bin/
```
This allows you to run the program from anywhere using rest_reminder.

### Step 3: Create a plist for Auto-Start
4.Create a .plist file for launchd to automatically start the program when the Mac wakes from sleep:
```bash
vim ~/Library/LaunchAgents/com.user.restreminder.plist
```
5.Add the following content to the file:
```xml
<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
<plist version="1.0">
<dict>
    <key>Label</key>
    <string>com.user.restreminder</string>
    <key>ProgramArguments</key>
    <array>
        <string>/usr/local/bin/rest_reminder</string>
    </array>
    <key>RunAtLoad</key>
    <true/>
    <key>KeepAlive</key>
    <true/>
</dict>
</plist>
```
6.Save the file and exit.
### Step 4: Load the plist File
7.Load the plist file using launchctl:

```bash
sudo launchctl bootstrap gui/$(id -u) ~/Library/LaunchAgents/com.user.restreminder.plist
```
This command loads the plist into launchd for the current user session, allowing the program to start when the Mac wakes from sleep or reboots.
###Step 5: Verify and Troubleshoot
8.Verify that the job is loaded:
```bash
launchctl list | grep com.user.restreminder
```
If the job is listed, it means the program is set to run automatically.
9.If you encounter errors like Load failed: 5: Input/output error, ensure that:
- The rest_reminder binary is located at the specified path (/usr/local/bin/rest_reminder).

- The plist file is correctly formatted.

- You can also try to unload the job and reload it:

```bash
launchctl unload ~/Library/LaunchAgents/com.user.restreminder.plist
sudo launchctl bootstrap gui/$(id -u) ~/Library/LaunchAgents/com.user.restreminder.plist
```
10.If further issues persist, check system logs for more details:

```bash
log show --predicate 'process == "launchd"' --info
```
### Step 6: Uninstall or Disable
1.To stop the program from running automatically:
```bash
launchctl unload ~/Library/LaunchAgents/com.user.restreminder.plist
```
2.(Optional) Remove the plist file:
```bash
rm ~/Library/LaunchAgents/com.user.restreminder.plist
```

### Usage
- Once configured, the rest_reminder program will run automatically every time your Mac starts or wakes from sleep.
- It will send a notification every hour, indicating the number of hours that have passed since the last notification.



