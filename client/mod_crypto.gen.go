package client

// DON'T EDIT THIS FILE! It is generated via 'task generate' at 13 Feb 21 14:36 UTC
//
// Mod crypto
//
// Crypto functions.

import (
	"encoding/json"

	"github.com/volatiletech/null"
)

type CryptoErrorCode uint32

const (
	InvalidPublicKeyCryptoErrorCode          CryptoErrorCode = 100
	InvalidSecretKeyCryptoErrorCode          CryptoErrorCode = 101
	InvalidKeyCryptoErrorCode                CryptoErrorCode = 102
	InvalidFactorizeChallengeCryptoErrorCode CryptoErrorCode = 106
	InvalidBigIntCryptoErrorCode             CryptoErrorCode = 107
	ScryptFailedCryptoErrorCode              CryptoErrorCode = 108
	InvalidKeySizeCryptoErrorCode            CryptoErrorCode = 109
	NaclSecretBoxFailedCryptoErrorCode       CryptoErrorCode = 110
	NaclBoxFailedCryptoErrorCode             CryptoErrorCode = 111
	NaclSignFailedCryptoErrorCode            CryptoErrorCode = 112
	Bip39InvalidEntropyCryptoErrorCode       CryptoErrorCode = 113
	Bip39InvalidPhraseCryptoErrorCode        CryptoErrorCode = 114
	Bip32InvalidKeyCryptoErrorCode           CryptoErrorCode = 115
	Bip32InvalidDerivePathCryptoErrorCode    CryptoErrorCode = 116
	Bip39InvalidDictionaryCryptoErrorCode    CryptoErrorCode = 117
	Bip39InvalidWordCountCryptoErrorCode     CryptoErrorCode = 118
	MnemonicGenerationFailedCryptoErrorCode  CryptoErrorCode = 119
	MnemonicFromEntropyFailedCryptoErrorCode CryptoErrorCode = 120
	SigningBoxNotRegisteredCryptoErrorCode   CryptoErrorCode = 121
	InvalidSignatureCryptoErrorCode          CryptoErrorCode = 122
)

type (
	SigningBoxHandle  uint32
	ParamsOfFactorize struct {
		// Hexadecimal representation of u64 composite number.
		Composite string `json:"composite"`
	}
)

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
	// Input data for CRC calculation.
	// Encoded with `base64`.
	Data string `json:"data"`
}

type ResultOfTonCrc16 struct {
	// Calculated CRC for input data.
	Crc uint16 `json:"crc"`
}

type ParamsOfGenerateRandomBytes struct {
	// Size of random byte array.
	Length uint32 `json:"length"`
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
	// Input data for hash calculation.
	// Encoded with `base64`.
	Data string `json:"data"`
}

type ResultOfHash struct {
	// Hash of input `data`.
	// Encoded with 'hex'.
	Hash string `json:"hash"`
}

type ParamsOfScrypt struct {
	// The password bytes to be hashed. Must be encoded with `base64`.
	Password string `json:"password"`
	// Salt bytes that modify the hash to protect against Rainbow table attacks. Must be encoded with `base64`.
	Salt string `json:"salt"`
	// CPU/memory cost parameter.
	LogN uint8 `json:"log_n"`
	// The block size parameter, which fine-tunes sequential memory read size and performance.
	R uint32 `json:"r"`
	// Parallelization parameter.
	P uint32 `json:"p"`
	// Intended output length in octets of the derived key.
	DkLen uint32 `json:"dk_len"`
}

