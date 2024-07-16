package decorator

import (
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
)

type queryLoggingDecorator[C any, R any] struct {
	base   QueryHandler[C, R]
	logger *logrus.Entry
}

func (ql queryLoggingDecorator[C, R]) Handle(ctx context.Context, cmd C) (result R, err error) {
	logger := ql.logger.WithFields(logrus.Fields{
		"query":      "Test",
		"query_body": fmt.Sprintf("%#v", cmd),
	})

	logger.Debug("Executing query")
	//defer func() {
	//	if err == nil {
	//		logger.Info("query executed successfully")
	//	} else {
	//		logger.WithError(err).Error("Failed to execute query")
	//	}
	//}()

	return ql.base.Handle(ctx, cmd)
}
