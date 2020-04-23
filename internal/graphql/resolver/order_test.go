package resolver

import (
	"context"
	"os"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func testLogger() *logrus.Entry {
	log := logrus.New()
	log.Out = os.Stderr
	log.Formatter = &logrus.JSONFormatter{
		TimestampFormat: "02/Jan/2006:15:04:05",
	}
	log.Level = logrus.PanicLevel
	return logrus.NewEntry(log)

}

func TestOrder(t *testing.T) {
	assert := assert.New(t)
	ord := &QueryResolver{
		Registry: &mockRegistry{},
		Logger:   testLogger(),
	}
	_, err := ord.Order(context.Background(), "999")
	assert.NoError(err, "expect no error from getting order information")
	// assert.Exactly(name, "DBS0236922", "should match systematic name")
}
