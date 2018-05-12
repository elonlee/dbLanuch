package main

import (
	"github.com/rivo/tview"
	"os/exec"
	"bytes"
	"fmt"
)

var (
	list *tview.List
	modal *tview.Modal
)

func main() {
	app := tview.NewApplication()

	modal = tview.NewModal().
		SetText("Do you want to quit the application?").
		AddButtons([]string{"Quit", "Cancel"}).
		SetDoneFunc(func(buttonIndex int, buttonLabel string) {
		if buttonLabel == "Quit" {
			app.Stop()
		}
		if buttonLabel == "Cancel"{
			app.SetRoot(list, true).Draw()
		}
	})

	list = tview.NewList().
		AddItem("List item 1", "Some explanatory text", 'a', nil).
		AddItem("List item 2", "Some explanatory text", 'b', nil).
		AddItem("Dosbox", "Some explanatory text", 'c', func(){
			exeCmd("dosbox")
	}).
		AddItem("Chrome", "Some explanatory text", 'd', func(){
			exeCmd("/usr/bin/google-chrome-stable")
	}).
		AddItem("Quit", "Press to exit", 'q', func() {
		app.SetRoot(modal,true).Draw()
	})
	if err := app.SetRoot(list, true).Run(); err != nil {
		panic(err)
	}
}


func exeCmd(cmdName string, arg ... string){
	cmd := exec.Command(cmdName, arg...)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Start()
	if err != nil {
		fmt.Print(err)
	}
	fmt.Printf("Waiting for command to finish...")
	fmt.Println(cmd.Args)
	err = cmd.Wait()
	if err != nil {
		fmt.Printf("Command finished with error: %v", err)
	}
	fmt.Println(out.String())
}