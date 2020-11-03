package client

// DON'T EDIT THIS FILE is generated 03 Nov 20 11:41 UTC
//
// Mod crypto
//
// Crypto functions.

import (
	"gopkg.in/guregu/null.v4"
)

type SigningBoxHandle struct {
}

type ParamsOfFactorize struct {
	// Hexadecimal representation of u64 composite number.
	Composite string `json:"composite"`
}

type ResultOfFactorize struct {
	// Two factors of composite or empty if composite can't be factorized.
	Factors []string `json:"factors"`
}

type ParamsOfModularPower struct {
	// `base` argument of calculation.
	Base string `json:"base"`
	// `exponent` argument of calculation.
	Exponent string `json:"exponent"`
	// `modulus` argument of calculation.
	Modulus string `json:"modulus"`
}

type ResultOfModularPower struct {
	// Result of modular exponentiation.
	ModularPower string `json:"modular_power"`
}

type ParamsOfTonCrc16 struct {
	// Input data for CRC calculation. Encoded with `base64`.
	Data string `json:"data"`
}

type ResultOfTonCrc16 struct {
	// Calculated CRC for input data.
	Crc int `json:"crc"`
}

type ParamsOfGenerateRandomBytes struct {
	// Size of random byte array.
	Length int `json:"length"`
}

type ResultOfGenerateRandomBytes struct {
	// Generated bytes encoded in `base64`.
	Bytes string `json:"bytes"`
}

type ParamsOfConvertPublicKeyToTonSafeFormat struct {
	// Public key - 64 symbols hex string.
	PublicKey string `json:"public_key"`
}

type ResultOfConvertPublicKeyToTonSafeFormat struct {
	// Public key represented in TON safe format.
	TonPublicKey string `json:"ton_public_key"`
}

type KeyPair struct {
	// Public key - 64 symbols hex string.
	Public string `json:"public"`
	// Private key - u64 symbols hex string.
	Secret string `json:"secret"`
}

type ParamsOfSign struct {
	// Data that must be signed encoded in `base64`.
	Unsigned string `json:"unsigned"`
	// Sign keys.
	Keys KeyPair `json:"keys"`
}

type ResultOfSign struct {
	// Signed data combined with signature encoded in `base64`.
	Signed string `json:"signed"`
	// Signature encoded in `hex`.
	Signature string `json:"signature"`
}

type ParamsOfVerifySignature struct {
	// Signed data that must be verified encoded in `base64`.
	Signed string `json:"signed"`
	// Signer's public key - 64 symbols hex string.
	Public string `json:"public"`
}

type ResultOfVerifySignature struct {
	// Unsigned data encoded in `base64`.
	Unsigned string `json:"unsigned"`
}

type ParamsOfHash struct {
	// Input data for hash calculation. Encoded with `base64`.
	Data string `json:"data"`
}

type ResultOfHash struct {
	// Hash of input `data`. Encoded with 'hex'.
	Hash string `json:"hash"`
}

type ParamsOfScrypt struct {
	// The password bytes to be hashed.
	// Must be encoded with `base64`.
	Password string `json:"password"`
	// A salt bytes that modifies the hash to protect against Rainbow table attacks.
	// Must be encoded with `base64`.
	Salt string `json:"salt"`
	// CPU/memory cost parameter.
	LogN int `json:"log_n"`
	// The block size parameter, which fine-tunes sequential memory read size and performance.
	R int `json:"r"`
	// Parallelization parameter.
	P int `json:"p"`
	// Intended output length in octets of the derived key.
	DkLen int `json:"dk_len"`
}

type ResultOfScrypt struct {
	// Derived key. Encoded with `hex`.
	Key string `json:"key"`
}

type ParamsOfNaclSignKeyPairFromSecret struct {
	// Secret key - unprefixed 0-padded to 64 symbols hex string.
	Secret string `json:"secret"`
}

