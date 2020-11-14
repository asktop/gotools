package astress

import (
    "fmt"
    "testing"
    "time"
)

func TestQuickStress(t *testing.T) {
    config := Config{StartNumber: 10, StepSecond: 1, StepNumber: 2, EndNumber: 20}
    stressFunc := func() error {
        fmt.Println(time.Now().UnixNano())
        return nil
    }
    QuickStress(config, stressFunc)
}
