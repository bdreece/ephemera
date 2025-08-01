package version

var (
	version  string
	metadata string

	gitCommit string
	gitSha    string
	gitTag    string
)

func Info() (info struct {
	Version  string
	Metadata string

	GitCommit string
	GitSHA    string
	GitTag    string
}) {
	info.Version = version
	info.Metadata = metadata
	info.GitCommit = gitCommit
	info.GitSHA = gitSha
	info.GitTag = gitTag
	return
}
