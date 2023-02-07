# Zipxf

A little program I hacked together to zip an archive and unzip an archive.
Really only works correctly with it is place within your $PATH which everything in your current working directory
Usage:
```
zipxf [-z] [archive_name] [-u] [archive_name]
        CANNOT USE BOTH -z and -u
    -h  - Shows this menu.
    -z  - Zips all files in Current Working Directory into an Archive. Default: archive.zip
    -u  - Unzips Archive into Current Working Directory

ex:
zipxf -z archive.zip
zipxf -u archive.zip
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
