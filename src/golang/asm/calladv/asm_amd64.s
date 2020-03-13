#include "textflag.h"

// caller的advance进阶版本
TEXT ·caller2(SB), $0-16
   MOVQ pc+0(SP), AX
   MOVQ AX, r+8(FP)
   MOVQ a+0(FP), AX
   MOVQ AX, ia+0(SP)
   CALL ·callee(SB)
   MOVQ r+8(FP), AX
   MOVQ AX, pc+0(SP)
   MOVQ ret+8(SP), AX
   MOVQ AX, r+8(FP)
   RET

TEXT ·callee(SB),NOSPLIT,$8-16

   MOVQ    $0, d-8(SP)         // 初始化d[0]
   MOVQ    a+0(FP), AX          // d[0] = a
   MOVQ    AX, d-8(SP)         //
   MOVQ    d-8(SP), AX         // d[0] = return [0]
   MOVQ    AX, r+8(FP)         //
   RET

/*
TEXT ·caller1(SB), $16-16
   MOVQ a+0(FP), AX
   MOVQ AX, ia-16(SP)
   CALL ·callee(SB)
   MOVQ z2-8(SP), AX
   MOVQ AX, r2+8(FP)
   RET
*/