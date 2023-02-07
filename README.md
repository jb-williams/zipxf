# Zipxf(Move to its own Repo, NO MAINTAINTED)
https://github.com/jb-williams/zipxf

A little program I hacked together to zip an archive and unzip an archive.
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
