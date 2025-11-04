package gateway

type Gateway interface{
	Process(amount int64) string
}