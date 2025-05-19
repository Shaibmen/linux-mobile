package handlers

import (
	"net/http"
	"server/logging"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/shirou/gopsutil/cpu"
)

var (
	cpuValue []float64
	mu       sync.Mutex
)

func CpuData() {
	for {
		percent, err := cpu.Percent(0, false)
		if err != nil || len(percent) == 0 {
			logging.Log.Info("Ошибка во время сбора информации о CPU")
			continue
		}

		mu.Lock()
		cpuValue = append(cpuValue, percent[0])
		if len(cpuValue) > 10 {
			cpuValue = cpuValue[1:]
		}
		mu.Unlock()

		time.Sleep(5 * time.Second)
	}
}

func GetCPU(c *gin.Context) {

	mu.Lock()
	values := append([]float64{}, cpuValue...)
	mu.Unlock()

	c.JSON(http.StatusOK, gin.H{"cpu": values})

}
