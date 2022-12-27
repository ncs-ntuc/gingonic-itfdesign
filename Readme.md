

### Goals for re-design 
1. Keep the design loosely coupled. respecting SRP
2. Keep the implementations open for extension and closed for modifications.
3. The middleware chains are alredy in a pattern with  context of the request being the visitor. Its grossly over engineering to fit a pattern there  ?
4. Cart data objects can be in interface based design using interfaces.
5. Error framework needs to be connected to http response and status codes.Logging needs to be connected to error frameworks
6. whats missing - we arent able to take advantage of Go's packages concept so that reusability across projects and u-services