# internal

Package internal provides version information for the software.

## Files

- internal/version.go

## Code Summary

The code defines two global variables, Version and Commit, which store the version and commit hash of the software, respectively. The PrintableVersion function returns a string containing the version and commit hash of the software in the format "Version (Commit)".

## Relations between code entities

The PrintableVersion function uses the Version and Commit variables to construct the output string.

