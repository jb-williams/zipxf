# Zipxf

A little program I hacked together to zip an archive and unzip an archive. \
Really only works correctly with if it is placed within your $PATH which everything in your current working directory.
Usage:
```
zipxf [-z archive_name] [-uz archive_name] [-t archive_name] [-ut archive_name]
        CAN ONLY USE ONE AT A TIME
    -h   - Shows this menu.
    -z   - Zips all files in Current Working Directory into an Archive.
    -uz  - Unzips Archive into Current Working Directory.
    -t   - TarGz's all files in Current Working Dirctory into an Archive.
    -ut

ex:
zipxf -z archive.zip
zipxf -u archive.zip
zipxf -t archive.zip
zipxf -ut archive.zip
```

### Install
Been messing with a make file to build locally
```
# Show values of vars
make echo

# Make Dir tree
make makedir

# Compile Executable, Rename it and move to /bin dir
make build

# Build all makdir/buil
make all

# Clean up
make clean

# Super Clean
make super

# Man (untested) moves man page to man1
make man
```
TODO:
[X] - Finish implementing tar.gz functionality
[ ] - Implement un-tar.gz functionality
