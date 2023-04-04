package client

import (
	"bufio"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/keys"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/spf13/cobra"
	"github.com/tendermint/tendermint/libs/cli"

	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	ethclient "github.com/evmos/ethermint/client"
	clientkeys "github.com/evmos/ethermint/client/keys"
	"github.com/evmos/ethermint/crypto/hd"
)

var ethFlag = "eth"

// KeyCommands registers a sub-tree of commands to interact with
// local private key storage.
func KeyCommands(defaultNodeHome string) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "keys",
		Short: "Manage your application's keys",
		Long: `Keyring management commands. These keys may be in any format supported by the
Tendermint crypto library and can be used by light-clients, full nodes, or any other application
that needs to sign with a private key.

The keyring supports the following backends:

    os          Uses the operating system's default credentials store.
    file        Uses encrypted file-based keystore within the app's configuration directory.
                This keyring will request a password each time it is accessed, which may occur
                multiple times in a single command resulting in repeated password prompts.
    kwallet     Uses KDE Wallet Manager as a credentials management application.
    pass        Uses the pass command line utility to store and retrieve keys.
    test        Stores keys insecurely to disk. It does not prompt for a password to be unlocked
                and it should be use only for testing purposes.

kwallet and pass backends depend on external tools. Refer to their respective documentation for more
information:
    KWallet     https://github.com/KDE/kwallet
    pass        https://www.passwordstore.org/

The pass backend requires GnuPG: https://gnupg.org/
`,
	}

	// support adding Ethereum supported keys
	addCmd := keys.AddKeyCommand()
	addCmd.Flags().Bool(ethFlag, false, "use default evm coin-type (60) and key signing algorithm (\"eth_secp256k1\")")

	addCmd.RunE = runAddCmd

	cmd.AddCommand(
		keys.MnemonicKeyCommand(),
		addCmd,
		keys.ExportKeyCommand(),
		keys.ImportKeyCommand(),
		keys.ListKeysCmd(),
		keys.ShowKeysCmd(),
		flags.LineBreak,
		keys.DeleteKeyCommand(),
		keys.ParseKeyStringCommand(),
		keys.MigrateCommand(),
		flags.LineBreak,
		ethclient.UnsafeExportEthKeyCommand(),
		ethclient.UnsafeImportKeyCommand(),
	)

	cmd.PersistentFlags().String(flags.FlagHome, defaultNodeHome, "The application home directory")
	cmd.PersistentFlags().String(flags.FlagKeyringDir, "", "The client Keyring directory; if omitted, the default 'home' directory will be used")
	cmd.PersistentFlags().String(flags.FlagKeyringBackend, keyring.BackendOS, "Select keyring's backend (os|file|test)")
	cmd.PersistentFlags().String(cli.OutputFlag, "text", "Output format (text|json)")
	return cmd
}

func runAddCmd(cmd *cobra.Command, args []string) error {
	buf := bufio.NewReader(cmd.InOrStdin())
	clientCtx := client.GetClientContextFromCmd(cmd)

	var (
		kr  keyring.Keyring
		err error
	)

	dryRun, _ := cmd.Flags().GetBool(flags.FlagDryRun)
	if dryRun {
		kr, err = keyring.New(
			sdk.KeyringServiceName(),
			keyring.BackendMemory,
			clientCtx.KeyringDir,
			buf,
			clientCtx.Codec,
			hd.EthSecp256k1Option(),
		)
		clientCtx = clientCtx.WithKeyring(kr)
	}
	eth, _ := cmd.Flags().GetBool(ethFlag)
	if eth {
		cmd.Print("eth flag specified: using coin-type 60 and signing algorithm eth_secp256k1\n")
		cmd.Flags().Set(flags.FlagKeyAlgorithm, string(hd.EthSecp256k1Type))
		cmd.Flags().Set("coin-type", "60")
	}

	if err != nil {
		return err
	}

	return clientkeys.RunAddCmd(clientCtx, cmd, args, buf)
}
