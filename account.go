package mt

import (
	"crypto/ecdsa"
	"crypto/sha256"
	"encoding/hex"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/google/uuid"
	"math/big"
	"sync"
)

type Account struct {
	PublicKey  common.Address
	PrivateKey *ecdsa.PrivateKey
	StoredTxs  map[uint64]*types.Transaction
	mux        sync.Mutex
}

func GenKeyStore(key *ecdsa.PrivateKey, randomPwdFirst, randomPwdSecond string) ([]byte, error) {
	ks := &keystore.Key{
		Id:         uuid.New(),
		Address:    crypto.PubkeyToAddress(key.PublicKey),
		PrivateKey: key,
	}
	auth := accountAuth(randomPwdFirst, randomPwdSecond)
	return keystore.EncryptKey(ks, auth, keystore.StandardScryptN, keystore.StandardScryptP)
}

func GetAccountFromKeyStore(keyJson []byte, randomPwdFirst, randomPwdSecond string) (*Account, error) {
	auth := accountAuth(randomPwdFirst, randomPwdSecond)
	key, err := keystore.DecryptKey(keyJson, auth)
	if err != nil {
		return nil, err
	}
	return &Account{
		PrivateKey: key.PrivateKey,
		PublicKey:  key.Address,
	}, nil
}

func accountAuth(randomPwdFirst, randomPwdSecond string) string {
	h := sha256.New()
	h.Write([]byte(randomPwdFirst))
	h.Write([]byte(randomPwdSecond))
	auth := hex.EncodeToString(h.Sum(nil))
	return auth
}

func (acc *Account) Balance() (*big.Int, error) {
	return C.BalanceAt(Ctx, acc.PublicKey, nil)
}

func (acc *Account) StoreTx(tx *types.Transaction) {
	acc.mux.Lock()
	defer acc.mux.Unlock()
	nonce := tx.Nonce()
	if acc.StoredTxs == nil {
		acc.StoredTxs = make(map[uint64]*types.Transaction)
	}
	acc.StoredTxs[nonce] = tx
}

func (acc *Account) GetTxAtNonce(nonce uint64) *types.Transaction {
	acc.mux.Lock()
	defer acc.mux.Unlock()
	if tx, ok := acc.StoredTxs[nonce]; ok {
		return tx
	}
	return nil
}
