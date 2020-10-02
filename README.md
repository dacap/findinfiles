# findinfiles

> Copyright (C) 2017-2020 David Capello
>
> This file is released under the terms of the MIT license.
> Read [LICENSE.txt](LICENSE.txt) for more information.

This little utility searches a piece of regular expression in all
files of the current directory and subdirectories. The idea is to
solve the most common case of a typical `find | grep` command.  It's
not a replacement. It's not fully configurable. It's not what you
want. It's what I need and it might be useful to you.

### Usage

    findinfiles [-1iv] [-. extension] PATTERN [PATTERN2...]

    -1      Do not recurse directories
    -i      Ignore case
    -v      Verbose mode
    -. ext  Search only in files with the given extension

### Examples

    findinfiles -. cpp -. h class

Finds all lines that contain `class` word in `.cpp` or `.h` files of
the current directory and subdirectories.

    findinfiles Copyright

Finds all lines that contain `Copyright` word with a capital `C`.
