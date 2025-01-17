package iplddecoders

import (
	"fmt"

	"github.com/ipld/go-ipld-prime"
	"github.com/ipld/go-ipld-prime/codec/dagcbor"
	"go.firedancer.io/radiance/cmd/radiance/car/createcar/ipld/ipldbindcode"
)

type Kind int

const (
	KindTransaction Kind = iota
	KindEntry
	KindBlock
	KindSubset
	KindEpoch
)

// String returns the string representation of the Kind.
func (k Kind) String() string {
	switch k {
	case KindTransaction:
		return "Transaction"
	case KindEntry:
		return "Entry"
	case KindBlock:
		return "Block"
	case KindSubset:
		return "Subset"
	case KindEpoch:
		return "Epoch"
	default:
		return "Unknown"
	}
}

func DecodeEpoch(epochRaw []byte) (*ipldbindcode.Epoch, error) {
	var epoch ipldbindcode.Epoch
	_, err := ipld.Unmarshal(epochRaw, dagcbor.Decode, &epoch, ipldbindcode.Prototypes.Epoch.Type())
	if err != nil {
		return nil, fmt.Errorf("failed to decode Epoch node: %w", err)
	}
	if epoch.Kind != int(KindEpoch) {
		return nil, fmt.Errorf("expected Epoch node, got %s", Kind(epoch.Kind))
	}
	return &epoch, nil
}

func DecodeSubset(subsetRaw []byte) (*ipldbindcode.Subset, error) {
	var subset ipldbindcode.Subset
	_, err := ipld.Unmarshal(subsetRaw, dagcbor.Decode, &subset, ipldbindcode.Prototypes.Subset.Type())
	if err != nil {
		return nil, fmt.Errorf("failed to decode Subset node: %w", err)
	}
	if subset.Kind != int(KindSubset) {
		return nil, fmt.Errorf("expected Subset node, got %s", Kind(subset.Kind))
	}
	return &subset, nil
}

func DecodeBlock(blockRaw []byte) (*ipldbindcode.Block, error) {
	var block ipldbindcode.Block
	_, err := ipld.Unmarshal(blockRaw, dagcbor.Decode, &block, ipldbindcode.Prototypes.Block.Type())
	if err != nil {
		return nil, fmt.Errorf("failed to decode Block node: %w", err)
	}
	if block.Kind != int(KindBlock) {
		return nil, fmt.Errorf("expected Block node, got %s", Kind(block.Kind))
	}
	return &block, nil
}

func DecodeEntry(entryRaw []byte) (*ipldbindcode.Entry, error) {
	var entry ipldbindcode.Entry
	_, err := ipld.Unmarshal(entryRaw, dagcbor.Decode, &entry, ipldbindcode.Prototypes.Entry.Type())
	if err != nil {
		return nil, fmt.Errorf("failed to decode Entry node: %w", err)
	}
	if entry.Kind != int(KindEntry) {
		return nil, fmt.Errorf("expected Entry node, got %s", Kind(entry.Kind))
	}
	return &entry, nil
}

func DecodeTransaction(transactionRaw []byte) (*ipldbindcode.Transaction, error) {
	var transaction ipldbindcode.Transaction
	_, err := ipld.Unmarshal(transactionRaw, dagcbor.Decode, &transaction, ipldbindcode.Prototypes.Transaction.Type())
	if err != nil {
		return nil, fmt.Errorf("failed to decode Transaction node: %w", err)
	}
	if transaction.Kind != int(KindTransaction) {
		return nil, fmt.Errorf("expected Transaction node, got %s", Kind(transaction.Kind))
	}
	return &transaction, nil
}
