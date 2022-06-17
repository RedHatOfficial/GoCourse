"".main STEXT size=293 args=0x0 locals=0x68 funcid=0x0 align=0x0
	0x0000 00000 (10_compare_type_parameters.go:15)	TEXT	"".main(SB), ABIInternal, $104-0
	0x0000 00000 (10_compare_type_parameters.go:15)	CMPQ	SP, 16(R14)
	0x0004 00004 (10_compare_type_parameters.go:15)	PCDATA	$0, $-2
	0x0004 00004 (10_compare_type_parameters.go:15)	JLS	282
	0x000a 00010 (10_compare_type_parameters.go:15)	PCDATA	$0, $-1
	0x000a 00010 (10_compare_type_parameters.go:15)	SUBQ	$104, SP
	0x000e 00014 (10_compare_type_parameters.go:15)	MOVQ	BP, 96(SP)
	0x0013 00019 (10_compare_type_parameters.go:15)	LEAQ	96(SP), BP
	0x0018 00024 (10_compare_type_parameters.go:15)	FUNCDATA	$0, gclocals·69c1753bd5f81501d95132d08af04464(SB)
	0x0018 00024 (10_compare_type_parameters.go:15)	FUNCDATA	$1, gclocals·232077072e4d4c4b841d7a2024b5b669(SB)
	0x0018 00024 (10_compare_type_parameters.go:15)	FUNCDATA	$2, "".main.stkobj(SB)
	0x0018 00024 (<unknown line number>)	NOP
	0x0018 00024 (10_compare_type_parameters.go:16)	MOVUPS	X15, ""..autotmp_31+80(SP)
	0x001e 00030 (10_compare_type_parameters.go:16)	LEAQ	type.bool(SB), DX
	0x0025 00037 (10_compare_type_parameters.go:16)	MOVQ	DX, ""..autotmp_31+80(SP)
	0x002a 00042 (10_compare_type_parameters.go:16)	MOVL	$1, R8
	0x0030 00048 (10_compare_type_parameters.go:16)	MOVBLZX	R8B, R8
	0x0034 00052 (10_compare_type_parameters.go:16)	LEAQ	runtime.staticuint64s(SB), R9
	0x003b 00059 (10_compare_type_parameters.go:16)	LEAQ	(R9)(R8*8), R8
	0x003f 00063 (10_compare_type_parameters.go:16)	MOVQ	R8, ""..autotmp_49+40(SP)
	0x0044 00068 (10_compare_type_parameters.go:16)	MOVQ	R8, ""..autotmp_31+88(SP)
	0x0049 00073 (<unknown line number>)	NOP
	0x0049 00073 ($GOROOT/src/fmt/print.go:274)	MOVQ	os.Stdout(SB), BX
	0x0050 00080 ($GOROOT/src/fmt/print.go:274)	LEAQ	go.itab.*os.File,io.Writer(SB), AX
	0x0057 00087 ($GOROOT/src/fmt/print.go:274)	LEAQ	""..autotmp_31+80(SP), CX
	0x005c 00092 ($GOROOT/src/fmt/print.go:274)	MOVL	$1, DI
	0x0061 00097 ($GOROOT/src/fmt/print.go:274)	MOVQ	DI, SI
	0x0064 00100 ($GOROOT/src/fmt/print.go:274)	PCDATA	$1, $1
	0x0064 00100 ($GOROOT/src/fmt/print.go:274)	CALL	fmt.Fprintln(SB)
	0x0069 00105 (10_compare_type_parameters.go:17)	MOVUPS	X15, ""..autotmp_33+64(SP)
	0x006f 00111 (10_compare_type_parameters.go:17)	LEAQ	type.bool(SB), DX
	0x0076 00118 (10_compare_type_parameters.go:17)	MOVQ	DX, ""..autotmp_33+64(SP)
	0x007b 00123 (10_compare_type_parameters.go:17)	MOVQ	""..autotmp_49+40(SP), R8
	0x0080 00128 (10_compare_type_parameters.go:17)	MOVQ	R8, ""..autotmp_33+72(SP)
	0x0085 00133 (<unknown line number>)	NOP
	0x0085 00133 ($GOROOT/src/fmt/print.go:274)	MOVQ	os.Stdout(SB), BX
	0x008c 00140 ($GOROOT/src/fmt/print.go:274)	LEAQ	go.itab.*os.File,io.Writer(SB), AX
	0x0093 00147 ($GOROOT/src/fmt/print.go:274)	LEAQ	""..autotmp_33+64(SP), CX
	0x0098 00152 ($GOROOT/src/fmt/print.go:274)	MOVL	$1, DI
	0x009d 00157 ($GOROOT/src/fmt/print.go:274)	MOVQ	DI, SI
	0x00a0 00160 ($GOROOT/src/fmt/print.go:274)	PCDATA	$1, $0
	0x00a0 00160 ($GOROOT/src/fmt/print.go:274)	CALL	fmt.Fprintln(SB)
	0x00a5 00165 (<unknown line number>)	NOP
	0x00a5 00165 (10_compare_type_parameters.go:12)	LEAQ	go.string."foo"(SB), AX
	0x00ac 00172 (10_compare_type_parameters.go:12)	MOVL	$3, BX
	0x00b1 00177 (10_compare_type_parameters.go:12)	LEAQ	go.string."bar"(SB), CX
	0x00b8 00184 (10_compare_type_parameters.go:12)	MOVQ	BX, DI
	0x00bb 00187 (10_compare_type_parameters.go:12)	NOP
	0x00c0 00192 (10_compare_type_parameters.go:12)	CALL	runtime.cmpstring(SB)
	0x00c5 00197 (10_compare_type_parameters.go:18)	MOVUPS	X15, ""..autotmp_35+48(SP)
	0x00cb 00203 (10_compare_type_parameters.go:18)	LEAQ	type.bool(SB), DX
	0x00d2 00210 (10_compare_type_parameters.go:18)	MOVQ	DX, ""..autotmp_35+48(SP)
	0x00d7 00215 (10_compare_type_parameters.go:12)	TESTQ	AX, AX
	0x00da 00218 (10_compare_type_parameters.go:12)	SETLT	DL
	0x00dd 00221 (10_compare_type_parameters.go:18)	MOVBLZX	DL, DX
	0x00e0 00224 (10_compare_type_parameters.go:18)	LEAQ	runtime.staticuint64s(SB), R8
	0x00e7 00231 (10_compare_type_parameters.go:18)	LEAQ	(R8)(DX*8), DX
	0x00eb 00235 (10_compare_type_parameters.go:18)	MOVQ	DX, ""..autotmp_35+56(SP)
	0x00f0 00240 (<unknown line number>)	NOP
	0x00f0 00240 ($GOROOT/src/fmt/print.go:274)	MOVQ	os.Stdout(SB), BX
	0x00f7 00247 ($GOROOT/src/fmt/print.go:274)	LEAQ	go.itab.*os.File,io.Writer(SB), AX
	0x00fe 00254 ($GOROOT/src/fmt/print.go:274)	LEAQ	""..autotmp_35+48(SP), CX
	0x0103 00259 ($GOROOT/src/fmt/print.go:274)	MOVL	$1, DI
	0x0108 00264 ($GOROOT/src/fmt/print.go:274)	MOVQ	DI, SI
	0x010b 00267 ($GOROOT/src/fmt/print.go:274)	CALL	fmt.Fprintln(SB)
	0x0110 00272 (10_compare_type_parameters.go:19)	MOVQ	96(SP), BP
	0x0115 00277 (10_compare_type_parameters.go:19)	ADDQ	$104, SP
	0x0119 00281 (10_compare_type_parameters.go:19)	RET
	0x011a 00282 (10_compare_type_parameters.go:19)	NOP
	0x011a 00282 (10_compare_type_parameters.go:15)	PCDATA	$1, $-1
	0x011a 00282 (10_compare_type_parameters.go:15)	PCDATA	$0, $-2
	0x011a 00282 (10_compare_type_parameters.go:15)	CALL	runtime.morestack_noctxt(SB)
	0x011f 00287 (10_compare_type_parameters.go:15)	PCDATA	$0, $-1
	0x011f 00287 (10_compare_type_parameters.go:15)	NOP
	0x0120 00288 (10_compare_type_parameters.go:15)	JMP	0
	0x0000 49 3b 66 10 0f 86 10 01 00 00 48 83 ec 68 48 89  I;f.......H..hH.
	0x0010 6c 24 60 48 8d 6c 24 60 44 0f 11 7c 24 50 48 8d  l$`H.l$`D..|$PH.
	0x0020 15 00 00 00 00 48 89 54 24 50 41 b8 01 00 00 00  .....H.T$PA.....
	0x0030 45 0f b6 c0 4c 8d 0d 00 00 00 00 4f 8d 04 c1 4c  E...L......O...L
	0x0040 89 44 24 28 4c 89 44 24 58 48 8b 1d 00 00 00 00  .D$(L.D$XH......
	0x0050 48 8d 05 00 00 00 00 48 8d 4c 24 50 bf 01 00 00  H......H.L$P....
	0x0060 00 48 89 fe e8 00 00 00 00 44 0f 11 7c 24 40 48  .H.......D..|$@H
	0x0070 8d 15 00 00 00 00 48 89 54 24 40 4c 8b 44 24 28  ......H.T$@L.D$(
	0x0080 4c 89 44 24 48 48 8b 1d 00 00 00 00 48 8d 05 00  L.D$HH......H...
	0x0090 00 00 00 48 8d 4c 24 40 bf 01 00 00 00 48 89 fe  ...H.L$@.....H..
	0x00a0 e8 00 00 00 00 48 8d 05 00 00 00 00 bb 03 00 00  .....H..........
	0x00b0 00 48 8d 0d 00 00 00 00 48 89 df 0f 1f 44 00 00  .H......H....D..
	0x00c0 e8 00 00 00 00 44 0f 11 7c 24 30 48 8d 15 00 00  .....D..|$0H....
	0x00d0 00 00 48 89 54 24 30 48 85 c0 0f 9c c2 0f b6 d2  ..H.T$0H........
	0x00e0 4c 8d 05 00 00 00 00 49 8d 14 d0 48 89 54 24 38  L......I...H.T$8
	0x00f0 48 8b 1d 00 00 00 00 48 8d 05 00 00 00 00 48 8d  H......H......H.
	0x0100 4c 24 30 bf 01 00 00 00 48 89 fe e8 00 00 00 00  L$0.....H.......
	0x0110 48 8b 6c 24 60 48 83 c4 68 c3 e8 00 00 00 00 90  H.l$`H..h.......
	0x0120 e9 db fe ff ff                                   .....
	rel 3+0 t=23 type.bool+0
	rel 3+0 t=23 type.*os.File+0
	rel 3+0 t=23 type.bool+0
	rel 3+0 t=23 type.*os.File+0
	rel 3+0 t=23 type.bool+0
	rel 3+0 t=23 type.*os.File+0
	rel 33+4 t=14 type.bool+0
	rel 55+4 t=14 runtime.staticuint64s+0
	rel 76+4 t=14 os.Stdout+0
	rel 83+4 t=14 go.itab.*os.File,io.Writer+0
	rel 101+4 t=7 fmt.Fprintln+0
	rel 114+4 t=14 type.bool+0
	rel 136+4 t=14 os.Stdout+0
	rel 143+4 t=14 go.itab.*os.File,io.Writer+0
	rel 161+4 t=7 fmt.Fprintln+0
	rel 168+4 t=14 go.string."foo"+0
	rel 180+4 t=14 go.string."bar"+0
	rel 193+4 t=7 runtime.cmpstring+0
	rel 206+4 t=14 type.bool+0
	rel 227+4 t=14 runtime.staticuint64s+0
	rel 243+4 t=14 os.Stdout+0
	rel 250+4 t=14 go.itab.*os.File,io.Writer+0
	rel 268+4 t=7 fmt.Fprintln+0
	rel 283+4 t=7 runtime.morestack_noctxt+0
"".compare[go.shape.int_0] STEXT dupok nosplit size=7 args=0x18 locals=0x0 funcid=0x0 align=0x0
	0x0000 00000 (10_compare_type_parameters.go:11)	TEXT	"".compare[go.shape.int_0](SB), DUPOK|NOSPLIT|ABIInternal, $0-24
	0x0000 00000 (10_compare_type_parameters.go:11)	FUNCDATA	$0, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x0000 00000 (10_compare_type_parameters.go:11)	FUNCDATA	$1, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x0000 00000 (10_compare_type_parameters.go:11)	FUNCDATA	$5, "".compare[go.shape.int_0].arginfo1(SB)
	0x0000 00000 (10_compare_type_parameters.go:11)	FUNCDATA	$6, "".compare[go.shape.int_0].argliveinfo(SB)
	0x0000 00000 (10_compare_type_parameters.go:11)	PCDATA	$3, $1
	0x0000 00000 (10_compare_type_parameters.go:12)	CMPQ	CX, BX
	0x0003 00003 (10_compare_type_parameters.go:12)	SETGT	AL
	0x0006 00006 (10_compare_type_parameters.go:12)	RET
	0x0000 48 39 d9 0f 9f c0 c3                             H9.....
"".compare[go.shape.float64_0] STEXT dupok nosplit size=8 args=0x18 locals=0x0 funcid=0x0 align=0x0
	0x0000 00000 (10_compare_type_parameters.go:11)	TEXT	"".compare[go.shape.float64_0](SB), DUPOK|NOSPLIT|ABIInternal, $0-24
	0x0000 00000 (10_compare_type_parameters.go:11)	FUNCDATA	$0, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x0000 00000 (10_compare_type_parameters.go:11)	FUNCDATA	$1, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x0000 00000 (10_compare_type_parameters.go:11)	FUNCDATA	$5, "".compare[go.shape.float64_0].arginfo1(SB)
	0x0000 00000 (10_compare_type_parameters.go:11)	FUNCDATA	$6, "".compare[go.shape.float64_0].argliveinfo(SB)
	0x0000 00000 (10_compare_type_parameters.go:11)	PCDATA	$3, $1
	0x0000 00000 (10_compare_type_parameters.go:12)	UCOMISD	X0, X1
	0x0004 00004 (10_compare_type_parameters.go:12)	SETHI	AL
	0x0007 00007 (10_compare_type_parameters.go:12)	RET
	0x0000 66 0f 2e c8 0f 97 c0 c3                          f.......
"".compare[go.shape.string_0] STEXT dupok size=120 args=0x28 locals=0x28 funcid=0x0 align=0x0
	0x0000 00000 (10_compare_type_parameters.go:11)	TEXT	"".compare[go.shape.string_0](SB), DUPOK|ABIInternal, $40-40
	0x0000 00000 (10_compare_type_parameters.go:11)	CMPQ	SP, 16(R14)
	0x0004 00004 (10_compare_type_parameters.go:11)	PCDATA	$0, $-2
	0x0004 00004 (10_compare_type_parameters.go:11)	JLS	63
	0x0006 00006 (10_compare_type_parameters.go:11)	PCDATA	$0, $-1
	0x0006 00006 (10_compare_type_parameters.go:11)	SUBQ	$40, SP
	0x000a 00010 (10_compare_type_parameters.go:11)	MOVQ	BP, 32(SP)
	0x000f 00015 (10_compare_type_parameters.go:11)	LEAQ	32(SP), BP
	0x0014 00020 (10_compare_type_parameters.go:11)	MOVQ	BX, "".x+56(FP)
	0x0019 00025 (10_compare_type_parameters.go:11)	MOVQ	DI, "".y+72(FP)
	0x001e 00030 (10_compare_type_parameters.go:11)	FUNCDATA	$0, gclocals·7a680c56c7799a8f60d071b2f2541840(SB)
	0x001e 00030 (10_compare_type_parameters.go:11)	FUNCDATA	$1, gclocals·69c1753bd5f81501d95132d08af04464(SB)
	0x001e 00030 (10_compare_type_parameters.go:11)	FUNCDATA	$5, "".compare[go.shape.string_0].arginfo1(SB)
	0x001e 00030 (10_compare_type_parameters.go:11)	FUNCDATA	$6, "".compare[go.shape.string_0].argliveinfo(SB)
	0x001e 00030 (10_compare_type_parameters.go:11)	PCDATA	$3, $1
	0x001e 00030 (10_compare_type_parameters.go:12)	MOVQ	BX, AX
	0x0021 00033 (10_compare_type_parameters.go:12)	MOVQ	CX, BX
	0x0024 00036 (10_compare_type_parameters.go:12)	MOVQ	DI, CX
	0x0027 00039 (10_compare_type_parameters.go:12)	MOVQ	SI, DI
	0x002a 00042 (10_compare_type_parameters.go:12)	PCDATA	$1, $1
	0x002a 00042 (10_compare_type_parameters.go:12)	CALL	runtime.cmpstring(SB)
	0x002f 00047 (10_compare_type_parameters.go:12)	TESTQ	AX, AX
	0x0032 00050 (10_compare_type_parameters.go:12)	SETLT	AL
	0x0035 00053 (10_compare_type_parameters.go:12)	MOVQ	32(SP), BP
	0x003a 00058 (10_compare_type_parameters.go:12)	ADDQ	$40, SP
	0x003e 00062 (10_compare_type_parameters.go:12)	RET
	0x003f 00063 (10_compare_type_parameters.go:12)	NOP
	0x003f 00063 (10_compare_type_parameters.go:11)	PCDATA	$1, $-1
	0x003f 00063 (10_compare_type_parameters.go:11)	PCDATA	$0, $-2
	0x003f 00063 (10_compare_type_parameters.go:11)	MOVQ	AX, 8(SP)
	0x0044 00068 (10_compare_type_parameters.go:11)	MOVQ	BX, 16(SP)
	0x0049 00073 (10_compare_type_parameters.go:11)	MOVQ	CX, 24(SP)
	0x004e 00078 (10_compare_type_parameters.go:11)	MOVQ	DI, 32(SP)
	0x0053 00083 (10_compare_type_parameters.go:11)	MOVQ	SI, 40(SP)
	0x0058 00088 (10_compare_type_parameters.go:11)	CALL	runtime.morestack_noctxt(SB)
	0x005d 00093 (10_compare_type_parameters.go:11)	MOVQ	8(SP), AX
	0x0062 00098 (10_compare_type_parameters.go:11)	MOVQ	16(SP), BX
	0x0067 00103 (10_compare_type_parameters.go:11)	MOVQ	24(SP), CX
	0x006c 00108 (10_compare_type_parameters.go:11)	MOVQ	32(SP), DI
	0x0071 00113 (10_compare_type_parameters.go:11)	MOVQ	40(SP), SI
	0x0076 00118 (10_compare_type_parameters.go:11)	PCDATA	$0, $-1
	0x0076 00118 (10_compare_type_parameters.go:11)	JMP	0
	0x0000 49 3b 66 10 76 39 48 83 ec 28 48 89 6c 24 20 48  I;f.v9H..(H.l$ H
	0x0010 8d 6c 24 20 48 89 5c 24 38 48 89 7c 24 48 48 89  .l$ H.\$8H.|$HH.
	0x0020 d8 48 89 cb 48 89 f9 48 89 f7 e8 00 00 00 00 48  .H..H..H.......H
	0x0030 85 c0 0f 9c c0 48 8b 6c 24 20 48 83 c4 28 c3 48  .....H.l$ H..(.H
	0x0040 89 44 24 08 48 89 5c 24 10 48 89 4c 24 18 48 89  .D$.H.\$.H.L$.H.
	0x0050 7c 24 20 48 89 74 24 28 e8 00 00 00 00 48 8b 44  |$ H.t$(.....H.D
	0x0060 24 08 48 8b 5c 24 10 48 8b 4c 24 18 48 8b 7c 24  $.H.\$.H.L$.H.|$
	0x0070 20 48 8b 74 24 28 eb 88                           H.t$(..
	rel 43+4 t=7 runtime.cmpstring+0
	rel 89+4 t=7 runtime.morestack_noctxt+0
""..dict.compare[int] SRODATA dupok size=8
	0x0000 00 00 00 00 00 00 00 00                          ........
	rel 0+8 t=1 type.int+0
	rel 0+0 t=23 type.int+0
""..dict.compare[float64] SRODATA dupok size=8
	0x0000 00 00 00 00 00 00 00 00                          ........
	rel 0+8 t=1 type.float64+0
	rel 0+0 t=23 type.float64+0
""..dict.compare[string] SRODATA dupok size=8
	0x0000 00 00 00 00 00 00 00 00                          ........
	rel 0+8 t=1 type.string+0
	rel 0+0 t=23 type.string+0
go.cuinfo.packagename. SDWARFCUINFO dupok size=0
	0x0000 6d 61 69 6e                                      main
go.info."".compare[go.shape.int_0]$abstract SDWARFABSFCN dupok size=45
	0x0000 05 2e 63 6f 6d 70 61 72 65 5b 67 6f 2e 73 68 61  ..compare[go.sha
	0x0010 70 65 2e 69 6e 74 5f 30 5d 00 01 01 13 78 00 00  pe.int_0]....x..
	0x0020 00 00 00 00 13 79 00 00 00 00 00 00 00           .....y.......
	rel 0+0 t=22 type.bool+0
	rel 0+0 t=22 type.go.shape.int_0+0
	rel 32+4 t=31 go.info.go.shape.int_0+0
	rel 40+4 t=31 go.info.go.shape.int_0+0
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
go.info."".compare[go.shape.float64_0]$abstract SDWARFABSFCN dupok size=49
	0x0000 05 2e 63 6f 6d 70 61 72 65 5b 67 6f 2e 73 68 61  ..compare[go.sha
	0x0010 70 65 2e 66 6c 6f 61 74 36 34 5f 30 5d 00 01 01  pe.float64_0]...
	0x0020 13 78 00 00 00 00 00 00 13 79 00 00 00 00 00 00  .x.......y......
	0x0030 00                                               .
	rel 0+0 t=22 type.bool+0
	rel 0+0 t=22 type.go.shape.float64_0+0
	rel 36+4 t=31 go.info.go.shape.float64_0+0
	rel 44+4 t=31 go.info.go.shape.float64_0+0
go.info."".compare[go.shape.string_0]$abstract SDWARFABSFCN dupok size=48
	0x0000 05 2e 63 6f 6d 70 61 72 65 5b 67 6f 2e 73 68 61  ..compare[go.sha
	0x0010 70 65 2e 73 74 72 69 6e 67 5f 30 5d 00 01 01 13  pe.string_0]....
	0x0020 78 00 00 00 00 00 00 13 79 00 00 00 00 00 00 00  x.......y.......
	rel 0+0 t=22 type.bool+0
	rel 0+0 t=22 type.go.shape.string_0+0
	rel 35+4 t=31 go.info.go.shape.string_0+0
	rel 43+4 t=31 go.info.go.shape.string_0+0
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
go.string."foo" SRODATA dupok size=3
	0x0000 66 6f 6f                                         foo
go.string."bar" SRODATA dupok size=3
	0x0000 62 61 72                                         bar
runtime.nilinterequal·f SRODATA dupok size=8
	0x0000 00 00 00 00 00 00 00 00                          ........
	rel 0+8 t=1 runtime.nilinterequal+0
runtime.memequal64·f SRODATA dupok size=8
	0x0000 00 00 00 00 00 00 00 00                          ........
	rel 0+8 t=1 runtime.memequal64+0
runtime.gcbits.01 SRODATA dupok size=1
	0x0000 01                                               .
type..namedata.*main.comparable- SRODATA dupok size=18
	0x0000 00 10 2a 6d 61 69 6e 2e 63 6f 6d 70 61 72 61 62  ..*main.comparab
	0x0010 6c 65                                            le
type.*"".comparable SRODATA dupok size=56
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 73 c6 88 be 08 08 08 36 00 00 00 00 00 00 00 00  s......6........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.memequal64·f+0
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*main.comparable-+0
	rel 48+8 t=1 type."".comparable+0
runtime.gcbits.02 SRODATA dupok size=1
	0x0000 02                                               .
type."".comparable SRODATA dupok size=96
	0x0000 10 00 00 00 00 00 00 00 10 00 00 00 00 00 00 00  ................
	0x0010 03 ac cf f0 07 08 08 14 00 00 00 00 00 00 00 00  ................
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0040 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0050 00 00 00 00 00 00 00 00 10 00 00 00 00 00 00 00  ................
	rel 24+8 t=1 runtime.nilinterequal·f+0
	rel 32+8 t=1 runtime.gcbits.02+0
	rel 40+4 t=5 type..namedata.*main.comparable-+0
	rel 44+4 t=5 type.*"".comparable+0
	rel 48+8 t=1 type..importpath."".+0
	rel 56+8 t=1 type."".comparable+96
	rel 80+4 t=5 type..importpath."".+0
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
runtime.f64equal·f SRODATA dupok size=8
	0x0000 00 00 00 00 00 00 00 00                          ........
	rel 0+8 t=1 runtime.f64equal+0
type..namedata.*go.shape.float64_0- SRODATA dupok size=21
	0x0000 00 13 2a 67 6f 2e 73 68 61 70 65 2e 66 6c 6f 61  ..*go.shape.floa
	0x0010 74 36 34 5f 30                                   t64_0
type.*go.shape.float64_0 SRODATA dupok size=56
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 de 48 d1 1b 08 08 08 36 00 00 00 00 00 00 00 00  .H.....6........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.memequal64·f+0
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*go.shape.float64_0-+0
	rel 48+8 t=1 type.go.shape.float64_0+0
runtime.gcbits. SRODATA dupok size=0
type..importpath.go.shape. SRODATA dupok size=10
	0x0000 00 08 67 6f 2e 73 68 61 70 65                    ..go.shape
type.go.shape.float64_0 SRODATA dupok size=64
	0x0000 08 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0010 81 6f 93 05 07 08 08 0e 00 00 00 00 00 00 00 00  .o..............
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00 10 00 00 00 00 00 00 00  ................
	rel 24+8 t=1 runtime.f64equal·f+0
	rel 32+8 t=1 runtime.gcbits.+0
	rel 40+4 t=5 type..namedata.*go.shape.float64_0-+0
	rel 44+4 t=5 type.*go.shape.float64_0+0
	rel 48+4 t=5 type..importpath.go.shape.+0
type..namedata.*go.shape.int_0- SRODATA dupok size=17
	0x0000 00 0f 2a 67 6f 2e 73 68 61 70 65 2e 69 6e 74 5f  ..*go.shape.int_
	0x0010 30                                               0
type.*go.shape.int_0 SRODATA dupok size=56
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 cb 09 08 b4 08 08 08 36 00 00 00 00 00 00 00 00  .......6........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.memequal64·f+0
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*go.shape.int_0-+0
	rel 48+8 t=1 type.go.shape.int_0+0
type.go.shape.int_0 SRODATA dupok size=64
	0x0000 08 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0010 89 63 aa a5 0f 08 08 02 00 00 00 00 00 00 00 00  .c..............
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00 10 00 00 00 00 00 00 00  ................
	rel 24+8 t=1 runtime.memequal64·f+0
	rel 32+8 t=1 runtime.gcbits.+0
	rel 40+4 t=5 type..namedata.*go.shape.int_0-+0
	rel 44+4 t=5 type.*go.shape.int_0+0
	rel 48+4 t=5 type..importpath.go.shape.+0
runtime.strequal·f SRODATA dupok size=8
	0x0000 00 00 00 00 00 00 00 00                          ........
	rel 0+8 t=1 runtime.strequal+0
type..namedata.*go.shape.string_0- SRODATA dupok size=20
	0x0000 00 12 2a 67 6f 2e 73 68 61 70 65 2e 73 74 72 69  ..*go.shape.stri
	0x0010 6e 67 5f 30                                      ng_0
type.*go.shape.string_0 SRODATA dupok size=56
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 54 6f 54 9a 08 08 08 36 00 00 00 00 00 00 00 00  ToT....6........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.memequal64·f+0
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*go.shape.string_0-+0
	rel 48+8 t=1 type.go.shape.string_0+0
type.go.shape.string_0 SRODATA dupok size=64
	0x0000 10 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 b4 93 c9 2b 07 08 08 18 00 00 00 00 00 00 00 00  ...+............
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00 10 00 00 00 00 00 00 00  ................
	rel 24+8 t=1 runtime.strequal·f+0
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*go.shape.string_0-+0
	rel 44+4 t=5 type.*go.shape.string_0+0
	rel 48+4 t=5 type..importpath.go.shape.+0
type..importpath.fmt. SRODATA dupok size=5
	0x0000 00 03 66 6d 74                                   ..fmt
gclocals·69c1753bd5f81501d95132d08af04464 SRODATA dupok size=8
	0x0000 02 00 00 00 00 00 00 00                          ........
gclocals·232077072e4d4c4b841d7a2024b5b669 SRODATA dupok size=10
	0x0000 02 00 00 00 07 00 00 00 00 01                    ..........
"".main.stkobj SRODATA static size=56
	0x0000 03 00 00 00 00 00 00 00 d0 ff ff ff 10 00 00 00  ................
	0x0010 10 00 00 00 00 00 00 00 e0 ff ff ff 10 00 00 00  ................
	0x0020 10 00 00 00 00 00 00 00 f0 ff ff ff 10 00 00 00  ................
	0x0030 10 00 00 00 00 00 00 00                          ........
	rel 20+4 t=5 runtime.gcbits.02+0
	rel 36+4 t=5 runtime.gcbits.02+0
	rel 52+4 t=5 runtime.gcbits.02+0
gclocals·33cdeccccebe80329f1fdbee7f5874cb SRODATA dupok size=8
	0x0000 01 00 00 00 00 00 00 00                          ........
"".compare[go.shape.int_0].arginfo1 SRODATA static dupok size=5
	0x0000 08 08 10 08 ff                                   .....
"".compare[go.shape.int_0].argliveinfo SRODATA static dupok size=2
	0x0000 00 00                                            ..
"".compare[go.shape.float64_0].arginfo1 SRODATA static dupok size=5
	0x0000 08 08 10 08 ff                                   .....
"".compare[go.shape.float64_0].argliveinfo SRODATA static dupok size=2
	0x0000 00 00                                            ..
gclocals·7a680c56c7799a8f60d071b2f2541840 SRODATA dupok size=10
	0x0000 02 00 00 00 04 00 00 00 0a 00                    ..........
"".compare[go.shape.string_0].arginfo1 SRODATA static dupok size=13
	0x0000 fe 08 08 10 08 fd fe 18 08 20 08 fd ff           ......... ...
"".compare[go.shape.string_0].argliveinfo SRODATA static dupok size=2
	0x0000 00 00                                            ..
