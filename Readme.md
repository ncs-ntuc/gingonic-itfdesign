1. The middleware chaining pattern is something that is desirable- why arent we switching to that ?
2. Cart implementations are just simple interface based design 
3. Error framework needs to be connected to http response and status codes 
4. Logging needs to be connected to error frameworks 
5. whats missing - we arent able to take advantage of Go's packages concept so that reusability across projects and u-services
6. testing data where and how is it populated?