package main

import (
    // internal/local imports: use an alias and always put in quotation marks
    //"cmd"
    "github.com/habitat-sh/arch-sample-job/cmd"
	
    // external imports
    "github.com/spf13/viper"

    // Go standard module imports
    "fmt"
    "time"
    "os"
)

func main() {
    fmt.Println("hello world")

        // read in environment configuration with Viper
        viper.SetConfigFile("./configuration/chef.yml")
        viper.ReadInConfig()
        // fmt.Println(viper.Get("PORT")) // read: https://blog.logrocket.com/handling-go-configuration-viper/
    
        // bind environment variables
        // read: https://www.developer.com/languages/viper-golang/
    
        // collect command-line flags (which may supercede environment config)
        // read: https://github.com/spf13/cobra README
        err := cmd.Execute()
        if err != nil {
            fmt.Println(err)
            os.Exit(1)
        }
        
        for i := 0; i < 5; i++ {
            time.Sleep(100 * time.Microsecond)
            fmt.Println("waking up... getting next job")
        }

        os.Exit(0)    
}