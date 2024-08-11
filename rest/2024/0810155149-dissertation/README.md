https://ics.uci.edu/~fielding/pubs/dissertation/rest_arch_style.htm

# deriving rest

design rationale behind web architecture is a set of constraints applied to elements within the arch

additional constraints can be applied to form new architectural styles that better reflect the desired properties of a modern web arch

## starting with the null style

on the process of arch design, two perspectives:
- designer starts with nothing, and builds-up to satisfy the needs of the intended system
- designer starts with the system without constraints, then incrementally identifies and applies constraints to elements of said system to differentiate the design spaces
  - allows the forces that influence the system behavior to flow naturally

first emphasizes creativity, second emphasizes restraint and understanding of the system context. REST implements the latter

null styler is an empty set of constraints, a system in which there are no distinguished boundaries between components; acts as a starting point for our desc of rest, e.g. `www`

## client-server

the first constraints added to the hybrid style are those of the client-server architecture, described [here](https://ics.uci.edu/~fielding/pubs/dissertation/net_arch_styles.htm#sec_3_4_1) in section 3.4.1

//

notes on section 3.4.1:
- server component, offering a set of services, listens for requests upon those services
- client component, desiring that a service be performed, sends a request to the server via a connector
- server either rejects or performs the request, sends a response back to the client
- client is a triggering process, server is a reactive process
- does not constrain how app state is partitioned, often it is referred to by the mechs used for the connector such as RPC or message-oriented middleware

//

intention is to seperate ui concerns from data storage concerns
- improves portability of the ui across platforms
- improves scalability of serrver components

most significantly, the separation allows the components to evolve interdependently

## stateless

communication must be stateless; client-stateless-server described [here](https://ics.uci.edu/~fielding/pubs/dissertation/net_arch_styles.htm#sec_3_4_3) in section 3.4.3

//

notes on section 3.4.3:
- client-server, but no session states allowed on the server component
- every request from client to server must contain all info neccesarry to understand the request + cannot take advantage of any stored context on the server
- session state is kept on client
- improves visibility, reliability & scalability
- decreases network performance by increasing the repetitive data (per-interaction-overhead) sent in requests

//

## cache

to improve network efficiency from the decrease in network performance via statelessness, we add cache constraints to form the client-cache-stateless-server model described [here](https://ics.uci.edu/~fielding/pubs/dissertation/net_arch_styles.htm#sec_3_4_4) in section 3.4.4

//

notes on section 3.4.4:
- cache constraints require that the data within a response be labeled as cacheable/non-cacheable
- if cacheable, client cache is given the right to reuse the response data for later requests
- cache acts as a mediator between client/server in which responses to prior requests can be reused in response to later requests

//

cache can decrease reliability if stale data within the cache differs from server

early web was defined on ccss models
- did not constrain the communcation interfaces to a consistent set of semantics

the design rationale presented for the web architecture prior to 1994 focused on stateless client-server interaction for the exchange of static docs

protocols for communicating had rudimentary support for non-shared caches; did not constrain the interface to a constraint set of semantics for all resources; instead relying on CERN's libwww to maintain consistency across web apps

![diagram](https://ics.uci.edu/~fielding/pubs/dissertation/early_web_arch.gif)

requests could identify services that deynamically generate responses such as image-maps and server-sides scripts

work began on intermediary components, through proxies and shared caches; extensions to the protocols were needed for them to function reliably 

## uniform interface

rest's central distinguishing feature from other network-based styles is the emphasis on a uniform interface between components

apply the component of generality to component interface, architecture is simplified and visibility is improved

implementations are decoupled from the services they provide, which encourages independent evolvability
trade-off; uniform interfaces degrade efficiency, info transferred in a standardized form tends to be far more robust than one specific to an app's needs
- rest is designed to be efficient for large-grain hypermedia data transfer, aka the common case of the web, but it's result is an interface that is not optimal for other forms of arch interaction 

to obtain a uniform interface, multiple architectural constraints are established:
- identification of resources
- manipulation of resources through representation
- self-descriptive messages
- hypermedia as the engine of application state 

## layered system

to improve behavior for internet-scale requirements - add layered system constraints

layered system style allows an architecture to be composed of hierarchical layers by constraining comp behavior such that it cannot see beyond it's immediate layer 

restricting knowledge of the system to a single layer places a bound on system complexity and promotes substrate independence

layers can thusly be used to encapsulate legacy services/protect new services from legacy clients. this simplifies comps to a shared intermediary

intermediaries can be used to improve system scalability by enabling load balancing of services across networks/processors

![diagram](https://ics.uci.edu/~fielding/pubs/dissertation/layered_uccss.gif)

disadvantages:
- they add overhead & latency to the processing of data

can be offset by the benefits of shared caching at intermediaries 
- shared caches at the boundaries of organizational domains can also result in significant performance benefits
- such layers also allow security policies to be enforced on data crossing the organizational boundary, as required by firewalls

layered systems + uniform interface constraints induce architectural properties similar to the uniform pipe-filter style 

rest interaction is two-way, but large-grain data flows of hypermedia interaction can be processed like a data-flow network, with filter comp selectively applied to the stream to transform the content as it passes
- within rest, intermediary components can transform the content of messages because they are self-descriptive, and their semantics are visible to intermediaries 

## code-on-demand

allows client functionality to be extended, by downloading/executing code in applets/scripts; simplifying clients by reducing the num of features required to be pre-implemented
- improves system extensibility, but reduces visibility 

this is an optional constraint

![diagram](https://ics.uci.edu/~fielding/pubs/dissertation/rest_style.gif)


