package generator

import (
	"fmt"
	"harvester/utils"
	"log"
	"net/url"
	"strings"
	"sync"

	"github.com/cheggaaa/pb/v3"
)

// Struct for Generate urls
type Generator struct {
	Urls            []string
	Extensions      []string
	NumberOfWorkers int
	OutputFilePath  string
}

// Struct for holding job
type Job struct {
	id        int
	url       string
	extension string
}

// Struct for holding results
type Result struct {
	job  Job
	urls []string
}

// jobsChannel is a channel that holds jobs
var jobsChannel chan Job

// resultChannel is a channel that holds results
var resultsChannel chan Result

// bar is a progress bar
var bar *pb.ProgressBar

// WriteUrls generates and writes urls to output file
func (g *Generator) WriteUrls() {

	jobsChannel = make(chan Job, g.NumberOfWorkers)
	resultsChannel = make(chan Result, g.NumberOfWorkers)
	done := make(chan bool)

	bar = pb.StartNew(len(g.Urls) * len(g.Extensions))

	go allocate(g.Urls, g.Extensions)
	go result(done, g.OutputFilePath)

	createWorkerPool(g.NumberOfWorkers)
	<-done

	bar.Finish()
	fmt.Println("Done")
}

// process generates urls
func process(job Job) []string {
	var results []string

	parsedUrl, err := url.Parse(job.url)
	if err != nil {
		log.Fatal(err)
	}

	var domain string
	var tld string

	dot := strings.LastIndexByte(parsedUrl.Host, '.')
	if dot != -1 {
		domain, tld = parsedUrl.Host[:dot], parsedUrl.Host[dot+1:]
	}

	results = append(results, fmt.Sprintf("%s://%s.%s/%s.%s", parsedUrl.Scheme, domain, tld, domain, job.extension))
	results = append(results, fmt.Sprintf("%s://%s.%s/%s.%s.%s", parsedUrl.Scheme, domain, tld, domain, tld, job.extension))
	results = append(results, fmt.Sprintf("%s://%s.%s/www.%s.%s", parsedUrl.Scheme, domain, tld, domain, job.extension))
	results = append(results, fmt.Sprintf("%s://%s.%s/www.%s.%s.%s", parsedUrl.Scheme, domain, tld, domain, tld, job.extension))

	bar.Increment()
	return results
}

// worker is a worker that generates urls and sends them to resultsChannel
func worker(wg *sync.WaitGroup) {
	for job := range jobsChannel {
		resultsChannel <- Result{job, process(job)}
	}
	wg.Done()
}

// createWorkerPool creates worker pool with specified number of workers
func createWorkerPool(numberOfWorkers int) {
	defer close(resultsChannel)
	var wg sync.WaitGroup
	for i := 0; i < numberOfWorkers; i++ {
		wg.Add(1)
		go worker(&wg)
	}
	wg.Wait()
}

// allocate allocates jobs to jobsChannel
func allocate(urls []string, extensions []string) {
	defer close(jobsChannel)
	for i, url := range urls {
		for j, extension := range extensions {
			index := i*len(extensions) + j
			jobsChannel <- Job{index, url, extension}
		}
	}
}

// result append urls to output file
func result(done chan bool, OutputFilePath string) {
	for resultChannel := range resultsChannel {
		utils.AppendLines(resultChannel.urls, OutputFilePath)
	}
	done <- true
}
