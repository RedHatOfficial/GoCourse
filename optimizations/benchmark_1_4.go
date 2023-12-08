func BenchmarkGetStorageConfigurationFunctionByValue(b *testing.B) {
        b.StopTimer()
        configuration := mustLoadBenchmarkConfiguration(b)
        b.StartTimer()

        for i := 0; i < b.N; i++ {
                conf.GetStorageConfigurationByValue(configuration)
        }
}

func BenchmarkGetStorageConfigurationFunctionByReference(b *testing.B) {
        b.StopTimer()
        configuration := mustLoadBenchmarkConfiguration(b)
        b.StartTimer()

        for i := 0; i < b.N; i++ {
                conf.GetStorageConfigurationByReference(&configuration)
        }
}
