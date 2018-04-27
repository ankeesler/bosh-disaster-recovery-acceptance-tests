package runner

type Config struct {
	BOSH          BOSHConfig
	BBRBinaryPath string
	ArtifactPath  string
}

type BOSHConfig struct {
	Hostname          string
	SSHUsername       string
	SSHPrivateKeyPath string
}
