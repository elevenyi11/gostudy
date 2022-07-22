import os.path


def readRsFile():
    zhlines = []
    folder = os.path.abspath('.')
    csfile = os.path.join(folder, 'ResultStatus.proto')
    with open(csfile, 'r', encoding="utf-8") as f:
        needConvert = False
        listLines = f.readlines()
        for i in range(len(listLines)):
            curLine = listLines[i]
            if curLine.find('enum ResultStatus {') > -1:
                needConvert = True

            if needConvert == False:
                continue

            beforeLine = listLines[i-1]
            newLine = replaceLine(beforeLine, curLine)
            if newLine == None:
                continue

            zhlines.append(newLine)

    return zhlines


def replaceLine(beforeLine, curLine):
    # 如果未找到，则认为不需要替换，直接加入结果
    index = curLine.find(' = ')
    if index == -1:
        return

    # 找到key
    kLeftIndex = curLine.find(' = ')+3
    key = curLine[kLeftIndex:len(curLine)-2]

    # 找到value
    vLeftIndex = beforeLine.find('// ')+3
    value = beforeLine[vLeftIndex:len(beforeLine)-1]

    newline = key+" -> "+value
    return newline


def outFile(cont):
    folder = os.path.abspath('.')
    newCsfile = os.path.join(folder, 'code.txt')
    with open(newCsfile, 'w', encoding="utf-8") as f:
        for line in cont:
            f.write(line+'\n')


allLine = readRsFile()
outFile(allLine)
