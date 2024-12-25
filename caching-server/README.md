# caching server

Solution for [caching-server](https://roadmap.sh/projects/caching-server)

## Usage

### Clone repo

```bash
git clone https://github.com/vinaykandagatla/backend-projects
```

### Build

```bash
cd backend-projects/caching-server
go get
go build -o caching-server
```

## Commands

Cache will be stored in `cache.db` file in current directory

### Starting the server

```bash
caching-server --origin "https://dummyjson.com"
```

### Starting the server on a specific port

```bash
caching-server --origin "https://dummyjson.com" --port 3000
```

### Clearing the cache

```bash
caching-server --clear-cache
```
