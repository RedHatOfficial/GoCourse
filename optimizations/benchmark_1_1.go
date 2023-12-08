package conf_test

// Benchmark for config module

import (
        "os"
        "testing"

        "config-struct/conf"
)

const (
        configFileEnvVariableName = "CCX_NOTIFICATION_SERVICE_CONFIG_FILE"
        defaultConfigFileName     = "./config"
)
