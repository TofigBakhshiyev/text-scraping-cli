### Text Scraping from Html Page
Getting text data from web page with URL and html tags
#### Requirements
- Go version 1.13.6 or later
#### Runing the app
from the main directory run:
```
go run main.go -tags "p" -url "https://medium.com/better-programming/modern-day-architecture-design-patterns-for-software-professionals-9056ee1ed977"
```
example output
<br/>
<p>[Many modern-day applications need to be built at an enterprise scale, sometimes even at an internet scale. Each application needs to meet scalability, availability, security, reliability, and resiliency demands. In this article, I’m going to talk about some design patterns that can help you achieve the above-mentioned abilities with ease. ...]</p>
<br/>
for more information run:
```
go run main.go -help
```