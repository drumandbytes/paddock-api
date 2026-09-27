// Command gentrackmaps pre-renders static/track_maps/<circuitId>.svg from
// bacinger/f1-circuits so /f1/next_map serves a static file. Renders in
// #e5d486; map.go swaps in TRACK_COLOUR at serve time.
package main

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"sort"
	"sync"
	"time"

	"github.com/maris/paddock-api/internal/trackmap"
)

const referenceColour = "#e5d486"

type result struct {
	circuitID string
	err       error
}

func main() {
	outputDir := "static/track_maps"
	if err := os.MkdirAll(outputDir, 0o755); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	// one goroutine per circuit; 26 fetches is nothing for raw.githubusercontent.com
	client := &http.Client{Timeout: 15 * time.Second}
	results := make(chan result, len(trackmap.CircuitGeometryIDs))
	var wg sync.WaitGroup
	for circuitID, geometryID := range trackmap.CircuitGeometryIDs {
		wg.Add(1)
		go func(circuitID, geometryID string) {
			defer wg.Done()
			results <- result{circuitID: circuitID, err: generate(client, outputDir, circuitID, geometryID)}
		}(circuitID, geometryID)
	}
	go func() {
		wg.Wait()
		close(results)
	}()

	var failures []string
	for r := range results {
		if r.err != nil {
			fmt.Printf("%s... FAILED: %v\n", r.circuitID, r.err)
			failures = append(failures, fmt.Sprintf("%s: %v", r.circuitID, r.err))
			continue
		}
		fmt.Printf("%s... ok\n", r.circuitID)
	}

	if len(failures) > 0 {
		sort.Strings(failures)
		fmt.Fprintln(os.Stderr, "\nFailed circuits (left as whatever was previously generated, if anything):")
		for _, failure := range failures {
			fmt.Fprintln(os.Stderr, " ", failure)
		}
		os.Exit(1)
	}
}

func generate(client *http.Client, outputDir, circuitID, geometryID string) error {
	coordinates, name, err := trackmap.FetchGeometry(client, trackmap.DefaultGeometryBase, geometryID)
	if err != nil {
		return err
	}
	svg, err := trackmap.RenderSVG(coordinates, name, referenceColour)
	if err != nil {
		return err
	}
	return os.WriteFile(filepath.Join(outputDir, circuitID+".svg"), svg, 0o644)
}
