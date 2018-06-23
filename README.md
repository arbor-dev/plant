# seedling
Plant your next project

Current Status: Only creates a main.go file in a given directory name. 

TODO: 

- [ ] Create Services Directory
- [ ] Create Config Directory

### Commands

#### root

```
seedling is a command line tool that can create and manage arbor projects

Usage:
  seedling [command]

Available Commands:
  help        Help about any command
  plant       Creates an arbor project.

Flags:
  -h, --help   help for seedling

Use "seedling [command] --help" for more information about a command.
```

#### plant

```
Creates an arbor project.

Usage:
  seedling plant [arbor project name] [flags]

Flags:
  -h, --help       help for plant
  -p, --port int   port for api-gateway (default 8000)
```

