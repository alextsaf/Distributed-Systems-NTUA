package main

type UTXOStack []TransactionOutput

func (s *UTXOStack) IsEmpty() bool {
	return len(*s) == 0;
}

func (s *UTXOStack) Push(UTXO TransactionOutput) {
	*s = append(*s, UTXO);
}

func (s *UTXOStack) Pop() (element TransactionOutput, isEmpty bool) {
	if s.IsEmpty() {
		return TransactionOutput{}, true;
	} else {
		index := len(*s) - 1;
		element = (*s)[index]; 
		*s = (*s)[:index]; 
		return element, false;
	}
}

func (s *UTXOStack) Contains(UTXOtoFind TransactionOutput) (uint,bool) {
	for _, utxo := range *s {
		if utxo == UTXOtoFind {
			return utxo.amount, true;
		}
	}
	return 0, false;
}

func (s *UTXOStack) SumAmounts() uint {
	var sum uint;
	for _, utxo := range *s {
		sum += utxo.amount;
	}
	return sum;
}
