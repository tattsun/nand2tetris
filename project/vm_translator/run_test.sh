BASEDIR=../..

TOOL_DIR=${BASEDIR}/original/tools
ASSEMBLER=${BASEDIR}/project/assembler/assembler
TRANSLATOR=${BASEDIR}/project/vm_translator/vm_translator
EMULATOR=${TOOL_DIR}/CPUEmulator.sh

CNT=0

RED="\e[31m"
GREEN="\e[32m"
CLEAR="\e[m"

do_test() {
    CNT=$(( ${CNT}+1 ))

    ${TRANSLATOR} ${1}.vm > /tmp/translator_output 2>&1
    if [ $? -ne 0 ]; then
        echo ${RED}Test ${CNT} Failed ${1}: Translator Failed${CLEAR}
        cat /tmp/translator_output
        echo
        return
    fi

    ${ASSEMBLER} ${1}.asm > /tmp/asm_output 2>&1
    if [ $? -ne 0 ]; then
        echo ${RED}Test ${CNT} Failed ${1}: Assembler Failed${CLEAR}
        cat /tmp/asm_output
        echo
        return
    fi

    sh ${EMULATOR} ${1}.tst > /tmp/emulator_output 2>&1
    if [ $? -ne 0 ]; then
        echo ${RED}Test ${CNT} Failed ${1}: Emulator Failed${CLEAR}
        cat /tmp/emulator_output
        echo
        return
    fi

    echo ${GREEN}Test ${CNT} OK ${1}${CLEAR}
}

echo --- Build
make build
if [ $? -ne 0 ]; then
    echo ${RED}Failed to build${CLEAR}
    exit 1
fi
echo


echo --- Test
do_test ${BASEDIR}/07/StackArithmetic/SimpleAdd/SimpleAdd
do_test ${BASEDIR}/07/StackArithmetic/StackTest/StackTest
do_test ${BASEDIR}/07/MemoryAccess/BasicTest/BasicTest
do_test ${BASEDIR}/07/MemoryAccess/PointerTest/PointerTest
do_test ${BASEDIR}/07/MemoryAccess/StaticTest/StaticTest
