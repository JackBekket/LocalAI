## Package: assets

Imports:
```
"embed"
"fmt"
"io/fs"
"os"
"path/filepath"

"github.com/mudler/LocalAI/pkg/library"
```

External data, input sources:
- Embedded files from the `embed.FS` object.

TODOs:
- None found.

### Function: ResolvePath
This function takes a directory and a list of paths as input and returns a string representing the resolved path. It appends the given paths to the directory and adds "backend-assets" as a prefix.

### Function: ExtractFiles
This function takes an `embed.FS` object and a target directory as input. It extracts all files from the embedded FS and writes them to the target directory, preserving the directory structure. It also handles the creation of directories and files in the target directory. After extracting the files, it calls the `LoadExtractedLibs` function from the `library` package to load any extracted libraries.

### Function: ListFiles
This function takes an embedded filesystem as input and returns a list of all files within it. The function uses the `fs.WalkDir` function to traverse the filesystem and collect the names of all files. If an error occurs during the traversal, the function logs an error message.

pkg/assets/extract.go
pkg/assets/list.go
