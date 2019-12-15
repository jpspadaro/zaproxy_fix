/////////////////////////////
// zaproxy_fix
// Jason Spadaro 2019
//
// Grabs the latest geckodriver and drops it into the zaproxy folder. Currently
// designed for 64 bit linux systems.
//
// Requires wget, bash, tar, and mv. Nothing too crazy.
/////////////////////////////

package main
	

import (
       "fmt"
       "net/http"
       "os"
       "os/exec"
       "io/ioutil"
       "regexp"
)
	

func main() {

     // Banner
     fmt.Println("===OWASP Zap Gecko Fixer===")

     // Pulls the page for the latest and greatest.
     resp, err := http.Get("https://github.com/mozilla/geckodriver/releases/latest")
     if err != nil {
     	fmt.Println("Could not contact github. Exiting...")
	os.Exit(3)
     }

     defer resp.Body.Close()

     // Grab the response.
     body, err := ioutil.ReadAll(resp.Body)

     // Regex for finding the download URL.
     r, _ := regexp.Compile("/mozilla/geckodriver/releases/download/v0.26.0/geckodriver-v.*linux64.*gz")

     //DOES NOT INCLUDE DOMAIN
     latesturl := r.FindString(string(body))

     fmt.Println("https://github.com" + latesturl)

     // Sets up shell commands.
     download := exec.Command("wget", "https://github.com" + latesturl)
     untar := exec.Command("bash", "-c", "tar xfz *.tar.gz")
     moveit := exec.Command("mv","geckodriver","$HOME/.ZAP/webdriver/linux/64/")

     // Runs the shell commands.
     download.Start()
     download.Wait()

    
     untar.Start()
     untar.Wait()

     moveit.Start()
     moveit.Wait()
}
