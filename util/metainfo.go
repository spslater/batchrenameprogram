package util

var Version string = "2.1.0"

var Usage string = `brp [-h] [-V] [[-a FILE] ...] filename [filename ...]

rename batches of files at one time

positional arguments:
  filename              list of files to rename

optional arguments:
  -h, --help            show this help message and exit
  -V, --version         print the brp version number and exit
  -a FILE, --auto FILE  automated file to run (default: None)
`
