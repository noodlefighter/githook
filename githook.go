package main

import (
	"path/filepath"
	"strings"
	"os"
    "os/exec"
	"log"
//    "bytes"
	"net/http"
	"gopkg.in/go-playground/webhooks.v5/github"
)

const (
	path = "/webhooks"
)

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func getCurrentDirectory() string {
    dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
    if err != nil {
        return ""
    }
    return strings.Replace(dir, "\\", "/", -1)
}

const ShellToUse = "bash"

func Shellout(command string) (error) {
    //var stdout bytes.Buffer
    cmd := exec.Command(ShellToUse, "-c", command)
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stdout
    err := cmd.Run()
    return err
}

func Worker(sh_file string) {
	log.Printf("exec script %s......", sh_file)
	err := Shellout(sh_file)
	if err != nil {
		log.Printf("exec script %s error: %v", sh_file, err)
	} else {
		log.Printf("exec script %s done.", sh_file)
	}
}

var secret string = "This is your Secret..."

func main() {
	
	//args
	if (len(os.Args) > 1) {
		secret = os.Args[1]
	}

	log.SetFlags(log.Ldate | log.Ltime)
	
	hook, _ := github.New(github.Options.Secret(secret))

	http.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
	
		// note: ugly hard coding, should use reflect??
		payload, err := hook.Parse(r, github.CheckRunEvent, github.CheckSuiteEvent, github.CommitCommentEvent, github.CreateEvent, github.DeleteEvent, github.DeploymentEvent, github.DeploymentStatusEvent, github.ForkEvent, github.GollumEvent, github.InstallationEvent, github.InstallationRepositoriesEvent, github.IntegrationInstallationEvent, github.IssueCommentEvent, github.IssuesEvent, github.LabelEvent, github.MemberEvent, github.MembershipEvent, github.MilestoneEvent, github.OrganizationEvent, github.OrgBlockEvent, github.PageBuildEvent, github.PingEvent, github.ProjectCardEvent, github.ProjectColumnEvent, github.ProjectEvent, github.PublicEvent, github.PullRequestEvent, github.PullRequestReviewEvent, github.PullRequestReviewCommentEvent, github.PushEvent, github.ReleaseEvent, github.RepositoryEvent, github.RepositoryVulnerabilityAlertEvent, github.SecurityAdvisoryEvent, github.StatusEvent, github.TeamEvent, github.TeamAddEvent, github.WatchEvent)
		if err != nil {
			if err == github.ErrEventNotFound {
				// ok event wasn;t one of the ones asked to be parsed				
				log.Printf("ErrEventNotFound")
			}
		}		
		 
		var filename string = ""
		switch payload.(type) {
		case github.PushPayload:
			req := payload.(github.PushPayload)			
			filename = req.Repository.FullName + "/push.sh"
		}
		if (filename == "") {
			log.Printf("invalid repo info, wrong Secret?")
			return
		}
		
		filename = getCurrentDirectory() + "/" + filename
		file_exist, err := PathExists(filename)
		if (!file_exist) {
			log.Printf("script %s notfound", filename)
			return
		}
		
		go Worker(filename)
	})
	log.Printf("githook server running... secret=\"%s\"", secret)
	http.ListenAndServe(":3000", nil)	
}