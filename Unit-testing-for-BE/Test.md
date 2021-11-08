go test -v
go test -cover
go test -coverprofile=coverage.out
go tool cover -html=coverage.out
go test -run Get
go test -run Get/oct

---

# The go test command accepts a -covermode flag to set the coverage mode to one of three settings:

set: did each statement run?
count: how many times did each statement run?
atomic: like count, but counts precisely in parallel programs
