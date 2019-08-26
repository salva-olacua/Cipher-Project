package xor

func XorToAllElements(input, key string)(output string){
    for i:=0; i<len(input); i++{
    	output+=string(input[i]^key[i])
    }
    return output
}