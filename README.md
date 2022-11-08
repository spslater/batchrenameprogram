# Batch Rename Program
The Batch Rename Program allows for renaming of multiple files from the command line.  
This is a rewrite in Go of the original program I wrote in Python.  

## Usage
```
usage: brp [-h] [-V] [[-a FILE]...]] filename [filename ...]

rename batches of files at one time

positional arguments:
  filename              list of files to rename

optional arguments:
  -h, --help            show this help message and exit
  -V, --version         show program's version number and exit
  -a [FILE ...], --auto [FILE ...]
                        automated file to run
```

## Operations
```
help (?, h) [-s] [commands ...]
  display help message
  positional arguments:
    commands   commands to get specific info on
  optional arguments:
    -s, --small
               display just the usage messages

save (s) [-c]
  save files with current changes
  optional arguments:
    -c, --confirm
               automatically confirm action

quit (q, exit) [-c]
  quit program, don't apply unsaved changes
  optional arguments:
    -c, --confirm
               automatically confirm action

write (w) [-c]
  write changes and quit program, same as save then quit
  optional arguments:
    -c, --confirm
               automatically confirm action

list (l, ls)
  lists current files being modified

history (hist, past) [-p]
  print history of changes for all files
  optional arguments:
    -p, --peak just show single file history

undo (u) number
  undo last change made
  positional arguments:
    number     number of changes to undo

reset (o, over) [-c]
  reset changes to original inputs, no undoing
  optional arguments:
    -c, --confirm
               automatically confirm action

automate (a, auto) [filenames ...]
  automate commands in order to speed up repetative tasks
  positional arguments:
    filenames  file names to run commands

replace (r, re, reg, regex) [find] [replace]
  find and replace based on a regex
  positional arguments:
    find       pattern to find
    replace    pattern to insert

append (ap) [-f FILENAMES [FILENAMES ...]] [-p PADDING] [append] [find]
  pattern and value to append to each file that matches, can be automated with a file
  positional arguments:
    append     value to append to filename
    find       regex pattern to match against
  optional arguments:
    -f FILENAMES [FILENAMES ...], --filenames FILENAMES [FILENAMES ...]
               file to load patterns from
    -p PADDING, --padding PADDING
               string to insert between the end of the filename and the value being appended

prepend (p, pre) [-f FILENAMES [FILENAMES ...]] [-p PADDING] [prepend] [find]
  tsv with pattern and value to prepend to each file that matches
  positional arguments:
    prepend    value to append to filename
    find       regex pattern to match against
  optional arguments:
    -f FILENAMES [FILENAMES ...], --filenames FILENAMES [FILENAMES ...]
               file to load patterns from
    -p PADDING, --padding PADDING
               string to insert between the end of the filename and the value being prepended

insert (i, in) [insert] [index] [-c]
  insert string, positive from begining, negative from ending
  positional arguments:
    insert     value to insert
    index      index (starting from 0) to insert at, negative numbers will insert counting from the end
  optional arguments:
    -c, --confirm
               automatically confirm action

case (c) [styles ...]
  change the case (title, upper, lower) of files
  positional arguments:
    styles     type of case style (lower, upper, title, camel, kebab, ect) to switch to

extension (x, ext) [-e] [-n] [new] [pattern]
  change the extension on all files or files that match pattern
  positional arguments:
    new        change the extension to this
    pattern    pattern to match against,
  optional arguments:
    -e, --ext  match against the extensions instead of filename
    -n, --np, --nopattern
               change for all files
```
## Links
* [Github](https://github.com/spslater/batchrenameprogram)

## Contributing
Help is greatly appreciated. First check if there are any issues open that relate to what you want
to help with. Also feel free to make a pull request with changes / fixes you make.

## License
[MIT License](https://opensource.org/licenses/MIT)
