workspace(name = "trpc_protocol")

load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")

http_archive(
    name = "com_google_protobuf",
    sha256 = "0cbdc9adda01f6d2facc65a22a2be5cecefbefe5a09e5382ee8879b522c04441",
    strip_prefix = "protobuf-3.15.8",
    urls = [
        "https://github.com/protocolbuffers/protobuf/archive/v3.15.8.tar.gz",
    ],
)

load("@com_google_protobuf//:protobuf_deps.bzl", "protobuf_deps")

protobuf_deps()
