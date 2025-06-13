package dataSrc

type DataSrc interface {
	ReadLine(dest chan string)
}