type ParamsOfNaclSign struct {
	// Data that must be signed encoded in `base64`.
	Unsigned string `json:"unsigned"`
	// Signer's secret key - unprefixed 0-padded to 64 symbols hex string.
	Secret string `json:"secret"`
}

type ResultOfNaclSign struct {
	// Signed data, encoded in `base64`.
	Signed string `json:"signed"`
}

type ParamsOfNaclSignOpen struct {
	// Signed data that must be unsigned. Encoded with `base64`.
	Signed string `json:"signed"`
	// Signer's public key - unprefixed 0-padded to 64 symbols hex string.
	Public string `json:"public"`
}

type ResultOfNaclSignOpen struct {
	// Unsigned data, encoded in `base64`.
	Unsigned string `json:"unsigned"`
}

type ResultOfNaclSignDetached struct {
	// Signature encoded in `hex`.
	Signature string `json:"signature"`
}

type ParamsOfNaclBoxKeyPairFromSecret struct {
	// Secret key - unprefixed 0-padded to 64 symbols hex string.
	Secret string `json:"secret"`
}

type ParamsOfNaclBox struct {
	// Data that must be encrypted encoded in `base64`.
	Decrypted string `json:"decrypted"`
	// Nonce, encoded in `hex`.
	Nonce string `json:"nonce"`
	// Receiver's public key - unprefixed 0-padded to 64 symbols hex string.
	TheirPublic string `json:"their_public"`
	// Sender's private key - unprefixed 0-padded to 64 symbols hex string.
	Secret string `json:"secret"`
}

type ResultOfNaclBox struct {
	// Encrypted data encoded in `base64`.
	Encrypted string `json:"encrypted"`
}

type ParamsOfNaclBoxOpen struct {
	// Data that must be decrypted. Encoded with `base64`.
	Encrypted string `json:"encrypted"`
	Nonce     string `json:"nonce"`
	// Sender's public key - unprefixed 0-padded to 64 symbols hex string.
	TheirPublic string `json:"their_public"`
	// Receiver's private key - unprefixed 0-padded to 64 symbols hex string.
	Secret string `json:"secret"`
}

type ResultOfNaclBoxOpen struct {
	// Decrypted data encoded in `base64`.
	Decrypted string `json:"decrypted"`
}

type ParamsOfNaclSecretBox struct {
	// Data that must be encrypted. Encoded with `base64`.
	Decrypted string `json:"decrypted"`
	// Nonce in `hex`.
	Nonce string `json:"nonce"`
	// Secret key - unprefixed 0-padded to 64 symbols hex string.
	Key string `json:"key"`
}

type ParamsOfNaclSecretBoxOpen struct {
	// Data that must be decrypted. Encoded with `base64`.
	Encrypted string `json:"encrypted"`
	// Nonce in `hex`.
	Nonce string `json:"nonce"`
	// Public key - unprefixed 0-padded to 64 symbols hex string.
	Key string `json:"key"`
}

type ParamsOfMnemonicWords struct {
	// Dictionary identifier.
	Dictionary null.Int `json:"dictionary"` // optional
}

type ResultOfMnemonicWords struct {
	// The list of mnemonic words.
	Words string `json:"words"`
}

type ParamsOfMnemonicFromRandom struct {
	// Dictionary identifier.
	Dictionary null.Int `json:"dictionary"` // optional
	// Mnemonic word count.
	WordCount null.Int `json:"word_count"` // optional
}

type ResultOfMnemonicFromRandom struct {
	// String of mnemonic words.
	Phrase string `json:"phrase"`
}

type ParamsOfMnemonicFromEntropy struct {
	// Entropy bytes. Hex encoded.
	Entropy string `json:"entropy"`
	// Dictionary identifier.
	Dictionary null.Int `json:"dictionary"` // optional
	// Mnemonic word count.
	WordCount null.Int `json:"word_count"` // optional
}

