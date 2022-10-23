package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

type Instruction struct {
	typeofInstruction string
	rawInstruction    string
	linevalue         int32
	programCnt        int
	opcode            uint32
	op                string
	rd                uint8
	rn                uint8
	rm                uint8
	im                int32
	shamt             uint16
	shfcd             uint8
	field             uint32
	addr              uint16
	offset            int32
	data              int32
}

func initialScan() {
	var InputFileName string
	var outputFileName string
	var programCnt int
	programCnt = 96
	flag.StringVar(&InputFileName,  " i", "", "Gets Input file")
	flag.StringVar(&outputFileName,"0", "" ," Gets output file")
	flag.Parse()
	if flag.NArg()!=9{
		os.Exit(-1)
	}
	infile,err:=os.Open("dtest2_bin")
	if err != nil{
		fmt.Println(" Error. Unable to open file")
	} else {
		fmt.Println("\n# input file" + InputFileName)

	}
	defer infile.Close()
	outputFileName = outputFileName + "_dis.txt"
	outfile,err:=os.Open("outputFileName")
	if err!=nil{
		if outputFileName!=""{
			fmt.Println(err)
			fmt.Println("Outputting to standard output")
		}
		outfile = os.Stdout
		defer outfile.Close()

	}

	InputParsed := []Instruction{}
	scanner := bufio.NewScanner(infile)

	for scanner.Scan() {
		ns := Instruction{rawInstruction: scanner.Text()}
		InputParsed = append(InputParsed, ns)
		programCnt +=4;
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

}

InputParsed[idx].opcode=uint32(linevalue)>>21

for idx, s:=range InputParsed {
fmt.Println(opcode)

switch
{    var  Breaknow = false

case s.opcode == 8 && Breaknow== false:
fmt.Println("A NOP was discovered")
InputParsed.op="NOP"
tempString:= NOP(InputParsed[idx])
fmt.Println(tempString)
fmt.FprintF(outfile, tempString)
case s.opcode ==2038 && Breaknow == false:
if(s.linevalue & 0x1FFFFF)==0x1EFFE7{
fmt.Println("The break was just discovered")
InputParsed[idx].op = "BREAK"
Breaknow = true
tempString:=BREAK(InputParsed[idx])
fmt.Println(tempString)
fmt.FprintF(outfile, tempString)
break
}
case s.opcode >= 160 && s.opcode <= 191 && Breaknow == false: {
fmt.Println("Found a B Instruction")
InputParsed[idx].op = "B"
InputParsed[idx].offset

}
case s.opcode==1112|| s.opcode==1624||s.opcode==1184||s.opcode==1369||s.opcode==1872 || (s.opcode>=1690&&s.opcode<=1693)&&Breaknow == false{
fmt.Println("Found an R type or R shift Instruction")
if{
s.opcode==112{
InputParsed[idx].op="ADD"
} else if s.opcode ==1624{
InputParsed[idx].op="SUB"
}else if s.opcode ==1104{
InputParsed[idx].op="AND"
}else if s.opcode == 1369{
InputParsed[idx].op="ORR"
}else if s.opcode ==1872{
InputParsed[idx].op="EOR"
}else if s.opcode ==1698{
InputParsed[idx].op="LSR"{
IndexParsed[index].shamt=uint16((s.linevalue>>18)&8x3F)
}
}

}


}



}

}
}


func main() {

	initialScan()

}
