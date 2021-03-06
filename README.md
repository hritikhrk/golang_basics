# This project is for personal learning and documentation with resources. 

## Reading Materials

### Basic Learning
*  https://www.golang-book.com/books/intro

### Concurrency in action 
*  http://divan.github.io/posts/go_concurrency_visualize/
*  https://codeburst.io/why-goroutines-are-not-lightweight-threads-7c460c1f155f
*  https://golangbot.com/channels/
*  https://golangbot.com/goroutines/

### Performance Caveats 
*  http://container-solutions.com/surprise-golang-thread-scheduling/
* Defer, Panic and Recover: https://blog.golang.org/defer-panic-and-recover
### Project Structure

![alt text](https://github.com/nanofaroque/golang_basics/blob/master/project_structure.png)

* https://golang.org/doc/code.html

## Golang microservice with Kafka and MongoDB 
* https://www.melvinvivas.com/developing-microservices-using-kafka-and-mongodb/

### Quick Note
#### Iterating over map:

   ```
   for k, v := range m { 
       fmt.Printf("key[%s] value[%s]\n", k, v)
   }
   or
   
   for k := range m {
       fmt.Printf("key[%s] value[%s]\n", k, m[k])
   }
   ```
   
#### Declaration of Map:
   ```
     m := make(map[int]bool)
   ```
   To get the data from map
   ```cassandraql
     m[121]
   ```
#### Linked List 
   Creation:
    
    list.New() 
    
   Accessing element from the list:
   ```cassandraql

   type Node struct {
	 row int
	 col int
   }

   for e := l.Front(); e != nil; e = e.Next() {
   		row := e.Value.(*Node).row
   		col := e.Value.(*Node).col
   }
   ``` 