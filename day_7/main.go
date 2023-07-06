package main

import (
	"fmt"
	"os/exec"
)

type File struct {
	fileSize int
	fileName string
}

type Path struct {
	parentDirectory string
	directoryName   string
	subDirectory    []Path
	containingFiles []File
}

func (p *Path) addSubDirectory(directoryToAdd string, path Path) {

}

func main() {
	// var currentDirectory *Path
	// filesAreBeingListed := false
	// fileData, err := os.Open("input.txt")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// scanner := bufio.NewScanner(fileData)

	// var rootDirectoryStructure Path

	// for scanner.Scan() {
	// 	str := scanner.Text()
	// 	// Commands based
	// 	if str[0] == '$' {
	// 		commands := strings.Split(str, " ")

	// 		if commands[1] == "cd" {
	// 			if commands[2] == "/" {
	// 				rootDirectoryStructure = Path{
	// 					parentDirectory: "", directoryName: "/",
	// 				}
	// 				currentDirectory = &rootDirectoryStructure
	// 			} else {
	// 				fmt.Println(rootDirectoryStructure)
	// 			}
	// 		}

	// 		if commands[1] == "ls" {
	// 			filesAreBeingListed = true
	// 			fmt.Println(rootDirectoryStructure.subDirectory)
	// 		}

	// 		//Data based
	// 	} else {
	// 		if filesAreBeingListed {
	// 			if strings.Contains(str, "dir") {
	// 				currentDirectory.subDirectory = append(currentDirectory.subDirectory, Path{directoryName: strings.Split(str, " ")[1]})
	// 			} else {
	// 				fileSize, err := strconv.Atoi(strings.Split(str, " ")[0])
	// 				if err != nil {
	// 					fmt.Println(err)
	// 				}
	// 				currentDirectory.containingFiles = append(currentDirectory.containingFiles, File{
	// 					fileName: strings.Split(str, " ")[1],
	// 					fileSize: fileSize,
	// 				})
	// 			}
	// 		}

	// 	}
	// }

	// defer fileData.Close()
	cmd := exec.Command("ls", "./")

	// The `Output` method executes the command and
	// collects the output, returning its value
	out, err := cmd.Output()
	if err != nil {
		// if there was any error, print it here
		fmt.Println("could not run command: ", err)
	}
	// otherwise, print the output from running the command
	fmt.Println("Output: ", string(out))
}
