package timestorage

import (
	"context"
	"github.com/influxdata/influxdb-client-go/v2/api/write"
)

func (c InfluxWriter) Write(point *write.Point) error {
	err := c.Writer.WritePoint(context.Background(), point)
	return err
}
