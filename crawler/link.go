package crawler

type Link struct {
	Title, Url string
	Links      *[]Link
}