type ResultOfScrypt struct {
	// Derived key.
	// Encoded with `hex`.
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
	// Signed data that must be unsigned.
	// Encoded with `base64`.
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

type ParamsOfNaclSignDetachedVerify struct {
	// Unsigned data that must be verified.
	// Encoded with `base64`.
	Unsigned string `json:"unsigned"`
	// Signature that must be verified.
	// Encoded with `hex`.
	Signature string `json:"signature"`
	// Signer's public key - unprefixed 0-padded to 64 symbols hex string.
	Public string `json:"public"`
}

type ResultOfNaclSignDetachedVerify struct {
	// `true` if verification succeeded or `false` if it failed.
	Succeeded bool `json:"succeeded"`
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
	// Data that must be decrypted.
	// Encoded with `base64`.
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
	// Data that must be encrypted.
	// Encoded with `base64`.
	Decrypted string `json:"decrypted"`
	// Nonce in `hex`.
	Nonce string `json:"nonce"`
	// Secret key - unprefixed 0-padded to 64 symbols hex string.
	Key string `json:"key"`
}

type ParamsOfNaclSecretBoxOpen struct {
	// Data that must be decrypted.
	// Encoded with `base64`.
	Encrypted string `json:"encrypted"`
	// Nonce in `hex`.
	Nonce string `json:"nonce"`
	// Public key - unprefixed 0-padded to 64 symbols hex string.
	Key string `json:"key"`
}

type ParamsOfMnemonicWords struct {
	// Dictionary identifier.
	Dictionary null.Uint8 `json:"dictionary"` // optional
}

type ResultOfMnemonicWords struct {
	// The list of mnemonic words.
	Words string `json:"words"`
}

type ParamsOfMnemonicFromRandom struct {
	// Dictionary identifier.
	Dictionary null.Uint8 `json:"dictionary"` // optional
	// Mnemonic word count.
	WordCount null.Uint8 `json:"word_count"` // optional
}

type ResultOfMnemonicFromRandom struct {
	// String of mnemonic words.
	Phrase string `json:"phrase"`
}

type ParamsOfMnemonicFromEntropy struct {
	// Entropy bytes.
	// Hex encoded.
	Entropy string `json:"entropy"`
	// Dictionary identifier.
	Dictionary null.Uint8 `json:"dictionary"` // optional
	// Mnemonic word count.
	WordCount null.Uint8 `json:"word_count"` // optional
}

type ResultOfMnemonicFromEntropy struct {
	// Phrase.
	Phrase string `json:"phrase"`
}

type ParamsOfMnemonicVerify struct {
	// Phrase.
	Phrase string `json:"phrase"`
	// Dictionary identifier.
	Dictionary null.Uint8 `json:"dictionary"` // optional
	// Word count.
	WordCount null.Uint8 `json:"word_count"` // optional
}

type ResultOfMnemonicVerify struct {
	// Flag indicating if the mnemonic is valid or not.
	Valid bool `json:"valid"`
}

type ParamsOfMnemonicDeriveSignKeys struct {
	// Phrase.
	Phrase string `json:"phrase"`
	// Derivation path, for instance "m/44'/396'/0'/0/0".
	Path null.String `json:"path"` // optional
	// Dictionary identifier.
	Dictionary null.Uint8 `json:"dictionary"` // optional
	// Word count.
	WordCount null.Uint8 `json:"word_count"` // optional
}

type ParamsOfHDKeyXPrvFromMnemonic struct {
	// String with seed phrase.
	Phrase string `json:"phrase"`
	// Dictionary identifier.
	Dictionary null.Uint8 `json:"dictionary"` // optional
	// Mnemonic word count.
	WordCount null.Uint8 `json:"word_count"` // optional
}

type ResultOfHDKeyXPrvFromMnemonic struct {
	// Serialized extended master private key.
	XPrv string `json:"xprv"`
}

type ParamsOfHDKeyDeriveFromXPrv struct {
	// Serialized extended private key.
	XPrv string `json:"xprv"`
	// Child index (see BIP-0032).
	ChildIndex uint32 `json:"child_index"`
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

type ParamsOfChaCha20 struct {
	// Source data to be encrypted or decrypted.
	// Must be encoded with `base64`.
	Data string `json:"data"`
	// 256-bit key.
	// Must be encoded with `hex`.
	Key string `json:"key"`
	// 96-bit nonce.
	// Must be encoded with `hex`.
	Nonce string `json:"nonce"`
}

type ResultOfChaCha20 struct {
	// Encrypted/decrypted data.
	// Encoded with `base64`.
	Data string `json:"data"`
}

type RegisteredSigningBox struct {
	// Handle of the signing box.
	Handle SigningBoxHandle `json:"handle"`
}

type ParamsOfAppSigningBoxType string

const (

	// Get signing box public key.
	GetPublicKeyParamsOfAppSigningBoxType ParamsOfAppSigningBoxType = "GetPublicKey"
	// Sign data.
	SignParamsOfAppSigningBoxType ParamsOfAppSigningBoxType = "Sign"
)

type ParamsOfAppSigningBox struct {
	Type ParamsOfAppSigningBoxType `json:"type"`
	// Data to sign encoded as base64.
	// presented in types:
	// "Sign".
	Unsigned string `json:"unsigned"`
}

type ResultOfAppSigningBoxType string

const (

	// Result of getting public key.
	GetPublicKeyResultOfAppSigningBoxType ResultOfAppSigningBoxType = "GetPublicKey"
	// Result of signing data.
	SignResultOfAppSigningBoxType ResultOfAppSigningBoxType = "Sign"
)

type ResultOfAppSigningBox struct {
	Type ResultOfAppSigningBoxType `json:"type"`
	// Signing box public key.
	// presented in types:
	// "GetPublicKey".
	PublicKey string `json:"public_key"`
	// Data signature encoded as hex.
	// presented in types:
	// "Sign".
	Signature string `json:"signature"`
}

type ResultOfSigningBoxGetPublicKey struct {
	// Public key of signing box.
	// Encoded with hex.
	Pubkey string `json:"pubkey"`
}

type ParamsOfSigningBoxSign struct {
	// Signing Box handle.
	SigningBox SigningBoxHandle `json:"signing_box"`
	// Unsigned user data.
	// Must be encoded with `base64`.
	Unsigned string `json:"unsigned"`
}

type ResultOfSigningBoxSign struct {
	// Data signature.
	// Encoded with `hex`.
	Signature string `json:"signature"`
}

// Performs prime factorization – decomposition of a composite number into a product of smaller prime integers (factors). See [https://en.wikipedia.org/wiki/Integer_factorization].
func (c *Client) CryptoFactorize(p *ParamsOfFactorize) (*ResultOfFactorize, error) {
	result := new(ResultOfFactorize)

	err := c.dllClient.waitErrorOrResultUnmarshal("crypto.factorize", p, result)

	return result, err
}

// Performs modular exponentiation for big integers (`base`^`exponent` mod `modulus`). See [https://en.wikipedia.org/wiki/Modular_exponentiation].
func (c *Client) CryptoModularPower(p *ParamsOfModularPower) (*ResultOfModularPower, error) {
	result := new(ResultOfModularPower)

	err := c.dllClient.waitErrorOrResultUnmarshal("crypto.modular_power", p, result)

	return result, err
}

// Calculates CRC16 using TON algorithm.
func (c *Client) CryptoTonCrc16(p *ParamsOfTonCrc16) (*ResultOfTonCrc16, error) {
	result := new(ResultOfTonCrc16)

	err := c.dllClient.waitErrorOrResultUnmarshal("crypto.ton_crc16", p, result)

	return result, err
}

// Generates random byte array of the specified length and returns it in `base64` format.
func (c *Client) CryptoGenerateRandomBytes(p *ParamsOfGenerateRandomBytes) (*ResultOfGenerateRandomBytes, error) {
	result := new(ResultOfGenerateRandomBytes)

	err := c.dllClient.waitErrorOrResultUnmarshal("crypto.generate_random_bytes", p, result)

	return result, err
}

// Converts public key to ton safe_format.
func (c *Client) CryptoConvertPublicKeyToTonSafeFormat(p *ParamsOfConvertPublicKeyToTonSafeFormat) (*ResultOfConvertPublicKeyToTonSafeFormat, error) {
	result := new(ResultOfConvertPublicKeyToTonSafeFormat)

	err := c.dllClient.waitErrorOrResultUnmarshal("crypto.convert_public_key_to_ton_safe_format", p, result)

	return result, err
}

// Generates random ed25519 key pair.
func (c *Client) CryptoGenerateRandomSignKeys() (*KeyPair, error) {
	result := new(KeyPair)

	err := c.dllClient.waitErrorOrResultUnmarshal("crypto.generate_random_sign_keys", nil, result)

	return result, err
}

// Signs a data using the provided keys.
func (c *Client) CryptoSign(p *ParamsOfSign) (*ResultOfSign, error) {
	result := new(ResultOfSign)

	err := c.dllClient.waitErrorOrResultUnmarshal("crypto.sign", p, result)

	return result, err
}

// Verifies signed data using the provided public key. Raises error if verification is failed.
func (c *Client) CryptoVerifySignature(p *ParamsOfVerifySignature) (*ResultOfVerifySignature, error) {
	result := new(ResultOfVerifySignature)

	err := c.dllClient.waitErrorOrResultUnmarshal("crypto.verify_signature", p, result)

	return result, err
}

// Calculates SHA256 hash of the specified data.
func (c *Client) CryptoSha256(p *ParamsOfHash) (*ResultOfHash, error) {
	result := new(ResultOfHash)

	err := c.dllClient.waitErrorOrResultUnmarshal("crypto.sha256", p, result)

	return result, err
}

// Calculates SHA512 hash of the specified data.
func (c *Client) CryptoSha512(p *ParamsOfHash) (*ResultOfHash, error) {
	result := new(ResultOfHash)

	err := c.dllClient.waitErrorOrResultUnmarshal("crypto.sha512", p, result)

	return result, err
}

// Derives key from `password` and `key` using `scrypt` algorithm. See [https://en.wikipedia.org/wiki/Scrypt].
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
	result := new(ResultOfScrypt)

	err := c.dllClient.waitErrorOrResultUnmarshal("crypto.scrypt", p, result)

	return result, err
}

// Generates a key pair for signing from the secret key.
func (c *Client) CryptoNaclSignKeypairFromSecretKey(p *ParamsOfNaclSignKeyPairFromSecret) (*KeyPair, error) {
	result := new(KeyPair)

	err := c.dllClient.waitErrorOrResultUnmarshal("crypto.nacl_sign_keypair_from_secret_key", p, result)

	return result, err
}

// Signs data using the signer's secret key.
func (c *Client) CryptoNaclSign(p *ParamsOfNaclSign) (*ResultOfNaclSign, error) {
	result := new(ResultOfNaclSign)

	err := c.dllClient.waitErrorOrResultUnmarshal("crypto.nacl_sign", p, result)

	return result, err
}

// Verifies the signature and returns the unsigned message.
// Verifies the signature in `signed` using the signer's public key `public`
// and returns the message `unsigned`.
//
// If the signature fails verification, crypto_sign_open raises an exception.
func (c *Client) CryptoNaclSignOpen(p *ParamsOfNaclSignOpen) (*ResultOfNaclSignOpen, error) {
	result := new(ResultOfNaclSignOpen)

	err := c.dllClient.waitErrorOrResultUnmarshal("crypto.nacl_sign_open", p, result)

	return result, err
}

// Signs the message using the secret key and returns a signature.
// Signs the message `unsigned` using the secret key `secret`
// and returns a signature `signature`.
func (c *Client) CryptoNaclSignDetached(p *ParamsOfNaclSign) (*ResultOfNaclSignDetached, error) {
	result := new(ResultOfNaclSignDetached)

	err := c.dllClient.waitErrorOrResultUnmarshal("crypto.nacl_sign_detached", p, result)

	return result, err
}

// Verifies the signature with public key and `unsigned` data.
func (c *Client) CryptoNaclSignDetachedVerify(p *ParamsOfNaclSignDetachedVerify) (*ResultOfNaclSignDetachedVerify, error) {
	result := new(ResultOfNaclSignDetachedVerify)

	err := c.dllClient.waitErrorOrResultUnmarshal("crypto.nacl_sign_detached_verify", p, result)

	return result, err
}

// Generates a random NaCl key pair.
func (c *Client) CryptoNaclBoxKeypair() (*KeyPair, error) {
	result := new(KeyPair)

	err := c.dllClient.waitErrorOrResultUnmarshal("crypto.nacl_box_keypair", nil, result)

	return result, err
}

// Generates key pair from a secret key.
func (c *Client) CryptoNaclBoxKeypairFromSecretKey(p *ParamsOfNaclBoxKeyPairFromSecret) (*KeyPair, error) {
	result := new(KeyPair)

	err := c.dllClient.waitErrorOrResultUnmarshal("crypto.nacl_box_keypair_from_secret_key", p, result)

	return result, err
}

// Public key authenticated encryption.
// Encrypt and authenticate a message using the senders secret key, the recievers public
// key, and a nonce.
func (c *Client) CryptoNaclBox(p *ParamsOfNaclBox) (*ResultOfNaclBox, error) {
	result := new(ResultOfNaclBox)

	err := c.dllClient.waitErrorOrResultUnmarshal("crypto.nacl_box", p, result)

	return result, err
}

// Decrypt and verify the cipher text using the recievers secret key, the senders public key, and the nonce.
func (c *Client) CryptoNaclBoxOpen(p *ParamsOfNaclBoxOpen) (*ResultOfNaclBoxOpen, error) {
	result := new(ResultOfNaclBoxOpen)

	err := c.dllClient.waitErrorOrResultUnmarshal("crypto.nacl_box_open", p, result)

	return result, err
}

// Encrypt and authenticate message using nonce and secret key.
func (c *Client) CryptoNaclSecretBox(p *ParamsOfNaclSecretBox) (*ResultOfNaclBox, error) {
	result := new(ResultOfNaclBox)

	err := c.dllClient.waitErrorOrResultUnmarshal("crypto.nacl_secret_box", p, result)

	return result, err
}

// Decrypts and verifies cipher text using `nonce` and secret `key`.
func (c *Client) CryptoNaclSecretBoxOpen(p *ParamsOfNaclSecretBoxOpen) (*ResultOfNaclBoxOpen, error) {
	result := new(ResultOfNaclBoxOpen)

	err := c.dllClient.waitErrorOrResultUnmarshal("crypto.nacl_secret_box_open", p, result)

	return result, err
}

// Prints the list of words from the specified dictionary.
func (c *Client) CryptoMnemonicWords(p *ParamsOfMnemonicWords) (*ResultOfMnemonicWords, error) {
	result := new(ResultOfMnemonicWords)

	err := c.dllClient.waitErrorOrResultUnmarshal("crypto.mnemonic_words", p, result)

	return result, err
}

// Generates a random mnemonic from the specified dictionary and word count.
func (c *Client) CryptoMnemonicFromRandom(p *ParamsOfMnemonicFromRandom) (*ResultOfMnemonicFromRandom, error) {
	result := new(ResultOfMnemonicFromRandom)

	err := c.dllClient.waitErrorOrResultUnmarshal("crypto.mnemonic_from_random", p, result)

	return result, err
}

// Generates mnemonic from pre-generated entropy.
func (c *Client) CryptoMnemonicFromEntropy(p *ParamsOfMnemonicFromEntropy) (*ResultOfMnemonicFromEntropy, error) {
	result := new(ResultOfMnemonicFromEntropy)

	err := c.dllClient.waitErrorOrResultUnmarshal("crypto.mnemonic_from_entropy", p, result)

	return result, err
}

// The phrase supplied will be checked for word length and validated according to the checksum specified in BIP0039.
func (c *Client) CryptoMnemonicVerify(p *ParamsOfMnemonicVerify) (*ResultOfMnemonicVerify, error) {
	result := new(ResultOfMnemonicVerify)

	err := c.dllClient.waitErrorOrResultUnmarshal("crypto.mnemonic_verify", p, result)

	return result, err
}

// Validates the seed phrase, generates master key and then derives the key pair from the master key and the specified path.
func (c *Client) CryptoMnemonicDeriveSignKeys(p *ParamsOfMnemonicDeriveSignKeys) (*KeyPair, error) {
	result := new(KeyPair)

	err := c.dllClient.waitErrorOrResultUnmarshal("crypto.mnemonic_derive_sign_keys", p, result)

	return result, err
}

// Generates an extended master private key that will be the root for all the derived keys.
func (c *Client) CryptoHdkeyXprvFromMnemonic(p *ParamsOfHDKeyXPrvFromMnemonic) (*ResultOfHDKeyXPrvFromMnemonic, error) {
	result := new(ResultOfHDKeyXPrvFromMnemonic)

	err := c.dllClient.waitErrorOrResultUnmarshal("crypto.hdkey_xprv_from_mnemonic", p, result)

	return result, err
}

// Returns extended private key derived from the specified extended private key and child index.
func (c *Client) CryptoHdkeyDeriveFromXprv(p *ParamsOfHDKeyDeriveFromXPrv) (*ResultOfHDKeyDeriveFromXPrv, error) {
	result := new(ResultOfHDKeyDeriveFromXPrv)

	err := c.dllClient.waitErrorOrResultUnmarshal("crypto.hdkey_derive_from_xprv", p, result)

	return result, err
}

// Derives the extended private key from the specified key and path.
func (c *Client) CryptoHdkeyDeriveFromXprvPath(p *ParamsOfHDKeyDeriveFromXPrvPath) (*ResultOfHDKeyDeriveFromXPrvPath, error) {
	result := new(ResultOfHDKeyDeriveFromXPrvPath)

	err := c.dllClient.waitErrorOrResultUnmarshal("crypto.hdkey_derive_from_xprv_path", p, result)

	return result, err
}

// Extracts the private key from the serialized extended private key.
func (c *Client) CryptoHdkeySecretFromXprv(p *ParamsOfHDKeySecretFromXPrv) (*ResultOfHDKeySecretFromXPrv, error) {
	result := new(ResultOfHDKeySecretFromXPrv)

	err := c.dllClient.waitErrorOrResultUnmarshal("crypto.hdkey_secret_from_xprv", p, result)

	return result, err
}

// Extracts the public key from the serialized extended private key.
func (c *Client) CryptoHdkeyPublicFromXprv(p *ParamsOfHDKeyPublicFromXPrv) (*ResultOfHDKeyPublicFromXPrv, error) {
	result := new(ResultOfHDKeyPublicFromXPrv)

	err := c.dllClient.waitErrorOrResultUnmarshal("crypto.hdkey_public_from_xprv", p, result)

	return result, err
}

// Performs symmetric `chacha20` encryption.
func (c *Client) CryptoChacha20(p *ParamsOfChaCha20) (*ResultOfChaCha20, error) {
	result := new(ResultOfChaCha20)

	err := c.dllClient.waitErrorOrResultUnmarshal("crypto.chacha20", p, result)

	return result, err
}

// Register an application implemented signing box.
func (c *Client) CryptoRegisterSigningBox(app AppSigningBox) (*RegisteredSigningBox, error) {
	result := new(RegisteredSigningBox)
	responses, err := c.dllClient.resultsChannel("crypto.register_signing_box", nil)
	if err != nil {
		return nil, err
	}

	response := <-responses
	if response.Code == ResponseCodeError {
		return nil, response.Error
	}

	if err := json.Unmarshal(response.Data, result); err != nil {
		return nil, err
	}

	go func() {
		for r := range responses {
			if r.Code == ResponseCodeAppRequest {
				c.dispatchRequestCryptoRegisterSigningBox(r.Data, app)
			}
			if r.Code == ResponseCodeAppNotify {
				c.dispatchNotifyCryptoRegisterSigningBox(r.Data, app)
			}
		}
	}()

	return result, nil
}

func (c *Client) dispatchRequestCryptoRegisterSigningBox(payload []byte, app AppSigningBox) {
	var appRequest ParamsOfAppRequest
	var appParams ParamsOfAppSigningBox
	err := json.Unmarshal(payload, &appRequest)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(appRequest.RequestData, &appParams)
	if err != nil {
		panic(err)
	}
	appResponse, err := app.Request(appParams)
	appRequestResult := AppRequestResult{}
	if err != nil {
		appRequestResult.Type = ErrorAppRequestResultType
		appRequestResult.Text = err.Error()
	} else {
		appRequestResult.Type = OkAppRequestResultType
		appRequestResult.Result, _ = json.Marshal(appResponse)
	}
	err = c.ClientResolveAppRequest(&ParamsOfResolveAppRequest{
		AppRequestID: appRequest.AppRequestID,
		Result:       appRequestResult,
	})
	if err != nil {
		panic(err)
	}
}

func (c *Client) dispatchNotifyCryptoRegisterSigningBox(payload []byte, app AppSigningBox) {
	var appParams ParamsOfAppSigningBox
	err := json.Unmarshal(payload, &appParams)
	if err != nil {
		panic(err)
	}
	app.Notify(appParams)
}

// Creates a default signing box implementation.
func (c *Client) CryptoGetSigningBox(p *KeyPair) (*RegisteredSigningBox, error) {
	result := new(RegisteredSigningBox)

	err := c.dllClient.waitErrorOrResultUnmarshal("crypto.get_signing_box", p, result)

	return result, err
}

// Returns public key of signing key pair.
func (c *Client) CryptoSigningBoxGetPublicKey(p *RegisteredSigningBox) (*ResultOfSigningBoxGetPublicKey, error) {
	result := new(ResultOfSigningBoxGetPublicKey)

	err := c.dllClient.waitErrorOrResultUnmarshal("crypto.signing_box_get_public_key", p, result)

	return result, err
}

// Returns signed user data.
func (c *Client) CryptoSigningBoxSign(p *ParamsOfSigningBoxSign) (*ResultOfSigningBoxSign, error) {
	result := new(ResultOfSigningBoxSign)

	err := c.dllClient.waitErrorOrResultUnmarshal("crypto.signing_box_sign", p, result)

	return result, err
}

// Removes signing box from SDK.
func (c *Client) CryptoRemoveSigningBox(p *RegisteredSigningBox) error {
	_, err := c.dllClient.waitErrorOrResult("crypto.remove_signing_box", p)

	return err
}
