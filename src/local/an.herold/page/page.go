package page

type Hi struct {
	Word string
}

func SayHello() string {
	word := Hi{}
	w := &word
	w.Word = "bl"

	return word.Word
}
