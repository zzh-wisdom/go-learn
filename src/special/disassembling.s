"".swap STEXT nosplit size=39 args=0x20 locals=0x0
	0x0000 00000 (disassembling.go:3)	TEXT	"".swap(SB), NOSPLIT|ABIInternal, $0-32
	0x0000 00000 (disassembling.go:3)	FUNCDATA	$0, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x0000 00000 (disassembling.go:3)	FUNCDATA	$1, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x0000 00000 (disassembling.go:3)	FUNCDATA	$2, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x0000 00000 (disassembling.go:3)	PCDATA	$0, $0
	0x0000 00000 (disassembling.go:3)	PCDATA	$1, $0
	0x0000 00000 (disassembling.go:3)	MOVQ	$0, "".x+24(SP)
	0x0009 00009 (disassembling.go:3)	MOVQ	$0, "".y+32(SP)
	0x0012 00018 (disassembling.go:4)	MOVQ	"".b+16(SP), AX
	0x0017 00023 (disassembling.go:4)	MOVQ	AX, "".x+24(SP)
	0x001c 00028 (disassembling.go:5)	MOVQ	"".a+8(SP), AX
	0x0021 00033 (disassembling.go:5)	MOVQ	AX, "".y+32(SP)
	0x0026 00038 (disassembling.go:6)	RET
	0x0000 48 c7 44 24 18 00 00 00 00 48 c7 44 24 20 00 00  H.D$.....H.D$ ..
	0x0010 00 00 48 8b 44 24 10 48 89 44 24 18 48 8b 44 24  ..H.D$.H.D$.H.D$
	0x0020 08 48 89 44 24 20 c3                             .H.D$ .
"".main STEXT size=68 args=0x0 locals=0x28
	0x0000 00000 (disassembling.go:9)	TEXT	"".main(SB), ABIInternal, $40-0
	0x0000 00000 (disassembling.go:9)	MOVQ	(TLS), CX
	0x0009 00009 (disassembling.go:9)	CMPQ	SP, 16(CX)
	0x000d 00013 (disassembling.go:9)	JLS	61
	0x000f 00015 (disassembling.go:9)	SUBQ	$40, SP
	0x0013 00019 (disassembling.go:9)	MOVQ	BP, 32(SP)
	0x0018 00024 (disassembling.go:9)	LEAQ	32(SP), BP
	0x001d 00029 (disassembling.go:9)	FUNCDATA	$0, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x001d 00029 (disassembling.go:9)	FUNCDATA	$1, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x001d 00029 (disassembling.go:9)	FUNCDATA	$2, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x001d 00029 (disassembling.go:10)	PCDATA	$0, $0
	0x001d 00029 (disassembling.go:10)	PCDATA	$1, $0
	0x001d 00029 (disassembling.go:10)	MOVQ	$10, (SP)
	0x0025 00037 (disassembling.go:10)	MOVQ	$20, 8(SP)
	0x002e 00046 (disassembling.go:10)	CALL	"".swap(SB)
	0x0033 00051 (disassembling.go:11)	MOVQ	32(SP), BP
	0x0038 00056 (disassembling.go:11)	ADDQ	$40, SP
	0x003c 00060 (disassembling.go:11)	RET
	0x003d 00061 (disassembling.go:11)	NOP
	0x003d 00061 (disassembling.go:9)	PCDATA	$1, $-1
	0x003d 00061 (disassembling.go:9)	PCDATA	$0, $-1
	0x003d 00061 (disassembling.go:9)	CALL	runtime.morestack_noctxt(SB)
	0x0042 00066 (disassembling.go:9)	JMP	0
	0x0000 64 48 8b 0c 25 00 00 00 00 48 3b 61 10 76 2e 48  dH..%....H;a.v.H
	0x0010 83 ec 28 48 89 6c 24 20 48 8d 6c 24 20 48 c7 04  ..(H.l$ H.l$ H..
	0x0020 24 0a 00 00 00 48 c7 44 24 08 14 00 00 00 e8 00  $....H.D$.......
	0x0030 00 00 00 48 8b 6c 24 20 48 83 c4 28 c3 e8 00 00  ...H.l$ H..(....
	0x0040 00 00 eb bc                                      ....
	rel 5+4 t=16 TLS+0
	rel 47+4 t=8 "".swap+0
	rel 62+4 t=8 runtime.morestack_noctxt+0
go.cuinfo.packagename. SDWARFINFO dupok size=0
	0x0000 6d 61 69 6e                                      main
go.loc."".swap SDWARFLOC size=0
go.info."".swap SDWARFINFO size=80
	0x0000 03 22 22 2e 73 77 61 70 00 00 00 00 00 00 00 00  ."".swap........
	0x0010 00 00 00 00 00 00 00 00 00 01 9c 00 00 00 00 01  ................
	0x0020 0f 61 00 00 03 00 00 00 00 01 9c 0f 62 00 00 03  .a..........b...
	0x0030 00 00 00 00 02 91 08 0f 78 00 01 03 00 00 00 00  ........x.......
	0x0040 02 91 10 0f 79 00 01 03 00 00 00 00 02 91 18 00  ....y...........
	rel 9+8 t=1 "".swap+0
	rel 17+8 t=1 "".swap+39
	rel 27+4 t=29 gofile../Users/zzh/Desktop/github/go-learn/src/special/disassembling.go+0
	rel 37+4 t=28 go.info.int+0
	rel 48+4 t=28 go.info.int+0
	rel 60+4 t=28 go.info.int+0
	rel 72+4 t=28 go.info.int+0
go.range."".swap SDWARFRANGE size=0
go.isstmt."".swap SDWARFMISC size=0
	0x0000 04 09 01 09 02 05 01 05 02 05 01 05 02 01 00     ...............
go.loc."".main SDWARFLOC size=0
go.info."".main SDWARFINFO size=33
	0x0000 03 22 22 2e 6d 61 69 6e 00 00 00 00 00 00 00 00  ."".main........
	0x0010 00 00 00 00 00 00 00 00 00 01 9c 00 00 00 00 01  ................
	0x0020 00                                               .
	rel 9+8 t=1 "".main+0
	rel 17+8 t=1 "".main+68
	rel 27+4 t=29 gofile../Users/zzh/Desktop/github/go-learn/src/special/disassembling.go+0
go.range."".main SDWARFRANGE size=0
go.isstmt."".main SDWARFMISC size=0
	0x0000 04 0f 04 0e 03 08 01 09 02 16 00                 ...........
""..inittask SNOPTRDATA size=24
	0x0000 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0010 00 00 00 00 00 00 00 00                          ........
gclocals·33cdeccccebe80329f1fdbee7f5874cb SRODATA dupok size=8
	0x0000 01 00 00 00 00 00 00 00                          ........
