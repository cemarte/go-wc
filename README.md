# My wc implementation in Go (WIP)

This is my work in progress implementation of the wc coding challenge in https://codingchallenges.fyi/challenges/challenge-wc/

# Lessons learned

## Golang standard library

- I tried my best the TDD approach but I spent more time that I would like to have spent on figuring out the best way to mock the file system interface in Go.
Between the os package and the io/fs package, I could achieve the same end goal of reading a file with both.
Although, the os package does not offer an easy interface to simply mock the underlying file system.
That's where io/fs and testing/fstest come to play.
I could open a file with fsys Open which returns a pointer to os.File, and then use bufio.NewScanner(file) and use the scanner.

- I used cobra for the command line utilities as I didn't want to spend time on command line argument parsing logic and go straight to the business logic.

- Good refresher on POSIX. Also, it wasn't immediately clear how the wc -l (count lines) work, but it's pretty straight forward.

# My local development setup

Code editor: [lazyvim](https://www.lazyvim.org/)
LSP: [gopls](https://github.com/golang/tools/tree/master/gopls)
Terminal: [Warp](https://www.warp.dev/)


