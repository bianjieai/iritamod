package utils

import (
	"bytes"
	"fmt"
	"os/exec"

	"github.com/tendermint/tendermint/crypto/algo"
)

func GenRootCert(keyPath, certPath, subj string) {
	switch algo.Algo {
	case algo.SM2:
		cmd := exec.Command("openssl", "ecparam", "-genkey", "-name", "SM2", "-out", keyPath)
		executeCmd(cmd)
	// ed25519
	default:
		cmd := exec.Command("openssl", "genpkey", "-algorithm", "ED25519", "-out", keyPath)
		executeCmd(cmd)
	}
	GenSelfSignCert(keyPath, certPath, subj)
}

func GenMultiRootCert(keyPath, certPath map[string]string, subj string) {
	// sm2
	cmd := exec.Command("openssl", "ecparam", "-genkey", "-name", "SM2", "-out", keyPath[algo.SM2])
	executeCmd(cmd)

	cmd = exec.Command(
		"openssl", "req", "-new", "-x509", "-sm3", "-sigopt", "distid:1234567812345678",
		"-key", keyPath[algo.SM2], "-subj", subj, "-out", certPath[algo.SM2], "-days", "365",
	)
	executeCmd(cmd)

	// ed25519
	cmd = exec.Command("openssl", "genpkey", "-algorithm", "ED25519", "-out", keyPath[algo.ED25519])
	executeCmd(cmd)

	cmd = exec.Command(
		"openssl", "req", "-new", "-x509",
		"-key", keyPath[algo.ED25519], "-subj", subj, "-out", certPath[algo.ED25519], "-days", "365",
	)
	executeCmd(cmd)
}

func GenSelfSignCert(keyPath, certPath, subj string) {
	switch algo.Algo {
	case algo.SM2:
		cmd := exec.Command(
			"openssl", "req", "-new", "-x509", "-sm3", "-sigopt", "distid:1234567812345678",
			"-key", keyPath, "-subj", subj, "-out", certPath, "-days", "365",
		)
		executeCmd(cmd)
	// ed25519
	default:
		cmd := exec.Command(
			"openssl", "req", "-new", "-x509",
			"-key", keyPath, "-subj", subj, "-out", certPath, "-days", "365",
		)
		executeCmd(cmd)
	}
}

func GenCertRequest(keyPath, cerPath, subj string) {
	switch algo.Algo {
	case algo.SM2:
		cmd := exec.Command(
			"openssl", "req", "-new", "-sm3", "-sigopt", "distid:1234567812345678",
			"-key", keyPath, "-subj", subj, "-out", cerPath,
		)
		executeCmd(cmd)
	// ed25519
	default:
		cmd := exec.Command(
			"openssl", "req", "-new",
			"-key", keyPath, "-subj", subj, "-out", cerPath,
		)
		executeCmd(cmd)
	}

}

func IssueCert(cerPath, caPath, caKeyPath, certPath string) {
	switch algo.Algo {
	case algo.SM2:
		cmd := exec.Command(
			"openssl", "x509", "-req", "-in", cerPath,
			"-CA", caPath, "-CAkey", caKeyPath, "-CAcreateserial", "-out", certPath, "-days", "365",
			"-sm3", "-sigopt", "distid:1234567812345678", "-vfyopt", "distid:1234567812345678",
		)
		executeCmd(cmd)
	// ed25519
	default:
		cmd := exec.Command(
			"openssl", "x509", "-req", "-in", cerPath,
			"-CA", caPath, "-CAkey", caKeyPath, "-CAcreateserial", "-out", certPath, "-days", "365",
		)
		executeCmd(cmd)
	}
}

func executeCmd(cmd *exec.Cmd) {
	var stdOut bytes.Buffer
	var stdErr bytes.Buffer
	cmd.Stderr = &stdErr
	cmd.Stdout = &stdOut
	if err := cmd.Run(); err != nil {
		fmt.Println(stdErr.String())
		panic(err)
	}
}
