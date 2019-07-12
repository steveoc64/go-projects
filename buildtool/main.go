package main

import (
	"flag"
	"fmt"
	"log"
	"os/exec"
)

// Use PGO by compiling with -fprofile-generate, running and than recompiling with -fprofile-use and -profile-correction
const compileflags = "-O3 -march=native -mtune=native -flto -ffast-math -pipe -Wall -v"

func runCompile(command string) {
    compile := exec.Command("sh", "-c", command)
    output, err := compile.CombinedOutput()
    if err != nil {
        log.Fatal(output, err)
    } else {
	print(output)
    }
}

func print(output []byte) {
    fmt.Printf("%s\n", output)
}

func compileCPP(file string) {
    command := "clang++ -std=c++14 -stdlib=libc++ " + file + " -o compiled-cpp " + compileflags
    runCompile(command)
}

func compileGO(file string) {
    command := `go build -ldflags="-s -w" -o compiled-go ` + file
    runCompile(command)
}

func checkLang(language, file string) {
    switch language {
    case "c++", "C++", "cpp":
        compileCPP(file)
    case "go", "Go", "GO":
        compileGO(file)
    case "c", "C":
        compileC(file)
    default:
	fmt.Println("Usage:\n   build [language] [file-to-build]\n",
            "     Valid languages:\n",
            "       Go  (Builds without debug and dwarf data)\n",
            "       C   (Uses perforamce optimized compiler flags with Clang compiler)\n",
            "       C++ (Uses same optimizing compiler flags as C in Clang++ compiler)\n",
            "\n       File to build:\n",
            "       Specifies the file to build.",
        )
    }
}

func main() {
    flag.Parse()
    language, file := flag.Arg(0), flag.Arg(1)
    checkLang(language, file)
}
