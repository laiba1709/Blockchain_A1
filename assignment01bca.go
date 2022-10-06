package assignment01bca

import (
	"crypto/sha256"
	"fmt"
	"strconv"
)

var globalcounter int=-1;

type Block struct {
	nonce  int
	transaction  string
	prevHash  string
	currHash string
}

func NewBlock(n int, t string, p string) *Block {
	s := new(Block)
	s.nonce = n
	s.transaction = t
	s.prevHash = p
	
	return s
}

type BlockList struct {
	list []*Block
}
var TransactionStringList []string;
var Hashlist []string;

func (ls *BlockList) CreateBlock(n int, t string, prevHash string) *Block {
	globalcounter++;
	var concatenation string = "";
	
	st := NewBlock(n, t, prevHash)
	ls.list = append(ls.list, st)
	

	if(globalcounter==0) {
		prevHash=""
	}
	if (globalcounter==1 || globalcounter>1){
		ls.list[globalcounter].prevHash= Hashlist[globalcounter-1]
		prevHash = Hashlist[globalcounter-1]
	}
	
		
	strnonce := strconv.Itoa(ls.list[globalcounter].nonce)
	concatenation += strnonce
	concatenation += t
	concatenation += prevHash
	output:= CalculateHash(concatenation)
	st.currHash = output
	Hashlist = append(Hashlist, output)

	return st
}


func (li *BlockList) ChangeBlock(nonce int, t string) {	
	l := len(li.list)
	for i:=0;i<l;i++{
		if li.list[i].nonce == nonce {
			li.list[i].transaction = t;
			strnonce := strconv.Itoa(li.list[globalcounter].nonce)
			var concatenation string="";
			concatenation += strnonce
			concatenation += t
			concatenation += li.list[i-1].currHash
			output:= CalculateHash(concatenation)
			Hashlist[i] = output
			li.list[i].currHash = output;
		}
	}
}

func (li *BlockList) VerifyChain() {	
	l := len(li.list)
	for i:=0;i<l;i++{
	
		if i==0{
			li.list[i].prevHash = " "
				fmt.Println("\nBlock ",i," is valid!")
		} else {
			if li.list[i-1].currHash == Hashlist[i-1]{
				fmt.Println("Block ",i," is valid!")
			}else{
				fmt.Println("Block ",i," is invalid!\n")
				fmt.Println("\nBlock li.list[i-1].currHash: ",li.list[i-1].currHash);
				fmt.Println("\nBlock li.list[i-1].prevHash: ",Hashlist[i-1]);
			}
			
			}
		}
	}

func (li *BlockList) Print() {

	l:=len(li.list)
	for i:=0; i<l;i++{
		fmt.Println("\n__________________________")
	fmt.Println("Block: ",i)
	fmt.Println("Block Nonce: ",li.list[i].nonce)
	fmt.Println("Transaction String: ",li.list[i].transaction)
	fmt.Println("Current Hash: ",li.list[i].currHash)

	if(i>0){
		fmt.Println("Previous Hash: ",li.list[i-1].currHash)	
		}
	}

	}


func CalculateHash(stringToHash string) string {
	//fmt.Println("String Recieved: ", stringToHash)
	return fmt.Sprintf("%x", sha256.Sum256([]byte(stringToHash)))
}