type ResultOfMnemonicFromEntropy struct {
	// Phrase.
	Phrase string `json:"phrase"`
}

type ParamsOfMnemonicVerify struct {
	// Phrase.
	Phrase string `json:"phrase"`
	// Dictionary identifier.
	Dictionary null.Int `json:"dictionary"` // optional
	// Word count.
	WordCount null.Int `json:"word_count"` // optional
}

type ResultOfMnemonicVerify struct {
	// Flag indicating the mnemonic is valid or not.
	Valid bool `json:"valid"`
}

type ParamsOfMnemonicDeriveSignKeys struct {
	// Phrase.
	Phrase string `json:"phrase"`
	// Derivation path, for instance "m/44'/396'/0'/0/0".
	Path null.String `json:"path"` // optional
	// Dictionary identifier.
	Dictionary null.Int `json:"dictionary"` // optional
	// Word count.
	WordCount null.Int `json:"word_count"` // optional
}

type ParamsOfHDKeyXPrvFromMnemonic struct {
	// String with seed phrase.
	Phrase string `json:"phrase"`
	// Dictionary identifier.
	Dictionary null.Int `json:"dictionary"` // optional
	// Mnemonic word count.
	WordCount null.Int `json:"word_count"` // optional
}

type ResultOfHDKeyXPrvFromMnemonic struct {
	// Serialized extended master private key.
	XPrv string `json:"xprv"`
}

type ParamsOfHDKeyDeriveFromXPrv struct {
	// Serialized extended private key.
	XPrv string `json:"xprv"`
	// Child index (see BIP-0032).
	ChildIndex int `json:"child_index"`
	// Indicates the derivation of hardened/not-hardened key (see BIP-0032).
	Hardened bool `json:"hardened"`
}

type ResultOfHDKeyDeriveFromXPrv struct {
	// Serialized extended private key.
	XPrv string `json:"xprv"`
}

type ParamsOfHDKeyDeriveFromXPrvPath struct {
	// Serialized extended private key.
	XPrv string `json:"xprv"`
	// Derivation path, for instance "m/44'/396'/0'/0/0".
	Path string `json:"path"`
}

type ResultOfHDKeyDeriveFromXPrvPath struct {
	// Derived serialized extended private key.
	XPrv string `json:"xprv"`
}

type ParamsOfHDKeySecretFromXPrv struct {
	// Serialized extended private key.
	XPrv string `json:"xprv"`
}

type ResultOfHDKeySecretFromXPrv struct {
	// Private key - 64 symbols hex string.
	Secret string `json:"secret"`
}

type ParamsOfHDKeyPublicFromXPrv struct {
	// Serialized extended private key.
	XPrv string `json:"xprv"`
}

type ResultOfHDKeyPublicFromXPrv struct {
	// Public key - 64 symbols hex string.
	Public string `json:"public"`
}

// Integer factorization.
// Performs prime factorization – decomposition of a composite number
// into a product of smaller prime integers (factors).
// See [https://en.wikipedia.org/wiki/Integer_factorization].
func (c *Client) CryptoFactorize(p *ParamsOfFactorize) (*ResultOfFactorize, error) {
	response := new(ResultOfFactorize)
	err := c.dllClient.waitErrorOrResultUnmarshal("crypto.factorize", p, response)

	return response, err
}

// Modular exponentiation.
// Performs modular exponentiation for big integers (`base`^`exponent` mod `modulus`).
// See [https://en.wikipedia.org/wiki/Modular_exponentiation].
func (c *Client) CryptoModularPower(p *ParamsOfModularPower) (*ResultOfModularPower, error) {
	response := new(ResultOfModularPower)
	err := c.dllClient.waitErrorOrResultUnmarshal("crypto.modular_power", p, response)

	return response, err
}

