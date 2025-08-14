package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"time"

	"github.com/pterm/pterm"
)

// UptimeRobotResponse reflects the structure of the JSON response from the Uptime Robot API.
type UptimeRobotResponse struct {
	Monitor struct {
		StatusClass string `json:"statusClass"`
	} `json:"monitor"`
}

func clearScreen() {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
}

func main() {
	apiURL := "https://stats.uptimerobot.com/api/getMonitor/vmM5ruWEAB?m=788139639"

	for {
		clearScreen()

		resp, err := http.Get(apiURL)
		if err != nil {
			pterm.Error.Println(fmt.Sprintf("Failed to fetch status from API: %v", err))
			time.Sleep(5 * time.Minute)
			continue
		}

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			pterm.Error.Println(fmt.Sprintf("Failed to read API response: %v", err))
			resp.Body.Close()
			time.Sleep(5 * time.Minute)
			continue
		}
		resp.Body.Close()

		var uptimeData UptimeRobotResponse
		err = json.Unmarshal(body, &uptimeData)
		if err != nil {
			pterm.Error.Println(fmt.Sprintf("Failed to parse API response: %v", err))
			time.Sleep(5 * time.Minute)
			continue
		}

		switch uptimeData.Monitor.StatusClass {
		case "success":
			aurUpText, _ := pterm.DefaultBigText.WithLetters(pterm.NewLettersFromString("AUR UP")).Srender()
			pterm.FgGreen.Println(aurUpText)
		case "danger":
			aurDownText, _ := pterm.DefaultBigText.WithLetters(pterm.NewLettersFromString("AUR DOWN")).Srender()
			pterm.FgRed.Println(aurDownText)
		default:
			unknownText, _ := pterm.DefaultBigText.WithLetters(pterm.NewLettersFromString("UNKNOWN")).Srender()
			pterm.FgYellow.Println(unknownText)
		}

		pterm.Info.Println("Next check in 5 minutes. Last checked at: " + time.Now().Format("15:04:05"))
		time.Sleep(5 * time.Minute)
	}
}
