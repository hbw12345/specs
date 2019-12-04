package message

import actor "github.com/filecoin-project/specs/systems/filecoin_vm/actor"
import addr "github.com/filecoin-project/specs/systems/filecoin_vm/actor/address"
import filcrypto "github.com/filecoin-project/specs/algorithms/crypto"
import util "github.com/filecoin-project/specs/util"

var IMPL_FINISH = util.IMPL_FINISH

type Serialization = util.Serialization

func UnsignedMessage_Make(
	from addr.Address,
	to addr.Address,
	method actor.MethodNum,
	params actor.MethodParams,
	callSeqNum actor.CallSeqNum,
	value actor.TokenAmount,
	gasPrice GasPrice,
	gasLimit GasAmount,
) UnsignedMessage {
	return &UnsignedMessage_I{
		From_:       from,
		To_:         to,
		Method_:     method,
		Params_:     params,
		CallSeqNum_: callSeqNum,
		Value_:      value,
		GasPrice_:   gasPrice,
		GasLimit_:   gasLimit,
	}
}

func SignedMessage_Make(message UnsignedMessage, signature filcrypto.Signature) SignedMessage {
	return &SignedMessage_I{
		Message_:   message,
		Signature_: signature,
	}
}

func Sign(message UnsignedMessage, keyPair filcrypto.SigKeyPair) (SignedMessage, error) {
	sig, err := filcrypto.Sign(keyPair, util.Bytes(Serialize_UnsignedMessage(message)))
	if err != nil {
		return nil, err
	}
	return SignedMessage_Make(message, sig), nil
}

func SignatureVerificationError() error {
	IMPL_FINISH()
	panic("")
}

func Verify(message SignedMessage, publicKey filcrypto.PublicKey) (UnsignedMessage, error) {
	m := util.Bytes(Serialize_UnsignedMessage(message.Message()))
	sigValid, err := filcrypto.Verify(publicKey, message.Signature(), m)
	if err != nil {
		return nil, err
	}
	if !sigValid {
		return nil, SignatureVerificationError()
	}
	return message.Message(), nil
}

func (x GasAmount) Add(y GasAmount) GasAmount {
	panic("TODO")
}

func (x GasAmount) Subtract(y GasAmount) GasAmount {
	panic("TODO")
}

func (x GasAmount) LessThan(y GasAmount) bool {
	panic("TODO")
}

func GasAmount_Zero() GasAmount {
	panic("TODO")
}