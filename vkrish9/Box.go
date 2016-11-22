package main 

import (
	"ethos/syscall"
	"ethos/ethos"
	ethosLog "ethos/log"
	"ethos/efmt"
)

func main () {
	me := syscall.GetUser()
	path := "/user/" + me + "/myDir/"
	status := ethosLog.RedirectToLog("Program2")
	if status != syscall.StatusOk {
		efmt.Fprintf(syscall.Stderr,"Error opening %v: %v\n", path, status)
	syscall.Exit(syscall.StatusOk)
	}


	data    := Box { 12, 121,32,43 }
	movebox := Box {}
	movebox.x1 = data.x1
	movebox.y1 = data.y2
	movebox.x2 = data.x2
	movebox.y2 = data.y1
	fd, status := ethos.OpenDirectoryPath(path)
	data.Write(fd)
	data.WriteVar(path +"foobar")
	efmt.Println("Co-ordinates of left bottom and right top"," x1:y1",data.x1,":",data.y1," x2:y2",data.x2,":",data.y2)
	efmt.Println("Calculated co-ordinates of left top and right bottom"," x3:y3",movebox.x1,":",movebox.y1," x4:y4",movebox.x2,":",movebox.y2) 
}
