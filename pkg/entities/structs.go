package entities

import "github.com/iotaledger/wasp/packages/vm/wasmlib/go/wasmlib"

type Crop struct {
	Country string // the country name
	Idx     int32  // crop index
	Name    string // name of the crop
}

func NewCropFromBytes(bytes []byte) *Crop {
	decode := wasmlib.NewBytesDecoder(bytes)
	data := &Crop{}
	data.Country = decode.String()
	data.Idx = decode.Int32()
	data.Name = decode.String()
	decode.Close()
	return data
}

func (o *Crop) Bytes() []byte {
	return wasmlib.NewBytesEncoder().
		String(o.Country).
		Int32(o.Idx).
		String(o.Name).
		Data()
}

type ImmutableCrop struct {
	objID int32
	keyID wasmlib.Key32
}

func (o ImmutableCrop) Exists() bool {
	return wasmlib.Exists(o.objID, o.keyID, wasmlib.TYPE_BYTES)
}

func (o ImmutableCrop) Value() *Crop {
	return NewCropFromBytes(wasmlib.GetBytes(o.objID, o.keyID, wasmlib.TYPE_BYTES))
}

type MutableCrop struct {
	objID int32
	keyID wasmlib.Key32
}

func (o MutableCrop) Exists() bool {
	return wasmlib.Exists(o.objID, o.keyID, wasmlib.TYPE_BYTES)
}

func (o MutableCrop) SetValue(value *Crop) {
	wasmlib.SetBytes(o.objID, o.keyID, wasmlib.TYPE_BYTES, value.Bytes())
}

func (o MutableCrop) Value() *Crop {
	return NewCropFromBytes(wasmlib.GetBytes(o.objID, o.keyID, wasmlib.TYPE_BYTES))
}

type Order struct {
	Amount    int64             // amount of crop
	CropIdx   int32             // crop index
	FullPrice int64             // amount * price
	Idx       int32             // crop index
	Owner     wasmlib.ScAgentID // agent id of the owner
	Price     int64             // price of the crop
	Status    string            // order status (opening, matched, canceled)
	Type      string            // sell, buy
}

func NewOrderFromBytes(bytes []byte) *Order {
	decode := wasmlib.NewBytesDecoder(bytes)
	data := &Order{}
	data.Amount = decode.Int64()
	data.CropIdx = decode.Int32()
	data.FullPrice = decode.Int64()
	data.Idx = decode.Int32()
	data.Owner = decode.AgentID()
	data.Price = decode.Int64()
	data.Status = decode.String()
	data.Type = decode.String()
	decode.Close()
	return data
}

func (o *Order) Bytes() []byte {
	return wasmlib.NewBytesEncoder().
		Int64(o.Amount).
		Int32(o.CropIdx).
		Int64(o.FullPrice).
		Int32(o.Idx).
		AgentID(o.Owner).
		Int64(o.Price).
		String(o.Status).
		String(o.Type).
		Data()
}

type ImmutableOrder struct {
	objID int32
	keyID wasmlib.Key32
}

func (o ImmutableOrder) Exists() bool {
	return wasmlib.Exists(o.objID, o.keyID, wasmlib.TYPE_BYTES)
}

func (o ImmutableOrder) Value() *Order {
	return NewOrderFromBytes(wasmlib.GetBytes(o.objID, o.keyID, wasmlib.TYPE_BYTES))
}

type MutableOrder struct {
	objID int32
	keyID wasmlib.Key32
}

func (o MutableOrder) Exists() bool {
	return wasmlib.Exists(o.objID, o.keyID, wasmlib.TYPE_BYTES)
}

func (o MutableOrder) SetValue(value *Order) {
	wasmlib.SetBytes(o.objID, o.keyID, wasmlib.TYPE_BYTES, value.Bytes())
}

func (o MutableOrder) Value() *Order {
	return NewOrderFromBytes(wasmlib.GetBytes(o.objID, o.keyID, wasmlib.TYPE_BYTES))
}

type Wallet struct {
	Amount  int64             // amount of crop
	CropIdx int32             // crop index
	Owner   wasmlib.ScAgentID // agent id of the owner
}

func NewWalletFromBytes(bytes []byte) *Wallet {
	decode := wasmlib.NewBytesDecoder(bytes)
	data := &Wallet{}
	data.Amount = decode.Int64()
	data.CropIdx = decode.Int32()
	data.Owner = decode.AgentID()
	decode.Close()
	return data
}

func (o *Wallet) Bytes() []byte {
	return wasmlib.NewBytesEncoder().
		Int64(o.Amount).
		Int32(o.CropIdx).
		AgentID(o.Owner).
		Data()
}

type ImmutableWallet struct {
	objID int32
	keyID wasmlib.Key32
}

func (o ImmutableWallet) Exists() bool {
	return wasmlib.Exists(o.objID, o.keyID, wasmlib.TYPE_BYTES)
}

func (o ImmutableWallet) Value() *Wallet {
	return NewWalletFromBytes(wasmlib.GetBytes(o.objID, o.keyID, wasmlib.TYPE_BYTES))
}

type MutableWallet struct {
	objID int32
	keyID wasmlib.Key32
}

func (o MutableWallet) Exists() bool {
	return wasmlib.Exists(o.objID, o.keyID, wasmlib.TYPE_BYTES)
}

func (o MutableWallet) SetValue(value *Wallet) {
	wasmlib.SetBytes(o.objID, o.keyID, wasmlib.TYPE_BYTES, value.Bytes())
}

func (o MutableWallet) Value() *Wallet {
	return NewWalletFromBytes(wasmlib.GetBytes(o.objID, o.keyID, wasmlib.TYPE_BYTES))
}