module github.com/hangingman/flutter-golang-integral

require (
	github.com/golang/protobuf v1.4.3
	github.com/hangingman/flutter-golang-integral/pb v0.0.0-00010101000000-000000000000
	github.com/hangingman/flutter-golang-integral/service v0.0.0-00010101000000-000000000000
	golang.org/x/net v0.0.0-20201110031124-69a78807bb2b // indirect
	golang.org/x/sys v0.0.0-20201112073958-5cba982894dd // indirect
	golang.org/x/text v0.3.4 // indirect
	google.golang.org/genproto v0.0.0-20201112120144-2985b7af83de // indirect
	google.golang.org/grpc v1.33.2
	google.golang.org/protobuf v1.25.0
)

go 1.13

replace (
	github.com/hangingman/flutter-golang-integral/pb => ./pb
	github.com/hangingman/flutter-golang-integral/service => ./service
)
