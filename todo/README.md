# todo

Solution for [task-tracker](https://roadmap.sh/projects/task-tracker)

## Usage

### Clone repo

```bash
git clone https://github.com/avimuser/smolprojects/todo
```

### Build

```bash
cd todo
go build -o todo
```

### Commands

Tasks will be stored in `tasks.json` file in current directory

#### Add Task

```bash
todo add "Write REDAME"
```

#### Listing Tasks

```bash
todo list # To list all
todo list done # in-progress or todo
```

#### Updating Tasks

```bash
todo update 1 "Write README"
```

#### Changing the status of a Task

```bash
todo mark 1 done # in-progress or todo
```

#### Deleting a Task

```bash
todo delete 1
```

#### Help

```bash
todo help
```
