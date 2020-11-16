package astress

import (
    "fmt"
    "testing"
    "time"
)

func TestQuickStress(t *testing.T) {
    config := Config{StartNumber: 100, EndSecond: 60 * 10, StepSecond: 10, StepNumber: 10, MaxNumber: 300}
    stressFunc := func() error {
        fmt.Println(time.Now().UnixNano())
        return nil
    }
    QuickStress(config, stressFunc)
}
