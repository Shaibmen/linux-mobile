package handlers

import (
	"net/http"
	"server/logging"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/shirou/gopsutil/v3/mem"
)

var (
	ramValue []float64
	ramAll   uint64
	muRam    sync.Mutex
)

func RamData() {
	for {
		memory, err := mem.VirtualMemory()
		if err != nil {
			logging.Log.Info("Ошибка во время сбора информации о RAM")
			continue
		}

		muRam.Lock()
		ramValue = append(ramValue, memory.UsedPercent)
		if len(ramValue) > 10 {
			ramValue = ramValue[1:]
		}
		ramAll = memory.Total
		muRam.Unlock()

		time.Sleep(5 * time.Second)
	}
}

func GetRAM(c *gin.Context) {
	muRam.Lock()
	value := append([]float64{}, ramValue...)
	all := ramAll
	muRam.Unlock()

	c.JSON(http.StatusOK, gin.H{
		"ram": value,
		"all": all,
	})

}
