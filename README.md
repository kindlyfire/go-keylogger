
# go-keylogger

This is a small go keylogger library for Windows. It has the advantage over *all the other* keyloggers I found that it correctly parses each keypress, no matter the keyboard layout. No keymaps are hard-coded.

See the `examples/` directory for examples.

Currently, the lib polls for keypresses every X milliseconds and checks the entire keyboard using `GetAsyncKeyState()`, which is not the optimal solution but it works well. Work is going on in the `win-events` branch to convert the lib to use the Windows events instead of polling.