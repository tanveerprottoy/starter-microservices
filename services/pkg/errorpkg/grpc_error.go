package errorpkg

type GRPCError struct {
	Code int
	Err  error
}