// Calculates CRC16 using TON algorithm.
func (c *Client) CryptoTonCrc16(p *ParamsOfTonCrc16) (*ResultOfTonCrc16, error) {
	response := new(ResultOfTonCrc16)
	err := c.dllClient.waitErrorOrResultUnmarshal("crypto.ton_crc16", p, response)

	return response, err
}

// Generates random byte array of the specified length and returns it in `base64` format.
func (c *Client) CryptoGenerateRandomBytes(p *ParamsOfGenerateRandomBytes) (*ResultOfGenerateRandomBytes, error) {
	response := new(ResultOfGenerateRandomBytes)
	err := c.dllClient.waitErrorOrResultUnmarshal("crypto.generate_random_bytes", p, response)

	return response, err
}

// Converts public key to ton safe_format.
func (c *Client) CryptoConvertPublicKeyToTonSafeFormat(p *ParamsOfConvertPublicKeyToTonSafeFormat) (*ResultOfConvertPublicKeyToTonSafeFormat, error) {
	response := new(ResultOfConvertPublicKeyToTonSafeFormat)
	err := c.dllClient.waitErrorOrResultUnmarshal("crypto.convert_public_key_to_ton_safe_format", p, response)

	return response, err
}

// Generates random ed25519 key pair.
func (c *Client) CryptoGenerateRandomSignKeys() (*KeyPair, error) {
	response := new(KeyPair)
	err := c.dllClient.waitErrorOrResultUnmarshal("crypto.generate_random_sign_keys", nil, response)

	return response, err
}

// Signs a data using the provided keys.
func (c *Client) CryptoSign(p *ParamsOfSign) (*ResultOfSign, error) {
	response := new(ResultOfSign)
	err := c.dllClient.waitErrorOrResultUnmarshal("crypto.sign", p, response)

	return response, err
}

// Verifies signed data using the provided public key.
// Raises error if verification is failed.
func (c *Client) CryptoVerifySignature(p *ParamsOfVerifySignature) (*ResultOfVerifySignature, error) {
	response := new(ResultOfVerifySignature)
	err := c.dllClient.waitErrorOrResultUnmarshal("crypto.verify_signature", p, response)

	return response, err
}

// Calculates SHA256 hash of the specified data.
func (c *Client) CryptoSha256(p *ParamsOfHash) (*ResultOfHash, error) {
	response := new(ResultOfHash)
	err := c.dllClient.waitErrorOrResultUnmarshal("crypto.sha256", p, response)

	return response, err
}

// Calculates SHA512 hash of the specified data.
func (c *Client) CryptoSha512(p *ParamsOfHash) (*ResultOfHash, error) {
	response := new(ResultOfHash)
	err := c.dllClient.waitErrorOrResultUnmarshal("crypto.sha512", p, response)

	return response, err
}

// Perform `scrypt` encryption.
// Derives key from `password` and `key` using `scrypt` algorithm.
// See [https://en.wikipedia.org/wiki/Scrypt].
//
// # Arguments
// - `log_n` - The log2 of the Scrypt parameter `N`
// - `r` - The Scrypt parameter `r`
// - `p` - The Scrypt parameter `p`
// # Conditions
// - `log_n` must be less than `64`
// - `r` must be greater than `0` and less than or equal to `4294967295`
// - `p` must be greater than `0` and less than `4294967295`
// # Recommended values sufficient for most use-cases
// - `log_n = 15` (`n = 32768`)
// - `r = 8`
// - `p = 1`.
func (c *Client) CryptoScrypt(p *ParamsOfScrypt) (*ResultOfScrypt, error) {
	response := new(ResultOfScrypt)
	err := c.dllClient.waitErrorOrResultUnmarshal("crypto.scrypt", p, response)

	return response, err
}

// Generates a key pair for signing from the secret key.
func (c *Client) CryptoNaclSignKeypairFromSecretKey(p *ParamsOfNaclSignKeyPairFromSecret) (*KeyPair, error) {
	response := new(KeyPair)
	err := c.dllClient.waitErrorOrResultUnmarshal("crypto.nacl_sign_keypair_from_secret_key", p, response)

	return response, err
}

