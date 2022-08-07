load("@crate_index//:defs.bzl", "aliases", "all_crate_deps")
load("@io_bazel_rules_docker//rust:image.bzl", "rust_image")

rust_image(
    name = "hello_rust",
    srcs = glob(["src/*.rs"]),
    aliases = aliases(),
    proc_macro_deps = all_crate_deps(
        proc_macro = True,
    ),
    visibility = ["//visibility:public"],
    deps = all_crate_deps(
        normal = True,
    ),
)
