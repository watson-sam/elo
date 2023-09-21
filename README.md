# Elo 

The Elo Rating System is a widely used method for calculating the relative skill levels of players in two-player games such as chess, go, and various sports competitions. This Go package provides essential functions and structs for implementing the Elo Rating System in your application.

## Installation

You can install the `elo` package using Go modules. Open your terminal and run the following command:

```bash
go get github.com/watson-sam/elo
```

## Usage 

Here's how you can use the elo package in your Go application:

```go
package main

import (
	"fmt"
	"github.com/watson-sam/elo"
)

func main() {
	// Create Elo settings with custom parameters
	settings := elo.New(
		elo.WithKFactor(32), // Specify the K-factor for rating adjustments
		elo.WithExpectedFunction(elo.ExpProbability), // Use a custom expected function (optional)
		elo.WithObservedFunction(elo.ObsWinLooseDraw), // Use a custom observed function (optional)
	)

	// Initial ratings for two players
	playerA := settings.InitRating
	playerB := settings.InitRating

	// Simulate a match result
	scoreA := 1.0 // Adjust this based on the actual match result
	scoreB := 0.0 // Adjust this based on the actual match result

	// Calculate updated ratings
	newRatingA := settings.UpdateRating(playerA, playerB, scoreA, scoreB)
	newRatingB := settings.UpdateRating(playerB, playerA, scoreB, scoreA)

	fmt.Printf("Player A's new rating: %.2f\n", newRatingA)
	fmt.Printf("Player B's new rating: %.2f\n", newRatingB)
}
```

This example demonstrates how to create Elo settings with custom parameters and use them to calculate updated ratings after a match. You can customize the package's behavior by adjusting the settings and using different update, expected, and observed functions.


## Contributing
If you'd like to contribute to this package or report issues, please visit the [GitHub repository](https://github.com/watson-sam/elo).

## License
This Elo Rating System Package is open-source software licensed under the MIT License.