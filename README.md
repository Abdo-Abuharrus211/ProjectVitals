# Project Vitalis
An in terminal PC stats monitor built in Go, using [Bubble Tea](https://github.com/charmbracelet/bubbletea) and [gopsutil](https://github.com/shirou/gopsutil) .

![thumbnail](thumbnail.png?raw=true)


## Idea
I wanted to build a terminal-based app to monitor PC vitals stats,
 such as CPU usage, memory usage, and disk usage.
 
 It's a simple, and dare I say cute, TUI to watch my system in real-time.

## Features
- System Info: displays PC name and OS + version.
- **CPU Usage Monitoring**: Display real-time CPU usage.
- **Memory Usage Monitoring**: Show current memory usage and available memory.
- **Disk Usage Monitoring**: Provide details on disk usage and available space.
  
#### Coming soon! 🤩
- **Cross-Platform**: Currently only tested on Linux, I'd like to make it compatible with Windows and macOS. _// might work on Mac since Unix..._
- **Customizable**: Allow users to customize the display and update intervals. 
- **Network Activity Monitoring**: Track network upload and download speeds.
- Introduce flow animation to progress bars.

## Issues
- The CPU temp readings aren't being read correctly, likely incorrect SensorKey being checked.
- Table borders don't align with table background (See thumbnail).
- 
## Instructions

I've yet to build this project as it's still a work in progress...

So to run it: 
1. Navigate to `ProjectVitalis/src/` in your terminal.
2. Enter `go run .`, then the TUI would appear in your terminal.
## License

This project is licensed under the MIT License. See the [License](/license.txt) file for details.

## Contributing

Contributions are welcome! Please fork the repository and submit a pull request.

## Acknowledgements

Special thanks to the authors of [Bubble Tea](https://github.com/charmbracelet/bubbletea), [Bubbles](https://github.com/charmbracelet/bubbles?tab=readme-ov-file), [Lip Gloss](https://github.com/charmbracelet/lipgloss), and [gopsutil](https://github.com/shirou/gopsutil) for their amazing libraries.