// Signs data using the signer's secret key.
func (c *Client) CryptoNaclSign(p *ParamsOfNaclSign) (*ResultOfNaclSign, error) {
	response := new(ResultOfNaclSign)
	err := c.dllClient.waitErrorOrResultUnmarshal("crypto.nacl_sign", p, response)

	return response, err
}

func (c *Client) CryptoNaclSignOpen(p *ParamsOfNaclSignOpen) (*ResultOfNaclSignOpen, error) {
	response := new(ResultOfNaclSignOpen)
	err := c.dllClient.waitErrorOrResultUnmarshal("crypto.nacl_sign_open", p, response)

	return response, err
}

func (c *Client) CryptoNaclSignDetached(p *ParamsOfNaclSign) (*ResultOfNaclSignDetached, error) {
	response := new(ResultOfNaclSignDetached)
	err := c.dllClient.waitErrorOrResultUnmarshal("crypto.nacl_sign_detached", p, response)

	return response, err
}

func (c *Client) CryptoNaclBoxKeypair() (*KeyPair, error) {
	response := new(KeyPair)
	err := c.dllClient.waitErrorOrResultUnmarshal("crypto.nacl_box_keypair", nil, response)

	return response, err
}

// Generates key pair from a secret key.
func (c *Client) CryptoNaclBoxKeypairFromSecretKey(p *ParamsOfNaclBoxKeyPairFromSecret) (*KeyPair, error) {
	response := new(KeyPair)
	err := c.dllClient.waitErrorOrResultUnmarshal("crypto.nacl_box_keypair_from_secret_key", p, response)

	return response, err
}

// Public key authenticated encryption
//
// Encrypt and authenticate a message using the senders secret key, the recievers public
// key, and a nonce.
func (c *Client) CryptoNaclBox(p *ParamsOfNaclBox) (*ResultOfNaclBox, error) {
	response := new(ResultOfNaclBox)
	err := c.dllClient.waitErrorOrResultUnmarshal("crypto.nacl_box", p, response)

	return response, err
}

// Decrypt and verify the cipher text using the recievers secret key, the senders public
// key, and the nonce.
func (c *Client) CryptoNaclBoxOpen(p *ParamsOfNaclBoxOpen) (*ResultOfNaclBoxOpen, error) {
	response := new(ResultOfNaclBoxOpen)
	err := c.dllClient.waitErrorOrResultUnmarshal("crypto.nacl_box_open", p, response)

	return response, err
}

// Encrypt and authenticate message using nonce and secret key.
func (c *Client) CryptoNaclSecretBox(p *ParamsOfNaclSecretBox) (*ResultOfNaclBox, error) {
	response := new(ResultOfNaclBox)
	err := c.dllClient.waitErrorOrResultUnmarshal("crypto.nacl_secret_box", p, response)

	return response, err
}

// Decrypts and verifies cipher text using `nonce` and secret `key`.
func (c *Client) CryptoNaclSecretBoxOpen(p *ParamsOfNaclSecretBoxOpen) (*ResultOfNaclBoxOpen, error) {
	response := new(ResultOfNaclBoxOpen)
	err := c.dllClient.waitErrorOrResultUnmarshal("crypto.nacl_secret_box_open", p, response)

	return response, err
}

// Prints the list of words from the specified dictionary.
func (c *Client) CryptoMnemonicWords(p *ParamsOfMnemonicWords) (*ResultOfMnemonicWords, error) {
	response := new(ResultOfMnemonicWords)
	err := c.dllClient.waitErrorOrResultUnmarshal("crypto.mnemonic_words", p, response)

	return response, err
}

