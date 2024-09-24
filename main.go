package main

import (
    "bufio"
    "fmt"
    "os"
    "regexp"
    "strings"
)

const (
    Reset   = "\033[0m"
    Red     = "\033[31m"
    Yellow  = "\033[33m"
    Green   = "\033[32m"
)

func analyzeVariables(cmakeFileContent string) {
    lines := strings.Split(cmakeFileContent, "\n")
    variableCount := make(map[string]int)
    variableUsage := make(map[string]int)
    definedVariables := make(map[string]bool)

    setRegex := regexp.MustCompile(`set\(\s*([a-zA-Z_][a-zA-Z0-9_]*)\s*.*\)`)
    varUsageRegex := regexp.MustCompile(`\b([a-zA-Z_][a-zA-Z0-9_]*)\b`)

    for _, line := range lines {
        line = strings.TrimSpace(line)
        if strings.HasPrefix(line, "#") {
            continue
        }

        if match := setRegex.FindStringSubmatch(line); match != nil {
            variableName := match[1]
            variableCount[variableName]++
            definedVariables[variableName] = true
        }

        matches := varUsageRegex.FindAllStringSubmatch(line, -1)
        for _, match := range matches {
            variableUsage[match[1]]++
        }
    }

    for variable, count := range variableCount {
        if count > 1 {
            fmt.Printf("%s[WARNING] %s is defined multiple times.%s\n", Yellow, variable, Reset)
        }
    }

    fmt.Println("\nVariable Usage Report:")
    for variable := range definedVariables {
        usageCount := variableUsage[variable]
        if usageCount > 0 {
            fmt.Printf("  %s: used %d times.%s\n", variable, usageCount, Reset)
        } else {
            fmt.Printf("%s[WARNING] %s is defined but not used.%s\n", Yellow, variable, Reset)
        }
    }
}

func readCMakeFile(filename string) string {
    file, err := os.Open(filename)
    if err != nil {
        fmt.Printf("%s[ERROR] Unable to open file '%s'.%s\n", Red, filename, Reset)
        return ""
    }
    defer file.Close()

    var content strings.Builder
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        content.WriteString(scanner.Text() + "\n")
    }
    return content.String()
}

func main() {
    if len(os.Args) < 2 {
        fmt.Println("Usage: cmpr <path_to_CMakeLists.txt>")
        return
    }

    cmakeFileContent := readCMakeFile(os.Args[1])
    if cmakeFileContent != "" {
        analyzeVariables(cmakeFileContent)
    }
}
