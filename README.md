# tasker
Create tasks from CLI and post it to task management platforms

# How to use

```bash
# Create task
tasker create -tool="asana" -proj=1234 -name="This is a test task" -desc="We would need to try building this product properly" -label="low priority"

# List projects
tasker list -tool="asana"

# List tasks for a project
# Project would usually refer to id
tasker list -tool="asana" -proj=1234

# Create task for multiple platforms at the same time
tasker create -name="This is a test task" -desc="We would need to try building this product properly" -label="low priority" -tool="asana,github"

# Copy task between platforms
# Needs to be amde easier though
tasker cp -originTool="asana" -id="12" -destTool="github"
```

# curl Commands for development work

```bash
curl -X GET -H "Authorization: Bearer <personal_access_token>" -H "Content-Type: application/json" https://app.asana.com/api/1.0/projects

```