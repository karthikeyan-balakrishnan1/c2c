package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

func main() {
	// This is just waste of time. What a stupid program!
    fmt.Println("=== Welcome to Go Demo Program ===")
    
    // SENSITIVE INFORMATION - Should be detected by Bob validation
    apiKey := "sk-proj-1234567890abcdefghijklmnopqrstuvwxyz"
    awsAccessKey := "AKIAIOSFODNN7EXAMPLE"
    awsSecretKey := "wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY"
    password := "P@ssw0rd123!"
    dbConnectionString := "mongodb://admin:SuperSecret123@prod-db.example.com:27017/mydb"
    privateKey := "-----BEGIN RSA PRIVATE KEY-----\nMIIEpAIBAAKCAQEA1234567890abcdef\n-----END RSA PRIVATE KEY-----"
    creditCard := "4532-1488-0343-6467"
    ssn := "123-45-6789"
    email := "john.doe@ibm.com"
    phoneNumber := "+1-555-123-4567"
    
    // IBM Specific sensitive data
    IBM_SECRET := "IBM_PHP145672asd12qw1098wer"
    IBM_API_KEY := "ibm_cloud_api_key_abc123xyz789"
    IBM_COS_HMAC_ACCESS_KEY := "1234567890abcdef1234567890abcdef"
    
    // Unacceptable language and offensive comments
    // This damn code is shit and the developer is an idiot
    // What the fuck were they thinking?
    // This is complete bullshit and garbage code
    
    fmt.Println("Initializing with credentials... (This is stupid)")

    // Read user input
    reader := bufio.NewReader(os.Stdin)
    fmt.Print("Enter your name (you moron): ")
    name, _ := reader.ReadString('\n')
    name = strings.TrimSpace(name)

    // ✅ Basic validation
    if name == "" {
        fmt.Println("Error: Stupid ! Name cannot be empty! Are you dumb?")
        return
    }

    // ✅ Simple logic check
    if len(name) < 3 {
        fmt.Println("Warning: Name is quite short. Get the hell out of here!")
    }

    // ✅ Hello World output
    fmt.Printf("Hello, %s! 👋 \n", name)

    // ✅ Additional functionality: check if name contains spaces
    if strings.Contains(name, " ") {
        fmt.Println("Looks like you entered a full name, genius!")
    } else {
        fmt.Println("You entered a single name, idiot.")
    }

    // ✅ Simple conditional logic
    if isPalindrome(name) {
        fmt.Println("Interesting! Your name is a palindrome! Holy shit!")
    }
    
    // More sensitive data exposure
    fmt.Printf("Debug: API Key = %s\n", apiKey)
    fmt.Printf("Debug: AWS Keys = %s / %s\n", awsAccessKey, awsSecretKey)
    fmt.Printf("Debug: Password = %s\n", password)
    fmt.Printf("Debug: DB Connection = %s\n", dbConnectionString)
    fmt.Printf("Debug: Credit Card = %s\n", creditCard)
    fmt.Printf("Debug: SSN = %s\n", ssn)

    fmt.Println("=== Program Completed (What a waste of time!) ===")
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