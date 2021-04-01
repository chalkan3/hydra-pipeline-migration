# Hydra - Pipelines migrations - CLI
 **This project is released for create a migration controller for gitlab with a lot projects inside** 

# Usage
## Installation 
### Golang
`go mod download && go build . -o o2b-hydra`

## Commands 

### Create pipeline template
`
    o2b-hydra pipelines create [new Template Name]
`

### Create a new migration folder
`
    o2b-hydra pipelines migration new
`
#### Description

This command will create a new timestamp folder to manage up/down on migration

### Create a new pipeline migration 
```
    o2b-hydra pipelines migration create \ 
    [name of file]  \
    [type of language like C#, GOLANG, NODE] \
    [project id from gitlab] \
    [targetBranch -> name of branch can create merge request] \
    [name branch for create a pr] \
    [commitMessage] 
 
```
#### Sample

```
    o2b-hydra pipelines migration create \ 
    pipeline-dotnet  \
    c# \
    12 \
    main \
    merge-it-on-main \
    "pipeline dotnet will be released" 
```

#### Description

This command will create a new timestamp folder to manage up/down on migration

Will create a lua file template where will controll custom pipelines data

File can be found in **/scripts/migrations/folderCreatedByNewCommand/here**


### Get current timestamp applyed 
`
    o2b-hydra pipelines migration current
`


### Migrate all pipelines 
`
    o2b-hydra pipelines migration migrate
`
