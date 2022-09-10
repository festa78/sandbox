load(
    "@rules_cc//cc:cc_toolchain_config_lib.bzl",
    "action_config",
    "feature",
    "flag_group",
    "flag_set",
    "tool",
    "with_feature_set",
)
load(
    "@bazel_tools//tools/build_defs/cc:action_names.bzl",
    "ACTION_NAMES",
    "CPP_COMPILE_ACTION_NAME",
    "CPP_LINK_EXECUTABLE_ACTION_NAME",
)

def _impl(ctx):
    return cc_common.create_cc_toolchain_config_info(
        ctx = ctx,
        toolchain_identifier = "custom-toolchain-identifier",
        host_system_name = "local",
        target_system_name = "local",
        target_cpu = "sample_cpu",
        target_libc = "unknown",
        compiler = "gcc",
        abi_version = "unknown",
        abi_libc_version = "unknown",
        action_configs = [
            action_config(
                action_name = CPP_COMPILE_ACTION_NAME,
                enabled = True,
                tools = [
                    tool(
                        path = "/usr/bin/g++",
                        with_features = [
                            with_feature_set(features = ["std-flag"]),
                        ],
                    ),
                ],
            ),
            action_config(
                action_name = CPP_LINK_EXECUTABLE_ACTION_NAME,
                enabled = True,
                tools = [tool(path = "/usr/bin/ld")],
            ),
        ],
        features = [
            feature(
                name = "std-flag",
                enabled = True,
                flag_sets = [
                    flag_set(
                        actions = [
                            ACTION_NAMES.cpp_compile,
                        ],
                        flag_groups = [
                            flag_group(
                                flags = [
                                    "-std=c++17",
                                ],
                            ),
                        ],
                    ),
                ],
            ),
        ],
    )

cc_toolchain_config = rule(
    implementation = _impl,
    attrs = {},
    provides = [CcToolchainConfigInfo],
)
