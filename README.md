# go-usrv-tmpl
golang microservice template

# explain structure
everything has been given a generalized name in order to make it as accomodating as possible while retaining meaning and being runnable.
below is an explanation of all directories and files included in this repository.
## domain
first is the domain directory where the core interest of the application is held.
let's look at the files (ordered by importance):
- `model.go` includes the main object, in the form of a struct, around which the service is built
- `service.go` includes the interface of the service that should be implemented
- `logic.go` includes the implementation of the above service
- `repo.go` includes the interface of what a storage repository should implement
- `serializer.go` includes the interface of what a format serializer should implement

the workflow is as follows:
- define the main structure of concern in `model.go`
- define the things that should be done with the above in `service.go`
- define the implementation of above things in `logic.go`
- if applicable define the actions related to a data store in `repo.go`
- if applicable define the actions related to a serialization format in `serializer.go`

that was the most difficult part, if you have this down then you're good to go.

## repo
in here there are sub-directories of the different types of data stores that can be used with no change to the rest of the code.
currently that is `mock` and `postgres` with more likely to be included in the future.
inside each of those directories there is a `repo.go` file which implements the interface defined in `domain/repo.go`.
that's it, whatever involves the specific data store is held here while the rest of the code just calls the methods define in `domain/repo.go` to perform a certain action.

## serializer
very similar to the above, under the `serializer` top-level directory there are sub-directories, one directory for each serialization format.
currently that includes `json`, `msgpack` and `xml` with more likely to be included in the future.
there are two files under each of those directories, `serializer.go` and `serializer_test.go`.
the former implements the interface defined in `domain/serializer.go` and the latter includes some tests.
whenever encoding or decoding functionality is needed, the code calls the appropriate method on the service and gets back what it needs.

## protos
this is somewhat optional and can be removed but grpc is pretty nice in my humble opinion.
the `thing.protos` file includes the definition for what is sent and what is received by the service.
golang code is then generated, by running `make protobufs` from the root of the repo,
afterwards, it can be included in order to create grpc servers and clients in order to perform remote procedure calls.

## api
this has proven problematic to implement as there is no good way to provide a common interface for rest, in the form of serialized format over http, as well as grpc.
the most likely solution at this point will be only a grpc api which can then be made available as rest through [a gateway](https://github.com/grpc-ecosystem/grpc-gateway).
