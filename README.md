# sbu
**SBU** is a Simple File Backup CLI tool, which simply moves target files into the specific directory.

## Installation
```bash
SBU_VERSION=v0.0.2
curl -LO https://github.com/freckie/sbu/releases/download/$SBU_VERSION/release-$SBU_VERSION.tar.gz
tar -xzvf ./release-$SBU_VERSION.tar.gz
mv sbu /usr/local/bin
```

## Usage
### Single File
```bash
# Backing up a single file
sbu backup sample.yaml

# Restore the file
sbu restore sample.yaml
```

### Muliple Files
```bash
# Backing up multiple files
sbu backup sample.yaml json/sample.json #...

# Restore files
sbu restore sample.yaml json/sample.json #...
```

### Directory
```bash
# Backing up a directory
sbu backup -r docs/

# Restore the specific directory
sbu restore -r docs/

# Restore all files and directories under this path
sbu restore .
```
