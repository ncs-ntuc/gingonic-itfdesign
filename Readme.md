

### My perspective on DDD when making u-services
----

In the context of u-services as is the amount of code that goes in to make the u-service is quite limited. __Strategically__ when split into packages of value models, controllers, aggregates, factories and services it can lead to code that is __more loosely coupled code than is required__. The code then appears to be over-engineered. On the other side the call stack is __bloated and often feels like a spaghetti.__ (Much difficult to understand after, debugging is tedious too.)

Tactically its a similar approach. - __Interfaces__ are at the center of everything. Implementations can be extended per change requests. If though here if I give the controllers, aggregates(u-service will seldom have aggregates since the object graph isnt that lengthy) a break (notice how I still use models, and services, factories) to design my code, it opens up an opportunity to make packages around the models. Packages that can be re-used in multiple projects and can be maintained independently with minimal imapct.

### Goals for re-design 
----

1. Keep the design loosely coupled. __respecting SRP__
2. Keep the implementations open for __extension and closed for modifications__.
3. The middleware chains are alredy in a pattern with  context of the request being the visitor. Its grossly over engineering to fit a pattern there  ?
4. Cart data objects can be in interface based design using interfaces.
5. Error framework needs to be connected to http response and status codes.Logging needs to be connected to error frameworks
6. Design of packages such that they can be used across projects