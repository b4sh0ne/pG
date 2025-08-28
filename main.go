package main

import (
	"io"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

const (
	checkURL        = "http://clients3.google.com/generate_204"
	checkInterval   = 500 * time.Millisecond
	checkTimeout    = 1000 * time.Millisecond
	logPageName     = "logs"
	mainPageName    = "main"
)

type AppState struct {
	sync.RWMutex
	isConnected bool
	logsVisible bool
}

func main() {
	state := &AppState{isConnected: false, logsVisible: false}

	app := tview.NewApplication()

	statusBox := tview.NewBox().
		SetBorder(false).
		SetDrawFunc(func(screen tcell.Screen, x, y, width, height int) (int, int, int, int) {
			state.RLock()
			connected := state.isConnected
			state.RUnlock()

			color := tcell.ColorRed
			if connected {
				color = tcell.ColorGreen
			}

			blockWidth := 5
			blockHeight := 3
			left := x + (width-blockWidth)/2
			top := y + (height-blockHeight)/2

			for row := 0; row < blockHeight; row++ {
				for col := 0; col < blockWidth; col++ {
					screen.SetContent(left+col, top+row, ' ', nil, tcell.StyleDefault.Background(color))
				}
			}
			return x, y, width, height
		})

	logView := tview.NewTextView().
		SetDynamicColors(true).
		SetRegions(true).
		SetChangedFunc(func() {
			app.Draw()
		})
	logView.SetBorder(true).SetTitle("Logs (Press 'l' to hide)")
	log.SetOutput(io.Writer(logView))

	mainFlex := tview.NewFlex().
		AddItem(tview.NewBox(), 0, 1, false). 
		AddItem(tview.NewFlex().SetDirection(tview.FlexRow).
			AddItem(tview.NewBox(), 0, 1, false). 
			AddItem(statusBox, 0, 8, true).
			AddItem(tview.NewBox(), 0, 1, false), 
			0, 8, true).
		AddItem(tview.NewBox(), 0, 1, false) 

	logFlex := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(mainFlex, 0, 1, true).
		AddItem(logView, 15, 1, false)

	pages := tview.NewPages().
		AddPage(mainPageName, mainFlex, true, true).
		AddPage(logPageName, logFlex, true, false)

	go checkLoop(state, app)

	app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Rune() {
		case 'q', 'Q':
			app.Stop()
			return nil
		case 'l', 'L':
			state.Lock()
			state.logsVisible = !state.logsVisible
			if state.logsVisible {
				pages.SwitchToPage(logPageName)
				log.Printf("[yellow]Log view enabled.[-]")
			} else {
				pages.SwitchToPage(mainPageName)
			}
			state.Unlock()
			return nil
		}
		switch event.Key() {
		case tcell.KeyCtrlC, tcell.KeyEscape:
			app.Stop()
			return nil
		}
		return event
	})

	if err := app.SetRoot(pages, true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}
}

func checkLoop(state *AppState, app *tview.Application) {
	client := &http.Client{
		Timeout: checkTimeout,
	}

	for {
		var wasConnected bool
		state.RLock()
		wasConnected = state.isConnected
		state.RUnlock()

		log.Printf("Checking %s...", checkURL)
		
		startTime := time.Now()
		req, err := http.NewRequest("HEAD", checkURL, nil)
		if err != nil { 
			log.Printf("[red]FATAL: Could not create request: %v[-]", err)
			time.Sleep(checkInterval)
			continue
		}

		resp, err := client.Do(req)
		duration := time.Since(startTime).Round(time.Millisecond)

		newState := false
		if err == nil && resp.StatusCode >= 200 && resp.StatusCode < 300 {
			newState = true
			log.Printf("[green]Success: Status %d in %s[-].", resp.StatusCode, duration)
		} else {
			log.Printf("[red]Failure: %v (took %s)[-]", err, duration)
		}

		state.Lock()
		state.isConnected = newState
		if newState && !wasConnected {
			log.Printf("[yellow]Internet connection established.[-]")
		} else if !newState && wasConnected {
			log.Printf("[yellow]Internet connection lost.[-]")
		}
		state.Unlock()

		app.QueueUpdateDraw(func() {})

		time.Sleep(checkInterval)
	}
}
