Used composite pattern:
https://refactoring.guru/design-patterns/composite

Workspace   == [ component container ] --> keeps inside Applications
Application == [   leaf component    ] --> ...

 - [x] Every config (workspace or application) is RawConfig and can be manipulated trough corresponding (c *RawFile) methods
 - [x] Workspace composes inside it Applications
 - [x] Can be used more than one workspaces
 - [x] Could be added more leaf-components to workspace 


To initialize your workspace you need to Load() and Parse() it:

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