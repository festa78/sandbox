load("@crate_index//:defs.bzl", "aliases", "all_crate_deps")
load("@io_bazel_rules_docker//rust:image.bzl", "rust_image")
load("@rules_rust//rust:defs.bzl", "rust_binary")
load("@rules_rust//rust/platform:platform.bzl", "declare_config_settings")

declare_config_settings()

platform(
    name = "my_host",
    constraint_values = [
        "@platforms//os:linux",
        "@platforms//cpu:x86_64",
        # "@platforms//cpu:aarch64",
    ],
)

platform(
    name = "my_target",
    constraint_values = [
        "@platforms//os:linux",
        "@platforms//cpu:x86_64",
        # "@platforms//cpu:aarch64",
    ],
)

# rust_binary(
#     name = "hello",
#     srcs = ["src/main.rs"],
#     aliases = aliases(),
#     edition = "2021",
#     proc_macro_deps = all_crate_deps(
#         proc_macro = True,
#     ),
#     visibility = ["//visibility:public"],
#     deps = all_crate_deps(
#         normal = True,
#     ),
# )

# rust_image(
#     name = "hello_image",
#     base = "@cc_image_base_arm64//image",
#     binary = ":hello",
# )

rust_image(
    name = "hello_image",
    srcs = glob(["src/*.rs"]),
    aliases = aliases(),
    # base = "@cc_image_base_arm64//image",
    edition = "2021",
    proc_macro_deps = all_crate_deps(
        proc_macro = True,
    ),
    target_compatible_with = [
        "@platforms//os:linux",
        "@platforms//cpu:x86_64",
    ],
    visibility = ["//visibility:public"],
    deps = all_crate_deps(
        normal = True,
    ),
)
