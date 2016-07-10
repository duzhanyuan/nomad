// +build darwin dragonfly freebsd netbsd openbsd solaris

package fingerprint

func initFingerprints(fps map[string]Factory) {
	fps["arch"] = NewArchFingerprint
	fps["cpu"] = NewCPUFingerprint
	fps["env_aws"] = NewEnvAWSFingerprint
	fps["env_gce"] = NewEnvGCEFingerprint
	fps["host"] = NewHostFingerprint
	fps["memory"] = NewMemoryFingerprint
	fps["network"] = NewNetworkFingerprint
	fps["nomad"] = NewNomadFingerprint
	fps["storage"] = NewStorageFingerprint
}