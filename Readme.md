

### Goals for re-design 
1. Keep the design loosely coupled. __respecting SRP__
2. Keep the implementations open for __extension and closed for modifications__.
3. The middleware chains are alredy in a pattern with  context of the request being the visitor. Its grossly over engineering to fit a pattern there  ?
4. Cart data objects can be in interface based design using interfaces.
5. Error framework needs to be connected to http response and status codes.Logging needs to be connected to error frameworks
6. Design of packages such that they can be used across projects