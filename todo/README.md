# todo

Solution for [task-tracker](https://roadmap.sh/projects/task-tracker)

## Usage

### Clone repo

```bash
git clone https://github.com/vinaykandagatla/backend-projects
```

### Build

```bash
cd backend-projects/todo
go build -o todo
```

## Commands

All operations will be performed on `tasks.json` file in current directory, it will be created if it does not exist

### Add

```bash
todo add "Write REDAME"
```

### List

```bash
todo list # To list all
todo list done # in-progress or todo
```

### Update

```bash
todo update 1 "Write README"
```

### Mark

```bash
todo mark 1 done # in-progress or todo
```

### Delete

```bash
todo delete 1
```

### Help

```bash
todo help
```
