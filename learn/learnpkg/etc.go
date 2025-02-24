package learn

import "fmt"

// SayHello returns a greeting message.
func SayHello(name string) string {
    return fmt.Sprintf("Hello, %s!", name)
}