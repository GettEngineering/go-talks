f, err := os.Create("trace.out")
// error checks, file closing, etc

if err := trace.Start(f); err != nil {
	log.Fatalf("failed to start trace: %v", err)
}
defer trace.Stop()

// your program here