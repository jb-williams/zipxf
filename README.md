# Zipxf

A little program I hacked together to zip an archive and unzip an archive. \
Really only works correctly with if it is placed within your $PATH which everything in your current working directory.
Usage:
```
zipxf [-z archive_name] [-uz archive_name] [-t archive_name] [-ut archive_name]
        CAN ONLY USE ONE AT A TIME
    -h   - Shows this menu.
    -z   - (Redundantly Zips itself)Zips all files in Current Working Directory into an Archive.
    -uz  - Unzips Archive into Current Working Directory.
    -t   - (Redundantly TarGz's itself)TarGz's all files in Current Working Dirctory into an Archive.
    -ut  - Un-TarGz Archive into Current Working Directory.

ex:
zipxf -z archive.zip
zipxf -uz archive.zip
zipxf -t archive.zip
zipxf -ut archive.zip
```

#### Makefile
Been messing with a makefile for me to build locally.
```
# Show values of vars
make echo

# Make Dir tree
make makedir

# Compile Executable, Rename it and move to /bin dir
make build

# Build all makdir/build
make all

# Clean up
make clean

# Super Clean
make super

# Man (untested) moves man page to man1
make man
```
BUGS: \
 \
TODO: \
[X] - Finish implementing tar.gz functionality \
[X] - Implement un-tar.gz functionality \
[X] - Fix Redudant Zipping \
[X] - Fix Redudant TarGziing \
