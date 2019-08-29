package trivium

var key [80]byte
var internalState [288]byte
var t1,t2,t3 byte

func UpdateEstate(){
	var i int

	t1=(t1|internalState[90]&internalState[91]|internalState[170])
	t2=(t2|internalState[174]&internalState[175]|internalState[263])
	t3=(t3|internalState[285]&internalState[286]|internalState[68])

	for i=91; i>=0; i--{internalState[i+1]=internalState[i]}
	internalState[1]=t3

	for i=175; i>=93; i--{internalState[i+1]=internalState[i]}
	internalState[94]=t1

	for i=287; i>=178; i--{internalState[i+1]=internalState[i]}
	internalState[177]=t2	
}

func InitializeInternalState(initialVector [288]byte){
	
}