// Generates a random mnemonic from the specified dictionary and word count.
func (c *Client) CryptoMnemonicFromRandom(p *ParamsOfMnemonicFromRandom) (*ResultOfMnemonicFromRandom, error) {
	response := new(ResultOfMnemonicFromRandom)
	err := c.dllClient.waitErrorOrResultUnmarshal("crypto.mnemonic_from_random", p, response)

	return response, err
}

// Generates mnemonic from the specified entropy.
// Generates mnemonic from pre-generated entropy.
func (c *Client) CryptoMnemonicFromEntropy(p *ParamsOfMnemonicFromEntropy) (*ResultOfMnemonicFromEntropy, error) {
	response := new(ResultOfMnemonicFromEntropy)
	err := c.dllClient.waitErrorOrResultUnmarshal("crypto.mnemonic_from_entropy", p, response)

	return response, err
}

// Validates a mnemonic phrase.
// The phrase supplied will be checked for word length and validated according to the checksum
// specified in BIP0039.
func (c *Client) CryptoMnemonicVerify(p *ParamsOfMnemonicVerify) (*ResultOfMnemonicVerify, error) {
	response := new(ResultOfMnemonicVerify)
	err := c.dllClient.waitErrorOrResultUnmarshal("crypto.mnemonic_verify", p, response)

	return response, err
}

// Derives a key pair for signing from the seed phrase.
// Validates the seed phrase, generates master key and then derives
// the key pair from the master key and the specified path.
func (c *Client) CryptoMnemonicDeriveSignKeys(p *ParamsOfMnemonicDeriveSignKeys) (*KeyPair, error) {
	response := new(KeyPair)
	err := c.dllClient.waitErrorOrResultUnmarshal("crypto.mnemonic_derive_sign_keys", p, response)

	return response, err
}

// Generates an extended master private key that will be the root for all the derived keys.
func (c *Client) CryptoHdkeyXprvFromMnemonic(p *ParamsOfHDKeyXPrvFromMnemonic) (*ResultOfHDKeyXPrvFromMnemonic, error) {
	response := new(ResultOfHDKeyXPrvFromMnemonic)
	err := c.dllClient.waitErrorOrResultUnmarshal("crypto.hdkey_xprv_from_mnemonic", p, response)

	return response, err
}

// Returns extended private key derived from the specified extended private key and child index.
func (c *Client) CryptoHdkeyDeriveFromXprv(p *ParamsOfHDKeyDeriveFromXPrv) (*ResultOfHDKeyDeriveFromXPrv, error) {
	response := new(ResultOfHDKeyDeriveFromXPrv)
	err := c.dllClient.waitErrorOrResultUnmarshal("crypto.hdkey_derive_from_xprv", p, response)

	return response, err
}

// Derives the exented private key from the specified key and path.
func (c *Client) CryptoHdkeyDeriveFromXprvPath(p *ParamsOfHDKeyDeriveFromXPrvPath) (*ResultOfHDKeyDeriveFromXPrvPath, error) {
	response := new(ResultOfHDKeyDeriveFromXPrvPath)
	err := c.dllClient.waitErrorOrResultUnmarshal("crypto.hdkey_derive_from_xprv_path", p, response)

	return response, err
}

// Extracts the private key from the serialized extended private key.
func (c *Client) CryptoHdkeySecretFromXprv(p *ParamsOfHDKeySecretFromXPrv) (*ResultOfHDKeySecretFromXPrv, error) {
	response := new(ResultOfHDKeySecretFromXPrv)
	err := c.dllClient.waitErrorOrResultUnmarshal("crypto.hdkey_secret_from_xprv", p, response)

	return response, err
}

// Extracts the public key from the serialized extended private key.
func (c *Client) CryptoHdkeyPublicFromXprv(p *ParamsOfHDKeyPublicFromXPrv) (*ResultOfHDKeyPublicFromXPrv, error) {
	response := new(ResultOfHDKeyPublicFromXPrv)
	err := c.dllClient.waitErrorOrResultUnmarshal("crypto.hdkey_public_from_xprv", p, response)

	return response, err
}
