package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/tyler-smith/go-bip39"
)

type Secret struct {
	Phrase string
}

func getSecrets(entropySize, amount int) ([]Secret, error) {
	secrets := make([]Secret, 0, amount)
	for i := 0; i < amount; i++ {
		entropy, err := bip39.NewEntropy(entropySize)
		if err != nil {
			return nil, fmt.Errorf("failed to generate entropy: %w", err)
		}

		mnemonic, err := bip39.NewMnemonic(entropy)
		if err != nil {
			return nil, fmt.Errorf("failed to generate mnemonic: %w", err)
		}

		secrets = append(secrets, Secret{Phrase: mnemonic})
	}
	return secrets, nil
}

func writeToFile(secrets []Secret, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer file.Close()

	for _, secret := range secrets {
		if _, err := file.WriteString(secret.Phrase + "\n"); err != nil {
			return fmt.Errorf("failed to write to file: %w", err)
		}
	}
	return nil
}

func main() {
	var amount int
	var filename string
	flag.IntVar(&amount, "amount", 10, "Number of secrets to generate -  default: 10")
	flag.StringVar(&filename, "output", "seedphrase.tx", "Output file name -  default: seedphrase.txt")
	flag.Parse()
	secrets, err := getSecrets(128, amount)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	if err = writeToFile(secrets, filename); err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}
	fmt.Printf("Secrets written to file %s successfully.\n", filename)
}
