# merkleHash
merkleHash is a simple CLI to calculate Merkle root hash (using sha256) for a given list of transactions.
Assuming the list of transactions is a file, where each line represents a single transaction.

This implementation carries the odd remainder to the next level until it's consumed, instead of duplicating and re-hashing.

### Note
This implementation is not suitable for very large files, as it uses `os.ReadFile`, which basically loads the 
entire file into memory for processing. While it's ok for the smaller files as it's actually faster than reading line
by line, for cases when it is necessary to consume larger files, containing `xxxMb` of data,
it is possible to take an alternative approach - reading in configured chunks and use workers for processing.

### Usage
```
./bin/merklehash 

                      _    _      _   _           _
 _ __ ___   ___ _ __| | _| | ___| | | | __ _ ___| |__
| '_   _ \ / _ \ '__| |/ / |/ _ \ |_| |/ _  / __| '_ \
| | | | | |  __/ |  |   <| |  __/  _  | (_| \__ \ | | |
|_| |_| |_|\___|_|  |_|\_\_|\___|_| |_|\__,_|___/_| |_|

Usage:
        merklehash [-flag]

Available flags:
  -path string
        path to the input file
```

## Running the application

### Building from source
1. Clone the repo
2. Build and run the executable
```
make build && ./bin/merklehash -path ./input.txt
```
To build locally, make sure you have required Go version installed.
Alternatively, use Docker image.

### Docker
You can build your own Docker image of merklehash from the Dockerfile with the following:

```
docker build -t merklehash .
```

Run your container:
```
docker run -t --rm -v "$(pwd)"/input.txt:/app/input.txt:ro merklehash -path /app/input.txt
```

## Benchmarks
```
BenchmarkTreeHash_CalculateRoot
BenchmarkTreeHash_CalculateRoot-8   166587   6514 ns/op   1112 B/op   12 allocs/op
PASS
```
