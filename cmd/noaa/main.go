package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/clairBuoyant/swellhub/pkg/noaa"
)

func main() {
	if len(os.Args) < 2 {
		printUsageAndExit()
	}

	activeCmd := flag.NewFlagSet("active", flag.ExitOnError)
	realtimeCmd := flag.NewFlagSet("realtime", flag.ExitOnError)

	stationID := realtimeCmd.String("stationID", "", "StationID is required for realtime")

	switch os.Args[1] {
	case "realtime":
		realtimeCmd.Parse(os.Args[2:])

		if *stationID == "" {
			fmt.Println("Error: stationID is required for realtime mode")
			printUsageAndExit()
		}

		mos, err := noaa.Realtime(*stationID, noaa.TXT)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error running Realtime: %v\n", err)
			os.Exit(1)
		}
		for _, mo := range mos {
			fmt.Printf("Observation Date: %s, Average Wave Period: %f, Wave Height: %f\n", mo.Datetime, mo.AverageWavePeriod, mo.WaveHeight)
		}

	case "active":
		activeCmd.Parse(os.Args[2:])

		stations, err := noaa.ActiveStations()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error running ActiveStations: %v\n", err)
			os.Exit(1)
		}
		for _, station := range stations.Stations {
			fmt.Printf("ID: %s, Name: %s, Lat: %f, Lon: %f\n", station.ID, station.Name, station.Lat, station.Lon)
		}

	default:
		fmt.Println("Error: invalid mode. Use 'realtime' or 'active'")
		printUsageAndExit()
	}
}

func printUsageAndExit() {
	fmt.Println("Usage: noaa <mode> [arguments]")
	fmt.Println("Modes:")
	fmt.Println("  realtime -stationID=<id>  Get real-time data for the specified station")
	fmt.Println("  active                    Get data for active stations")
	os.Exit(1)
}
