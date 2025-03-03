package ports

type IRandomStringGenerator interface {
	Generate(length int) string
}
