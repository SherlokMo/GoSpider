package link

type Link struct {
	title, Url string
	links      *[]Link
}

func NewLink(url string) *Link {
	return &Link{
		Url: url,
	}
}
