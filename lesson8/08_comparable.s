"".compare STEXT nosplit size=7 args=0x10 locals=0x0 funcid=0x0 align=0x0
	0x0000 00000 (08_comparable.go:7)	TEXT	"".compare(SB), NOSPLIT|ABIInternal, $0-16
	0x0000 00000 (08_comparable.go:7)	FUNCDATA	$0, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x0000 00000 (08_comparable.go:7)	FUNCDATA	$1, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x0000 00000 (08_comparable.go:7)	FUNCDATA	$5, "".compare.arginfo1(SB)
	0x0000 00000 (08_comparable.go:7)	FUNCDATA	$6, "".compare.argliveinfo(SB)
	0x0000 00000 (08_comparable.go:7)	PCDATA	$3, $1
	0x0000 00000 (08_comparable.go:8)	CMPQ	BX, AX
	0x0003 00003 (08_comparable.go:8)	SETGT	AL
	0x0006 00006 (08_comparable.go:8)	RET
	0x0000 48 39 c3 0f 9f c0 c3                             H9.....
"".main STEXT size=111 args=0x0 locals=0x40 funcid=0x0 align=0x0
	0x0000 00000 (08_comparable.go:11)	TEXT	"".main(SB), ABIInternal, $64-0
	0x0000 00000 (08_comparable.go:11)	CMPQ	SP, 16(R14)
	0x0004 00004 (08_comparable.go:11)	PCDATA	$0, $-2
	0x0004 00004 (08_comparable.go:11)	JLS	104
	0x0006 00006 (08_comparable.go:11)	PCDATA	$0, $-1
	0x0006 00006 (08_comparable.go:11)	SUBQ	$64, SP
	0x000a 00010 (08_comparable.go:11)	MOVQ	BP, 56(SP)
	0x000f 00015 (08_comparable.go:11)	LEAQ	56(SP), BP
	0x0014 00020 (08_comparable.go:11)	FUNCDATA	$0, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x0014 00020 (08_comparable.go:11)	FUNCDATA	$1, gclocals·f207267fbf96a0178e8758c6e3e0ce28(SB)
	0x0014 00020 (08_comparable.go:11)	FUNCDATA	$2, "".main.stkobj(SB)
	0x0014 00020 (<unknown line number>)	NOP
	0x0014 00020 (08_comparable.go:12)	MOVUPS	X15, ""..autotmp_12+40(SP)
	0x001a 00026 (08_comparable.go:12)	LEAQ	type.bool(SB), DX
	0x0021 00033 (08_comparable.go:12)	MOVQ	DX, ""..autotmp_12+40(SP)
	0x0026 00038 (08_comparable.go:12)	MOVL	$1, DX
	0x002b 00043 (08_comparable.go:12)	MOVBLZX	DL, DX
	0x002e 00046 (08_comparable.go:12)	LEAQ	runtime.staticuint64s(SB), R8
	0x0035 00053 (08_comparable.go:12)	LEAQ	(R8)(DX*8), DX
	0x0039 00057 (08_comparable.go:12)	MOVQ	DX, ""..autotmp_12+48(SP)
	0x003e 00062 (<unknown line number>)	NOP
	0x003e 00062 ($GOROOT/src/fmt/print.go:274)	MOVQ	os.Stdout(SB), BX
	0x0045 00069 ($GOROOT/src/fmt/print.go:274)	LEAQ	go.itab.*os.File,io.Writer(SB), AX
	0x004c 00076 ($GOROOT/src/fmt/print.go:274)	LEAQ	""..autotmp_12+40(SP), CX
	0x0051 00081 ($GOROOT/src/fmt/print.go:274)	MOVL	$1, DI
	0x0056 00086 ($GOROOT/src/fmt/print.go:274)	MOVQ	DI, SI
	0x0059 00089 ($GOROOT/src/fmt/print.go:274)	PCDATA	$1, $0
	0x0059 00089 ($GOROOT/src/fmt/print.go:274)	CALL	fmt.Fprintln(SB)
	0x005e 00094 (08_comparable.go:13)	MOVQ	56(SP), BP
	0x0063 00099 (08_comparable.go:13)	ADDQ	$64, SP
	0x0067 00103 (08_comparable.go:13)	RET
	0x0068 00104 (08_comparable.go:13)	NOP
	0x0068 00104 (08_comparable.go:11)	PCDATA	$1, $-1
	0x0068 00104 (08_comparable.go:11)	PCDATA	$0, $-2
	0x0068 00104 (08_comparable.go:11)	CALL	runtime.morestack_noctxt(SB)
	0x006d 00109 (08_comparable.go:11)	PCDATA	$0, $-1
	0x006d 00109 (08_comparable.go:11)	JMP	0
	0x0000 49 3b 66 10 76 62 48 83 ec 40 48 89 6c 24 38 48  I;f.vbH..@H.l$8H
	0x0010 8d 6c 24 38 44 0f 11 7c 24 28 48 8d 15 00 00 00  .l$8D..|$(H.....
	0x0020 00 48 89 54 24 28 ba 01 00 00 00 0f b6 d2 4c 8d  .H.T$(........L.
	0x0030 05 00 00 00 00 49 8d 14 d0 48 89 54 24 30 48 8b  .....I...H.T$0H.
	0x0040 1d 00 00 00 00 48 8d 05 00 00 00 00 48 8d 4c 24  .....H......H.L$
	0x0050 28 bf 01 00 00 00 48 89 fe e8 00 00 00 00 48 8b  (.....H.......H.
	0x0060 6c 24 38 48 83 c4 40 c3 e8 00 00 00 00 eb 91     l$8H..@........
	rel 2+0 t=23 type.bool+0
	rel 2+0 t=23 type.*os.File+0
	rel 29+4 t=14 type.bool+0
	rel 49+4 t=14 runtime.staticuint64s+0
	rel 65+4 t=14 os.Stdout+0
	rel 72+4 t=14 go.itab.*os.File,io.Writer+0
	rel 90+4 t=7 fmt.Fprintln+0
	rel 105+4 t=7 runtime.morestack_noctxt+0
go.cuinfo.packagename. SDWARFCUINFO dupok size=0
	0x0000 6d 61 69 6e                                      main
go.info."".compare$abstract SDWARFABSFCN dupok size=29
	0x0000 05 2e 63 6f 6d 70 61 72 65 00 01 01 13 78 00 00  ..compare....x..
	0x0010 00 00 00 00 13 79 00 00 00 00 00 00 00           .....y.......
	rel 0+0 t=22 type.bool+0
	rel 0+0 t=22 type.int+0
	rel 16+4 t=31 go.info.int+0
	rel 24+4 t=31 go.info.int+0
go.info.fmt.Println$abstract SDWARFABSFCN dupok size=42
	0x0000 05 66 6d 74 2e 50 72 69 6e 74 6c 6e 00 01 01 13  .fmt.Println....
	0x0010 61 00 00 00 00 00 00 13 6e 00 01 00 00 00 00 13  a.......n.......
	0x0020 65 72 72 00 01 00 00 00 00 00                    err.......
	rel 0+0 t=22 type.[]interface {}+0
	rel 0+0 t=22 type.error+0
	rel 0+0 t=22 type.int+0
	rel 19+4 t=31 go.info.[]interface {}+0
	rel 27+4 t=31 go.info.int+0
	rel 37+4 t=31 go.info.error+0
""..inittask SNOPTRDATA size=32
	0x0000 00 00 00 00 00 00 00 00 01 00 00 00 00 00 00 00  ................
	0x0010 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	rel 24+8 t=1 fmt..inittask+0
go.itab.*os.File,io.Writer SRODATA dupok size=32
	0x0000 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0010 44 b5 f3 33 00 00 00 00 00 00 00 00 00 00 00 00  D..3............
	rel 0+8 t=1 type.io.Writer+0
	rel 8+8 t=1 type.*os.File+0
	rel 24+8 t=-32767 os.(*File).Write+0
runtime.nilinterequal·f SRODATA dupok size=8
	0x0000 00 00 00 00 00 00 00 00                          ........
	rel 0+8 t=1 runtime.nilinterequal+0
runtime.memequal64·f SRODATA dupok size=8
	0x0000 00 00 00 00 00 00 00 00                          ........
	rel 0+8 t=1 runtime.memequal64+0
runtime.gcbits.01 SRODATA dupok size=1
	0x0000 01                                               .
type..namedata.*interface {}- SRODATA dupok size=15
	0x0000 00 0d 2a 69 6e 74 65 72 66 61 63 65 20 7b 7d     ..*interface {}
type.*interface {} SRODATA dupok size=56
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 4f 0f 96 9d 08 08 08 36 00 00 00 00 00 00 00 00  O......6........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.memequal64·f+0
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*interface {}-+0
	rel 48+8 t=1 type.interface {}+0
runtime.gcbits.02 SRODATA dupok size=1
	0x0000 02                                               .
type.interface {} SRODATA dupok size=80
	0x0000 10 00 00 00 00 00 00 00 10 00 00 00 00 00 00 00  ................
	0x0010 e7 57 a0 18 02 08 08 14 00 00 00 00 00 00 00 00  .W..............
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0040 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	rel 24+8 t=1 runtime.nilinterequal·f+0
	rel 32+8 t=1 runtime.gcbits.02+0
	rel 40+4 t=5 type..namedata.*interface {}-+0
	rel 44+4 t=-32763 type.*interface {}+0
	rel 56+8 t=1 type.interface {}+80
type..namedata.*[]interface {}- SRODATA dupok size=17
	0x0000 00 0f 2a 5b 5d 69 6e 74 65 72 66 61 63 65 20 7b  ..*[]interface {
	0x0010 7d                                               }
type.*[]interface {} SRODATA dupok size=56
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 f3 04 9a e7 08 08 08 36 00 00 00 00 00 00 00 00  .......6........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.memequal64·f+0
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*[]interface {}-+0
	rel 48+8 t=1 type.[]interface {}+0
type.[]interface {} SRODATA dupok size=56
	0x0000 18 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 70 93 ea 2f 02 08 08 17 00 00 00 00 00 00 00 00  p../............
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00                          ........
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*[]interface {}-+0
	rel 44+4 t=-32763 type.*[]interface {}+0
	rel 48+8 t=1 type.interface {}+0
type..importpath.fmt. SRODATA dupok size=5
	0x0000 00 03 66 6d 74                                   ..fmt
gclocals·33cdeccccebe80329f1fdbee7f5874cb SRODATA dupok size=8
	0x0000 01 00 00 00 00 00 00 00                          ........
"".compare.arginfo1 SRODATA static dupok size=5
	0x0000 00 08 08 08 ff                                   .....
"".compare.argliveinfo SRODATA static dupok size=2
	0x0000 00 00                                            ..
gclocals·f207267fbf96a0178e8758c6e3e0ce28 SRODATA dupok size=9
	0x0000 01 00 00 00 02 00 00 00 00                       .........
"".main.stkobj SRODATA static size=24
	0x0000 01 00 00 00 00 00 00 00 f0 ff ff ff 10 00 00 00  ................
	0x0010 10 00 00 00 00 00 00 00                          ........
	rel 20+4 t=5 runtime.gcbits.02+0
