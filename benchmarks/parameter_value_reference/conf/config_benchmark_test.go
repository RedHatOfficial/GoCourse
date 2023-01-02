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

// loadConfiguration function loads configuration prepared to be used by
// benchmarks
func loadConfiguration() (conf.ConfigStruct, error) {
	os.Clearenv()

	err := os.Setenv(configFileEnvVariableName, defaultConfigFileName)
	if err != nil {
		return conf.ConfigStruct{}, err
	}

	config, err := conf.LoadConfiguration(configFileEnvVariableName, defaultConfigFileName)
	if err != nil {
		return conf.ConfigStruct{}, err
	}

	return config, nil
}

func mustLoadBenchmarkConfiguration(b *testing.B) conf.ConfigStruct {
	configuration, err := loadConfiguration()
	if err != nil {
		b.Fatal(err)
	}
	return configuration
}

func BenchmarkGetStorageConfigurationFunctionByValue(b *testing.B) {
	b.StopTimer()
	configuration := mustLoadBenchmarkConfiguration(b)
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		// call benchmarked function
		conf.GetStorageConfigurationByValue(configuration)
	}

}

func BenchmarkGetStorageConfigurationFunctionByReference(b *testing.B) {
	b.StopTimer()
	configuration := mustLoadBenchmarkConfiguration(b)
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		// call benchmarked function
		conf.GetStorageConfigurationByReference(&configuration)
	}

}

func BenchmarkGetStorageConfigurationMethodByValue(b *testing.B) {
	b.StopTimer()
	configuration := mustLoadBenchmarkConfiguration(b)
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		// call benchmarked function
		configuration.GetStorageConfigurationByValue()
	}

}

func BenchmarkGetStorageConfigurationMethodByReference(b *testing.B) {
	b.StopTimer()
	configuration := mustLoadBenchmarkConfiguration(b)
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		// call benchmarked function
		configuration.GetStorageConfigurationByReference()
	}

}
