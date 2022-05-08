# GoSpider

GoSpider is a side project to study golang and concurrency; it crawls a website that you provide the gophere and logs every site.

# How to run ?

Just type in your console

```bash
go run main.go https://example.com
```

# Todo list
- [ ] Using pool of workers to avoid memory leak.
- [ ] Create HTTP interface to communicate through (push pull archtecture)

# Notes:
- This project is not done yet, it contains some bugs
- To Change depth of your search just change the constant in `main.go` file; might add a flag parameter later.