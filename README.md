# tasker
Create tasks from CLI and post it to task management platforms

# How to use

```bash
# Create task
tasker create task -tool="asana" -proj=1234 -name="This is a test task" -desc="We would need to try building this product properly"

# List projects
tasker list projects -tool="asana"

# List tasks for a project
# Project would usually refer to id
tasker list tasks -tool="asana" -proj=1234

# Create task for multiple platforms at the same time
tasker create -name="This is a test task" -desc="We would need to try building this product properly" -tool="asana,github"
```

# curl Commands for development work

```bash
curl -X GET -H "Authorization: Bearer <personal_access_token>" -H "Content-Type: application/json" https://app.asana.com/api/1.0/projects

```

# Contributing to the project

**Warning**: Project is still not stable - there will be plenty of internal API changes required in order to ensure compatability and smoother experience for users

## Quick start

To try and use the development version

```bash
# Go build command would create the binary of the CLI Tool
# It would generate a binary called tasker
go build

# Run the above tasker command if one so wishes (Not all are implemented yet)
./tasker
```
