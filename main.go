package main

import (
	"fmt"
)

func main()  {

	const numbersOfUrl = 1000
	const buffer = 5	

	done := make(chan bool)
	result := make(chan string,buffer)
	urlsMap := make(map[string]string)

	urlsMap = makeUrl(numbersOfUrl)
	
	go crawl(urlsMap,result)
	go printResult(result,done)

	<-done
	close(done)

}

func makeUrl(n int) map[string]string {
	ulrs := make(map[string]string)
	const exampleUrl = "http://example.com/"
	for i := 1; i <= n ; i++{
		url := fmt.Sprintf("%s%v",exampleUrl,i)
		data := fmt.Sprintf("%s: data", url)
		ulrs[url] = data
	}
	return ulrs
}

func crawl(urls map[string]string, result chan string)  {
	for key, value := range urls{
		result <- value
		fmt.Println("-----getting data from url:",key )
	}
	close(result)
}

func printResult(result chan string,done chan bool)  {
	for{
		data,next := <-result
		if(next){
			fmt.Println(data)
		}else{
			fmt.Println("*********** Done *********")
			done <- true
			return
		}
	}

}
	