load("//.macros:go_library.bzl", "go_library")
load("//.macros:go_linter.bzl", "go_linter")

go_library(
    name = "lib",
    deps = ["//third-party/github.com/uber-go/tally:lib"],
    package_name = "go.uber.org/cff",
    srcs = glob(
        include = ["*.go"],
        exclude = ["*_test.go"],
    ),
    visibility = ["PUBLIC"],
)

go_linter(
    name = "linter",
    srcs = glob(
        include = ["*.go"],
        exclude = [
            "mock_*.go",
            "mocks.go",
            "*_mock.go",
            "*_mock_test.go",
            "*_string.go",
        ],
    ),
)