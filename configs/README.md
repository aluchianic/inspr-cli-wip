Used composite pattern:
https://refactoring.guru/design-patterns/composite

Workspace   == [ component container ] --> keeps inside Applications
Application == [   leaf component    ] --> ...

 - [x] Every config (workspace or application) is RawConfig and can be manipulated trough corresponding (c *RawFile) methods
 - [x] Workspace composes inside it Applications
 - [x] Can be used more than one workspaces
 - [x] Could be added more leaf-components to workspace 


First of all you need to `Load()` your workspace, this will locate all files that are included in workspace based on its' `Root` (that can be changed on start).
On Load `configs` package will initialize default values, if no Root were set - assuming current working directory. During `Load()` package will only locate files,
and does not parse them. To parse files you need to call `Parse()` method, that will take all configs that are included in `workspace` and
unmarshal & parse them. Only after `Parse()` call you can retrieve data from configs.
 

Example of using multiple workspaces:
```go
    package main
    
    import 	(
        "fmt"
        "inspr-cli/configs"
    )

    func main() {
        // By default Workspace root is current pwd
	    w1 := configs.WorkspaceFiles{}
        configs.ShowAndExistIfErrorExists(w1.Load())
        configs.ShowAndExistIfErrorExists(w1.Parse())
	    
        // Workspace root can be changed on initialization
        w2 := configs.WorkspaceFiles{
            Root: "/Users/aluchianic/go/src/test-cli-workspace",
        }
        configs.ShowAndExistIfErrorExists(w2.Load())
        configs.ShowAndExistIfErrorExists(w2.Parse())
        
        fmt.Printf("Workspace-1 config: \n\t%v\nWorkspace-2 config: \n\t%v\n", w1, w2)
   
}
```