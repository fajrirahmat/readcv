# READCV
Read documents or files in directory and find keyword in it.

## How to Use
First you need to clone this repository or download it to your local PC
```
git clone https://github.com/fajrirahmat/readcv.git
```

Next install it or build it.

### Install
```
cd readcv
go install
```

### Build
```
cd readcv
go build
```

### Running
```
readcv -source=[PATH_TO_SOURCE_DIRECTORY] -target=[PATH_TO_TARGET_DIRECTORY] -keyword=[PATH_TO_KEYWORD_FILE]
```

* `-source` Source directory
* `-target` Target directory that file has keyword in it save.
* `-keyword` location keyword file saved.
