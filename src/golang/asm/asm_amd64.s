#include "textflag.h"

// caller的advance进阶版本
TEXT ·callerAdv(SB), $0-16
   MOVQ pc+0(SP), AX
   MOVQ AX, r+8(FP)
   MOVQ a+0(FP), AX
   MOVQ AX, ia+0(SP)
   CALL ·callee(SB)
   MOVQ r+8(FP), AX
   MOVQ AX, pc+0(SP)
   MOVQ z2+8(SP), AX
   MOVQ AX, r2+8(FP)
   RET

TEXT ·caller(SB), $16-16
   MOVQ a+0(FP), AX
   MOVQ AX, ia-16(SP)
   CALL ·callee(SB)
   MOVQ z2-8(SP), AX
   MOVQ AX, r2+8(FP)
   RET

TEXT ·callee(SB),NOSPLIT,$8-16    //
                                // 在编译后，栈空间会被+8用于存储BP寄存器，这步骤由编译器自动添加
   MOVQ    $0, d-8(SP)         // 初始化d[0]
   MOVQ    a+0(FP), AX          // d[0] = a
   MOVQ    AX, d-8(SP)         //
   MOVQ    d-8(SP), AX         // d[0] = return [0]
   MOVQ    AX, r+8(FP)         //
   RET

/*
// yyy0的优化版
TEXT ·yyy00(SB),NOSPLIT,$0-48
   MOVQ pc+0(SP),          AX            // 将PC寄存器中的值暂时保存在最后一个返回值的位置，因为在
                                         // 调用zzz时，该位置不会参与计算
   MOVQ AX,                ret_2+40(FP)  //
   MOVQ a+0(FP),           AX            // 将输入参数a，放置在栈顶
   MOVQ AX,                z_a+0(SP)     //
   MOVQ b+8(FP),           AX            // 将输入参数b，放置在栈顶+8
   MOVQ AX,                z_b+8(SP)     //
   MOVQ c+16(FP),          AX            // 将输入参数c，放置在栈顶+16
   MOVQ AX,                z_c+16(SP)    //
   CALL ·zzz(SB)                         // 调用函数zzz
   MOVQ ret_2+40(FP),      AX            // 将PC寄存器恢复
   MOVQ AX,                pc+0(SP)      //
   MOVQ z_ret_2+40(SP),    AX            // 将zzz的返回值[2]防止在yyy返回值[2]的位置
   MOVQ AX,                ret_2+40(FP)  //
   MOVQ z_ret_1+32(SP),    AX            // 将zzz的返回值[1]防止在yyy返回值[1]的位置
   MOVQ AX,                ret_1+32(FP)  //
   MOVQ z_ret_0+24(SP),    AX            // 将zzz的返回值[0]防止在yyy返回值[0]的位置
   MOVQ AX,                ret_0+24(FP)  //
   RET                                   // return

TEXT ·yyy0(SB), $48-48
   MOVQ a+0(FP), AX
   MOVQ AX, ia-48(SP)
   MOVQ b+8(FP), AX
   MOVQ AX, ib-40(SP)
   MOVQ c+16(FP), AX
   MOVQ AX, ic-32(SP)
   CALL ·zzz(SB)
   MOVQ z2-24(SP), AX
   MOVQ AX, r2+24(FP)
   MOVQ z1-16(SP), AX
   MOVQ AX, r1+32(FP)
   MOVQ z1-8(SP), AX
   MOVQ AX, r2+40(FP)
   RET

TEXT ·zzz(SB),NOSPLIT,$24-48    // $24值栈空间24byte，- 后面的48跟上面的含义一样，
                                // 在编译后，栈空间会被+8用于存储BP寄存器，这步骤由编译器自动添加
   MOVQ    $0, d-24(SP)         // 初始化d[0]
   MOVQ    $0, d-16(SP)         // 初始化d[1]
   MOVQ    $0, d-8(SP)          // 初始化d[2]
   MOVQ    a+0(FP), AX          // d[0] = a
   MOVQ    AX, d-24(SP)         //
   MOVQ    b+8(FP), AX          // d[1] = b
   MOVQ    AX, d-16(SP)         //
   MOVQ    c+16(FP), AX         // d[2] = c
   MOVQ    AX, d-8(SP)          //
   MOVQ    d-24(SP), AX         // d[0] = return [0]
   MOVQ    AX, r+24(FP)         //
   MOVQ    d-16(SP), AX         // d[1] = return [1]
   MOVQ    AX, r+32(FP)         //
   MOVQ    d-8(SP), AX          // d[2] = return [2]
   MOVQ    AX, r+40(FP)         //
   RET                          // return
*/

/*
TEXT ·asmain(SB), $16-0
    MOVQ ·helloworld+0(SB), AX; MOVQ AX, 0(SP)
    MOVQ ·helloworld+8(SB), BX; MOVQ BX, 8(SP)
    CALL runtime·printstring(SB)
    CALL runtime·printnl(SB)
    RET
*/
/*
GLOBL ·Name(SB),$24

DATA ·Name+0(SB)/8,$·Name+16(SB)
DATA ·Name+8(SB)/8,$6
DATA ·Name+16(SB)/8,$"gopher"
*/
/*
GLOBL ·Id(SB),$8
DATA ·Id+0(SB)/1,$0x37
DATA ·Id+1(SB)/1,$0x25
DATA ·Id+2(SB)/1,$0x00
DATA ·Id+3(SB)/1,$0x00
DATA ·Id+4(SB)/1,$0x00
DATA ·Id+5(SB)/1,$0x00
DATA ·Id+6(SB)/1,$0x00
DATA ·Id+7(SB)/1,$0x00
*/
