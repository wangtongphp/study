package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"sync"
)

func main() {

	//main1()
	main2()

}

// map race 无法recover
func main1() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("recover", err)
		}
	}()
	m := map[int]int{
		1: 1,
		2: 2,
	}
	go func() {
		defer func() {
			if err := recover(); err != nil {
				fmt.Println("recover", err)
			}
		}()
		for i := 0; i < 1111; i++ {
			v, ok := m[i]
			if ok != true || v == 0 {
				fmt.Println("err")
			}
		}
	}()
	for i := 0; i < 1111; i++ {
		m[i] = i
	}

	fmt.Println("vim-go",
		string(byte(01)), byte(01))
}

func main2() {
	s := []string{"a", "b", "c", "a", "b", "c", "a", "b", "c", "a", "b", "c", "a", "b", "c", "a", "b", "c", "a", "b", "c", "a", "b", "c", "a", "b", "c", "a", "b", "c", "a", "b", "c", "a", "b", "c"}
	ConcurrencyMapSearch(s)
}

//
type SearchInput struct {
	Query string
}

type SearchOutput struct {
	Query string
	Data  map[string]interface{}
}

func Search(input *SearchInput) (*SearchOutput, error) {
	output := new(SearchOutput)
	output.Query = input.Query

	// NOTE: This should retrieve from other service, such as service.Get(query). And I
	// realized that it use cached data with shared map by pointer among the same query, such as
	//
	//	type Experiment struct {
	//		Data map[string]interface{}
	//	}
	//
	//	type WhiteList map[string]*Experiment
	//
	//	func (list WhiteList) Get(query string) (*Experiment, bool) {
	//		exp, ok := list[query]
	//		return exp, ok
	//	}
	//
	//	exp, ok : whiteList.Get(query)
	//	if ok {
	//		output.Data = exp.Data
	//	}
	//
	output.Data = make(map[string]interface{})
	output.Data["filter_key"] = "value1"

	return output, nil
}

func ConcurrencyMapSearch(queries []string) {
	var (
		maxWorkers   = 10
		workerMapKey = "row_key" // for input row_key pass around

		workerInputs  = make(chan *SearchInput, 100)
		workerInputWg sync.WaitGroup

		workerOutputs  = make(chan *SearchOutput, 100)
		workerOutputWg sync.WaitGroup
	)

	// multiple workers for input processing
	workerInputWg.Add(maxWorkers)
	for i := 0; i < maxWorkers; i++ {
		go func() {
			defer workerInputWg.Done()

			for input := range workerInputs {
				output, err := Search(input)
				if err != nil {
					log.Println("search ", input.Query, " err:", err.Error())
				} else {
					output.Data[workerMapKey] = input.Query

					workerOutputs <- output
				}
			}
		}()
	}

	// NOTE: only one worker for output processing cause of file writer is not safe with concurrency!
	workerOutputWg.Add(1)
	go func(filterValues []string, filterKey string) {
		defer workerOutputWg.Done()

		for output := range workerOutputs {
			workerMapValue, ok := output.Data[workerMapKey].(string)
			if !ok {
				log.Println("invalid output")
				continue
			}

			for _, value := range filterValues {
				if value != output.Data[filterKey] {
					continue
				}

				// NOTE: It should write to some io.WriteCloser
				log.Println("search:", value)
				_, err := ioutil.Discard.Write([]byte(workerMapValue))
				if err != nil {
					log.Println("/dev/null:", err.Error())
				}
			}
		}
	}([]string{"value1", "value2"}, "filter_key")

	// trigger search
	for _, query := range queries {
		input := &SearchInput{
			Query: query,
		}

		workerInputs <- input
	}

	close(workerInputs)
	workerInputWg.Wait()

	close(workerOutputs)
	workerOutputWg.Wait()

	// do something else
}
