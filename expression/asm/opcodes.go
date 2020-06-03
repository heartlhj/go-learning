package asm

type Opcodes int

const (
	SAM4               Opcodes = 262144
	ASM4               Opcodes = 262144
	ASM5               Opcodes = 327680
	ASM6               Opcodes = 393216
	ASM7               Opcodes = 458752
	V1_1               Opcodes = 196653
	V1_2               Opcodes = 46
	V1_3               Opcodes = 47
	V1_4               Opcodes = 48
	V1_5               Opcodes = 49
	V1_6               Opcodes = 50
	V1_7               Opcodes = 51
	V1_8               Opcodes = 52
	V9                 Opcodes = 53
	V10                Opcodes = 54
	V11                Opcodes = 55
	V12                Opcodes = 56
	V_PREVIEW          Opcodes = -65536
	ACC_PUBLIC         Opcodes = 1
	ACC_PRIVATE        Opcodes = 2
	ACC_PROTECTED      Opcodes = 4
	ACC_STATIC         Opcodes = 8
	ACC_FINAL          Opcodes = 16
	ACC_SUPER          Opcodes = 32
	ACC_SYNCHRONIZED   Opcodes = 32
	ACC_OPEN           Opcodes = 32
	ACC_TRANSITIVE     Opcodes = 32
	ACC_VOLATILE       Opcodes = 64
	ACC_BRIDGE         Opcodes = 64
	ACC_STATIC_PHASE   Opcodes = 64
	ACC_VARARGS        Opcodes = 128
	ACC_TRANSIENT      Opcodes = 128
	ACC_NATIVE         Opcodes = 256
	ACC_INTERFACE      Opcodes = 512
	ACC_ABSTRACT       Opcodes = 1024
	ACC_STRICT         Opcodes = 2048
	ACC_SYNTHETIC      Opcodes = 4096
	ACC_ANNOTATION     Opcodes = 8192
	ACC_ENUM           Opcodes = 16384
	ACC_MANDATED       Opcodes = 32768
	ACC_MODULE         Opcodes = 32768
	ACC_DEPRECATED     Opcodes = 131072
	T_BOOLEAN          Opcodes = 4
	T_CHAR             Opcodes = 5
	T_FLOAT            Opcodes = 6
	T_DOUBLE           Opcodes = 7
	T_BYTE             Opcodes = 8
	T_SHORT            Opcodes = 9
	T_INT              Opcodes = 10
	T_LONG             Opcodes = 11
	H_GETFIELD         Opcodes = 1
	H_GETSTATIC        Opcodes = 2
	H_PUTFIELD         Opcodes = 3
	H_PUTSTATIC        Opcodes = 4
	H_INVOKEVIRTUAL    Opcodes = 5
	H_INVOKESTATIC     Opcodes = 6
	H_INVOKESPECIAL    Opcodes = 7
	H_NEWINVOKESPECIAL Opcodes = 8
	H_INVOKEINTERFACE  Opcodes = 9
	F_NEW              Opcodes = -1
	F_FULL             Opcodes = 0
	F_APPEND           Opcodes = 1
	F_CHOP             Opcodes = 2
	F_SAME             Opcodes = 3
	F_SAME1            Opcodes = 4
	NOP                Opcodes = 0
	ACONST_NULL        Opcodes = 1
	ICONST_M1          Opcodes = 2
	ICONST_0           Opcodes = 3
	ICONST_1           Opcodes = 4
	ICONST_2           Opcodes = 5
	ICONST_3           Opcodes = 6
	ICONST_4           Opcodes = 7
	ICONST_5           Opcodes = 8
	LCONST_0           Opcodes = 9
	LCONST_1           Opcodes = 10
	FCONST_0           Opcodes = 11
	FCONST_1           Opcodes = 12
	FCONST_2           Opcodes = 13
	DCONST_0           Opcodes = 14
	DCONST_1           Opcodes = 15
	BIPUSH             Opcodes = 16
	SIPUSH             Opcodes = 17
	LDC                Opcodes = 18
	ILOAD              Opcodes = 21
	LLOAD              Opcodes = 22
	FLOAD              Opcodes = 23
	DLOAD              Opcodes = 24
	ALOAD              Opcodes = 25
	IALOAD             Opcodes = 46
	LALOAD             Opcodes = 47
	FALOAD             Opcodes = 48
	DALOAD             Opcodes = 49
	AALOAD             Opcodes = 50
	BALOAD             Opcodes = 51
	CALOAD             Opcodes = 52
	SALOAD             Opcodes = 53
	ISTORE             Opcodes = 54
	LSTORE             Opcodes = 55
	FSTORE             Opcodes = 56
	DSTORE             Opcodes = 57
	ASTORE             Opcodes = 58
	IASTORE            Opcodes = 79
	LASTORE            Opcodes = 80
	FASTORE            Opcodes = 81
	DASTORE            Opcodes = 82
	AASTORE            Opcodes = 83
	BASTORE            Opcodes = 84
	CASTORE            Opcodes = 85
	SASTORE            Opcodes = 86
	POP                Opcodes = 87
	POP2               Opcodes = 88
	DUP                Opcodes = 89
	DUP_X1             Opcodes = 90
	DUP_X2             Opcodes = 91
	DUP2               Opcodes = 92
	DUP2_X1            Opcodes = 93
	DUP2_X2            Opcodes = 94
	SWAP               Opcodes = 95
	IADD               Opcodes = 96
	LADD               Opcodes = 97
	FADD               Opcodes = 98
	DADD               Opcodes = 99
	ISUB               Opcodes = 100
	LSUB               Opcodes = 101
	FSUB               Opcodes = 102
	DSUB               Opcodes = 103
	IMUL               Opcodes = 104
	LMUL               Opcodes = 105
	FMUL               Opcodes = 106
	DMUL               Opcodes = 107
	IDIV               Opcodes = 108
	LDIV               Opcodes = 109
	FDIV               Opcodes = 110
	DDIV               Opcodes = 111
	IREM               Opcodes = 112
	LREM               Opcodes = 113
	FREM               Opcodes = 114
	DREM               Opcodes = 115
	INEG               Opcodes = 116
	LNEG               Opcodes = 117
	FNEG               Opcodes = 118
	DNEG               Opcodes = 119
	ISHL               Opcodes = 120
	LSHL               Opcodes = 121
	ISHR               Opcodes = 122
	LSHR               Opcodes = 123
	IUSHR              Opcodes = 124
	LUSHR              Opcodes = 125
	IAND               Opcodes = 126
	LAND               Opcodes = 127
	IOR                Opcodes = 128
	LOR                Opcodes = 129
	IXOR               Opcodes = 130
	LXOR               Opcodes = 131
	IINC               Opcodes = 132
	I2L                Opcodes = 133
	I2F                Opcodes = 134
	I2D                Opcodes = 135
	L2I                Opcodes = 136
	L2F                Opcodes = 137
	L2D                Opcodes = 138
	F2I                Opcodes = 139
	F2L                Opcodes = 140
	F2D                Opcodes = 141
	D2I                Opcodes = 142
	D2L                Opcodes = 143
	D2F                Opcodes = 144
	I2B                Opcodes = 145
	I2C                Opcodes = 146
	I2S                Opcodes = 147
	LCMP               Opcodes = 148
	FCMPL              Opcodes = 149
	FCMPG              Opcodes = 150
	DCMPL              Opcodes = 151
	DCMPG              Opcodes = 152
	IFEQ               Opcodes = 153
	IFNE               Opcodes = 154
	IFLT               Opcodes = 155
	IFGE               Opcodes = 156
	IFGT               Opcodes = 157
	IFLE               Opcodes = 158
	IF_ICMPEQ          Opcodes = 159
	IF_ICMPNE          Opcodes = 160
	IF_ICMPLT          Opcodes = 161
	IF_ICMPGE          Opcodes = 162
	IF_ICMPGT          Opcodes = 163
	IF_ICMPLE          Opcodes = 164
	IF_ACMPEQ          Opcodes = 165
	IF_ACMPNE          Opcodes = 166
	GOTO               Opcodes = 167
	JSR                Opcodes = 168
	RET                Opcodes = 169
	TABLESWITCH        Opcodes = 170
	LOOKUPSWITCH       Opcodes = 171
	IRETURN            Opcodes = 172
	LRETURN            Opcodes = 173
	FRETURN            Opcodes = 174
	DRETURN            Opcodes = 175
	ARETURN            Opcodes = 176
	RETURN             Opcodes = 177
	GETSTATIC          Opcodes = 178
	PUTSTATIC          Opcodes = 179
	GETFIELD           Opcodes = 180
	PUTFIELD           Opcodes = 181
	INVOKEVIRTUAL      Opcodes = 182
	INVOKESPECIAL      Opcodes = 183
	INVOKESTATIC       Opcodes = 184
	INVOKEINTERFACE    Opcodes = 185
	INVOKEDYNAMIC      Opcodes = 186
	NEW                Opcodes = 187
	NEWARRAY           Opcodes = 188
	ANEWARRAY          Opcodes = 189
	ARRAYLENGTH        Opcodes = 190
	ATHROW             Opcodes = 191
	CHECKCAST          Opcodes = 192
	INSTANCEOF         Opcodes = 193
	MONITORENTER       Opcodes = 194
	MONITOREXIT        Opcodes = 195
	MULTIANEWARRAY     Opcodes = 197
	IFNULL             Opcodes = 198
	IFNONNULL          Opcodes = 199
)