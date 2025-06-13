package dataSrc

import "context"

type DataSrc interface {
	ReadLine(ctx context.Context, dest chan string)
}
