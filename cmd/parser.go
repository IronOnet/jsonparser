package main

import (
	"fmt"
	"os"

	parser "github.com/irononet/jsonparser/pkg"
	"github.com/spf13/cobra"
)


var description string = `
JPPARSER(1) User Commands JPARSER(1)

NAME 
	jparser parse - parse json file and output contents`

var rootCmd *cobra.Command = &cobra.Command{
	Use: "parser", 
	Short: "a json parser", 
	Long: description, 
	Run: ParseJsonCmdF,
}

func ParseJsonCmdF(cmd *cobra.Command, args []string){
	if len(args) < 1{
		fmt.Println("Error: please provide a file path as the first argument")
		return
	}

	file := args[0] // Access the first argument
	res := parser.IsValidJSON(file)
	if res == 1{
		fmt.Printf("%s is a valid JSON file!\n", file)
	} else{
		fmt.Printf("%s is not a valid JSON file\n", file)
	}
}


func main(){
	if err := rootCmd.Execute(); err != nil{
		fmt.Println(err)
		os.Exit(1)
	}
}