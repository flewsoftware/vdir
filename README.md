# Vdir
A simple but useful go library that makes it easy for developers to handle virtual file systems. It contains functions to change directories.

## example
````go
package main
import (
	"github.com/flew-software/vdir"
    "os"
)

func main() {
    // create a new vPath object to keep track of the dir
	v := vdir.NewVPath(func(s string) bool {
        // we will be using our os file system for the demo
        // you will be able to do this using any other in memory filesystem library
        // this function has to return false if the dir doesnt exist
        // true if it exists.
		f, err := os.Stat(s)
		if err != nil {
			return false
		}
		return f.IsDir()
	})
   
    // then we have to set the starting point of the dir
    // change this to your own dir
	v.ForceChangeDir(vdir.Format("C:\\"))
	
	if v.ChangeDir("..") == false {
		println("cant change dir")
	} else {
		println(v.GetDir())
	}
}
````
