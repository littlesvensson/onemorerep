package pr

import (
	"errors"
	"math"
	"strconv"
)

/* At the moment, only Brzycki is implemented. There is more than one way
to estimate 1RM, and they differ more at higher reps. Keepinng it simple for now,
as it's like with watches - if you have more of them, you are not sure what time it is :)
*/
func Brzycki1RM(weight float64, reps int) (float64, error) { 
    den := 37.0 - float64(reps)
    if den <= 0 {
        return 0, errors.New("Cannot estimate PRs for reps above 36")
    }
    return weight * 36.0 / den, nil
}

func WeightAtReps(oneRM float64, reps int) float64 {
    return oneRM * ((37.0 - float64(reps)) / 36.0)
}

func RepTable(oneRM float64, maxReps int) map[string]float64 {
    if maxReps < 1 {
        maxReps = 1
    }
    if maxReps > 30 {
        maxReps = 30
    }
    out := make(map[string]float64, maxReps)
    for r := 1; r <= maxReps; r++ {
        out[strconv.Itoa(r)] = round2(WeightAtReps(oneRM, r)) // Itoa = integer to ascii
    }
    return out
}

func round2(x float64) float64 {
    return math.Round(x*100) / 100
}
