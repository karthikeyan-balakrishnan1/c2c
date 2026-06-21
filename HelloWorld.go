package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

func main() {
    fmt.Println("=== Welcome to Go Demo Program ===")

    // Read user input
    reader := bufio.NewReader(os.Stdin)
    fmt.Print("Enter your name: ")
    name, _ := reader.ReadString('\n')
    name = strings.TrimSpace(name)

    // ✅ Basic validation
    if name == "" {
        fmt.Println("Error: Name cannot be empty!")
        return
    }

    // ✅ Simple logic check
    if len(name) < 3 {
        fmt.Println("Warning: Name is quite short.")
    }

    // ✅ Hello World output
    fmt.Printf("Hello, %s! 👋\n", name)

    // ✅ Additional functionality: check if name contains spaces
    if strings.Contains(name, " ") {
        fmt.Println("Looks like you entered a full name!")
    } else {
        fmt.Println("You entered a single name.")
    }

    // ✅ Simple conditional logic
    if isPalindrome(name) {
        fmt.Println("Interesting! Your name is a palindrome!")
    }

    fmt.Println("=== Program Completed ===")
}

// Function to check if a string is a palindrome
func isPalindrome(s string) bool {
    s = strings.ToLower(strings.ReplaceAll(s, " ", ""))
    n := len(s)

    for i := 0; i < n/2; i++ {
        if s[i] != s[n-i-1] {
            return false
        }
    }
    return true
}