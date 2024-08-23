package timestorage

import (
	"context"
	"github.com/influxdata/influxdb-client-go/v2/api/write"
	"github.com/labstack/gommon/log"
)

func (c *InfluxWriter) Write(point *write.Point) error {
	c.Wmu.Lock()
	err := c.Writer.WritePoint(context.Background(), point)
	c.Wmu.Unlock()
	if err != nil {
		log.Error(err)
		return err // 需要返回吗，一个没写上而已
	}
	return nil
}
