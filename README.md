# Zipxf
*first program, no tutorials, unix only atm*
A little program I hacked together to zip an archive and unzip an archive. \
Really Only Works correctly with if it is placed within your `$PATH`, unless ran in current directory with `./zipxf`.
Does not work when there are non-compressed/archived directories in current working directory, only with files/compressed/archived items in current working directory.
Usage:
```
zipxf [-z archive_name] [-uz archive_name] [-t archive_name] [-ut archive_name]
        CAN ONLY USE ONE AT A TIME
    -h   - Shows this menu.
    -z   - Zips all files in Current Working Directory into an Archive.
    -uz  - Unzips Archive into Current Working Directory.
    -t   - TarGz's all files in Current Working Dirctory into an Archive.
    -ut  - Un-TarGz Archive into Current Working Directory.

ex:
zipxf -z archive.zip
zipxf -uz archive.zip
zipxf -t archive.tar.gz
zipxf -ut archive.tar.gz
```

#### Makefile
Been messing with a makefile for me to build locally.


BUGS: \
Does not work when there are non-compressed/archived directories in current working directory, only with files/compressed/archived items in current working directory.
 \
TODO: \
[X] - Finish implementing tar.gz functionality \
[X] - Implement un-tar.gz functionality \
[X] - Fix Redudant Zipping \
[X] - Fix Redudant TarGziing \
[ ] - Make it include directories(won't work with directories in pwd)
[/] - Make it OS agnostic(NOT TESTED)
